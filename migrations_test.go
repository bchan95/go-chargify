package chargify

import (
	"bytes"
	"github.com/bchan95/go-chargify/test"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestMigration_Create(t *testing.T) {
	type fields struct {
		Id        int64
		Migration *MigrationBody
	}
	type args struct {
		client Client
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	tests := []struct {
		name         string
		fields       fields
		args         args
		stub         func()
		wantResponse *Migration
		wantErr      bool
	}{
		{
			name: "create",
			fields: fields{
				Id: 1,
				Migration: &MigrationBody{
					ProductId:            1,
					ProductPricePointId:  1,
					IncludeTrial:         false,
					IncludeInitialCharge: false,
					IncludeCoupons:       false,
					PreservePeriod:       false,
				},
			},
			args: args{
				client: client,
			},
			stub: func() {
				res, _ := json.Marshal(&Migration{
					Migration: &MigrationBody{
						ProductId:            1,
						ProductPricePointId:  1,
						IncludeTrial:         false,
						IncludeInitialCharge: false,
						IncludeCoupons:       false,
						PreservePeriod:       false,
					},
				})
				client.EXPECT().Post(gomock.Any(), "subscriptions/1/migrations.json").Return(
					&http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewBuffer(res)),
					}, nil)
			},
			wantResponse: &Migration{
				Migration: &MigrationBody{
					ProductId:            1,
					ProductPricePointId:  1,
					IncludeTrial:         false,
					IncludeInitialCharge: false,
					IncludeCoupons:       false,
					PreservePeriod:       false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.stub != nil {
				tt.stub()
			}
			m := &Migration{
				Id:        tt.fields.Id,
				Migration: tt.fields.Migration,
			}
			gotResponse, err := m.Create(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Create() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestMigration_Preview(t *testing.T) {
	type fields struct {
		Id        int64
		Migration *MigrationBody
	}
	type args struct {
		client Client
	}
	ctrl := gomock.NewController(t)
	client := test.NewMockClient(ctrl)
	tests := []struct {
		name         string
		fields       fields
		args         args
		stub         func()
		wantResponse *MigrationResponse
		wantErr      bool
	}{
		{
			name: "create",
			fields: fields{
				Id: 1,
				Migration: &MigrationBody{
					ProductId:            1,
					ProductPricePointId:  1,
					IncludeTrial:         false,
					IncludeInitialCharge: false,
					IncludeCoupons:       false,
					PreservePeriod:       false,
				},
			},
			args: args{
				client: client,
			},
			stub: func() {
				res, _ := json.Marshal(&MigrationResponse{
					Migration: &MigrationPreview{
						ProratedAdjustmentInCents: 10,
						ChargeInCents:             200,
						PaymentDueInCents:         180,
						CreditAppliedInCents:      10,
					},
				})
				client.EXPECT().Post(gomock.Any(), "subscriptions/1/migrations/preview.json").Return(
					&http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewBuffer(res)),
					}, nil)
			},
			wantResponse: &MigrationResponse{
				Migration: &MigrationPreview{
					ProratedAdjustmentInCents: 10,
					ChargeInCents:             200,
					PaymentDueInCents:         180,
					CreditAppliedInCents:      10,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.stub != nil {
				tt.stub()
			}
			m := &Migration{
				Id:        tt.fields.Id,
				Migration: tt.fields.Migration,
			}
			gotResponse, err := m.Preview(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("Preview() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Preview() gotResponse = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
