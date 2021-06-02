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
	TransactionObject []TransactionBlock `json:"transactionObject"  validate:"required"` 
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

//Post request to get archive pagewise
type GetArchiveRequestPagewise struct {
	
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

	// Page requested (must start with 1)
	//
	// required: true
	Page int64 `json:"page" validate:"required"`
}

//Post response to get archive pagewise
type GetArchiveResponsePagewise struct{
		//Archive data object
		//
		Data PagewiseArchiveReponseFromDB `json:"data"`

		//Response message
		//
		Response string `json:"response"`

		//status code
		//
		Status int64 `json:"status"`
}

type PagewiseArchiveReponseFromDB struct {
	// Date the transaction took place
	//
	TransactionDate string `json:"transactionDate"`

	// Ranking index of the transaction
	//
	RankingIndex int64 `json:"rankingIndex"`

	// Transaction array of objects
	//
	TransactionObject []struct{
		// Response of bot in a block of transaction
		//
		Bot string `json:"bot"`

		// Response of user in a block of transaction (NULL for a terminating block)
		//
		User string `json:"user"`
	} `json:"transactionObject"` 
}

func (d *CreateArchiveRequest) ValidateCreateArchiveRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}

func (d *GetArchiveRequestPagewise) ValidateGetArchiveRequestPagewise() error {
	validate := validator.New()
	return validate.Struct(d)
}