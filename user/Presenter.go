package user

// Response for every returned data
type Response struct {
	Data interface{} `json:"data"`
}

// ErrorResponse for every returned error data
type ErrorResponse struct {
	Message string `json: "message"`
}
