package stripe

type Discount struct {
	Object   string  `json:"object"`
	Coupon   *Coupon `json:"coupon"`
	Customer string  `json:"customer"`
	Start    int64   `json:"start"`
	End      int64   `json:"end"`
}

type DiscountClient struct{}

// Delete deletes a customers discount.
//
// For more information: https://stripe.com/docs/api#delete_discount
func (c *DiscountClient) Delete(customerId string) (*DeleteResponse, error) {
	response := DeleteResponse{}
	err := delete("/customers/"+customerId+"/discount", nil, &response)
	return &response, err
}
