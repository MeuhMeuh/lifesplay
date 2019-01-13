package communication

// OutboundPayload defines a payload that's goal is to be transmitted to the UI.
type OutboundPayload interface {
	SetBody(body interface{})
	SetError(err error)
}

// OutboundPayloadImpl defines a structure that is used to send data to the Electron app.
type OutboundPayloadImpl struct {
	Body  interface{} `json:"body"`
	Error string      `json:"error"`
}

// SetBody sets a body on the payload to be returned.
func (payload *OutboundPayloadImpl) SetBody(body interface{}) {
	payload.Body = body
}

// SetError sets an error on the payload to be returned.
func (payload *OutboundPayloadImpl) SetError(err error) {
	payload.Error = err.Error()
}
