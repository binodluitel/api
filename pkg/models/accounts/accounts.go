package accounts

import (
	"github.com/binodluitel/api/pkg/models/common"
)

// Account Statuses
const (
	// AccountStatusPending indicates that the account is pending approval or activation.
	AccountStatusPending = "pending"
)

// Account information
type Account struct {
	common.Identifier
	Name   string            `json:"name" binding:"omitempty"`
	Owner  common.Identifier `json:"owner" binding:"omitempty"`
	Status string            `json:"status" binding:"omitempty"`
}
