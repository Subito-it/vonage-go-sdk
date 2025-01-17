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

// UpdateCallRequestTransferNccoDestination struct for UpdateCallRequestTransferNccoDestination
type UpdateCallRequestTransferNccoDestination struct {
	// Always `ncco`
	Type string `json:"type"`
	// Refer to the [NCCO Guide](https://developer.nexmo.com/voice/voice-api/ncco-reference) for a description of possible NCCO parameters.
	Ncco []map[string]interface{} `json:"ncco"`
}
