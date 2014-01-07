package stripe

import "net/url"

type Dispute struct {
	Object             string `json:"object"`
	Livemode           bool   `json:"livemode"`
	Amount             int64  `json:"amount"`
	BalanceTransaction string `json:"balance_transaction"`
	Charge             string `json:"charge"`
	Currency           string `json:"currency"`
	Reason             string `json:"reason"`
	Status             string `json:"status"`
	Evidence           string `json:"evidence"`
	Created            int64  `json:"created"`
	EvidenceDueBy      int64  `json:"evidence_due_by"`
}

type DisputeClient struct {
	client Client
}

// Update updates a dispute.
//
// For more information: https://stripe.com/docs/api#update_dispute
func (c *DisputeClient) Update(chargeId, evidence string) (*Dispute, error) {
	dispute := Dispute{}
	values := url.Values{"evidence": {evidence}}
	err := c.client.post("/charges/"+chargeId+"/dispute", values, &dispute)
	return &dispute, err
}

// Close closes a dispute.
//
// For more information: https://stripe.com/docs/api#close_dispute
func (c *DisputeClient) Close(chargeId string) (*Dispute, error) {
	dispute := Dispute{}
	err := c.client.post("/charges/"+chargeId+"/dispute/close", nil, &dispute)
	return &dispute, err
}
