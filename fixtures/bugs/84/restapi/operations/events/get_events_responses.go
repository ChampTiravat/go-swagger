package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"

	"github.com/go-swagger/go-swagger/fixtures/bugs/84/models"
)

/*GetEventsOK Successful response

swagger:response getEventsOK
*/
type GetEventsOK struct {

	// In: body
	Payload []*models.Event `json:"body,omitempty"`
}

// WriteResponse to the client
func (o *GetEventsOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetEventsDefault Generic Error

swagger:response getEventsDefault
*/
type GetEventsDefault struct {
}

// WriteResponse to the client
func (o *GetEventsDefault) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(500)
}