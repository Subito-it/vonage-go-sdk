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

// CreateCallRequestAnswerUrlAllOf struct for CreateCallRequestAnswerUrlAllOf
type CreateCallRequestAnswerUrlAllOf struct {
	// The webhook endpoint where you provide the [Nexmo Call Control Object](/voice/voice-api/ncco-reference) that governs this call.
	AnswerUrl []string `json:"answer_url,omitempty"`
	// The HTTP method used to send event information to answer_url.
	AnswerMethod string `json:"answer_method,omitempty"`
}