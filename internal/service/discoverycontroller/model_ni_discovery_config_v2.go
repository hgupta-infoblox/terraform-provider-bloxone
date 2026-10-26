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

type NIDiscoveryConfigV2Model struct {
	BlackoutSchedule   types.List   `tfsdk:"blackout_schedule"`
	ScanDomains        types.List   `tfsdk:"scan_domains"`
	SdnConfigurations  types.List   `tfsdk:"sdn_configurations"`
	SdnPollingSettings types.Object `tfsdk:"sdn_polling_settings"`
}

var NIDiscoveryConfigV2AttrTypes = map[string]attr.Type{
	"blackout_schedule":    types.ListType{ElemType: types.ObjectType{AttrTypes: ScheduleTimeAttrTypes}},
	"scan_domains":         types.ListType{ElemType: types.ObjectType{AttrTypes: ScanDomainAttrTypes}},
	"sdn_configurations":   types.ListType{ElemType: types.ObjectType{AttrTypes: SDNConfigurationAttrTypes}},
	"sdn_polling_settings": types.ObjectType{AttrTypes: SDNPollingSettingsAttrTypes},
}

var NIDiscoveryConfigV2ResourceSchemaAttributes = map[string]schema.Attribute{
	"blackout_schedule": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ScheduleTimeResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Discovery blackout settings.",
	},
	"scan_domains": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ScanDomainResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Scan domains grouping scan networks and seed routers by interface.",
	},
	"sdn_configurations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SDNConfigurationResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "List of SDN connection credentials.",
	},
	"sdn_polling_settings": schema.SingleNestedAttribute{
		Attributes:          SDNPollingSettingsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Settings for SDN endpoint polling behavior.",
	},
}

func ExpandNIDiscoveryConfigV2(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.NIDiscoveryConfigV2 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NIDiscoveryConfigV2Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NIDiscoveryConfigV2Model) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.NIDiscoveryConfigV2 {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.NIDiscoveryConfigV2{
		BlackoutSchedule:   flex.ExpandFrameworkListNestedBlock(ctx, m.BlackoutSchedule, diags, ExpandScheduleTime),
		ScanDomains:        flex.ExpandFrameworkListNestedBlock(ctx, m.ScanDomains, diags, ExpandScanDomain),
		SdnConfigurations:  flex.ExpandFrameworkListNestedBlock(ctx, m.SdnConfigurations, diags, ExpandSDNConfiguration),
		SdnPollingSettings: ExpandSDNPollingSettings(ctx, m.SdnPollingSettings, diags),
	}
	return to
}

func FlattenNIDiscoveryConfigV2(ctx context.Context, from *discoverycontroller.NIDiscoveryConfigV2, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NIDiscoveryConfigV2AttrTypes)
	}
	m := NIDiscoveryConfigV2Model{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NIDiscoveryConfigV2AttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NIDiscoveryConfigV2Model) Flatten(ctx context.Context, from *discoverycontroller.NIDiscoveryConfigV2, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NIDiscoveryConfigV2Model{}
	}
	m.BlackoutSchedule = flex.FlattenFrameworkListNestedBlock(ctx, from.BlackoutSchedule, ScheduleTimeAttrTypes, diags, FlattenScheduleTime)
	m.ScanDomains = flex.FlattenFrameworkListNestedBlock(ctx, from.ScanDomains, ScanDomainAttrTypes, diags, FlattenScanDomain)
	m.SdnConfigurations = flex.FlattenFrameworkListNestedBlock(ctx, from.SdnConfigurations, SDNConfigurationAttrTypes, diags, FlattenSDNConfiguration)
	m.SdnPollingSettings = FlattenSDNPollingSettings(ctx, from.SdnPollingSettings, diags)
}
