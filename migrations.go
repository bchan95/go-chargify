package chargify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Migration struct {
	Id        int64          `json:"-"`
	Migration *MigrationBody `json:"migration,omitempty"`
}

type MigrationBody struct {
	ProductId               int64  `json:"product_id,omitempty"`
	ProductHandle           string `json:"product_handle,omitempty"`
	ProductPricePointId     int64  `json:"product_price_point_id,omitempty"`
	ProductPricePointHandle string `json:"product_price_point_handle,omitempty"`
	IncludeTrial            bool   `json:"include_trial,omitempty"`
	IncludeInitialCharge    bool   `json:"include_initial_charge,omitempty"`
	IncludeCoupons          bool   `json:"include_coupons,omitempty"`
	PreservePeriod          bool   `json:"preserve_period,omitempty"`
}

type MigrationResponse struct {
	Migration *MigrationPreview `json:"migration"`
}

type MigrationPreview struct {
	ProratedAdjustmentInCents int64 `json:"prorated_adjustment_in_cents,omitempty"`
	ChargeInCents             int64 `json:"charge_in_cents,omitempty"`
	PaymentDueInCents         int64 `json:"payment_due_in_cents,omitempty"`
	CreditAppliedInCents      int64 `json:"credit_applied_in_cents,omitempty"`
}

// Migrate a subscription from one product to another.
// NOTE: This will not update any component price points associated with the product, so update those first.
func (m *Migration) Create(client Client) (response *Migration, err error) {
	if m.Migration == nil {
		return nil, errors.New("missing request")
	}
	// have to nest this because chargify is a mess
	var jsonReq []byte
	jsonReq, err = json.Marshal(m)
	if err != nil {
		return
	}
	var res *http.Response
	res, err = client.Post(jsonReq, fmt.Sprintf("subscriptions/%d/migrations.json", m.Id))
	if err != nil {
		return
	}
	if err = checkError(res); err != nil {
		return
	}
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	response = new(Migration)
	err = json.Unmarshal(body, response)
	return
}

// Preview a migration before creating it.
func (m *Migration) Preview(client Client) (response *MigrationResponse, err error) {
	if m.Migration == nil {
		return nil, errors.New("missing request")
	}
	// have to nest this because chargify is a mess
	var jsonReq []byte
	jsonReq, err = json.Marshal(m)
	if err != nil {
		return
	}
	var res *http.Response
	res, err = client.Post(jsonReq, fmt.Sprintf("subscriptions/%d/migrations/preview.json", m.Id))
	if err != nil {
		return
	}
	if err = checkError(res); err != nil {
		return
	}
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	response = new(MigrationResponse)
	err = json.Unmarshal(body, response)
	return
}
