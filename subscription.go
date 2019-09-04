package chargify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SubscriptionRequest struct {
	Request       *SubscriptionCreate
	CancelRequest *SubscriptionCancel
}

type SubscriptionCreate struct {
	ProductHandle                 string                  `json:"product_handle,omitempty"`
	ProductID                     string                  `json:"product_id,omitempty"`
	Ref                           string                  `json:"ref,omitempty"`
	CouponCode                    string                  `json:"coupon_code,omitempty"`
	PaymentCollectionMethod       string                  `json:"payment_collection_method,omitempty"`
	ReceivesInvoiceEmails         bool                    `json:"receives_invoice_emails,omitempty"`
	NetTerms                      string                  `json:"net_terms,omitempty"`
	CustomerID                    string                  `json:"customer_id,omitempty"`
	NextBillingAt                 string                  `json:"next_billing_at,omitempty"`
	StoredCredentialTransactionID int64                   `json:"stored_credential_transaction_id,omitempty"`
	PaymentProfileID              string                  `json:"payment_profile_id,omitempty"`
	CustomerAttributes            *CustomerBody           `json:"customer_attributes,omitempty"`
	CreditCardAttributes          *CreditCard             `json:"credit_card_attributes,omitempty"`
	BankAccountAttributes         *BankAccount            `json:"bank_account_attributes,omitempty"`
	Components                    []*Component            `json:"components,omitempty"`
	CalendarBilling               *CalendarBilling        `json:"calendar_billing,omitempty"`
	Metafields                    *SubscriptionMetafields `json:"metafields,omitempty"`
}

type SubscriptionResponse struct {
	ID                            int64         `json:"id,omitempty"`
	State                         string        `json:"state,omitempty"`
	BalanceInCents                int64         `json:"balance_in_cents,omitempty"`
	TotalRevenueInCents           int64         `json:"total_revenue_in_cents,omitempty"`
	ProductPriceInCents           int64         `json:"product_price_in_cents,omitempty"`
	ProductVersionNumber          int64         `json:"product_version_number,omitempty"`
	CurrentPeriodEndsAt           string        `json:"current_period_ends_at,omitempty"`
	NextAssessmentAt              string        `json:"next_assessment_at,omitempty"`
	TrialStartedAt                string        `json:"trial_started_at,omitempty"`
	TrialEndedAt                  string        `json:"trial_ended_at,omitempty"`
	ActivatedAt                   string        `json:"activated_at,omitempty"`
	CreatedAt                     string        `json:"created_at,omitempty"`
	UpdatedAt                     string        `json:"updated_at,omitempty"`
	CancellationMessage           string        `json:"cancellation_message,omitempty"`
	CancellationMethod            string        `json:"cancellation_method,omitempty"`
	CancelAtEndOfPeriod           bool          `json:"cancel_at_end_of_period,omitempty"`
	CanceledAt                    string        `json:"canceled_at,omitempty"`
	CurrentPeriodStartedAt        string        `json:"current_period_started_at,omitempty"`
	PreviousState                 string        `json:"previous_state,omitempty"`
	SignupPaymentID               int64         `json:"signup_payment_id,omitempty"`
	SignupRevenue                 string        `json:"signup_revenue,omitempty"`
	DelayedCancelAt               string        `json:"delayed_cancel_at,omitempty"`
	CouponCode                    string        `json:"coupon_code,omitempty"`
	PaymentCollectionMethod       string        `json:"payment_collection_method,omitempty"`
	SnapDay                       string        `json:"snap_day,omitempty"`
	ReasonCode                    string        `json:"reason_code,omitempty"`
	ReceivesInvoiceEmails         bool          `json:"receives_invoice_emails,omitempty"`
	Customer                      *CustomerBody `json:"customer,omitempty"`
	Product                       *Product      `json:"product,omitempty"`
	CreditCard                    *CreditCard   `json:"credit_card,omitempty"`
	PaymentType                   string        `json:"payment_type,omitempty"`
	ReferralCode                  string        `json:"referral_code,omitempty"`
	NextProductID                 int64         `json:"next_product_id,omitempty"`
	CouponUseCount                int64         `json:"coupon_use_count,omitempty"`
	CouponUsesAllowed             int64         `json:"coupon_uses_allowed,omitempty"`
	NextProductHandle             string        `json:"next_product_handle,omitempty"`
	StoredCredentialTransactionID int64         `json:"stored_credential_transaction_id,omitempty"`
}

type SubscriptionCancel struct {
	SubscriptionID      string
	CancellationMessage string `json:"cancellation_message,omitempty"`
	CancellationMethod  string `json:"cancellation_method,omitempty"`
	ReasonCode          string `json:"reason_code,omitempty"`
}

