/*
 * Project X
 *
 * OpenAPI definition for project X endpoint and resources
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ApiApiController binds http requests to an api service and writes the service results to the http response
type ApiApiController struct {
	service      ApiApiServicer
	errorHandler ErrorHandler
}

// ApiApiOption for how the controller is set up.
type ApiApiOption func(*ApiApiController)

// WithApiApiErrorHandler inject ErrorHandler into controller
func WithApiApiErrorHandler(h ErrorHandler) ApiApiOption {
	return func(c *ApiApiController) {
		c.errorHandler = h
	}
}

// NewApiApiController creates a default api controller
func NewApiApiController(s ApiApiServicer, opts ...ApiApiOption) Router {
	controller := &ApiApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the ApiApiController
func (c *ApiApiController) Routes() Routes {
	return Routes{
		{
			"CreateTest",
			strings.ToUpper("Post"),
			"/tests",
			c.CreateTest,
		},
		{
			"GetTests",
			strings.ToUpper("Get"),
			"/tests",
			c.GetTests,
		},
	}
}

// CreateTest - Create new test
func (c *ApiApiController) CreateTest(w http.ResponseWriter, r *http.Request) {
	testParam := Test{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&testParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTestRequired(testParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateTest(r.Context(), testParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTests - Create new test
func (c *ApiApiController) GetTests(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetTests(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}