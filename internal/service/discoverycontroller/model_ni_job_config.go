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

type NIJobConfigModel struct {
	ConfigId                 types.String      `tfsdk:"config_id"`
	CreatedAt                timetypes.RFC3339 `tfsdk:"created_at"`
	ForceServiceListOverride types.Bool        `tfsdk:"force_service_list_override"`
	Id                       types.String      `tfsdk:"id"`
	NiJobConfigName          types.String      `tfsdk:"ni_job_config_name"`
	Ophid                    types.String      `tfsdk:"ophid"`
	PoolId                   types.String      `tfsdk:"pool_id"`
	Properties               types.Object      `tfsdk:"properties"`
	ServiceId                types.String      `tfsdk:"service_id"`
	ServiceName              types.String      `tfsdk:"service_name"`
	Services                 types.List        `tfsdk:"services"`
	UpdatedAt                timetypes.RFC3339 `tfsdk:"updated_at"`
}

var NIJobConfigAttrTypes = map[string]attr.Type{
	"config_id":                   types.StringType,
	"created_at":                  timetypes.RFC3339Type{},
	"force_service_list_override": types.BoolType,
	"id":                          types.StringType,
	"ni_job_config_name":          types.StringType,
	"ophid":                       types.StringType,
	"pool_id":                     types.StringType,
	"properties":                  types.ObjectType{AttrTypes: NIPropertiesAttrTypes},
	"service_id":                  types.StringType,
	"service_name":                types.StringType,
	"services":                    types.ListType{ElemType: types.ObjectType{AttrTypes: ServiceAttrTypes}},
	"updated_at":                  timetypes.RFC3339Type{},
}

var NIJobConfigResourceSchemaAttributes = map[string]schema.Attribute{
	"config_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Read-only. Timestamp created_at, equals to time when particular object was created.",
	},
	"force_service_list_override": schema.BoolAttribute{
		Optional: true,
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"ni_job_config_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "ni_job_config_name represents the name of the NI Job Configuration.",
	},
	"ophid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Deprecated. On-Prem Host ID on which the job is to run.",
	},
	"pool_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Unique Identifier of the pool to which the service and host belongs.",
	},
	"properties": schema.SingleNestedAttribute{
		Attributes: NIPropertiesResourceSchemaAttributes,
		Optional:   true,
	},
	"service_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Deprecated. Unique Identifier of the discovery service associated with this job config.",
	},
	"service_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Deprecated. Read-only. Name of the discovery service.",
	},
	"services": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ServiceResourceSchemaAttributes,
		},
		Optional: true,
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at.",
	},
}

func ExpandNIJobConfig(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.NIJobConfig {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NIJobConfigModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NIJobConfigModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.NIJobConfig {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.NIJobConfig{
		ConfigId:                 flex.ExpandStringPointer(m.ConfigId),
		ForceServiceListOverride: flex.ExpandBoolPointer(m.ForceServiceListOverride),
		Ophid:                    flex.ExpandStringPointer(m.Ophid),
		PoolId:                   flex.ExpandStringPointer(m.PoolId),
		Properties:               ExpandNIProperties(ctx, m.Properties, diags),
		ServiceId:                flex.ExpandStringPointer(m.ServiceId),
		Services:                 flex.ExpandFrameworkListNestedBlock(ctx, m.Services, diags, ExpandService),
	}
	return to
}

func FlattenNIJobConfig(ctx context.Context, from *discoverycontroller.NIJobConfig, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NIJobConfigAttrTypes)
	}
	m := NIJobConfigModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NIJobConfigAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NIJobConfigModel) Flatten(ctx context.Context, from *discoverycontroller.NIJobConfig, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NIJobConfigModel{}
	}
	m.ConfigId = flex.FlattenStringPointer(from.ConfigId)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.ForceServiceListOverride = types.BoolPointerValue(from.ForceServiceListOverride)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.NiJobConfigName = flex.FlattenStringPointer(from.NiJobConfigName)
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.PoolId = flex.FlattenStringPointer(from.PoolId)
	m.Properties = FlattenNIProperties(ctx, from.Properties, diags)
	m.ServiceId = flex.FlattenStringPointer(from.ServiceId)
	m.ServiceName = flex.FlattenStringPointer(from.ServiceName)
	m.Services = flex.FlattenFrameworkListNestedBlock(ctx, from.Services, ServiceAttrTypes, diags, FlattenService)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
