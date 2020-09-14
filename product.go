package chargify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type ProductBody struct {
	ID                      int64             `json:"id,omitempty"`
	Name                    string            `json:"name,omitempty"`
	Handle                  string            `json:"handle,omitempty"`
	Description             string            `json:"description,omitempty"`
	AccountingCode          string            `json:"accounting_code,omitempty"`
	PriceInCents            int64             `json:"price_in_cents,omitempty"`
	Interval                int64             `json:"interval,omitempty"`
	IntervalUnit            string            `json:"interval_unit,omitempty"`
	InitialChargeInCents    int64             `json:"initial_charge_in_cents,omitempty"`
	ExpirationInterval      int64             `json:"expiration_interval,omitempty"`
	ExpirationIntervalUnit  string            `json:"expiration_interval_unit,omitempty"`
	TrialPriceInCents       int64             `json:"trial_price_in_cents,omitempty"`
	TrialInterval           int64             `json:"trial_interval,omitempty"`
	TrialIntervalUnit       string            `json:"trial_interval_unit,omitempty"`
	InitialChargeAfterTrial bool              `json:"initial_charge_after_trial,omitempty"`
	ReturnParams            string            `json:"return_params,omitempty"`
	RequestCreditCard       bool              `json:"request_credit_card,omitempty"`
	RequireCreditCard       bool              `json:"require_credit_card,omitempty"`
	CreatedAt               string            `json:"created_at,omitempty"`
	UpdatedAt               string            `json:"updated_at,omitempty"`
	ArchivedAt              string            `json:"archived_at,omitempty"`
	UpdateReturnURL         string            `json:"update_return_url,omitempty"`
	UpdateReturnParams      string            `json:"update_return_params,omitempty"`
	ProductFamily           *ProductFamily    `json:"product_family,omitempty"`
	ProductPricePointHandle string            `json:"product_price_point_handle,omitempty"`
	PublicSignupPage        *PublicSignupPage `json:"public_signup_page,omitempty"`
	Taxable                 bool              `json:"taxable,omitempty"`
	VersionNumber           int64             `json:"version_number,omitempty"`
}

type ProductFamily struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Handle      string `json:"handle,omitempty"`
	AccountCode string `json:"account_code,omitempty"`
	Description string `json:"description,omitempty"`
}

type Product struct {
	Product *ProductBody `json:"product"`
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

func GetProductsByFamily(client Client, familyID int64) (products []*Product, err error) {
	if familyID == 0 {
		return nil, NoID()
	}
	uri := fmt.Sprintf("product_families/%d/products.json", familyID)
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
	err = json.Unmarshal(body, &products)
	if err != nil {
		return
	}
	// return sorted by price
	sort.Slice(products, func(i, j int) bool {
		return products[i].Product.PriceInCents < products[j].Product.PriceInCents
	})
	return
}

func CreateProduct(client Client, familyId int64, product *Product) (response *Product, err error) {
	if product.Product == nil {
		return nil, errors.New("missing request")
	}
	// have to nest this because chargify is a mess
	var jsonReq []byte
	jsonReq, err = json.Marshal(product)
	if err != nil {
		return
	}
	var res *http.Response
	res, err = client.Post(jsonReq, fmt.Sprintf("product_families/%d/products.json", familyId))
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
	response = new(Product)
	err = json.Unmarshal(body, response)
	return
}
