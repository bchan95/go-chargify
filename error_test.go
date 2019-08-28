package chargify

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func Test_checkError(t *testing.T) {
	type args struct {
		res *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "200",
			args: args{
				res: &http.Response{
					StatusCode: 200,
				},
			},
		},
		{
			name: "422 no cause",
			args: args{
				res: &http.Response{
					StatusCode: 422,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte{})),
				},
			},
			wantErr: &Error{
				Errors: []string{"unprocessable entity"},
			},
		},
		{
			name: "422 with cause",
			args: args{
				res: &http.Response{
					StatusCode: 422,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"errors": ["mock error"]}`))),
				},
			},
			wantErr: &Error{
				Errors: []string{"mock error"},
			},
		},
		{
			name: "422 with causes",
			args: args{
				res: &http.Response{
					StatusCode: 422,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"errors": ["mock error", "new mock error"]}`))),
				},
			},
			wantErr: &Error{
				Errors: []string{"mock error", "new mock error"},
			},
		},
		{
			name: "404",
			args: args{
				res: &http.Response{
					StatusCode: 404,
				},
			},
			wantErr: errors.New("not found"),
		},
		{
			name: "503",
			args: args{
				res: &http.Response{
					StatusCode: 503,
				},
			},
			wantErr: errors.New("unrecognized response code"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkError(tt.args.res); !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("checkError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
