package types

type MessageRequest struct {
	ReceiverUsername string `json:"ReceiverUsername"`
	Body             string `json:"Body"`
}

type UserRegisterRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type UserLoginRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}
