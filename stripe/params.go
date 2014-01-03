package stripe

// BankAccountParams hold all of the parameters used for creating and updating
// BankAccounts.
type BankAccountParams struct {
	Country       string
	RoutingNumber string
	AccountNumber string
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
	Amount         int
	Currency       string
	Customer       string
	CardParams     *CardParams
	Description    string
	Capture        bool
	ApplicationFee int
	Metadata       map[string]string
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
	AccountBalance int
	CardParams     *CardParams
	Coupon         string
	Description    string
	Email          string
	Plan           string
	Quantity       int
	TrialEnd       int
	Metadata       map[string]string
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
	Metadata    map[string]string
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
	Metadata        map[string]string
}

// RecipientParams hold all of the parameters used for creating and updating
// Recipients.
type RecipientParams struct {
	Name              string
	Type              string
	TaxId             string
	BankAccountParams *BankAccountParams
	Email             string
	Description       string
	Metadata          map[string]string
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
	Prorate               bool
	TrialEnd              int
	CardParams            *CardParams
	Quantity              int
	ApplicationFeePercent float64
	AtPeriodEnd           bool
}

// TokenParams hold all of the parameters used for creating Tokens.
type TokenParams struct {
	BankAccountParams *BankAccountParams
	CardParams        *CardParams
	Customer          string
}

// TransferParams hold all of the parameters used for creating and updating
// Transfers.
type TransferParams struct {
	Amount              int
	Currency            string
	Recipient           string
	Description         string
	StatementDescriptor string
	Metadata            map[string]string
}
