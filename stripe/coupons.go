package stripe

import (
	"net/url"
)

type Coupon struct {
	Id               string `json:"id"`
	Object           string `json:"object"`
	Livemode         bool   `json:"livemode"`
	Duration         string `json:"duration"`
	AmountOff        int64  `json:"amount_off"`
	Currency         string `json:"currency"`
	DurationInMonths int64  `json:"duration_in_months"`
	MaxRedemptions   int64  `json:"max_redemptions"`
	PercentOff       int64  `json:"percent_off"`
	RedeemBy         int64  `json:"redeem_by"`
	TimesRedeemed    int64  `json:"times_redeemed"`
	Valid            bool   `json:"valid"`
}

type CouponListResponse struct {
	Object string    `json:"object"`
	Url    string    `json:"url"`
	Count  int       `json:"count"`
	Data   []*Coupon `json:"data"`
}

// Delete deletes a coupon.
//
// For more information: https://stripe.com/docs/api#delete_coupon
func (c *Coupon) Delete() (*DeleteResponse, error) {
	response := DeleteResponse{}
	err := delete("/coupons/"+c.Id, nil, &response)
	return &response, err
}

type CouponClient struct{}

// Create creates a coupon.
//
// For more information: https://stripe.com/docs/api#create_coupon
func (c *CouponClient) Create(params *CouponParams) (*Coupon, error) {
	coupon := Coupon{}
	values := url.Values{}
	parseCouponParams(params, &values)
	err := post("/coupons", values, &coupon)
	return &coupon, err
}

// Retrieve loads a coupon.
//
// For more information: https://stripe.com/docs/api#retrieve_coupon
func (c *CouponClient) Retrieve(id string) (*Coupon, error) {
	coupon := Coupon{}
	err := get("/coupons/"+id, nil, &coupon)
	return &coupon, err
}

// Delete deletes a coupon.
//
// For more information: https://stripe.com/docs/api#delete_coupon
func (c *CouponClient) Delete(id string) (*DeleteResponse, error) {
	response := DeleteResponse{}
	err := delete("/coupons/"+id, nil, &response)
	return &response, err
}

// All lists the first 10 coupons. It calls AllWithFilters with a blank Filters
// so all defaults are used.
//
// For more information: https://stripe.com/docs/api#list_coupons
func (c *CouponClient) All() (*CouponListResponse, error) {
	return c.AllWithFilters(Filters{})
}

// AllWithFilters takes a Filters and applies all valid filters for the action.
//
// For more information: https://stripe.com/docs/api#list_coupons
func (c *CouponClient) AllWithFilters(filters Filters) (*CouponListResponse, error) {
	response := CouponListResponse{}
  values := url.Values{}
  addFiltersToValues([]string{"count", "offset"}, filters, &values)
	err := get("/coupons", values, &response)
	return &response, err
}

// parseCouponParams takes a pointer to a CouponParams and a pointer to a url.Values,
// it iterates over everything in the CouponParams struct and Adds what is there
// to the url.Values.
func parseCouponParams(params *CouponParams, values *url.Values) {
	addParamsToValues(params, values)
}
