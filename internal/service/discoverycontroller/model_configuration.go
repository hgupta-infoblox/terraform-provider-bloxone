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

type ConfigurationModel struct {
	AssetsUnmanaged types.Bool        `tfsdk:"assets_unmanaged"`
	CreatedAt       timetypes.RFC3339 `tfsdk:"created_at"`
	Description     types.String      `tfsdk:"description"`
	Destination     types.String      `tfsdk:"destination"`
	DestinationId   types.String      `tfsdk:"destination_id"`
	Enable          types.Bool        `tfsdk:"enable"`
	Id              types.String      `tfsdk:"id"`
	Name            types.String      `tfsdk:"name"`
	SourceId        types.String      `tfsdk:"source_id"`
	SourceName      types.String      `tfsdk:"source_name"`
	SourceType      types.String      `tfsdk:"source_type"`
	UpdatedAt       timetypes.RFC3339 `tfsdk:"updated_at"`
}

var ConfigurationAttrTypes = map[string]attr.Type{
	"assets_unmanaged": types.BoolType,
	"created_at":       timetypes.RFC3339Type{},
	"description":      types.StringType,
	"destination":      types.StringType,
	"destination_id":   types.StringType,
	"enable":           types.BoolType,
	"id":               types.StringType,
	"name":             types.StringType,
	"source_id":        types.StringType,
	"source_name":      types.StringType,
	"source_type":      types.StringType,
	"updated_at":       timetypes.RFC3339Type{},
}

var ConfigurationResourceSchemaAttributes = map[string]schema.Attribute{
	"assets_unmanaged": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates whether discovered assets should be tagged as unmanaged.",
	},
	"created_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Read-only. Timestamp created_at, equals to time when particular object was created.",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Optional. Description of the configuration.",
	},
	"destination": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Read-only. Name of the Read-Write IPSpace where discovered objects would be created.",
	},
	"destination_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"enable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Flag to indicate if this Configuration is enabled. Defaults to _true_.",
	},
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Required. Name of the configuration.",
	},
	"source_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Source id is the unique identifier of the source. Example: in case of NGC, it would be NIOS grid id.",
	},
	"source_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "In case of NGC, it would be the grid name. The grid name can have conflicts for multiple grids in same account. So grid-name should be suffixed with last 4 chars of grid_id. Example: Infoblox-c35f This can be done only if name is default (Infoblox) In case of Active Discovery, it would be NMAP.",
	},
	"source_type": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Required. Defines the source type for the configuration. Valid values are:  * _ngc_: NIOS Grid Connector source type.  * _active_: Active Discovery (NMAP) source type. **Deprecated** - retained for backwards compatibility only. Use _ni_ for new deployments.  * _ni_: Network Insight active discovery source type.  * _passive_: Passive Discovery source type.",
	},
	"updated_at": schema.StringAttribute{
		CustomType:          timetypes.RFC3339Type{},
		Computed:            true,
		MarkdownDescription: "Read-only. Timestamp updated_at, equals to time when particular object was updated, by default equals to created_at.",
	},
}

func ExpandConfiguration(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.Configuration {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ConfigurationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ConfigurationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.Configuration {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.Configuration{
		AssetsUnmanaged: flex.ExpandBoolPointer(m.AssetsUnmanaged),
		Description:     flex.ExpandStringPointer(m.Description),
		Destination:     flex.ExpandStringPointer(m.Destination),
		DestinationId:   flex.ExpandStringPointer(m.DestinationId),
		Enable:          flex.ExpandBoolPointer(m.Enable),
		Name:            flex.ExpandStringPointer(m.Name),
		SourceId:        flex.ExpandStringPointer(m.SourceId),
		SourceName:      flex.ExpandStringPointer(m.SourceName),
		SourceType:      flex.ExpandStringPointer(m.SourceType),
	}
	return to
}

func FlattenConfiguration(ctx context.Context, from *discoverycontroller.Configuration, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ConfigurationAttrTypes)
	}
	m := ConfigurationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ConfigurationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ConfigurationModel) Flatten(ctx context.Context, from *discoverycontroller.Configuration, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ConfigurationModel{}
	}
	m.AssetsUnmanaged = types.BoolPointerValue(from.AssetsUnmanaged)
	m.CreatedAt = timetypes.NewRFC3339TimePointerValue(from.CreatedAt)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Destination = flex.FlattenStringPointer(from.Destination)
	m.DestinationId = flex.FlattenStringPointer(from.DestinationId)
	m.Enable = types.BoolPointerValue(from.Enable)
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.SourceId = flex.FlattenStringPointer(from.SourceId)
	m.SourceName = flex.FlattenStringPointer(from.SourceName)
	m.SourceType = flex.FlattenStringPointer(from.SourceType)
	m.UpdatedAt = timetypes.NewRFC3339TimePointerValue(from.UpdatedAt)
}
