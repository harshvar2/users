package domain

import "errors"

const (
	//InternalServerError indicates that the server encountered an unexpected condition that prevented it from fulfilling the request
	InternalServerError = "Internal Server Error"
	//BadRequest indicates that the server cannot or will not process the request due to something that is perceived to be a client error
	BadRequest = "Bad Request"
	//InvalidUserID represents an error when the user ID is invalid
	InvalidUserID = "User ID is invalid"
	//MIssingUserID represents an error when the user ID is missing from the request
	MIssingUserID = "User ID is missing"
)

//ErrUserNotFound represents an error when the user ID is not found in the request
var ErrUserNotFound = errors.New("Requested User ID was not found")