type PublicSignupPage struct {
	ID  int64  `json:"id,omitempty"`
	URL string `json:"url,omitempty"`
}

type CreditCard struct {
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
	CustomerID         string `json:"customer_id,omitempty"`
	PaypalEmail        string `json:"paypal_email,omitempty"`
	PaypalMethodNonce  string `json:"paypal_method_nonce,omitempty"`
}

type BankAccount struct {
	ChargifyToken         string `json:"chargify_token,omitempty"`
	BankName              string `json:"bank_name,omitempty"`
	BankRoutingNumber     string `json:"bank_routing_number,omitempty"`
	BankAccountNumber     string `json:"bank_account_number,omitempty"`
	BankAccountType       string `json:"bank_account_type,omitempty"`
	BankBranchCode        string `json:"bank_branch_code,omitempty"`
	BankIBAN              string `json:"bank_iban,omitempty"`
	BankAccountHolderType string `json:"bank_account_holder_type,omitempty"`
}

type Component struct {
	ComponentID       int64 `json:"component_id,omitempty"`
	Enabled           bool  `json:"enabled,omitempty"`
	UnitBalance       int64 `json:"unit_balance,omitempty"`
	AllocatedQuantity int64 `json:"allocated_quantity,omitempty"`
	PricePointID      int64 `json:"price_point_id,omitempty"`
}

type CalendarBilling struct {
	SnapDay                    int64  `json:"snap_day,omitempty"`
	CalendarBillingFirstCharge string `json:"calendar_billing_first_charge,omitempty"`
}
type SubscriptionMetafields struct {
	Color    string `json:"color,omitempty"`
	Comments string `json:"comments,omitempty"`
}

func (req *SubscriptionRequest) Create(client Client) (response *SubscriptionResponse, err error) {
	if req.Request == nil {
		return nil, errors.New("missing request")
	}
	// have to nest this because chargify is a mess
	var jsonReq []byte
	jsonReq, err = json.Marshal(req.wrap())
	if err != nil {
		return
	}
	log.Println(string(jsonReq))
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
	err = json.Unmarshal(body, response.wrap())
	return
}

func (req *SubscriptionRequest) Update(client Client, subscriptionID string) (response *SubscriptionResponse, err error) {
	if subscriptionID == "" {
		return nil, NoID()
	}
	if req.Request == nil {
		return nil, errors.New("missing request")
	}
	var jsonReq []byte
	jsonReq, err = json.Marshal(req.wrap())
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
	err = json.Unmarshal(body, response.wrap())
	return
}

func GetSubscription(client Client, subscriptionID string) (response *SubscriptionResponse, err error) {
	if subscriptionID == "" {
		return nil, NoID()
	}
	var res *http.Response
	uri := fmt.Sprintf("subscriptions/%s.json", subscriptionID)
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
	err = json.Unmarshal(body, response.wrap())
	return
}

func (req *SubscriptionRequest) CancelDelayed(client Client) (err error) {
	if req.CancelRequest == nil {
		return errors.New("missing request")
	}
	if req.CancelRequest.SubscriptionID == "" {
		return NoID()
	}
	var jsonReq []byte
	jsonReq, err = json.Marshal(req.wrap())
	if err != nil {
		return
	}
	uri := fmt.Sprintf("subscriptions/%s/delayed_cancel.json", req.CancelRequest.SubscriptionID)
	_, err = client.Post(jsonReq, uri)
	return
}

func (req *SubscriptionRequest) CancelNow(client Client) (response *SubscriptionResponse, err error) {
	if req.CancelRequest == nil {
		return nil, errors.New("missing request")
	}
	if req.CancelRequest.SubscriptionID == "" {
		return nil, NoID()
	}
	var jsonReq []byte
	jsonReq, err = json.Marshal(req.wrap())
	if err != nil {
		return
	}
	var res *http.Response

	uri := fmt.Sprintf("subscriptions/%s.json", req.CancelRequest.SubscriptionID)
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
	err = json.Unmarshal(body, response.wrap())
	return
}

func (req *SubscriptionRequest) wrap() interface{} {
	if req.Request != nil {
		return &struct {
			Subscription *SubscriptionCreate `json:"subscription"`
		}{
			Subscription: req.Request,
		}
	}
	if req.CancelRequest != nil {
		return &struct {
			Subscription *SubscriptionCancel `json:"subscription"`
		}{
			Subscription: req.CancelRequest,
		}
	}
	return nil
}

func (res *SubscriptionResponse) wrap() interface{} {
	return &struct {
		Subscription *SubscriptionResponse `json:"subscription"`
	}{
		Subscription: res,
	}
}
