package chargify

import (
	"github.com/bchan95/go-chargify/test"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

var mockErr = errors.New("mock error")

func TestSubscriptionRequest_Create(t *testing.T) {
	type fields struct {
		Request       *SubscriptionCreate
		CancelRequest *SubscriptionCancel
	}
	type args struct {
		client Client
		stub   func()
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	req := &SubscriptionCreate{
		ProductHandle: "basic",
		CustomerAttributes: &CustomerBody{
			FirstName: "test1",
			LastName:  "test2",
			Email:     "test@talkatoo.ai",
		},
		CreditCardAttributes: &CreditCard{
			PaymentType:     "Credit",
			FirstName:       "test1",
			LastName:        "test2",
			FullNumber:      "0000000000000001",
			ExpirationMonth: 1,
			ExpirationYear:  2020,
		},
	}
	res := &SubscriptionResponse{
		ID:       123456789,
		Customer: req.CustomerAttributes,
	}
	body, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse *SubscriptionResponse
		wantErr      error
	}{
		{
			name: "create",
			fields: fields{
				Request: req,
			},
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Post(gomock.Any(), "subscriptions.json").Return(&http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewReader(body)),
					}, nil)
				},
			},
			wantResponse: res,
		},
		{
			name:    "create no req",
			wantErr: errors.New("missing request"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.stub != nil {
				tt.args.stub()
			}
			req := &SubscriptionRequest{
				Request:       tt.fields.Request,
				CancelRequest: tt.fields.CancelRequest,
			}
			gotResponse, err := req.Create(tt.args.client)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SubscriptionRequest.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("SubscriptionRequest.Create() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestSubscriptionRequest_Update(t *testing.T) {
	type fields struct {
		Request       *SubscriptionCreate
		CancelRequest *SubscriptionCancel
	}
	type args struct {
		client         Client
		subscriptionID int64
		stub           func()
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	req := &SubscriptionCreate{
		ProductHandle: "basic",
		CustomerAttributes: &CustomerBody{
			FirstName: "test1",
			LastName:  "test2",
			Email:     "test@talkatoo.ai",
		},
		CreditCardAttributes: &CreditCard{
			PaymentType:     "Credit",
			FirstName:       "test1",
			LastName:        "test2",
			FullNumber:      "0000000000000001",
			ExpirationMonth: 1,
			ExpirationYear:  2020,
		},
	}
	res := &SubscriptionResponse{
		ID:       123456789,
		Customer: req.CustomerAttributes,
	}
	body, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse *SubscriptionResponse
		wantErr      error
	}{
		{
			name: "update",
			fields: fields{
				Request: req,
			},
			args: args{
				client:         client,
				subscriptionID: 123456789,
				stub: func() {
					client.EXPECT().Put(gomock.Any(), "subscriptions/123456789.json").Return(&http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewReader(body)),
					}, nil)
				},
			},
			wantResponse: res,
		},
		{
			name: "create no req",
			args: args{
				subscriptionID: 123456789,
			},
			wantErr: errors.New("missing request"),
		},
		{
			name:    "create no id",
			wantErr: NoID(),
		},
	}
	for _, tt := range tests {
		if tt.args.stub != nil {
			tt.args.stub()
		}
		t.Run(tt.name, func(t *testing.T) {
			req := &SubscriptionRequest{
				Request:       tt.fields.Request,
				CancelRequest: tt.fields.CancelRequest,
			}
			gotResponse, err := req.Update(tt.args.client, tt.args.subscriptionID)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SubscriptionRequest.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("SubscriptionRequest.Update() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestGetSubscription(t *testing.T) {
	type args struct {
		client         Client
		subscriptionID int64
		stub           func()
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	res := &SubscriptionResponse{
		ID: 123456789,
		Customer: &CustomerBody{
			FirstName: "test1",
			LastName:  "test2",
			Email:     "test@talkatoo.ai",
		},
	}
	body, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name         string
		args         args
		wantResponse *SubscriptionResponse
		wantErr      error
	}{
		{
			name: "get",
			args: args{
				client:         client,
				subscriptionID: 123456789,
				stub: func() {
					client.EXPECT().Get("subscriptions/123456789.json").Return(&http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewReader(body)),
					}, nil)
				},
			},
			wantResponse: res,
		},
		{
			name:    "create no id",
			wantErr: NoID(),
		},
	}
	for _, tt := range tests {
		if tt.args.stub != nil {
			tt.args.stub()
		}
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := GetSubscription(tt.args.client, tt.args.subscriptionID)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("GetSubscription() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("GetSubscription() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestSubscriptionRequest_CancelNow(t *testing.T) {
	type fields struct {
		Request       *SubscriptionCreate
		CancelRequest *SubscriptionCancel
	}
	type args struct {
		client Client
		stub   func()
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	res := &SubscriptionResponse{
		ID: 123456789,
		Customer: &CustomerBody{
			FirstName: "test1",
			LastName:  "test2",
			Email:     "test@talkatoo.ai",
		},
	}
	body, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantResponse *SubscriptionResponse
		wantErr      error
	}{
		{
			name: "cancel",
			fields: fields{
				CancelRequest: &SubscriptionCancel{
					SubscriptionID:      "123456789",
					ReasonCode:          "r1",
					CancellationMessage: "GOOD DAY SIR",
				},
			},
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Delete(gomock.Any(), "subscriptions/123456789.json").Return(
						&http.Response{
							StatusCode: 200,
							Body:       ioutil.NopCloser(bytes.NewReader(body)),
						}, nil)
				},
			},
			wantResponse: res,
		},
		{
			name:    "create no req",
			wantErr: errors.New("missing request"),
		},
		{
			name: "create no id",
			fields: fields{
				CancelRequest: &SubscriptionCancel{
					ReasonCode:          "r1",
					CancellationMessage: "GOOD DAY SIR",
				},
			},
			wantErr: NoID(),
		},
	}
	for _, tt := range tests {
		if tt.args.stub != nil {
			tt.args.stub()
		}
		t.Run(tt.name, func(t *testing.T) {
			req := &SubscriptionRequest{
				Request:       tt.fields.Request,
				CancelRequest: tt.fields.CancelRequest,
			}
			gotResponse, err := req.CancelNow(tt.args.client)
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SubscriptionRequest.Cancel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("SubscriptionRequest.Cancel() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestSubscriptionRequest_CancelDelayed(t *testing.T) {
	type fields struct {
		Request       *SubscriptionCreate
		CancelRequest *SubscriptionCancel
	}
	type args struct {
		client Client
		stub   func()
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "cancel delayed",
			fields: fields{
				CancelRequest: &SubscriptionCancel{
					SubscriptionID:      "123456789",
					ReasonCode:          "r1",
					CancellationMessage: "GOOD DAY SIR",
				},
			},
			args: args{
				client: client,
				stub: func() {
					client.EXPECT().Post(gomock.Any(), "subscriptions/123456789/delayed_cancel.json").Return(nil, nil)
				},
			},
		},
		{
			name:    "create no req",
			wantErr: errors.New("missing request"),
		},
		{
			name: "create no id",
			fields: fields{
				CancelRequest: &SubscriptionCancel{
					ReasonCode:          "r1",
					CancellationMessage: "GOOD DAY SIR",
				},
			},
			wantErr: NoID(),
		},
	}
	for _, tt := range tests {
		if tt.args.stub != nil {
			tt.args.stub()
		}
		t.Run(tt.name, func(t *testing.T) {
			req := &SubscriptionRequest{
				Request:       tt.fields.Request,
				CancelRequest: tt.fields.CancelRequest,
			}
			if err := req.CancelDelayed(tt.args.client); !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("SubscriptionRequest.CancelDelayed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
