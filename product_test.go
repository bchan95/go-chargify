package chargify

import (
	"github.com/bchan95/go-chargify/test"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestGetProductByID(t *testing.T) {
	type args struct {
		client    Client
		stub      func()
		productID int64
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	res := &Product{
		ID:          123456789,
		Name:        "test",
		Handle:      "test_handle",
		Description: "a test product",
	}
	body, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name        string
		args        args
		wantProduct *Product
		wantErr     error
	}{
		{
			name: "get by id",
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Get("products/123456789.json").Return(
						&http.Response{
							StatusCode: 200,
							Body:       ioutil.NopCloser(bytes.NewReader(body)),
						}, nil)
				},
				productID: 123456789,
			},
			wantProduct: res,
		},
		{
			name:    "get no id",
			wantErr: NoID(),
		},
	}
	for _, tt := range tests {
		if tt.args.stub != nil {
			tt.args.stub()
		}
		t.Run(tt.name, func(t *testing.T) {
			gotProduct, err := GetProductByID(tt.args.client, tt.args.productID)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetProductByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProduct, tt.wantProduct) {
				t.Errorf("GetProductByID() = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}

func TestGetProductByHandle(t *testing.T) {
	type args struct {
		client Client
		stub   func()
		handle string
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	res := &Product{
		ID:          123456789,
		Name:        "test",
		Handle:      "test_handle",
		Description: "a test product",
	}
	body, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name        string
		args        args
		wantProduct *Product
		wantErr     error
	}{
		{
			name: "get by id",
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Get("products/handle/test_handle.json").Return(
						&http.Response{
							StatusCode: 200,
							Body:       ioutil.NopCloser(bytes.NewReader(body)),
						}, nil)
				},
				handle: "test_handle",
			},
			wantProduct: res,
		},
		{
			name:    "get no id",
			wantErr: errors.New("no handle provided"),
		},
	}
	for _, tt := range tests {
		if tt.args.stub != nil {
			tt.args.stub()
		}
		t.Run(tt.name, func(t *testing.T) {
			gotProduct, err := GetProductByHandle(tt.args.client, tt.args.handle)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetProductByHandle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProduct, tt.wantProduct) {
				t.Errorf("GetProductByHandle() = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}
