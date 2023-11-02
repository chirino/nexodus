/*
Nexodus API

This is the Nexodus API Server.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package public

// ModelsAddDevice struct for ModelsAddDevice
type ModelsAddDevice struct {
	AdvertiseCidrs  []string         `json:"advertise_cidrs,omitempty"`
	Discovery       bool             `json:"discovery,omitempty"`
	Endpoints       []ModelsEndpoint `json:"endpoints,omitempty"`
	Hostname        string           `json:"hostname,omitempty"`
	Os              string           `json:"os,omitempty"`
	PublicKey       string           `json:"public_key,omitempty"`
	Relay           bool             `json:"relay,omitempty"`
	SecurityGroupId string           `json:"security_group_id,omitempty"`
	SymmetricNat    bool             `json:"symmetric_nat,omitempty"`
	TunnelIp        string           `json:"tunnel_ip,omitempty"`
	TunnelIpV6      string           `json:"tunnel_ip_v6,omitempty"`
	VpcId           string           `json:"vpc_id,omitempty"`
}
