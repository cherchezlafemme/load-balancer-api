package client

import (
	"errors"
)

var (
	// ErrUnauthorized returned when the request is not authorized
	ErrUnauthorized = errors.New("client is unauthorized")

	// ErrPermissionDenied returned when the subject does not permissions to access the resource
	ErrPermissionDenied = errors.New("client does not have permissions")

	// ErrLBNotfound returned when the load balancer ID not found
	ErrLBNotfound = errors.New("loadbalancer ID not found")

	// ErrHTTPError returned when the http response is an error
	ErrHTTPError = errors.New("loadbalancer api http error")
)
