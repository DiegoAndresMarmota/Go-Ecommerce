package model

//Respuesta con codigo y mensaje especifico
type Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

// Response[]
type Responses []Response

//Respuesta general
type MessageResponse struct {
	Data interface{} `json:"data"`
	Erros Responses `json:"errors"`
	Messages Responses `json:"messages"`
}
