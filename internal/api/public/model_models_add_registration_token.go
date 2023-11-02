/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ModelsAddRegistrationToken struct for ModelsAddRegistrationToken
type ModelsAddRegistrationToken struct {
	Description string `json:"description,omitempty"`
	// Expiration is optional, if set the registration token is only valid until the Expiration time.
	Expiration string `json:"expiration,omitempty"`
	// SingleUse only allows the registration token to be used once.
	SingleUse bool   `json:"single_use,omitempty"`
	VpcId     string `json:"vpc_id,omitempty"`
}
