package request

type MessageRequest struct {
	SenderUserId   string `json:"SenderUserId"`
	ReceiverUserId string `json:"ReceiverUserId"`
	Body           string `json:"Body"`
}
