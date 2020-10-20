package resterr

import (
	"errors"
	"net/http"
	"testing"

	"github.com/matryer/is"
)

func TestInternalServer(t *testing.T) {
	is := is.New(t)
	err := InternalServerError("message", errors.New("database error"))

	is.Equal(err.Status, http.StatusInternalServerError)
	is.Equal(err.Message, "message")
	is.Equal(err.Type, "internal_server_error")

	is.True(err.Causes != nil)
	is.Equal(len(err.Causes), 1)
	is.Equal(err.Causes[0], "database error")
}

func TestNotFound(t *testing.T) {
	is := is.New(t)
	err := NotFound("message")
	is.Equal(err.Status, http.StatusNotFound)
	is.Equal(err.Message, "message")
	is.Equal(err.Type, "not_found")
}

func TestBadRequest(t *testing.T) {
	is := is.New(t)
	err := BadRequest("message")
	is.Equal(err.Status, http.StatusBadRequest)
	is.Equal(err.Message, "message")
	is.Equal(err.Type, "bad_request")
}
