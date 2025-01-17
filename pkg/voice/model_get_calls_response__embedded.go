/*
 * Voice API
 *
 * The Voice API lets you create outbound calls, control in-progress calls and get information about historical calls. More information about the Voice API can be found at <https://developer.nexmo.com/voice/voice-api/overview>.
 *
 * API version: 1.3.2
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package voice

// GetCallsResponseEmbedded A list of call objects. See the [get details of a specific call](#getCall) response fields for a description of the nested objects
type GetCallsResponseEmbedded struct {
	Calls []GetCallResponse `json:"calls,omitempty"`
}
