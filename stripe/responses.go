package stripe

// ListResponse is a part of what is returned from the Stripe API after a GET
// for a collection.
type ListResponse struct {
	Object string `json:"object"`
	Url    string `json:"url"`
	Count  int    `json:"count"`
}

// DeleteResponse is what is returned from the Stripe API after a DELETE.
type DeleteResponse struct {
	Id      string `json:"id"`
	Deleted bool   `json:"deleted"`
}

// ErrorResponse is what is returned from the Stripe API after an error.
type ErrorResponse struct {
	Err struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Code    string `json:"code,omitempty"`
		Param   string `json:"param,omitempty"`
	} `json:"error"`
}

// ErrorResponse must implement an Error() method to satisfy the error interface.
func (e *ErrorResponse) Error() string {
	// TODO: Do more than just return the Message?
	return e.Err.Message
}
