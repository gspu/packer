// Code generated by "mapstructure-to-hcl2 -type WaitIpConfig"; DO NOT EDIT.

package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatWaitIpConfig is an auto-generated flat version of WaitIpConfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatWaitIpConfig struct {
	WaitTimeout   *string `mapstructure:"ip_wait_timeout" cty:"ip_wait_timeout" hcl:"ip_wait_timeout"`
	SettleTimeout *string `mapstructure:"ip_settle_timeout" cty:"ip_settle_timeout" hcl:"ip_settle_timeout"`
	WaitAddress   *string `mapstructure:"ip_wait_address" cty:"ip_wait_address" hcl:"ip_wait_address"`
}

// FlatMapstructure returns a new FlatWaitIpConfig.
// FlatWaitIpConfig is an auto-generated flat version of WaitIpConfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*WaitIpConfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatWaitIpConfig)
}

// HCL2Spec returns the hcl spec of a WaitIpConfig.
// This spec is used by HCL to read the fields of WaitIpConfig.
// The decoded values from this spec will then be applied to a FlatWaitIpConfig.
func (*FlatWaitIpConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"ip_wait_timeout":   &hcldec.AttrSpec{Name: "ip_wait_timeout", Type: cty.String, Required: false},
		"ip_settle_timeout": &hcldec.AttrSpec{Name: "ip_settle_timeout", Type: cty.String, Required: false},
		"ip_wait_address":   &hcldec.AttrSpec{Name: "ip_wait_address", Type: cty.String, Required: false},
	}
	return s
}