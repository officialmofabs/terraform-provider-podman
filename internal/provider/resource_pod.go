package provider

import (
	"context"
	"fmt"

	"github.com/containers/podman/v4/pkg/domain/entities"
	"github.com/containers/podman/v4/pkg/specgen"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/project0/terraform-provider-podman/internal/modifier"
	"github.com/project0/terraform-provider-podman/internal/provider/shared"
	"github.com/project0/terraform-provider-podman/internal/utils"
)

type (
	podResource struct {
		genericResource
	}
	podResourceData struct {
		ID     types.String `tfsdk:"id"`
		Name   types.String `tfsdk:"name"`
		Labels types.Map    `tfsdk:"labels"`

		CgroupParent types.String `tfsdk:"cgroup_parent"`
		Hostname     types.String `tfsdk:"hostname"`

		Mounts shared.Mounts `tfsdk:"mounts"`
	}
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &podResource{}
	_ resource.ResourceWithConfigure   = &podResource{}
	_ resource.ResourceWithImportState = &podResource{}
)

// NewPodResource creates a new pod resource.
func NewPodResource() resource.Resource {
	return &podResource{}
}

// Configure adds the provider configured client to the resource.
func (r *podResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	r.genericResource.Configure(ctx, req, resp)
}

// Metadata returns the resource type name.
func (r podResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pod"
}

// GetSchema returns the resource schema.
func (r podResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	mountsAttr := make(shared.Mounts, 0)
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		Description: "Manage pods for containers",
		Attributes: withGenericAttributes(map[string]tfsdk.Attribute{
			"cgroup_parent": {
				MarkdownDescription: "Path to cgroups under which the cgroup for the pod will be created. " +
					"If the path is not absolute, the path is considered to be relative to the cgroups path of the init process. " +
					"Cgroups will be created if they do not already exist.",
				Required: false,
				Optional: true,
				Computed: true,
				Type:     types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
					modifier.RequiresReplaceComputed(),
				},
			},

			"hostname": {
				Description: "Hostname is the pod's hostname. " +
					"If not set, the name of the pod will be used (if a name was not provided here, the name auto-generated for the pod will be used). " +
					"This will be used by the infra container and all containers in the pod as long as the UTS namespace is shared.",
				Required: false,
				Optional: true,
				Computed: false,
				Type:     types.StringType,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
					resource.RequiresReplace(),
				},
			},

			"mounts": mountsAttr.AttributeSchema(),

			// infra container settings
			/*			"infra": {
							Description: "Infra tells the pod to create an infra container or not. " +
								"If this is false, many networking-related options will become unavailabl. Defaults to true.",
							Required: false,
							Optional: true,
							Computed: true,
							Type:     types.BoolType,
							PlanModifiers: tfsdk.AttributePlanModifiers{
								resource.UseStateForUnknown(),
								resource.RequiresReplace(),
							},
						},

						"infra_command":         {},
						"infra_conmon_pid_file": {},
						"infra_image":           {},
						"infra_name":            {},
			*/
		}),
	}, nil
}

func toPodmanPodSpecGenerator(ctx context.Context, d podResourceData, diags *diag.Diagnostics) *specgen.PodSpecGenerator {
	s := specgen.NewPodSpecGenerator()
	p := &entities.PodCreateOptions{
		Name:         d.Name.ValueString(),
		CgroupParent: d.CgroupParent.ValueString(),
		Hostname:     d.Hostname.ValueString(),
		Infra:        true,
	}

	diags.Append(d.Labels.ElementsAs(ctx, &p.Labels, true)...)
	sp, err := entities.ToPodSpecGen(*s, p)
	if err != nil {
		diags.AddError("Invalid pod configuration", fmt.Sprintf("Cannot build pod configuration: %q", err.Error()))
	}
	// add storage
	sp.Volumes, sp.Mounts = d.Mounts.ToPodmanSpec(diags)
	if err := sp.Validate(); err != nil {
		diags.AddError("Invalid pod configuration", fmt.Sprintf("Cannot build pod configuration: %q", err.Error()))
	}
	return sp
}

func fromPodResponse(p *entities.PodInspectReport, diags *diag.Diagnostics) *podResourceData {

	d := &podResourceData{
		ID:           types.String{Value: p.ID},
		Name:         types.String{Value: p.Name},
		Labels:       utils.MapStringToMapType(p.Labels),
		Mounts:       shared.FromPodmanToMounts(diags, p.Mounts),
		CgroupParent: types.String{Value: p.CgroupParent},
		Hostname:     types.String{Value: p.Hostname, Null: p.Hostname == ""},
	}

	return d
}
