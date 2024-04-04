// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package addrs

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPluginParseSourceString(t *testing.T) {
	tests := []struct {
		name      string
		str       string
		want      *Plugin
		wantDiags bool
	}{
		{"invalid: only one component, rejected", "potato", nil, true},
		{"valid: two components in name", "hashicorp/azr", &Plugin{"hashicorp/azr"}, false},
		{"valid: three components, nothing superfluous", "github.com/hashicorp/azr", &Plugin{"github.com/hashicorp/azr"}, false},
		{"valid: trailing slash, will be removed", "github.com/hashicorp/azr/", &Plugin{"github.com/hashicorp/azr"}, false},
		{"invalid: reject because scheme specified", "https://github.com/hashicorp/azr", nil, true},
		{"invalid: reject because query non nil", "github.com/hashicorp/azr?arg=1", nil, true},
		{"invalid: reject because fragment present", "github.com/hashicorp/azr#anchor", nil, true},
		{"valid: leading and trailing slashes are removed", "/github.com/hashicorp/azr/", &Plugin{"github.com/hashicorp/azr"}, false},
		{"valid: leading slashes are removed", "/github.com/hashicorp/azr", &Plugin{"github.com/hashicorp/azr"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotDiags := ParsePluginSourceString(tt.str)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePluginSourceString() got = %v, want %v", got, tt.want)
			}
			if tt.wantDiags && len(gotDiags) == 0 {
				t.Errorf("Expected diags, but got none")
			}
			if !tt.wantDiags && len(gotDiags) != 0 {
				t.Errorf("Unexpected diags: %s", gotDiags)
			}
		})
	}
}

func TestPluginName(t *testing.T) {
	tests := []struct {
		name         string
		pluginString string
		expectName   string
	}{
		{
			"valid minimal name",
			"github.com/hashicorp/amazon",
			"amazon",
		},
		{
			"valid name with prefix",
			"github.com/hashicorp/packer-plugin-amazon",
			"amazon",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plug := &Plugin{
				Source: tt.pluginString,
			}

			name := plug.Name()
			if name != tt.expectName {
				t.Errorf("Expected plugin %q to have %q as name, got %q", tt.pluginString, tt.expectName, name)
			}
		})
	}
}

func TestPluginParts(t *testing.T) {
	tests := []struct {
		name          string
		pluginSource  string
		expectedParts []string
	}{
		{
			"valid with two parts",
			"factiartory.com/packer",
			[]string{"factiartory.com", "packer"},
		},
		{
			"valid with four parts",
			"factiartory.com/hashicrop/fields/packer",
			[]string{"factiartory.com", "hashicrop", "fields", "packer"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plugin := &Plugin{tt.pluginSource}
			diff := cmp.Diff(plugin.Parts(), tt.expectedParts)
			if diff != "" {
				t.Errorf("Difference found between expected and computed parts: %s", diff)
			}
		})
	}
}
