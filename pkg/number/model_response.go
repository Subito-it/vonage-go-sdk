/*
 * Numbers API
 *
 * The Numbers API enables you to manage your existing numbers and buy new virtual numbers for use with Nexmo's APIs. Further information is here: <https://developer.nexmo.com/numbers/overview>
 *
 * API version: 1.0.18
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package number

// Response struct for Response
type Response struct {
	// The status code of the response. `200` indicates a successful request.
	ErrorCode string `json:"error-code,omitempty"`
	// The status code description
	ErrorCodeLabel string `json:"error-code-label,omitempty"`
}
