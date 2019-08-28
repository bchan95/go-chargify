package chargify

import "context"

type SubscriptionRequest struct {
	Request       SubscriptionCreate
	CancelRequest SubscriptionCancel
}

type SubscriptionCreate struct {
	ProductHandle                 string                 `json:"product_handle"`
	ProductID                     string                 `json:"product_id"`
	Ref                           string                 `json:"ref"`
	CouponCode                    string                 `json:"coupon_code"`
	PaymentCollectionMethod       string                 `json:"payment_collection_method"`
	ReceivesInvoiceEmails         bool                   `json:"receives_invoice_emails"`
	NetTerms                      string                 `json:"net_terms"`
	CustomerID                    string                 `json:"customer_id"`
	NextBillingAt                 string                 `json:"next_billing_at"`
	StoredCredentialTransactionID int64                  `json:"stored_credential_transaction_id"`
	PaymentProfileID              string                 `json:"payment_profile_id"`
	CustomerAttributes            Customer               `json:"customer_attributes"`
	CreditCardAttributes          CreditCard             `json:"credit_card_attributes"`
	BankAccountAttributes         BankAccount            `json:"bank_account_attributes"`
	Components                    []*Component           `json:"components"`
	CalendarBilling               CalendarBilling        `json:"calendar_billing"`
	Metafields                    SubscriptionMetafields `json:"metafields"`
}

type SubscriptionResponse struct {
	ID                            int64      `json:"id"`
	State                         string     `json:"state"`
	BalanceInCents                int64      `json:"balance_in_cents"`
	TotalRevenueInCents           int64      `json:"total_revenue_in_cents"`
	ProductPriceInCents           int64      `json:"product_price_in_cents"`
	ProductVersionNumber          int64      `json:"product_version_number"`
	CurrentPeriodEndsAt           string     `json:"current_period_ends_at"`
	NextAssessmentAt              string     `json:"next_assessment_at"`
	TrialStartedAt                string     `json:"trial_started_at"`
	TrialEndedAt                  string     `json:"trial_ended_at"`
	ActivatedAt                   string     `json:"activated_at"`
	CreatedAt                     string     `json:"created_at"`
	UpdatedAt                     string     `json:"updated_at"`
	CancellationMessage           string     `json:"cancellation_message"`
	CancellationMethod            string     `json:"cancellation_method"`
	CancelAtEndOfPeriod           bool       `json:"cancel_at_end_of_period"`
	CanceledAt                    string     `json:"canceled_at"`
	CurrentPeriodStartedAt        string     `json:"current_period_started_at"`
	PreviousState                 string     `json:"previous_state"`
	SignupPaymentID               int64      `json:"signup_payment_id"`
	SignupRevenue                 string     `json:"signup_revenue"`
	DelayedCancelAt               string     `json:"delayed_cancel_at"`
	CouponCode                    string     `json:"coupon_code"`
	PaymentCollectionMethod       string     `json:"payment_collection_method"`
	SnapDay                       string     `json:"snap_day"`
	ReasonCode                    string     `json:"reason_code"`
	ReceivesInvoiceEmails         bool       `json:"receives_invoice_emails"`
	Customer                      Customer   `json:"customer"`
	Product                       Product    `json:"product"`
	CreditCard                    CreditCard `json:"credit_card"`
	PaymentType                   string     `json:"payment_type"`
	ReferralCode                  string     `json:"referral_code"`
	NextProductID                 int64      `json:"next_product_id"`
	CouponUseCount                int64      `json:"coupon_use_count"`
	CouponUsesAllowed             int64      `json:"coupon_uses_allowed"`
	NextProductHandle             string     `json:"next_product_handle"`
	StoredCredentialTransactionID int64      `json:"stored_credential_transaction_id"`
}

type SubscriptionCancel struct {
	subscriptionID      string
	CancellationMessage string `json:"cancellation_message"`
	CancellationMethod  string `json:"cancellation_method"`
	ReasonCode          string `json:"reason_code"`
}

type Product struct {
	ID                      int64            `json:"id"`
	Name                    string           `json:"name"`
	Handle                  string           `json:"handle"`
	Description             string           `json:"description"`
	AccountingCode          string           `json:"accounting_code"`
	PriceInCents            int64            `json:"price_in_cents"`
	Interval                int64            `json:"interval"`
	IntervalUnit            string           `json:"interval_unit"`
	InitialChargeInCents    int64            `json:"initial_charge_in_cents"`
	ExpirationInterval      int64            `json:"expiration_interval"`
	ExpirationIntervalUnit  string           `json:"expiration_interval_unit"`
	TrialPriceInCents       int64            `json:"trial_price_in_cents"`
	TrialInterval           int64            `json:"trial_interval"`
	TrialIntervalUnit       string           `json:"trial_interval_unit"`
	InitialChargeAfterTrial bool             `json:"initial_charge_after_trial"`
	ReturnParams            string           `json:"return_params"`
	RequestCreditCard       bool             `json:"request_credit_card"`
	RequireCreditCard       bool             `json:"require_credit_card"`
	CreatedAt               string           `json:"created_at"`
	UpdatedAt               string           `json:"updated_at"`
	ArchivedAt              string           `json:"archived_at"`
	UpdateReturnURL         string           `json:"update_return_url"`
	UpdateReturnParams      string           `json:"update_return_params"`
	ProductFamily           ProductFamily    `json:"product_family"`
	PublicSignupPage        PublicSignupPage `json:"public_signup_page"`
	Taxable                 bool             `json:"taxable"`
	VersionNumber           int64            `json:"version_number"`
}

type ProductFamily struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Handle      string `json:"handle"`
	AccountCode string `json:"account_code"`
	Description string `json:"description"`
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

func (req *SubscriptionRequest) Create(ctx context.Context, client *Client) (response *SubscriptionResponse, err error) {
	return
}

func (req *SubscriptionRequest) Update(ctx context.Context, client *Client) (response *SubscriptionResponse, err error) {
	return
}

func GetSubscription(ctx context.Context, client *Client, subscriptionID string) (response *SubscriptionResponse, err error) {
	return
}

func (req SubscriptionRequest) Cancel(ctx context.Context, client *Client) (response *SubscriptionResponse, err error) {
	return
}
