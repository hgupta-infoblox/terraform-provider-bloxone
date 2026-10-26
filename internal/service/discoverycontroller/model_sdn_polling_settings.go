package discoverycontroller

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/discoverycontroller"

	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
)

type SDNPollingSettingsModel struct {
	DefaultSdnVirtualNetwork      types.String `tfsdk:"default_sdn_virtual_network"`
	EndHostPollingEnabled         types.Bool   `tfsdk:"end_host_polling_enabled"`
	EndHostPollingIntervalMinutes types.Int64  `tfsdk:"end_host_polling_interval_minutes"`
	SdnDiscovery                  types.Bool   `tfsdk:"sdn_discovery"`
}

var SDNPollingSettingsAttrTypes = map[string]attr.Type{
	"default_sdn_virtual_network":       types.StringType,
	"end_host_polling_enabled":          types.BoolType,
	"end_host_polling_interval_minutes": types.Int64Type,
	"sdn_discovery":                     types.BoolType,
}

var SDNPollingSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"default_sdn_virtual_network": schema.StringAttribute{
		Optional: true,
	},
	"end_host_polling_enabled": schema.BoolAttribute{
		Optional: true,
	},
	"end_host_polling_interval_minutes": schema.Int64Attribute{
		Optional: true,
	},
	"sdn_discovery": schema.BoolAttribute{
		Optional: true,
	},
}

func ExpandSDNPollingSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.SDNPollingSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SDNPollingSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SDNPollingSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.SDNPollingSettings {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.SDNPollingSettings{
		DefaultSdnVirtualNetwork:      flex.ExpandStringPointer(m.DefaultSdnVirtualNetwork),
		EndHostPollingEnabled:         flex.ExpandBoolPointer(m.EndHostPollingEnabled),
		EndHostPollingIntervalMinutes: flex.ExpandInt64Pointer(m.EndHostPollingIntervalMinutes),
		SdnDiscovery:                  flex.ExpandBoolPointer(m.SdnDiscovery),
	}
	return to
}

func FlattenSDNPollingSettings(ctx context.Context, from *discoverycontroller.SDNPollingSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SDNPollingSettingsAttrTypes)
	}
	m := SDNPollingSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SDNPollingSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SDNPollingSettingsModel) Flatten(ctx context.Context, from *discoverycontroller.SDNPollingSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SDNPollingSettingsModel{}
	}
	m.DefaultSdnVirtualNetwork = flex.FlattenStringPointer(from.DefaultSdnVirtualNetwork)
	m.EndHostPollingEnabled = types.BoolPointerValue(from.EndHostPollingEnabled)
	m.EndHostPollingIntervalMinutes = flex.FlattenInt64Pointer(from.EndHostPollingIntervalMinutes)
	m.SdnDiscovery = types.BoolPointerValue(from.SdnDiscovery)
}
