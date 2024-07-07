package users

import (
	"github.com/binodluitel/api/pkg/models/accounts"
	"github.com/binodluitel/api/pkg/models/common"
)

// User Statuses
const (
	UserStatusPending = "pending"
)

// User is information of a user
type User struct {
	common.Identifier `binding:"omitempty"`
	Account           *accounts.Account `json:"account" binding:"omitempty"`
	Address           []common.Address  `json:"address,omitempty"`
	Email             string            `json:"email" binding:"required,email"`
	FirstName         string            `json:"first_name,omitempty" binding:"omitempty,alpha"`
	LastName          string            `json:"last_name,omitempty" binding:"omitempty,alpha"`
	Password          string            `json:"password,omitempty" binding:"required"`
	Phone             []common.Phone    `json:"phone,omitempty"`
	Status            string            `json:"status" binding:"omitempty"`
}

// CreateRequest for creating a new user
type CreateRequest struct {
	User `json:",flatten"`
}

// UpdateRequest for updating an existing user
type UpdateRequest struct {
	User `json:",flatten"`
}
