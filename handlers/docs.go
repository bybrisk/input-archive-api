// Package classification of Input-Archive API
//
// Documentation for Input-Archive API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta

package handlers
import "github.com/bybrisk/input-archive-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// Message on creating archive of a transaction.
// swagger:response createArchiveResponse
type createArchivePostResponseWrapper struct {
	// response on creating archive of a transaction
	// in: body
	Body data.CreateArchiveResponse
}

// swagger:parameters createNewArchive
type createArchiveParmsWrapper struct {
	// Data structure to create an archive of a transaction
	// in: body
	// required: true
	Body data.CreateArchiveRequest
}

// Message on getting archives pagewise.
// swagger:response getArchiveResponsePagewise
type getArchivePagewisePostResponseWrapper struct {
	// response for fetching archives pagewise
	// in: body
	Body data.GetArchiveResponsePagewise
}

// swagger:parameters getArchivePagewise
type getArchivePagewiseParmsWrapper struct {
	// Data structure to fetch archives pagewise
	// in: body
	// required: true
	Body data.GetArchiveRequestPagewise
}