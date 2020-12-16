package chargify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Statement struct {
	Id                       int64          `json:"id,omitempty"`
	BasicHtmlView            string         `json:"basic_html_view,omitempty"`
	ClosedAt                 string         `json:"closed_at,omitempty"`
	CreatedAt                string         `json:"created_at,omitempty"`
	CustomerBillingAddress   string         `json:"customer_billing_address,omitempty"`
	CustomerBillingAddress2  string         `json:"customer_billing_address_2,omitempty"`
	CustomerBillingCity      string         `json:"customer_billing_city,omitempty"`
	CustomerBillingCountry   string         `json:"customer_billing_country,omitempty"`
	CustomerBillingState     string         `json:"customer_billing_state,omitempty"`
	CustomerBillingZip       string         `json:"customer_billing_zip,omitempty"`
	CustomerFirstName        string         `json:"customer_first_name,omitempty"`
	CustomerLastName         string         `json:"customer_last_name,omitempty"`
	CustomerOrganization     string         `json:"customer_organization,omitempty"`
	CustomerShippingAddress  string         `json:"customer_shipping_address,omitempty"`
	CustomerShippingAddress2 string         `json:"customer_shipping_address_2,omitempty"`
	CustomerShippingCity     string         `json:"customer_shipping_city,omitempty"`
	CustomerShippingCountry  string         `json:"customer_shipping_country,omitempty"`
	CustomerShippingState    string         `json:"customer_shipping_state,omitempty"`
	CustomerShippingZip      string         `json:"customer_shipping_zip,omitempty"`
	EndingBalanceInCents     int64          `json:"ending_balance_in_cents,omitempty"`
	HtmlView                 string         `json:"html_view,omitempty"`
	Memo                     string         `json:"memo,omitempty"`
	OpenedAt                 string         `json:"opened_at,omitempty"`
	SettledAt                string         `json:"settled_at,omitempty"`
	StartingBalanceInCents   int64          `json:"starting_balance_in_cents,omitempty"`
	SubscriptionId           int64          `json:"subscription_id,omitempty"`
	TextView                 string         `json:"text_view,omitempty"`
	UpdatedAt                string         `json:"updated_at,omitempty"`
	TotalInCents             int64          `json:"total_in_cents,omitempty"`
	Transactions             []*Transaction `json:"transactions,omitempty"`
	Events                   []*Event       `json:"events,omitempty"`
}

type Transaction struct {
	Id                     int64       `json:"id,omitempty"`
	SubscriptionId         int64       `json:"subscription_id,omitempty"`
	Type                   string      `json:"type,omitempty"`
	Kind                   string      `json:"kind,omitempty"`
	TransactionType        string      `json:"transaction_type,omitempty"`
	Success                bool        `json:"success,omitempty"`
	AmountInCents          int64       `json:"amount_in_cents,omitempty"`
	Memo                   string      `json:"memo,omitempty"`
	CreatedAt              string      `json:"created_at,omitempty"`
	StartingBalanceInCents int64       `json:"starting_balance_in_cents,omitempty"`
	EndingBalanceInCents   int64       `json:"ending_balance_in_cents,omitempty"`
	GatewayUsed            string      `json:"gateway_used,omitempty"`
	GatewayTransactionId   string      `json:"gateway_transaction_id,omitempty"`
	GatewayOrderId         string      `json:"gateway_order_id,omitempty"`
	PaymentId              int64       `json:"payment_id,omitempty"`
	ProductId              int64       `json:"product_id,omitempty"`
	TaxId                  string      `json:"tax_id,omitempty"`
	ComponentId            int64       `json:"component_id,omitempty"`
	StatementId            int64       `json:"statement_id,omitempty"`
	CustomerId             int64       `json:"customer_id,omitempty"`
	ItemName               string      `json:"item_name,omitempty"`
	OriginalAmountInCents  int64       `json:"original_amount_in_cents,omitempty"`
	DiscountAmountInCents  int64       `json:"discount_amount_in_cents,omitempty"`
	TaxableAmountInCents   int64       `json:"taxable_amount_in_cents,omitempty"`
	Taxations              []*Taxation `json:"taxations,omitempty"`
}

type Taxation struct {
	TaxId                 string     `json:"tax_id,omitempty"`
	TaxChargeId           int64      `json:"tax_charge_id,omitempty"`
	TaxName               string     `json:"tax_name,omitempty"`
	Rate                  string     `json:"rate,omitempty"`
	TaxAmountInCents      int64      `json:"tax_amount_in_cents,omitempty"`
	TaxRules              []*TaxRule `json:"tax_rules,omitempty"`
	CardNumber            string     `json:"card_number,omitempty"`
	CardExpiration        string     `json:"card_expiration,omitempty"`
	CardType              string     `json:"card_type,omitempty"`
	RefundedAmountInCents int64      `json:"refunded_amount_in_cents,omitempty"`
}

type TaxRule struct {
	TaxRuleId        string `json:"tax_rule_id,omitempty"`
	CountryCode      string `json:"country_code,omitempty"`
	SubdivisionName  string `json:"subdivision_name,omitempty"`
	Rate             string `json:"rate,omitempty"`
	TaxAmountInCents int64  `json:"tax_amount_in_cents,omitempty"`
	Description      string `json:"description,omitempty"`
}

type Event struct {
	Id             int64  `json:"id,omitempty"`
	Key            string `json:"key,omitempty"`
	Message        string `json:"message,omitempty"`
	SubscriptionId int64  `json:"subscription_id,omitempty"`
	CreatedAt      string `json:"created_at,omitempty"`
}

func GetSubscriptionStatements(client Client, subscriptionId int64, pageNumber int32, perPage int32) ([]*Statement, error) {
	if subscriptionId == 0 {
		return nil, NoID()
	}
	uri := fmt.Sprintf("subscriptions/%d/statements.json?direction=desc&per_page=%d&page=%d", subscriptionId, perPage, pageNumber)
	res, err := client.Get(uri)
	if err != nil {
		return nil, err
	}
	if err = checkError(res); err != nil {
		return nil, err
	}
	nestedStatements := make([]*struct {
		Statement *Statement
	}, 0)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &nestedStatements); err != nil {
		return nil, err
	}
	var statements []*Statement
	for _, s := range nestedStatements {
		statements = append(statements, s.Statement)
	}
	return statements, nil
}

func GetStatement(client Client, statementId int64) (statement *Statement, err error) {
	if statementId == 0 {
		return nil, NoID()
	}
	uri := fmt.Sprintf("statements/%d.json", statementId)
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
	statement = new(Statement)
	err = json.Unmarshal(body, statement)
	return
}
