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
	Id               string
	Duration         string
	AmountOff        int
	Currency         string
	DurationInMonths int
	MaxRedemptions   int
	PercentOff       int
	RedeemBy         int
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
	Customer       string
	ApplicationFee int
	Closed         bool
}

// InvoiceItemParams hold all of the parameters used for creating and updating
// InvoiceItems.
type InvoiceItemParams struct {
	Customer    string
	Amount      int
	Currency    string
	Invoice     string
	Description string
	Metadata
}

// PlanParams hold all of the parameters used for creating and updating Plans.
type PlanParams struct {
	Id              string
	Amount          int
	Currency        string
	Interval        string
	IntervalCount   int
	Name            string
	TrialPeriodDays int
	Metadata
}

// RecipientParams hold all of the parameters used for creating and updating
// Recipients.
type RecipientParams struct {
	Name        string
	Type        string
	TaxId       string
	Email       string
	Description string
	*BankAccountParams
	Metadata
}

// RefundParams hold all of the parameters used for refunding Charges.
type RefundParams struct {
	Amount               int
	RefundApplicationFee bool
}

// SubscriptionParams hold all of the parameters used for updating and
// canceling Subscriptions.
type SubscriptionParams struct {
	Plan                  string
	Coupon                string
	DisableProrate        bool
	TrialEnd              int
	Quantity              int
	ApplicationFeePercent float64
	AtPeriodEnd           bool
	*CardParams
}

// TokenParams hold all of the parameters used for creating Tokens.
type TokenParams struct {
	Customer string
	*BankAccountParams
	*CardParams
}

// TransferParams hold all of the parameters used for creating and updating
// Transfers.
type TransferParams struct {
	Amount              int
	Currency            string
	Recipient           string
	Description         string
	StatementDescriptor string
	Metadata
}
