package user

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Type    string `json: "type"`
	Message string `json: "message"`
}
