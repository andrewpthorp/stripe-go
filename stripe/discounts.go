package stripe

// The DiscountClient is the receiver for most standard discount related
// endpoints.
type DiscountClient struct{}

// Delete deletes a customers discount.
//
// For more information: https://stripe.com/docs/api#delete_discount
func (c *DiscountClient) Delete(customerId string) (*DeleteResponse, error) {
  response := DeleteResponse{}
  err := delete("/customers/" + customerId + "/discount", nil, &response)
  return &response, err
}
