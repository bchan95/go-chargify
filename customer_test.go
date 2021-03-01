package chargify

import (
	"github.com/bchan95/go-chargify/test"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestGetCustomer(t *testing.T) {
	type args struct {
		client     Client
		stub       func()
		customerID int64
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	res := &Customer{
		FirstName: "First",
		LastName:  "McName",
		Email:     "hello@email.com",
	}
	body, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name         string
		args         args
		wantCustomer *Customer
		wantErr      error
	}{
		{
			name: "get customer",
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Get("customers/123456789.json").Return(&http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewReader(body)),
					}, nil)
				},
				customerID: 123456789,
			},
			wantCustomer: res,
		},
		{
			name:    "get customer, no id",
			args:    args{},
			wantErr: NoID(),
		},
		{
			name: "get customer, err",
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Get("customers/123456789.json").Return(nil, mockErr)
				},
				customerID: 123456789,
			},
			wantErr: mockErr,
		},
	}
	for _, tt := range tests {
		if tt.args.stub != nil {
			tt.args.stub()
		}
		t.Run(tt.name, func(t *testing.T) {
			gotCustomer, err := GetCustomer(tt.args.client, tt.args.customerID)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomer, tt.wantCustomer) {
				t.Errorf("GetCustomer() = %v, want %v", gotCustomer, tt.wantCustomer)
			}
		})
	}
}

func TestGetAllCustomers(t *testing.T) {
	type args struct {
		client Client
		stub   func()
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	res := []*Customer{
		{
			FirstName: "First",
			LastName:  "McName",
			Email:     "hello@email.com",
		},
		{
			FirstName: "Second",
			LastName:  "McName",
			Email:     "helloagain@email.com",
		},
	}
	body, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name          string
		args          args
		wantCustomers []*Customer
		wantErr       error
	}{
		{
			name: "get all customers",
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Get("customers.json").Return(&http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewReader(body)),
					}, nil)
				},
			},
			wantCustomers: res,
		},
		{
			name: "get all customers, err",
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Get("customers.json").Return(nil, mockErr)
				},
			},
			wantErr: mockErr,
		},
	}
	for _, tt := range tests {
		if tt.args.stub != nil {
			tt.args.stub()
		}
		t.Run(tt.name, func(t *testing.T) {
			gotCustomers, err := GetAllCustomers(tt.args.client)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetAllCustomers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCustomers, tt.wantCustomers) {
				t.Errorf("GetAllCustomers() = %v, want %v", gotCustomers, tt.wantCustomers)
			}
		})
	}
}

func TestGetCustomerSubscriptions(t *testing.T) {
	type args struct {
		client     Client
		stub       func()
		customerID int64
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	res := []*SubscriptionResponse{
		{
			ID: 9876543210,
		},
		{
			ID: 9876543211,
		},
	}
	body, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name              string
		args              args
		wantSubscriptions []*SubscriptionResponse
		wantErr           error
	}{
		{
			name: "get customer subscriptions",
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Get("customers/123456789/subscriptions.json").Return(&http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewReader(body)),
					}, nil)
				},
				customerID: 123456789,
			},
			wantSubscriptions: res,
		},
		{
			name:    "get customer subscription, no id",
			wantErr: NoID(),
		},
		{
			name: "get customer subscriptions, err",
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Get("customers/123456789/subscriptions.json").Return(nil, mockErr)
				},
				customerID: 123456789,
			},
			wantErr: mockErr,
		},
	}
	for _, tt := range tests {
		if tt.args.stub != nil {
			tt.args.stub()
		}
		t.Run(tt.name, func(t *testing.T) {
			gotSubscriptions, err := GetCustomerSubscriptions(tt.args.client, tt.args.customerID)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetCustomerSubscriptions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotSubscriptions, tt.wantSubscriptions) {
				t.Errorf("GetCustomerSubscriptions() = %v, want %v", gotSubscriptions, tt.wantSubscriptions)
			}
		})
	}
}

func TestCustomer_Update(t *testing.T) {
	type fields struct {
		c *Customer
	}
	type args struct {
		client     Client
		stub       func()
		customerID int64
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	res := &Customer{
		FirstName: "NewName",
		LastName:  "McName",
		Email:     "hello@email.com",
		Address2:  "new address",
	}
	body, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantCustomer *Customer
		wantErr      error
	}{
		{
			name: "update customer",
			fields: fields{
				c: &Customer{
					FirstName: "NewName",
					Address2:  "new address",
				},
			},
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Put(gomock.Any(), "customers/123456789.json").Return(&http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewReader(body)),
					}, nil)
				},
				customerID: 123456789,
			},
			wantCustomer: res,
		},
		{
			name:    "update customer no id",
			wantErr: NoID(),
		},
		{
			name: "update customer, err",
			fields: fields{
				c: &Customer{
					FirstName: "NewName",
					Address2:  "new address",
				},
			},
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Put(gomock.Any(), "customers/123456789.json").Return(nil, mockErr)
				},
				customerID: 123456789,
			},
			wantErr: mockErr,
		},
	}
	for _, tt := range tests {
		if tt.args.stub != nil {
			tt.args.stub()
		}
		t.Run(tt.name, func(t *testing.T) {
			gotCustomer, err := tt.fields.c.Update(tt.args.client, tt.args.customerID)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("Customer.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantCustomer, gotCustomer) {
				t.Errorf("UpdateCustomer() got = %v, want %v", gotCustomer, tt.wantCustomer)
			}
		})
	}
}

func TestDeleteCustomer(t *testing.T) {
	type args struct {
		client     Client
		stub       func()
		customerID int64
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "delete customer",
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Delete(nil, "customers/123456789.json").Return(&http.Response{
						StatusCode: 204,
					}, nil)
				},
				customerID: 123456789,
			},
		},
		{
			name:    "delete customer, no id",
			wantErr: NoID(),
		},
		{
			name: "delete customer, err",
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Delete(nil, "customers/123456789.json").Return(nil, mockErr)
				},
				customerID: 123456789,
			},
			wantErr: mockErr,
		},
	}
	for _, tt := range tests {
		if tt.args.stub != nil {
			tt.args.stub()
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteCustomer(tt.args.client, tt.args.customerID); !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("DeleteCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
