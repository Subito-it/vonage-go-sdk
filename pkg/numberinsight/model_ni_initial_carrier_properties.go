/*
 * Number Insight API
 *
 * Nexmo's Number Insight API delivers real-time intelligence about the validity, reachability and roaming status of a phone number and tells you how to format the number correctly in your application. There are three levels of Number Insight API available: [Basic, Standard and Advanced](https://developer.nexmo.com/number-insight/overview#basic-standard-and-advanced-apis). The advanced API is available asynchronously as well as synchronously.
 *
 * API version: 1.0.7
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package numberinsight

// NiInitialCarrierProperties Information about the network `number` is currently connected to.
type NiInitialCarrierProperties struct {
	// The [https://en.wikipedia.org/wiki/Mobile_country_code](https://en.wikipedia.org/wiki/Mobile_country_code) for the carrier`number` is associated with. Unreal numbers are marked as`unknown` and the request is rejected altogether if the number is impossible according to the [E.164](https://en.wikipedia.org/wiki/E.164) guidelines.
	NetworkCode string `json:"network_code,omitempty"`
	// The full name of the carrier that `number` is associated with.
	Name string `json:"name,omitempty"`
	// The country that `number` is associated with. This is in ISO 3166-1 alpha-2   format.
	Country string `json:"country,omitempty"`
	// The type of network that `number` is associated with.
	NetworkType string `json:"network_type,omitempty"`
}
