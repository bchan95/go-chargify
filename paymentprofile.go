package chargify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PaymentProfile struct {
	CustomerID         int64  `json:"customer_id"`
	ChargifyToken      string `json:"chargify_token,omitempty"`
	ID                 int64  `json:"id,omitempty"`
	PaymentType        string `json:"payment_type,omitempty"`
	FirstName          string `json:"first_name,omitempty"`
	LastName           string `json:"last_name,omitempty"`
	MaskedCardNumber   string `json:"masked_card_number,omitempty"`
	FullNumber         string `json:"full_number,omitempty"`
	CardType           string `json:"card_type,omitempty"`
	ExpirationMonth    int64  `json:"expiration_month,omitempty"`
	ExpirationYear     int64  `json:"expiration_year,omitempty"`
	BillingAddress     string `json:"billing_address,omitempty"`
	BillingAddress2    string `json:"billing_address_2,omitempty"`
	BillingCity        string `json:"billing_city,omitempty"`
	BillingState       string `json:"billing_state,omitempty"`
	BillingCountry     string `json:"billing_country,omitempty"`
	BillingZIP         string `json:"billing_zip,omitempty"`
	CurrentVault       string `json:"current_vault,omitempty"`
	CustomerVaultToken string `json:"customer_vault_token,omitempty"`
	PaypalEmail        string `json:"paypal_email,omitempty"`
	PaypalMethodNonce  string `json:"paypal_method_nonce,omitempty"`
}

type PaymentProfileRequest struct {
	PaymentProfile *PaymentProfile `json:"payment_profile"`
}

type PaymentProfileResponse struct {
	Payment *PaymentProfileResponseBody `json:"payment"`
}

type PaymentProfileResponseBody struct {
	ID                     int64  `json:"id"`
	SubscriptionID         int64  `json:"subscription_id"`
	Type                   string `json:"type"`
	Kind                   string `json:"kind"`
	TransactionType        string `json:"transaction_type"`
	Success                bool   `json:"success"`
	AmountInCents          int64  `json:"amount_in_cents"`
	Memo                   string `json:"memo"`
	CreatedAt              string `json:"created_at"`
	StartingBalanceInCents int64  `json:"starting_balance_in_cents"`
	EndingBalanceInCents   int64  `json:"ending_balance_in_cents"`
	GatewayUsed            string `json:"gateway_used"`
	GatewayTransactionID   string `json:"gateway_transaction_id"`
	GatewayOrderID         string `json:"gateway_order_id"`
	PaymentID              string `json:"payment_id"`
	ProductID              int64  `json:"product_id"`
	TaxID                  string `json:"tax_id"`
	ComponentID            int64  `json:"component_id"`
	StatementID            int64  `json:"statement_id"`
	CustomerID             int64  `json:"customer_id"`
	CardNumber             string `json:"card_number"`
	CardExpiration         string `json:"card_expiration"`
	CardType               string `json:"card_type"`
	RefundedAmountInCents  string `json:"refunded_amount_in_cents"`
}

func (pp *PaymentProfile) Update(client Client) (response *PaymentProfileResponse, err error) {
	if pp.ID == 0 {
		return nil, errors.New("no payment profile id present")
	}
	if pp.ChargifyToken != "" && pp.CustomerID == 0 {
		return nil, errors.New("cannot pass chargify token without customer id")
	}
	uri := fmt.Sprintf("payment_profile/%d.json", pp.ID)
	updateRequest := &PaymentProfileRequest{
		PaymentProfile: pp,
	}
	var jsonReq []byte
	jsonReq, err = json.Marshal(updateRequest)
	if err != nil {
		return
	}
	var res *http.Response
	res, err = client.Put(jsonReq, uri)
	if err != nil {
		return
	}
	if err = checkError(res); err != nil {
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	response = new(PaymentProfileResponse)
	err = json.Unmarshal(body, response)
	return

}
