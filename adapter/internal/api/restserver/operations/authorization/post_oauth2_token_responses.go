// Code generated by go-swagger; DO NOT EDIT.

//
// This product is licensed by WSO2 Inc. under Apache License 2.0. The license
// can be downloaded from the following locations:
// 	http://www.apache.org/licenses/LICENSE-2.0.html
// 	http://www.apache.org/licenses/LICENSE-2.0.txt
//
// This product also contains software under different licenses. This table below
// all the contained libraries (jar files) and the license under which they are
// provided to you.
//
// At the bottom of this file is a table that shows what each license indicated
// below is and where the actual text of the license can be found.
//
//
// Dependency								License
// github.com/PuerkitoBio/purell						BSD 3-Clause "New" or "Revised" License
// github.com/PuerkitoBio/urlesc						BSD 3-Clause "New" or "Revised" License
// github.com/asaskevich/govalidator					MIT License
// github.com/census-instrumentation/opencensus-proto			Apache License 2.0
// github.com/cncf/udpa/go							Apache License 2.0
// github.com/decred/dcrd/dcrec/secp256k1					ISC License
// github.com/docker/go-units						Apache License 2.0
// github.com/envoyproxy/go-control-plane					Apache License 2.0
// github.com/envoyproxy/protoc-gen-validate				Apache License 2.0
// github.com/fsnotify/fsnotify						BSD 3-Clause "New" or "Revised" License
// github.com/getkin/kin-openapi						MIT License
// github.com/ghodss/yaml							MIT License
// github.com/go-openapi/analysis						Apache License 2.0
// github.com/go-openapi/errors						Apache License 2.0
// github.com/go-openapi/jsonpointer					Apache License 2.0
// github.com/go-openapi/jsonreference					Apache License 2.0
// github.com/go-openapi/loads						Apache License 2.0
// github.com/go-openapi/runtime						Apache License 2.0
// github.com/go-openapi/spec						Apache License 2.0
// github.com/go-openapi/strfmt						Apache License 2.0
// github.com/go-openapi/swag						Apache License 2.0
// github.com/go-openapi/validate						Apache License 2.0
// github.com/go-stack/stack						MIT License
// github.com/golang/protobuf						BSD 3-Clause "New" or "Revised" License
// github.com/google/uuid							BSD 3-Clause "New" or "Revised" License
// github.com/jessevdk/go-flags						BSD 3-Clause "New" or "Revised" License
// github.com/lestrrat-go/backoff						MIT License
// github.com/lestrrat-go/httpcc						MIT License
// github.com/lestrrat-go/iter						MIT License
// github.com/lestrrat-go/jwx						MIT License
// github.com/lestrrat-go/option						MIT License
// github.com/mailru/easyjson						MIT License
// github.com/mitchellh/mapstructure					MIT License
// github.com/pelletier/go-toml						MIT License
// github.com/pkg/errors							BSD 2-Clause "Simplified" License
// github.com/sirupsen/logrus						MIT License
// github.com/streadway/amqp						BSD 2-Clause "Simplified" License
// go.mongodb.org/mongo-driver						Apache License 2.0
// golang.org/x/crypto							BSD 3-Clause "New" or "Revised" License
// golang.org/x/net							BSD 3-Clause "New" or "Revised" License
// golang.org/x/sys							BSD 3-Clause "New" or "Revised" License
// golang.org/x/text							BSD 3-Clause "New" or "Revised" License
// google.golang.org/genproto						Apache License 2.0
// google.golang.org/grpc							Apache License 2.0
// google.golang.org/protobuf						BSD 3-Clause "New" or "Revised" License
// gopkg.in/natefinch/lumberjack.v2					MIT License
// gopkg.in/yaml.v2							Apache License 2.0
//
//
//
//
// The license types used by the above libraries and their information is given below:
//
// apache2        Apache License Version 2.0
//                http://www.apache.org/licenses/LICENSE-2.0.html
// mit            MIT License
//                http://www.opensource.org/licenses/mit-license.php
// bsd2           Berkeley License - 2
//                https://opensource.org/licenses/BSD-2-Clause
// bsd3           Berkeley License - 3
//                http://opensource.org/licenses/BSD-3-Clause
// isc	       Internet Systems Consortium
// 	       https://opensource.org/licenses/ISC
//
//

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/wso2/adapter/internal/api/models"
)

// PostOauth2TokenOKCode is the HTTP code returned for type PostOauth2TokenOK
const PostOauth2TokenOKCode int = 200

/*PostOauth2TokenOK Authentication successful.
Returns an access token.


swagger:response postOauth2TokenOK
*/
type PostOauth2TokenOK struct {

	/*
	  In: Body
	*/
	Payload *PostOauth2TokenOKBody `json:"body,omitempty"`
}

// NewPostOauth2TokenOK creates PostOauth2TokenOK with default headers values
func NewPostOauth2TokenOK() *PostOauth2TokenOK {

	return &PostOauth2TokenOK{}
}

// WithPayload adds the payload to the post oauth2 token o k response
func (o *PostOauth2TokenOK) WithPayload(payload *PostOauth2TokenOKBody) *PostOauth2TokenOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post oauth2 token o k response
func (o *PostOauth2TokenOK) SetPayload(payload *PostOauth2TokenOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOauth2TokenOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOauth2TokenUnauthorizedCode is the HTTP code returned for type PostOauth2TokenUnauthorized
const PostOauth2TokenUnauthorizedCode int = 401

/*PostOauth2TokenUnauthorized Unauthorized. Invalid authentication credentials.

swagger:response postOauth2TokenUnauthorized
*/
type PostOauth2TokenUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostOauth2TokenUnauthorized creates PostOauth2TokenUnauthorized with default headers values
func NewPostOauth2TokenUnauthorized() *PostOauth2TokenUnauthorized {

	return &PostOauth2TokenUnauthorized{}
}

// WithPayload adds the payload to the post oauth2 token unauthorized response
func (o *PostOauth2TokenUnauthorized) WithPayload(payload *models.Error) *PostOauth2TokenUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post oauth2 token unauthorized response
func (o *PostOauth2TokenUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOauth2TokenUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOauth2TokenInternalServerErrorCode is the HTTP code returned for type PostOauth2TokenInternalServerError
const PostOauth2TokenInternalServerErrorCode int = 500

/*PostOauth2TokenInternalServerError Internal Server Error.

swagger:response postOauth2TokenInternalServerError
*/
type PostOauth2TokenInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostOauth2TokenInternalServerError creates PostOauth2TokenInternalServerError with default headers values
func NewPostOauth2TokenInternalServerError() *PostOauth2TokenInternalServerError {

	return &PostOauth2TokenInternalServerError{}
}

// WithPayload adds the payload to the post oauth2 token internal server error response
func (o *PostOauth2TokenInternalServerError) WithPayload(payload *models.Error) *PostOauth2TokenInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post oauth2 token internal server error response
func (o *PostOauth2TokenInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOauth2TokenInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
