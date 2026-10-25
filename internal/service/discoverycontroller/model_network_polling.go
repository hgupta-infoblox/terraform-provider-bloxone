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

type NetworkPollingModel struct {
	ArpCollectionPriority    types.String  `tfsdk:"arp_collection_priority"`
	AutoArpRefresh           types.Bool    `tfsdk:"auto_arp_refresh"`
	CliCollection            types.Bool    `tfsdk:"cli_collection"`
	NetBiosScanning          types.Bool    `tfsdk:"net_bios_scanning"`
	PollingFrequencyModifier types.Float64 `tfsdk:"polling_frequency_modifier"`
	RouteCollectionPriority  types.String  `tfsdk:"route_collection_priority"`
	SmartSubnetPingSweep     types.Bool    `tfsdk:"smart_subnet_ping_sweep"`
	SnmpCollection           types.Bool    `tfsdk:"snmp_collection"`
}

var NetworkPollingAttrTypes = map[string]attr.Type{
	"arp_collection_priority":    types.StringType,
	"auto_arp_refresh":           types.BoolType,
	"cli_collection":             types.BoolType,
	"net_bios_scanning":          types.BoolType,
	"polling_frequency_modifier": types.Float64Type,
	"route_collection_priority":  types.StringType,
	"smart_subnet_ping_sweep":    types.BoolType,
	"snmp_collection":            types.BoolType,
}

var NetworkPollingResourceSchemaAttributes = map[string]schema.Attribute{
	"arp_collection_priority": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Instructs whether to perform ARP collection using SNMP or CLI.",
	},
	"auto_arp_refresh": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "When true, enables automatic ARP Refresh across all switch ports.",
	},
	"cli_collection": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Instructs whether data collection for discovery is to be carried out using CLI.",
	},
	"net_bios_scanning": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "When true, enables NetBios Scanning.",
	},
	"polling_frequency_modifier": schema.Float64Attribute{
		Optional: true,
	},
	"route_collection_priority": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Instructs whether to perform Route collection using SNMP or CLI.",
	},
	"smart_subnet_ping_sweep": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Instructs whether data collection for discovery is to be carried out during a Smart Subnet Ping Sweep.",
	},
	"snmp_collection": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Instructs whether data collection for discovery is to be carried out using SNMP protocol.",
	},
}

func ExpandNetworkPolling(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.NetworkPolling {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkPollingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NetworkPollingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.NetworkPolling {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.NetworkPolling{
		ArpCollectionPriority:    flex.ExpandStringPointer(m.ArpCollectionPriority),
		AutoArpRefresh:           flex.ExpandBoolPointer(m.AutoArpRefresh),
		CliCollection:            flex.ExpandBoolPointer(m.CliCollection),
		NetBiosScanning:          flex.ExpandBoolPointer(m.NetBiosScanning),
		PollingFrequencyModifier: flex.ExpandFloat64Pointer(m.PollingFrequencyModifier),
		RouteCollectionPriority:  flex.ExpandStringPointer(m.RouteCollectionPriority),
		SmartSubnetPingSweep:     flex.ExpandBoolPointer(m.SmartSubnetPingSweep),
		SnmpCollection:           flex.ExpandBoolPointer(m.SnmpCollection),
	}
	return to
}

func FlattenNetworkPolling(ctx context.Context, from *discoverycontroller.NetworkPolling, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkPollingAttrTypes)
	}
	m := NetworkPollingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkPollingAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NetworkPollingModel) Flatten(ctx context.Context, from *discoverycontroller.NetworkPolling, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NetworkPollingModel{}
	}
	m.ArpCollectionPriority = flex.FlattenStringPointer(from.ArpCollectionPriority)
	m.AutoArpRefresh = types.BoolPointerValue(from.AutoArpRefresh)
	m.CliCollection = types.BoolPointerValue(from.CliCollection)
	m.NetBiosScanning = types.BoolPointerValue(from.NetBiosScanning)
	m.PollingFrequencyModifier = flex.FlattenFloat64Pointer(from.PollingFrequencyModifier)
	m.RouteCollectionPriority = flex.FlattenStringPointer(from.RouteCollectionPriority)
	m.SmartSubnetPingSweep = types.BoolPointerValue(from.SmartSubnetPingSweep)
	m.SnmpCollection = types.BoolPointerValue(from.SnmpCollection)
}
