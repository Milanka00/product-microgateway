/*
 * Choreo-Connect Interceptor Service
 *
 * Interceptor Service
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ResponseHandlerRequestBody struct {
	ResponseCode int32 `json:"responseCode"`

	RequestHeaders *map[string]string `json:"requestHeaders,omitempty"`

	RequestTrailers *map[string]string `json:"requestTrailers,omitempty"`

	RequestBody string `json:"requestBody,omitempty"`

	ResponseHeaders *map[string]string `json:"responseHeaders,omitempty"`

	ResponseTrailers *map[string]string `json:"responseTrailers,omitempty"`

	ResponseBody string `json:"responseBody,omitempty"`

	InvocationContext *InvocationContext `json:"invocationContext,omitempty"`

	InterceptorContext *map[string]string `json:"interceptorContext,omitempty"`
}