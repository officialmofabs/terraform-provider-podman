package provider

import (
	"context"
	"fmt"
	"net"

	ntypes "github.com/containers/common/libnetwork/types"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/project0/terraform-provider-podman/internal/modifier"
	"github.com/project0/terraform-provider-podman/internal/utils"
	"github.com/project0/terraform-provider-podman/internal/validator"
)

type (
	networkResource struct {
		genericResource
	}

	networkResourceData struct {
		ID     types.String `tfsdk:"id"`
		Name   types.String `tfsdk:"name"`
		Labels types.Map    `tfsdk:"labels"`

		DNS      types.Bool `tfsdk:"dns"`
		IPv6     types.Bool `tfsdk:"ipv6"`
		Internal types.Bool `tfsdk:"internal"`

		Driver     types.String `tfsdk:"driver"`
		IPAMDriver types.String `tfsdk:"ipam_driver"`
		Options    types.Map    `tfsdk:"options"`

		Subnets []networkResourceSubnetData `tfsdk:"subnets"`
	}

	networkResourceSubnetData struct {
		Subnet  types.String `tfsdk:"subnet"`
		Gateway types.String `tfsdk:"gateway"`
	}
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &networkResource{}
	_ resource.ResourceWithConfigure   = &networkResource{}
	_ resource.ResourceWithModifyPlan  = &networkResource{}
	_ resource.ResourceWithImportState = &networkResource{}
)

// NewNetworkResource creates a new network resource
func NewNetworkResource() resource.Resource {
	return &networkResource{}
}

// Configure adds the provider configured client to the data source.
func (r *networkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.genericResource.Configure(ctx, req, resp)
}

// Configure adds the provider configured client to the data source.
func (r *networkResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	r.genericResource.ModifyPlan(ctx, req, resp)
}

// Metadata returns the data source type name.
func (r *networkResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_network"
}

func (r *networkResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		Description: "Manage networks for containers and pods",
		Attributes: withGenericAttributes(map[string]tfsdk.Attribute{
			"dns": {
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "Enable the DNS plugin for this network which if enabled, can perform container to container name resolution. Defaults to `false`.",
				Type:                types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
					modifier.RequiresReplaceComputed(),
				},
			},

			"ipv6": {
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "Enable IPv6 (Dual Stack) networking. If no subnets are given it will allocate a ipv4 and ipv6 subnet. Defaults to `false`.",
				Type:                types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
					modifier.RequiresReplaceComputed(),
				},
			},

			"internal": {
				Computed:            true,
				Optional:            true,
				MarkdownDescription: "Internal is whether the Network should not have external routes to public or other Networks. Defaults to `false`.",
				Type:                types.BoolType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
					modifier.RequiresReplaceComputed(),
				},
			},

			"driver": {
				Computed: true,
				Optional: true,
				MarkdownDescription: fmt.Sprintf(
					"Driver to manage the network. One of `%s`, `%s`, `%s` are currently supported. By podman defaults to `bridge`.",
					ntypes.BridgeNetworkDriver,
					ntypes.MacVLANNetworkDriver,
					ntypes.IPVLANNetworkDriver,
				),
				Type: types.StringType,
				Validators: []tfsdk.AttributeValidator{
					validator.OneOf([]string{
						ntypes.BridgeNetworkDriver,
						ntypes.MacVLANNetworkDriver,
						ntypes.IPVLANNetworkDriver,
					},
					)},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
					modifier.RequiresReplaceComputed(),
				},
			},

			"ipam_driver": {
				Computed: true,
				Optional: true, // TODO
				MarkdownDescription: fmt.Sprintf(
					"Set the ipam driver (IP Address Management Driver) for the network. Valid values are `%s`, `%s`, `%s`. When unset podman will choose an ipam driver automatically based on the network driver.",
					ntypes.HostLocalIPAMDriver,
					ntypes.DHCPIPAMDriver,
					"none",
				),
				Type: types.StringType,
				Validators: []tfsdk.AttributeValidator{
					validator.OneOf([]string{
						ntypes.HostLocalIPAMDriver,
						ntypes.DHCPIPAMDriver,
						"none",
					},
					)},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
					modifier.RequiresReplaceComputed(),
				},
			},

			"options": {
				Description: "Driver specific options.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Type: types.MapType{
					ElemType: types.StringType,
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
					modifier.UseDefaultModifier(types.Map{ElemType: types.StringType, Null: false}),
					modifier.RequiresReplaceComputed(),
				},
			},

			"subnets": {
				Description: "Subnets for this network.",
				Required:    false,
				Optional:    true,
				Computed:    true,
				Attributes: tfsdk.SetNestedAttributes(
					map[string]tfsdk.Attribute{
						"subnet": {
							Required:            true,
							Optional:            false,
							MarkdownDescription: "The subnet in CIDR notation.",
							Type:                types.StringType,
							Validators:          []tfsdk.AttributeValidator{validator.IsCIDR()},
							PlanModifiers: tfsdk.AttributePlanModifiers{
								resource.UseStateForUnknown(),
								modifier.RequiresReplaceComputed(),
							},
						},
						"gateway": {
							Computed:            true,
							Optional:            true,
							MarkdownDescription: "Gateway IP for this Network.",
							Type:                types.StringType,
							Validators:          []tfsdk.AttributeValidator{validator.IsIpAdress()},
							PlanModifiers: tfsdk.AttributePlanModifiers{
								resource.UseStateForUnknown(),
								modifier.RequiresReplaceComputed(),
							},
						},
					},
				),
			},
		}),
	}, nil
}

