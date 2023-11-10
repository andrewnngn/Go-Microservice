package main

import (
	"fmt"
	"log"
	"mailer-service/api"
	"mailer-service/mailer"
	"os"
)

const webPort = 8080

func main() {

	//port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(os.Getenv("MAIL_DOMAIN"))
	//marks, err := strconv.ParseInt(os.Getenv("MAIL_SERVER_PORT"), 10, 0)
	//if err != nil {
	//	log.Fatal(err)
	//}

	m := mailer.Mail{
		Domain:      os.Getenv("MAIL_DOMAIN"),
		Host:        os.Getenv("MAIL_HOST"),
		Port:        1025,
		Username:    os.Getenv("MAIL_USERNAME"),
		Password:    os.Getenv("MAIL_PASSWORD"),
		Encryption:  os.Getenv("ENCRYPTION"),
		FromAddress: os.Getenv("FROM_ADDRESS"),
		FromName:    os.Getenv("FORM_NAME"),
	}

	//---	DEVELOPMENT
	// m := mailer.Mail{
	// 	Domain:      "localhost",
	// 	Host:        "localhost",
	// 	Port:        1025,
	// 	Username:    "sontung",
	// 	Password:    "mtp",
	// 	Encryption:  "none",
	// 	FromAddress: "chacaidoseven@gmailcom",
	// 	FromName:    "emcuangayhomqua",
	// }
	//log.Println(os.Getenv("MAIL_DOMAIN"))

	server, err := api.NewServer(m)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		StartRabbitConsumer()
	}()

	err = server.Start(webPort)
	if err != nil {
		return
	}

	fmt.Printf("Running the mail server at port %v\n", os.Getenv("UI_PORT"))
	fmt.Printf("Running the mail client at port %v\n", os.Getenv("UI_PORT"))

}
