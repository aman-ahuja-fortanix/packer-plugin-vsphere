// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package vsphere

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName     *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType   *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion   *string           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug         *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce         *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError       *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars      map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Cluster             *string           `mapstructure:"cluster" cty:"cluster" hcl:"cluster"`
	Datacenter          *string           `mapstructure:"datacenter" cty:"datacenter" hcl:"datacenter"`
	Datastore           *string           `mapstructure:"datastore" cty:"datastore" hcl:"datastore"`
	DiskMode            *string           `mapstructure:"disk_mode" cty:"disk_mode" hcl:"disk_mode"`
	Host                *string           `mapstructure:"host" cty:"host" hcl:"host"`
	ESXiHost            *string           `mapstructure:"esxi_host" cty:"esxi_host" hcl:"esxi_host"`
	Insecure            *bool             `mapstructure:"insecure" cty:"insecure" hcl:"insecure"`
	Options             []string          `mapstructure:"options" cty:"options" hcl:"options"`
	Overwrite           *bool             `mapstructure:"overwrite" cty:"overwrite" hcl:"overwrite"`
	Password            *string           `mapstructure:"password" cty:"password" hcl:"password"`
	ResourcePool        *string           `mapstructure:"resource_pool" cty:"resource_pool" hcl:"resource_pool"`
	Username            *string           `mapstructure:"username" cty:"username" hcl:"username"`
	VMFolder            *string           `mapstructure:"vm_folder" cty:"vm_folder" hcl:"vm_folder"`
	VMName              *string           `mapstructure:"vm_name" cty:"vm_name" hcl:"vm_name"`
	VMNetwork           *string           `mapstructure:"vm_network" cty:"vm_network" hcl:"vm_network"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"cluster":                    &hcldec.AttrSpec{Name: "cluster", Type: cty.String, Required: false},
		"datacenter":                 &hcldec.AttrSpec{Name: "datacenter", Type: cty.String, Required: false},
		"datastore":                  &hcldec.AttrSpec{Name: "datastore", Type: cty.String, Required: false},
		"disk_mode":                  &hcldec.AttrSpec{Name: "disk_mode", Type: cty.String, Required: false},
		"host":                       &hcldec.AttrSpec{Name: "host", Type: cty.String, Required: false},
		"esxi_host":                  &hcldec.AttrSpec{Name: "esxi_host", Type: cty.String, Required: false},
		"insecure":                   &hcldec.AttrSpec{Name: "insecure", Type: cty.Bool, Required: false},
		"options":                    &hcldec.AttrSpec{Name: "options", Type: cty.List(cty.String), Required: false},
		"overwrite":                  &hcldec.AttrSpec{Name: "overwrite", Type: cty.Bool, Required: false},
		"password":                   &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"resource_pool":              &hcldec.AttrSpec{Name: "resource_pool", Type: cty.String, Required: false},
		"username":                   &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"vm_folder":                  &hcldec.AttrSpec{Name: "vm_folder", Type: cty.String, Required: false},
		"vm_name":                    &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"vm_network":                 &hcldec.AttrSpec{Name: "vm_network", Type: cty.String, Required: false},
	}
	return s
}
