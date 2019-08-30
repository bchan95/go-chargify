package chargify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type SubscriptionRequest struct {
	Request       *SubscriptionCreate
	CancelRequest *SubscriptionCancel
}

type SubscriptionCreate struct {
	ProductHandle                 string                  `json:"product_handle"`
	ProductID                     string                  `json:"product_id"`
	Ref                           string                  `json:"ref"`
	CouponCode                    string                  `json:"coupon_code"`
	PaymentCollectionMethod       string                  `json:"payment_collection_method"`
	ReceivesInvoiceEmails         bool                    `json:"receives_invoice_emails"`
	NetTerms                      string                  `json:"net_terms"`
	CustomerID                    string                  `json:"customer_id"`
	NextBillingAt                 string                  `json:"next_billing_at"`
	StoredCredentialTransactionID int64                   `json:"stored_credential_transaction_id"`
	PaymentProfileID              string                  `json:"payment_profile_id"`
	CustomerAttributes            *Customer               `json:"customer_attributes"`
	CreditCardAttributes          *CreditCard             `json:"credit_card_attributes"`
	BankAccountAttributes         *BankAccount            `json:"bank_account_attributes"`
	Components                    []*Component            `json:"components"`
	CalendarBilling               *CalendarBilling        `json:"calendar_billing"`
	Metafields                    *SubscriptionMetafields `json:"metafields"`
}

type SubscriptionResponse struct {
	ID                            int64       `json:"id"`
	State                         string      `json:"state"`
	BalanceInCents                int64       `json:"balance_in_cents"`
	TotalRevenueInCents           int64       `json:"total_revenue_in_cents"`
	ProductPriceInCents           int64       `json:"product_price_in_cents"`
	ProductVersionNumber          int64       `json:"product_version_number"`
	CurrentPeriodEndsAt           string      `json:"current_period_ends_at"`
	NextAssessmentAt              string      `json:"next_assessment_at"`
	TrialStartedAt                string      `json:"trial_started_at"`
	TrialEndedAt                  string      `json:"trial_ended_at"`
	ActivatedAt                   string      `json:"activated_at"`
	CreatedAt                     string      `json:"created_at"`
	UpdatedAt                     string      `json:"updated_at"`
	CancellationMessage           string      `json:"cancellation_message"`
	CancellationMethod            string      `json:"cancellation_method"`
	CancelAtEndOfPeriod           bool        `json:"cancel_at_end_of_period"`
	CanceledAt                    string      `json:"canceled_at"`
	CurrentPeriodStartedAt        string      `json:"current_period_started_at"`
	PreviousState                 string      `json:"previous_state"`
	SignupPaymentID               int64       `json:"signup_payment_id"`
	SignupRevenue                 string      `json:"signup_revenue"`
	DelayedCancelAt               string      `json:"delayed_cancel_at"`
	CouponCode                    string      `json:"coupon_code"`
	PaymentCollectionMethod       string      `json:"payment_collection_method"`
	SnapDay                       string      `json:"snap_day"`
	ReasonCode                    string      `json:"reason_code"`
	ReceivesInvoiceEmails         bool        `json:"receives_invoice_emails"`
	Customer                      *Customer   `json:"customer"`
	Product                       *Product    `json:"product"`
	CreditCard                    *CreditCard `json:"credit_card"`
	PaymentType                   string      `json:"payment_type"`
	ReferralCode                  string      `json:"referral_code"`
	NextProductID                 int64       `json:"next_product_id"`
	CouponUseCount                int64       `json:"coupon_use_count"`
	CouponUsesAllowed             int64       `json:"coupon_uses_allowed"`
	NextProductHandle             string      `json:"next_product_handle"`
	StoredCredentialTransactionID int64       `json:"stored_credential_transaction_id"`
}

type SubscriptionCancel struct {
	subscriptionID      string
	CancellationMessage string `json:"cancellation_message"`
	CancellationMethod  string `json:"cancellation_method"`
	ReasonCode          string `json:"reason_code"`
}

type PublicSignupPage struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}

