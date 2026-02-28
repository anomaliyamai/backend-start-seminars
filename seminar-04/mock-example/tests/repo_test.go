package tests

import (
	"errors"
	"testing"

	"example/repo"
	"example/repo/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNameWithMockery(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	mockRepo.On("GetUserByID", 1).Return(&repo.User{ID: 1, Name: "John Doe"}, nil)

	mockRepo.On("GetUserByID", 0).Return(nil, errors.New("user not found"))

	tests := []struct {
		name     string
		id       int
		expected string
		wantErr  bool
	}{
		{"valid user", 1, "John Doe", false},
		{"invalid user", 0, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, err := repo.GetUserName(mockRepo, tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("unexpected error: %v", err)
			}
			assert.Equal(t, tt.expected, name)
		})
	}

	mockRepo.AssertExpectations(t)
}
