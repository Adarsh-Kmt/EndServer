package types

type MessageRequest struct {
	ReceiverUserId string `json:"ReceiverUserId"`
	Body           string `json:"Body"`
}

type UserRegisterRequest struct {
	UserId   string `json:"UserId"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	UserId   string `json:"UserId"`
	Password string `json:"password"`
}