type CreditCard struct {
	ChargifyToken      string `json:"chargify_token"`
	ID                 int64  `json:"id"`
	PaymentType        string `json:"payment_type"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	MaskedCardNumber   string `json:"masked_card_number"`
	FullNumber         string `json:"full_number"`
	CardType           string `json:"card_type"`
	ExpirationMonth    int64  `json:"expiration_month"`
	ExpirationYear     int64  `json:"expiration_year"`
	BillingAddress     string `json:"billing_address"`
	BillingAddress2    string `json:"billing_address_2"`
	BillingCity        string `json:"billing_city"`
	BillingState       string `json:"billing_state"`
	BillingCountry     string `json:"billing_country"`
	BillingZIP         string `json:"billing_zip"`
	CurrentVault       string `json:"current_vault"`
	CustomerVaultToken string `json:"customer_vault_token"`
	CustomerID         string `json:"customer_id"`
	PaypalEmail        string `json:"paypal_email"`
	PaypalMethodNonce  string `json:"paypal_method_nonce"`
}

type BankAccount struct {
	ChargifyToken         string `json:"chargify_token"`
	BankName              string `json:"bank_name"`
	BankRoutingNumber     string `json:"bank_routing_number"`
	BankAccountNumber     string `json:"bank_account_number"`
	BankAccountType       string `json:"bank_account_type"`
	BankBranchCode        string `json:"bank_branch_code"`
	BankIBAN              string `json:"bank_iban"`
	BankAccountHolderType string `json:"bank_account_holder_type"`
}

type Component struct {
	ComponentID       int64 `json:"component_id"`
	Enabled           bool  `json:"enabled"`
	UnitBalance       int64 `json:"unit_balance"`
	AllocatedQuantity int64 `json:"allocated_quantity"`
	PricePointID      int64 `json:"price_point_id"`
}

type CalendarBilling struct {
	SnapDay                    int64  `json:"snap_day"`
	CalendarBillingFirstCharge string `json:"calendar_billing_first_charge"`
}
type SubscriptionMetafields struct {
	Color    string `json:"color"`
	Comments string `json:"comments"`
}

func (req *SubscriptionRequest) Create(client Client) (response *SubscriptionResponse, err error) {
	if req.Request == nil {
		return nil, errors.New("missing request")
	}
	var jsonReq []byte
	jsonReq, err = json.Marshal(req.Request)
	if err != nil {
		return
	}
	var res *http.Response
	res, err = client.Post(jsonReq, "subscriptions.json")
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
	response = new(SubscriptionResponse)
	err = json.Unmarshal(body, response)
	return
}

func (req *SubscriptionRequest) Update(client Client, subscriptionID string) (response *SubscriptionResponse, err error) {
	if subscriptionID == "" {
		return nil, errors.New("no id")
	}
	if req.Request == nil {
		return nil, errors.New("missing request")
	}
	var jsonReq []byte
	jsonReq, err = json.Marshal(req.Request)
	if err != nil {
		return
	}
	var res *http.Response
	uri := fmt.Sprintf("subscriptions/%s.json", subscriptionID)
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
	response = new(SubscriptionResponse)
	err = json.Unmarshal(body, response)
	return
}

func GetSubscription(client Client, subscriptionID string) (response *SubscriptionResponse, err error) {
	if subscriptionID == "" {
		return nil, errors.New("no id")
	}
	var res *http.Response
	uri := fmt.Sprintf("subscriptions/%s.json", uri)
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
	response = new(SubscriptionResponse)
	err = json.Unmarshal(body, response)
	return
}

func (req *SubscriptionRequest) CancelDelayed(client Client) (err error) {
	if req.CancelRequest == nil {
		return errors.New("missing request")
	}
	if req.CancelRequest.subscriptionID == "" {
		return errors.New("no id")
	}
	var jsonReq []byte
	jsonReq, err = json.Marshal(req.CancelRequest)
	if err != nil {
		return
	}

	uri := fmt.Sprintf("subscriptions/%s/delayed_cancel.json", req.CancelRequest.subscriptionID)
	_, err = client.Post(jsonReq, uri)
	return
}

func (req *SubscriptionRequest) CancelNow(client Client) (response *SubscriptionResponse, err error) {
	if req.CancelRequest == nil {
		return nil, errors.New("missing request")
	}
	if req.CancelRequest.subscriptionID == "" {
		return nil, errors.New("no id")
	}
	var jsonReq []byte
	jsonReq, err = json.Marshal(req.CancelRequest)
	if err != nil {
		return
	}
	var res *http.Response

	uri := fmt.Sprintf("subscriptions/%s.json", req.CancelRequest.subscriptionID)
	res, err = client.Delete(jsonReq, uri)
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
	response = new(SubscriptionResponse)
	err = json.Unmarshal(body, response)
	return
}
