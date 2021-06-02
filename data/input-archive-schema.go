package data

import (
	"github.com/go-playground/validator/v10"
)

//response after archiving a transaction
type CreateArchiveResponse struct{
	//Response message
	Response string `json:"response"`

	//status code
	//
	Status int64 `json:"status"`
}

//post request for creating archive of a transaction
type CreateArchiveRequest struct{
	// UserID of the user 
	//
	// required: true
	// max length: 1000
	UserID string `json: "userID" validate:"required"`

	// BusinessID of the business user is subscribing to
	//
	// required: true
	// max length: 1000
	BusinessID string `json: "businessid" validate:"required"`

	// Date the transaction took place (provided by the server)
	//
	// required: false
	// max length: 1000
	TransactionDate string `json:"transactionDate"`

	// Ranking index of the transaction (provided by the server)
	//
	// required: false
	// max length: 1000
	RankingIndex int64 `json:"rankingIndex"`

	// Transaction object to be sent as a payload
	//
	// required: true
	TransactionOnject []TransactionBlock `json:"transactionOnject"  validate:"required"` 
}

type TransactionBlock struct {

	// Response of bot in a block of transaction
	//
	// required: true
	Bot string `json:"bot" validate:"required"`

	// Response of user in a block of transaction (NULL for a terminating block)
	//
	// required: true
	User string `json:"user" validate:"required"`
}

func (d *CreateArchiveRequest) ValidateCreateArchiveRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}