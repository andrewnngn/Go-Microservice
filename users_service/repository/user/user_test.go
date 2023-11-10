package user

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang-boilerplate/ent"
	"golang-boilerplate/internal/mocks/repomocks"
	"golang-boilerplate/model"
	"testing"
)

func TestCreateUser(t *testing.T) {
	// Create a new instance of the mock repository.
	mockRepo := repomocks.NewRepository(t)
	ctx := context.Background()
	// Create an instance of your service, passing the mock repository.
	service := New(nil, nil, nil) // Replace with your actual dependencies
	service = mockRepo            // Assuming your service variable is of type *impl

	// Define the input for the CreateUser function.
	input := ent.CreateUserInput{
		// Provide necessary input data here.
	}

	// Set up expectations for the CreateUser method.
	mockRepo.On("CreateUser", ctx, input).Return(&model.UserReturn{
		Data:          nil,
		IsError:       false,
		MsgFromDev:    "h12",
		MsgFromServer: "h12",
	}, nil)

	// Call the CreateUser function.
	userReturn, err := service.CreateUser(ctx, input)

	// Assert that the expectations were met.
	//mockRepo.AssertExpectations(t)

	// Assert the result.
	assert.EqualValues(t, userReturn, &model.UserReturn{
		Data:          nil,
		IsError:       false,
		MsgFromDev:    "h11",
		MsgFromServer: "h11",
	})
	assert.NoError(t, err)
	// Add additional assertions based on your use case.
}
