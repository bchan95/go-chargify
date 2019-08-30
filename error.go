package chargify

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Error struct {
	Errors []string `json:"errors"`
}

func checkError(res *http.Response) error {
	code := res.StatusCode
	switch code {
	case 200:
		return nil
	case 204:
		return nil
	case 422:
		return extractErr(res.Body)
	case 404:
		return errors.New("not found")
	default:
		return errors.New("unrecognized response code")
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
