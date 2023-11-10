package group

import (
	"github.com/stretchr/testify/mock"
	"golang-boilerplate/ent"
	"golang-boilerplate/internal/mocks/repomocks"
	"golang-boilerplate/model"
	"testing"
	"time"
)

func TestGroupRepo_CreateGroup(t *testing.T) {
	repo := &repomocks.Repository{}

	t.Run("success", func(t *testing.T) {
		user1 := ent.CreateGroupInput{
			Name:      "a",
			Address:   "a",
			Email:     "a",
			Role:      "a",
			CreatedAt: nil,
			UpdatedAt: nil,
			UserIDs:   nil,
		}

		expRes := &model.GroupReturn{
			Data: &ent.Group{
				ID:        0,
				Name:      "a",
				Address:   "a",
				Email:     "a",
				Role:      "a",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				Edges:     ent.GroupEdges{},
			},
			IsError:       false,
			MsgFromDev:    "",
			MsgFromServer: "",
		}
		repo.On("CreateGroup", mock.Anything, user1).Return(expRes, nil).Once()
		repo.AssertExpectations(t)
	})
}
