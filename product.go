package chargify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	ID                      int64             `json:"id"`
	Name                    string            `json:"name"`
	Handle                  string            `json:"handle"`
	Description             string            `json:"description"`
	AccountingCode          string            `json:"accounting_code"`
	PriceInCents            int64             `json:"price_in_cents"`
	Interval                int64             `json:"interval"`
	IntervalUnit            string            `json:"interval_unit"`
	InitialChargeInCents    int64             `json:"initial_charge_in_cents"`
	ExpirationInterval      int64             `json:"expiration_interval"`
	ExpirationIntervalUnit  string            `json:"expiration_interval_unit"`
	TrialPriceInCents       int64             `json:"trial_price_in_cents"`
	TrialInterval           int64             `json:"trial_interval"`
	TrialIntervalUnit       string            `json:"trial_interval_unit"`
	InitialChargeAfterTrial bool              `json:"initial_charge_after_trial"`
	ReturnParams            string            `json:"return_params"`
	RequestCreditCard       bool              `json:"request_credit_card"`
	RequireCreditCard       bool              `json:"require_credit_card"`
	CreatedAt               string            `json:"created_at"`
	UpdatedAt               string            `json:"updated_at"`
	ArchivedAt              string            `json:"archived_at"`
	UpdateReturnURL         string            `json:"update_return_url"`
	UpdateReturnParams      string            `json:"update_return_params"`
	ProductFamily           *ProductFamily    `json:"product_family"`
	PublicSignupPage        *PublicSignupPage `json:"public_signup_page"`
	Taxable                 bool              `json:"taxable"`
	VersionNumber           int64             `json:"version_number"`
}

type ProductFamily struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Handle      string `json:"handle"`
	AccountCode string `json:"account_code"`
	Description string `json:"description"`
}

func GetProductByID(client Client, productID int64) (product *Product, err error) {
	if productID == 0 {
		return nil, NoID()
	}
	uri := fmt.Sprintf("products/%d.json", productID)
	var res *http.Response
	res, err = client.Get(uri)
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
	product = new(Product)
	err = json.Unmarshal(body, product)
	return
}

func GetProductByHandle(client Client, handle string) (product *Product, err error) {
	if handle == "" {
		return nil, errors.New("no handle provided")
	}
	uri := fmt.Sprintf("products/handle/%s.json", handle)
	var res *http.Response
	res, err = client.Get(uri)
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
	product = new(Product)
	err = json.Unmarshal(body, product)
	return
}
