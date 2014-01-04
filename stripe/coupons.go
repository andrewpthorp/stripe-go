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

// List lists the first 10 coupons. It calls ListCount with 10 as
// the count and 0 as the offset, which are the defaults in the Stripe API.
//
// For more information: https://stripe.com/docs/api#list_coupons
func (c *CouponClient) List() (*CouponListResponse, error) {
	return c.ListCount(10, 0)
}

// ListCount lists `count` coupons starting at `offset`.
//
// For more information: https://stripe.com/docs/api#list_coupons
func (c *CouponClient) ListCount(count, offset int) (*CouponListResponse, error) {
	response := CouponListResponse{}

	params := url.Values{
		"count":  {strconv.Itoa(count)},
		"offset": {strconv.Itoa(offset)},
	}

	err := get("/coupons", params, &response)
	return &response, err
}

// parseCouponParams takes a pointer to a CouponParams and a pointer to a url.Values,
// it iterates over everything in the CouponParams struct and Adds what is there
// to the url.Values.
func parseCouponParams(params *CouponParams, values *url.Values) {
	addParamsToValues(params, values)
}
