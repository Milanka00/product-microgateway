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
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// PostOauth2TokenURL generates an URL for the post oauth2 token operation
type PostOauth2TokenURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostOauth2TokenURL) WithBasePath(bp string) *PostOauth2TokenURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PostOauth2TokenURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PostOauth2TokenURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/oauth2/token"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/mgw/adapter/0.1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PostOauth2TokenURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PostOauth2TokenURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PostOauth2TokenURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PostOauth2TokenURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PostOauth2TokenURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *PostOauth2TokenURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
