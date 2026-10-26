package discoverycontroller

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/discoverycontroller"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type NIDiscoveryServiceV2Model struct {
	CreatedAt         timetypes.RFC3339 `tfsdk:"created_at"`
	HostId            types.String      `tfsdk:"host_id"`
	Id                types.String      `tfsdk:"id"`
	NiDiscoveryConfig types.Object      `tfsdk:"ni_discovery_config"`
	NiJobConfigId     types.String      `tfsdk:"ni_job_config_id"`
	NiJobConfigName   types.String      `tfsdk:"ni_job_config_name"`
	Ophid             types.String      `tfsdk:"ophid"`
	PoolId            types.String      `tfsdk:"pool_id"`
	ServiceId         types.String      `tfsdk:"service_id"`
	UpdatedAt         timetypes.RFC3339 `tfsdk:"updated_at"`
}

var NIDiscoveryServiceV2AttrTypes = map[string]attr.Type{
	"created_at":          timetypes.RFC3339Type{},
	"host_id":             types.StringType,
	"id":                  types.StringType,
	"ni_discovery_config": types.ObjectType{AttrTypes: NIDiscoveryConfigV2AttrTypes},
	"ni_job_config_id":    types.StringType,
	"ni_job_config_name":  types.StringType,
	"ophid":               types.StringType,
	"pool_id":             types.StringType,
	"service_id":          types.StringType,
	"updated_at":          timetypes.RFC3339Type{},
}

var NIDiscoveryServiceV2ResourceSchemaAttributes = map[string]schema.Attribute{
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Read-only. Timestamp created_at.",
	},
	"host_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Required. Id of the onprem host.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"ni_discovery_config": schema.SingleNestedAttribute{
		Attributes:          NIDiscoveryConfigV2ResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "V2 Configuration of the NI Discovery Service with scan domains.",
	},
	"ni_job_config_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Reference to the NI Job Configuration.",
	},
	"ni_job_config_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Read-only. Name of the Job Configuration.",
	},
	"ophid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Ophid of Host on which the job is to run.",
	},
	"pool_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Id of the pool of which the service and host belongs.",
	},
	"service_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Read-only. Timestamp updated_at.",
	},
}

func ExpandNIDiscoveryServiceV2(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.NIDiscoveryServiceV2 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NIDiscoveryServiceV2Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NIDiscoveryServiceV2Model) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.NIDiscoveryServiceV2 {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.NIDiscoveryServiceV2{
		HostId:            flex.ExpandStringPointer(m.HostId),
		NiDiscoveryConfig: ExpandNIDiscoveryConfigV2(ctx, m.NiDiscoveryConfig, diags),
		NiJobConfigId:     flex.ExpandStringPointer(m.NiJobConfigId),
		Ophid:             flex.ExpandStringPointer(m.Ophid),
		PoolId:            flex.ExpandStringPointer(m.PoolId),
		ServiceId:         flex.ExpandStringPointer(m.ServiceId),
	}
	return to
}

func FlattenNIDiscoveryServiceV2(ctx context.Context, from *discoverycontroller.NIDiscoveryServiceV2, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NIDiscoveryServiceV2AttrTypes)
	}
	m := NIDiscoveryServiceV2Model{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NIDiscoveryServiceV2AttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NIDiscoveryServiceV2Model) Flatten(ctx context.Context, from *discoverycontroller.NIDiscoveryServiceV2, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NIDiscoveryServiceV2Model{}
	}
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.HostId = flex.FlattenStringPointer(from.HostId)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.NiDiscoveryConfig = FlattenNIDiscoveryConfigV2(ctx, from.NiDiscoveryConfig, diags)
	m.NiJobConfigId = flex.FlattenStringPointer(from.NiJobConfigId)
	m.NiJobConfigName = flex.FlattenStringPointer(from.NiJobConfigName)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.PoolId = flex.FlattenStringPointer(from.PoolId)
	m.ServiceId = flex.FlattenStringPointer(from.ServiceId)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
