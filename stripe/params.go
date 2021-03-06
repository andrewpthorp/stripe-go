package stripe

// BankAccountParams hold all of the parameters used for creating and updating
// BankAccounts.
type BankAccountParams struct {
	Country       string `stripe_field:"bank_account[country]"`
	RoutingNumber string `stripe_field:"bank_account[routing_number]"`
	AccountNumber string `stripe_field:"bank_account[account_number]"`
}

// CardParams hold all of the parameters used for creating and updating Cards.
type CardParams struct {
	Number         string
	ExpMonth       int
	ExpYear        int
	CVC            string
	Name           string
	AddressLine1   string
	AddressLine2   string
	AddressCity    string
	AddressZip     string
	AddressState   string
	AddressCountry string
	Token          string
}

// ChargeParams hold all of the parameters used for creating Charges.
type ChargeParams struct {
	Amount         int    `stripe_field:"amount"`
	Currency       string `stripe_field:"currency"`
	Customer       string `stripe_field:"customer"`
	Description    string `stripe_field:"description"`
	DisableCapture bool   `stripe_field:"capture" opposite:"true"`
	ApplicationFee int    `stripe_field:"application_fee"`
	*CardParams
	Metadata
}

// CouponParams hold all of the parameters used for creating Coupons.
type CouponParams struct {
	Id               string `stripe_field:"id"`
	Duration         string `stripe_field:"duration"`
	AmountOff        int    `stripe_field:"amount_off"`
	Currency         string `stripe_field:"currency"`
	DurationInMonths int    `stripe_field:"duration_in_months"`
	MaxRedemptions   int    `stripe_field:"max_redemptions"`
	PercentOff       int    `stripe_field:"percent_off"`
	RedeemBy         int    `stripe_field:"redeem_by"`
}

// CustomerParams hold all of the parameters used for creating and updating
// Customers.
type CustomerParams struct {
	AccountBalance int    `stripe_field:"account_balance"`
	Coupon         string `stripe_field:"coupon"`
	Description    string `stripe_field:"description"`
	Email          string `stripe_field:"email"`
	Plan           string `stripe_field:"plan"`
	Quantity       int    `stripe_field:"quantity"`
	TrialEnd       int    `stripe_field:"trial_end"`
	*CardParams
	Metadata
}

// InvoiceParams hold all of the parameters used for creating and updating
// Invoices.
type InvoiceParams struct {
	Customer       string `stripe_field:"customer"`
	ApplicationFee int    `stripe_field:"application_fee"`
	Closed         bool   `stripe_field:"closed"`
}

// InvoiceItemParams hold all of the parameters used for creating and updating
// InvoiceItems.
type InvoiceItemParams struct {
	Customer    string `stripe_field:"customer"`
	Amount      int    `stripe_field:"amount"`
	Currency    string `stripe_field:"currency"`
	Invoice     string `stripe_field:"invoice"`
	Description string `stripe_field:"description"`
	Metadata
}

// PlanParams hold all of the parameters used for creating and updating Plans.
type PlanParams struct {
	Id              string `stripe_field:"id"`
	Amount          int    `stripe_field:"amount"`
	Currency        string `stripe_field:"currency"`
	Interval        string `stripe_field:"interval"`
	IntervalCount   int    `stripe_field:"interval_count"`
	Name            string `stripe_field:"name"`
	TrialPeriodDays int    `stripe_field:"trial_period_days"`
	Metadata
}

// RecipientParams hold all of the parameters used for creating and updating
// Recipients.
type RecipientParams struct {
	Name        string `stripe_field:"name"`
	Type        string `stripe_field:"type"`
	TaxId       string `stripe_field:"tax_id"`
	Email       string `stripe_field:"email"`
	Description string `stripe_field:"description"`
	*BankAccountParams
	Metadata
}

// RefundParams hold all of the parameters used for refunding Charges.
type RefundParams struct {
	Amount               int  `stripe_field:"amount"`
	RefundApplicationFee bool `stripe_field:"refund_application_fee"`
}

// SubscriptionParams hold all of the parameters used for creating, updating,
// and canceling Subscriptions.
type SubscriptionParams struct {
	Plan                  string  `stripe_field:"plan"`
	Coupon                string  `stripe_field:"coupon"`
	DisableProrate        bool    `stripe_field:"prorate" opposite:"true"`
	TrialEnd              int     `stripe_field:"trial_end"`
	Quantity              int     `stripe_field:"quantity"`
	ApplicationFeePercent float64 `stripe_field:"application_fee_percent"`
	AtPeriodEnd           bool    `stripe_field:"at_period_end"`
	*CardParams
}

// TokenParams hold all of the parameters used for creating Tokens.
type TokenParams struct {
	Customer string `stripe_field:"customer"`
	*BankAccountParams
	*CardParams
}

// TransferParams hold all of the parameters used for creating and updating
// Transfers.
type TransferParams struct {
	Amount              int    `stripe_field:"amount"`
	Currency            string `stripe_field:"currency"`
	Recipient           string `stripe_field:"recipient"`
	Description         string `stripe_field:"description"`
	StatementDescription string `stripe_field:"statement_description"`
	Metadata
}
