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

type SwitchPortManagementModel struct {
	Enabled               types.Bool   `tfsdk:"enabled"`
	PollingInterval       types.Int64  `tfsdk:"polling_interval"`
	PortControlPreference types.String `tfsdk:"port_control_preference"`
}

var SwitchPortManagementAttrTypes = map[string]attr.Type{
	"enabled":                 types.BoolType,
	"polling_interval":        types.Int64Type,
	"port_control_preference": types.StringType,
}

var SwitchPortManagementResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled": schema.BoolAttribute{
		Optional: true,
	},
	"polling_interval": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Polling interval in minutes.",
	},
	"port_control_preference": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Port control preference corresponds to one of SNMP or CLI.",
	},
}

func ExpandSwitchPortManagement(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.SwitchPortManagement {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SwitchPortManagementModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SwitchPortManagementModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.SwitchPortManagement {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.SwitchPortManagement{
		Enabled:               flex.ExpandBoolPointer(m.Enabled),
		PollingInterval:       flex.ExpandInt64Pointer(m.PollingInterval),
		PortControlPreference: flex.ExpandStringPointer(m.PortControlPreference),
	}
	return to
}

func FlattenSwitchPortManagement(ctx context.Context, from *discoverycontroller.SwitchPortManagement, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SwitchPortManagementAttrTypes)
	}
	m := SwitchPortManagementModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SwitchPortManagementAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SwitchPortManagementModel) Flatten(ctx context.Context, from *discoverycontroller.SwitchPortManagement, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SwitchPortManagementModel{}
	}
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.PollingInterval = flex.FlattenInt64Pointer(from.PollingInterval)
	m.PortControlPreference = flex.FlattenStringPointer(from.PortControlPreference)
}
