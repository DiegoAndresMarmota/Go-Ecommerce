package response

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"e-commerce/model"
)

//const Code del struct Response
const (
	BindFailed      = "bind_failed"
	Ok              = "ok"
	RecordCreated   = "record_created"
	RecordUpdated   = "record_updated"
	RecordDeleted   = "record_deleted"
	UnexpectedError = "unexpected_error"
	AuthError       = "authorization_error"
)

type API struct{}

//API
func New() API {
	return API{}
}

//Response --> Cuando la API reciba una Data, y este OK, se responde con StatusOk del type Response
func (a API) OK(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: Ok, Message: "¡listo!"}},
	}
}


//Response --> Cuando la API reciba una petición para crear(POST), y este RecordCreated, se responde con un StatusCreated del type Response
func (a API) Created(data interface{}) (int, model.MessageResponse) {
	return http.StatusCreated, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordCreated, Message: "¡listo!"}},
	}
}


//Response --> Cuando la API reciba una petición para actualizar(PUT), y este RecordUpdated, se responde con StatusOk del type Response
func (a API) Updated(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordUpdated, Message: "¡listo!"}},
	}
}


//Response --> Cuando la API reciba una petición para eliminar(DELETE), y este RecordDeleted, se responde con StatusOk del type Response
func (a API) Deleted(data interface{}) (int, model.MessageResponse) {
	return http.StatusOK, model.MessageResponse{
		Data:     data,
		Messages: model.Responses{{Code: RecordDeleted, Message: "¡listo!"}},
	}
}


//Response --> Cuando la API recibe un error del JSON recibido.
func (a API) BindFailed(err error) error {
	e := model.NewError()
	e.Err = err
	e.Code = BindFailed
	e.StatusHTTP = http.StatusBadRequest
	e.Who = "c.Bind()"

	log.Warnf("%s", e.Error())
	return &e
}


//Response --> Cuando la API recibe un error interno.
func (a API) Error(c echo.Context, who string, err error) *model.Error {
	e := model.NewError()
	e.Err = err
	e.APIMessage = "Error interno del servidor!"
	e.Code = UnexpectedError
	e.StatusHTTP = http.StatusInternalServerError
	e.Who = who

	userID, ok := c.Get("userID").(uuid.UUID)

	if !ok {
		log.Errorf("cannot get/parse uuid from userID")
	}
	e.UserID = userID.String()

	log.Errorf("%s", e.Error())
	return &e
}