// toPodmanNetwork converts a resource data to a podman network
func toPodmanNetwork(ctx context.Context, d networkResourceData, diags *diag.Diagnostics) *ntypes.Network {
	var nw = &ntypes.Network{
		Name:        d.Name.Value,
		Driver:      d.Driver.Value,
		IPv6Enabled: d.IPv6.Value,
		DNSEnabled:  d.DNS.Value,
		Internal:    d.Internal.Value,
	}

	// Convert map types
	diags.Append(d.Labels.ElementsAs(ctx, &nw.Labels, true)...)
	diags.Append(d.Options.ElementsAs(ctx, &nw.Options, true)...)

	if !d.IPAMDriver.Null {
		ipam := map[string]string{
			"driver": d.IPAMDriver.Value,
		}
		nw.IPAMOptions = ipam
	}

	// subnet
	for _, s := range d.Subnets {
		_, ipNet, err := net.ParseCIDR(s.Subnet.Value)
		if err != nil {
			diags.AddError("Cannot parse subnet CIDR", err.Error())
			continue
		}
		subnet := &ntypes.Subnet{
			Subnet: ntypes.IPNet{
				IPNet: *ipNet,
			},
		}
		if !s.Gateway.Null {
			subnet.Gateway = net.ParseIP(s.Gateway.Value)
		}
		nw.Subnets = append(nw.Subnets, *subnet)
	}

	return nw
}

// fromNetwork converts a podman network to a resource data
func fromPodmanNetwork(n ntypes.Network, diags *diag.Diagnostics) *networkResourceData {
	d := &networkResourceData{
		ID:       types.String{Value: n.ID},
		Name:     types.String{Value: n.Name},
		DNS:      types.Bool{Value: n.DNSEnabled},
		IPv6:     types.Bool{Value: n.IPv6Enabled},
		Internal: types.Bool{Value: n.Internal},
		Driver:   types.String{Value: n.Driver},
		Labels:   utils.MapStringToMapType(n.Labels),
		Options:  utils.MapStringToMapType(n.Options),
	}

	d.IPAMDriver = utils.MapStringValueToStringType(n.IPAMOptions, "driver")

	for _, s := range n.Subnets {
		subnet := networkResourceSubnetData{
			Subnet:  types.String{Value: s.Subnet.String()},
			Gateway: types.String{Value: s.Gateway.String()},
		}
		d.Subnets = append(d.Subnets, subnet)
	}
	return d
}
