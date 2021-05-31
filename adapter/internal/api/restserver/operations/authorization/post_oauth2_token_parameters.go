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
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	"github.com/wso2/adapter/internal/api/models"
)

// NewPostOauth2TokenParams creates a new PostOauth2TokenParams object
//
// There are no default values defined in the spec.
func NewPostOauth2TokenParams() PostOauth2TokenParams {

	return PostOauth2TokenParams{}
}

// PostOauth2TokenParams contains all the bound params for the post oauth2 token operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostOauth2Token
type PostOauth2TokenParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Credentials of the microgateway REST API user
	  Required: true
	  In: body
	*/
	Credentials *models.Credentials
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostOauth2TokenParams() beforehand.
func (o *PostOauth2TokenParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Credentials
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("credentials", "body", ""))
			} else {
				res = append(res, errors.NewParseError("credentials", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Credentials = &body
			}
		}
	} else {
		res = append(res, errors.Required("credentials", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
