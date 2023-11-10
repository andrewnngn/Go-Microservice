package api

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"golang-boilerplate/config"
	"golang-boilerplate/ent"
	"golang-boilerplate/ent/migrate"
	grpcClient "golang-boilerplate/internal/grpc/client"
	"golang-boilerplate/resolver"
	"net/http"
	"os"
)

func NewServerCmd(configs *config.Configurations, logger *zap.Logger, grpcConn *grpcClient.GRPCClientInit, db *ent.Client) *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "run api server",
		Long:  "run api server with graphql",
		Run: func(cmd *cobra.Command, args []string) {
			// Connect to postgresql database
			defer db.Close()

			// Run the automation migration tool

			if err := db.Schema.Create(
				context.Background(),
				migrate.WithGlobalUniqueID(false), // add here diff
			); err != nil {
				logger.Error("Failed to creating db schema from the automation migration tool", zap.Error(err))
				os.Exit(1)
			}

			//if err := db.Schema.Create(
			//	context.Background(),
			//	schema.WithAtlas(true), // add here diff
			//); err != nil {
			//	logger.Error("Failed to creating db schema from the automation migration tool", zap.Error(err))
			//	os.Exit(1)
			//}

			// End migration config

			// Create validator
			validator := validator.New()
			// Add translator for validator
			en := en.New()
			uni := ut.New(en, en)
			validationTranslator, _ := uni.GetTranslator("en")
			// Register default translation for validator
			err := en_translations.RegisterDefaultTranslations(validator, validationTranslator)
			if err != nil {
				logger.Error("Getting error from register default translation", zap.Error(err))
				os.Exit(1)
			}

			// GraphQL schema resolver handler.
			resolverHandler := handler.NewDefaultServer(resolver.NewExecutableSchema(db, logger, grpcConn))
			// Handler for GraphQL Playground
			playgroundHandler := playground.Handler(configs.Trivial.PlaygroundName, "/graphql")
			// Handle a method not allowed.

			app := fiber.New(fiber.Config{
				CaseSensitive: false,
				StrictRouting: false,
				AppName:       configs.Trivial.AppName,
				ServerHeader:  configs.Trivial.ServerHeader,
			})

			// Use middlewares
			app.Use(
				recover.New(),
				requestid.New(requestid.Config{
					Header: "X-Request-ID",
					Generator: func() string {
						return uuid.New().String()
					},
				}),
				cors.New(cors.Config{
					AllowOrigins:     "*",
					AllowMethods:     "GET,POST,HEAD",
					AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Request-ID",
					AllowCredentials: true,
				}),
				compress.New(compress.Config{
					Level: compress.LevelBestCompression,
				}),
			)

			// Create a new GraphQL
			app.Post("/graphql", adaptor.HTTPHandler(resolverHandler))

			app.Options("/graphql", func(c *fiber.Ctx) error {
				return c.SendStatus(http.StatusOK)
			})

			// Enable playground for query/testing
			app.Get("/", adaptor.HTTPHandler(playgroundHandler))

			graphqlPort := fmt.Sprintf(":%v", configs.Port.GraphqlPort)
			if err := app.Listen(graphqlPort); err != nil {
				logger.Error("Get error from run server", zap.Error(err))
			}
		},
	}
}
