package user

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json: "message"`
}
