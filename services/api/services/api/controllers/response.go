package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

// Response is the base HTTP response used by all controllers
type Response struct {
	Error      error `json:"-"`
	statusCode int
	data       map[string]interface{}

	raw, rawType string
	html         string
	mfaEnabled   bool
}

// Render will render the response object over HTTP with the specified response data
func (r *Response) Render(w http.ResponseWriter, req *http.Request) error {
	if r.html != "" {
		render.Status(req, r.statusCode)
		render.HTML(w, req, r.html)
		return nil
	}

	if r.raw != "" {
		switch r.rawType {
		case "json":
			render.Status(req, r.statusCode)
			render.JSON(w, req, r.data)
			return nil
		}
	}

	if r.data == nil {
		r.data = map[string]interface{}{}
	}

	if r.statusCode >= 200 && r.statusCode <= 299 {
		render.Status(req, r.statusCode)
		render.JSON(w, req, r.data)
		return nil
	}

	if r.mfaEnabled {
		type errorData struct {
			Code   string `json:"code"`
			Status int    `json:"status"`
		}
		res := errorData{
			Code:   "MFA_ENABLED",
			Status: http.StatusUnauthorized,
		}
		r.Set("error", res)
	}

	if r.Error != nil {
		type errorData struct {
			Error string `json:"error"`
		}

		res := errorData{r.Error.Error()}
		r.Set("errors", []interface{}{res})
	}

	render.Status(req, r.statusCode)
	render.JSON(w, req, r.data)

	return nil
}

// BadRequest retuns the HTTP status code 400, bad request
func (r *Response) BadRequest(w http.ResponseWriter, req *http.Request, message string) {
	r.statusCode = http.StatusBadRequest
	r.Error = fmt.Errorf(message)
	r.Render(w, req)
}

// InternalServerError retuns the HTTP status code 500, internal server error
func (r *Response) InternalServerError(w http.ResponseWriter, req *http.Request, message string) {
	r.statusCode = http.StatusInternalServerError
	r.Error = fmt.Errorf(message)
	r.Render(w, req)
}

// JSON returns a JSON response
func (r *Response) JSON(w http.ResponseWriter, req *http.Request, data any, statusCode int) {
	render.Status(req, statusCode)
	render.JSON(w, req, data)
}

// Set takes a key, value pair and adds them to the Response data map
func (r *Response) Set(key string, value interface{}) {
	if r.data == nil {
		r.data = map[string]interface{}{
			key: value,
		}
	}

	r.data[key] = value
}

// NotFound retuns the HTTP status code 404, page not found
func (r *Response) NotFound(w http.ResponseWriter, req *http.Request, message string) {
	r.statusCode = http.StatusNotFound
	r.Error = fmt.Errorf(message)
	r.Render(w, req)
}
