package chargify

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Error struct {
	Errors []string `json:"errors"`
}

var (
	NotFound     = errors.New("not found")
	Unrecognized = errors.New("unrecognized response code")
)

func checkError(res *http.Response) error {
	code := res.StatusCode
	switch code {
	case 200:
		return nil
	case 201:
		return nil
	case 204:
		return nil
	case 422:
		return extractErr(res.Body)
	case 404:
		return NotFound
	default:
		log.Printf("Unrecognized: %d", res.StatusCode)
		return Unrecognized
	}
}

func (e *Error) Error() string {
	return strings.Join(e.Errors, ", ")
}

// extracts the description of an error from the body of a response
func extractErr(body io.ReadCloser) *Error {
	defer body.Close()
	byteBody, err := ioutil.ReadAll(body)
	if err != nil {
		return &Error{
			Errors: []string{err.Error()},
		}
	}
	if len(byteBody) < 1 {
		return &Error{
			Errors: []string{"unprocessable entity"},
		}
	}
	extractedErrs := new(Error)
	if err = json.Unmarshal(byteBody, extractedErrs); err != nil {
		return &Error{
			Errors: []string{err.Error()},
		}
	}
	return extractedErrs
}

func NoID() error {
	return errors.New("no ID provided")
}
