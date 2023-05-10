package product

import (
	"e-commerce/domain/product"
	"e-commerce/infrastructure/handler/response"
)

//
type handler struct {
	useCase product.UserCase
	response response.API
}

//
func newHandler(useCase product.UserCase) handler {
	return handler{useCase: useCase}
}

