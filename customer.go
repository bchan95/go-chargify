package chargify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Customer struct {
	FirstName                string       `json:"first_name"`
	LastName                 string       `json:"last_name"`
	Email                    string       `json:"email"`
	CCEmails                 string       `json:"cc_emails"`
	Organization             string       `json:"organization"`
	Reference                string       `json:"reference"`
	ID                       int64        `json:"id"`
	CreatedAt                string       `json:"created_at"`
	UpdatedAt                string       `json:"updated_at"`
	Address                  string       `json:"address"`
	Address2                 string       `json:"address_2"`
	City                     string       `json:"city"`
	State                    string       `json:"state"`
	Country                  string       `json:"country"`
	Phone                    string       `json:"phone"`
	Verfied                  bool         `json:"verfied"`
	PortalCustomerCreatedAt  string       `json:"portal_customer_created_at"`
	PortalInviteLastSend     string       `json:"portal_invite_last_send"`
	PortalInviteLastAccepted string       `json:"portal_invite_last_accepted"`
	TaxExampt                bool         `json:"tax_exampt"`
	VatNumber                string       `json:"vat_number"`
	ParentID                 int64        `json:"parent_id"`
	Metafields               []*Metafield `json:"metafields"`
}

type Metafield struct {
	MetafieldName string `json:"metafield_name"`
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
	if err = checkError(res); err != nil {
		return
	}
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	customer = new(Customer)
	err = json.Unmarshal(body, customer)
	return
}

func GetAllCustomers(client Client) (customers []*Customer, err error) {
	uri := "customers.json"
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
	if err = checkError(res); err != nil {
		return
	}
	defer res.Body.Close()
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
	if err = checkError(res); err != nil {
		return
	}
	defer res.Body.Close()
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
	err = checkError(res)
	return
}
