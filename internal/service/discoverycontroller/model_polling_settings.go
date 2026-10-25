package discoverycontroller

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/universal-ddi-go-client/discoverycontroller"
)

type PollingSettingsModel struct {
	Polling              types.Object `tfsdk:"polling"`
	SwitchPortManagement types.Object `tfsdk:"switch_port_management"`
}

var PollingSettingsAttrTypes = map[string]attr.Type{
	"polling":                types.ObjectType{AttrTypes: NetworkPollingAttrTypes},
	"switch_port_management": types.ObjectType{AttrTypes: SwitchPortManagementAttrTypes},
}

var PollingSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"polling": schema.SingleNestedAttribute{
		Attributes: NetworkPollingResourceSchemaAttributes,
		Optional:   true,
	},
	"switch_port_management": schema.SingleNestedAttribute{
		Attributes: SwitchPortManagementResourceSchemaAttributes,
		Optional:   true,
	},
}

func ExpandPollingSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.PollingSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PollingSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *PollingSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.PollingSettings {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.PollingSettings{
		Polling:              ExpandNetworkPolling(ctx, m.Polling, diags),
		SwitchPortManagement: ExpandSwitchPortManagement(ctx, m.SwitchPortManagement, diags),
	}
	return to
}

func FlattenPollingSettings(ctx context.Context, from *discoverycontroller.PollingSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PollingSettingsAttrTypes)
	}
	m := PollingSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PollingSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *PollingSettingsModel) Flatten(ctx context.Context, from *discoverycontroller.PollingSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = PollingSettingsModel{}
	}
	m.Polling = FlattenNetworkPolling(ctx, from.Polling, diags)
	m.SwitchPortManagement = FlattenSwitchPortManagement(ctx, from.SwitchPortManagement, diags)
}
