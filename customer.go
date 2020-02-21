package chargify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CustomerBody struct {
	FirstName                string       `json:"first_name,omitempty"`
	LastName                 string       `json:"last_name,omitempty"`
	Email                    string       `json:"email,omitempty"`
	CCEmails                 string       `json:"cc_emails,omitempty"`
	Organization             string       `json:"organization,omitempty"`
	Reference                string       `json:"reference,omitempty"`
	ID                       int64        `json:"id,omitempty"`
	CreatedAt                string       `json:"created_at,omitempty"`
	UpdatedAt                string       `json:"updated_at,omitempty"`
	Address                  string       `json:"address,omitempty"`
	Address2                 string       `json:"address_2,omitempty"`
	City                     string       `json:"city,omitempty"`
	State                    string       `json:"state,omitempty"`
	Zip                      string       `json:"zip"`
	Country                  string       `json:"country,omitempty"`
	Phone                    string       `json:"phone,omitempty"`
	Verfied                  bool         `json:"verfied,omitempty"`
	PortalCustomerCreatedAt  string       `json:"portal_customer_created_at,omitempty"`
	PortalInviteLastSend     string       `json:"portal_invite_last_send,omitempty"`
	PortalInviteLastAccepted string       `json:"portal_invite_last_accepted,omitempty"`
	TaxExampt                bool         `json:"tax_exampt,omitempty"`
	VatNumber                string       `json:"vat_number,omitempty"`
	ParentID                 int64        `json:"parent_id,omitempty"`
	Metafields               []*Metafield `json:"metafields,omitempty"`
}

type Customer struct {
	Customer *CustomerBody `json:"customer"`
}

type Metafield struct {
	MetafieldName string `json:"metafield_name,omitempty"`
}

func GetCustomer(client Client, customerID int64) (customer *Customer, err error) {
	if customerID == 0 {
		return nil, NoID()
	}
	var res *http.Response
	uri := fmt.Sprintf("customers/%d.json", customerID)
	res, err = client.Get(uri)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if err = checkError(res); err != nil {
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	customer = new(Customer)
	err = json.Unmarshal(body, customer)
	return
}

func GetCustomerByEmail(client Client, email string) (customers []*Customer, err error) {
	if email == "" {
		return nil, errors.New("no email specified")
	}
	uri := fmt.Sprintf("customers.json?q=%s", email)
	res, err := client.Get(uri)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if err = checkError(res); err != nil {
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &customers)
	return
}

func GetAllCustomers(client Client) (customers []*Customer, err error) {
	uri := "customers.json"
	var res *http.Response
	res, err = client.Get(uri)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if err = checkError(res); err != nil {
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &customers)
	return
}

func GetCustomerSubscriptions(client Client, customerID int64) (subscriptions []*SubscriptionResponse, err error) {
	if customerID == 0 {
		return nil, NoID()
	}
	uri := fmt.Sprintf("customers/%d/subscriptions.json", customerID)
	var res *http.Response
	res, err = client.Get(uri)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if err = checkError(res); err != nil {
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &subscriptions)
	return
}

func (c *Customer) Update(client Client, customerID int64) (customer *Customer, err error) {
	if customerID == 0 {
		return nil, NoID()
	}
	var jsonReq []byte
	jsonReq, err = json.Marshal(c)
	if err != nil {
		return
	}
	uri := fmt.Sprintf("customers/%d.json", customerID)
	var res *http.Response
	res, err = client.Put(jsonReq, uri)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if err = checkError(res); err != nil {
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	customer = new(Customer)
	err = json.Unmarshal(body, customer)
	return
}

func DeleteCustomer(client Client, customerID int64) (err error) {
	if customerID == 0 {
		return NoID()
	}
	uri := fmt.Sprintf("customers/%d.json", customerID)
	var res *http.Response
	res, err = client.Delete(nil, uri)
	if err != nil {
		return
	}
	defer res.Body.Close()
	return checkError(res)
}
