package stripe

import (
	"net/url"
	"strconv"
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

// Delete deletes a coupon.
//
// For more information: https://stripe.com/docs/api#delete_coupon
func (c *Coupon) Delete() (*DeleteResponse, error) {
	response := DeleteResponse{}
	err := delete("/coupons/"+c.Id, nil, &response)
	return &response, err
}

// The CouponClient is the receiver for most standard coupon related endpoints.
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

// List lists the first 10 coupons. It calls ListCount with 10 as
// the count and 0 as the offset, which are the defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#list_coupons
func (c *CouponClient) List() ([]*Coupon, error) {
	return c.ListCount(10, 0)
}

// ListCount lists `count` coupons starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_coupons
func (c *CouponClient) ListCount(count, offset int) ([]*Coupon, error) {
	type coupons struct{ Data []*Coupon }
	list := coupons{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/coupons", params, &list)
	return list.Data, err
}

// parseCouponParams takes a pointer to a CouponParams and a pointer to a url.Values,
// it iterates over everything in the CouponParams struct and Adds what is there
// to the url.Values.
func parseCouponParams(params *CouponParams, values *url.Values) {
	if params.Id != "" {
		values.Add("id", params.Id)
	}

	if params.Duration != "" {
		values.Add("duration", params.Duration)
	}

	if params.AmountOff != 0 {
		values.Add("amount_off", strconv.Itoa(params.AmountOff))
	}

	if params.Currency != "" {
		values.Add("currency", params.Currency)
	}

	if params.DurationInMonths != 0 {
		values.Add("duration_in_months", strconv.Itoa(params.DurationInMonths))
	}

	if params.MaxRedemptions != 0 {
		values.Add("max_redemptions", strconv.Itoa(params.MaxRedemptions))
	}

	if params.PercentOff != 0 {
		values.Add("percent_off", strconv.Itoa(params.PercentOff))
	}

	if params.RedeemBy != 0 {
		values.Add("redeem_by", strconv.Itoa(params.RedeemBy))
	}
}
