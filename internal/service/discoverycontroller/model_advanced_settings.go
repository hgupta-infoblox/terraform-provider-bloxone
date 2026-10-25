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

type AdvancedSettingsModel struct {
	ArpAggregateLimit            types.Int64  `tfsdk:"arp_aggregate_limit"`
	ArpCacheRefreshPeriod        types.Int64  `tfsdk:"arp_cache_refresh_period"`
	DeviceExpirationDays         types.Int64  `tfsdk:"device_expiration_days"`
	DnsLookupOptions             types.String `tfsdk:"dns_lookup_options"`
	DnsLookupThrottle            types.Int64  `tfsdk:"dns_lookup_throttle"`
	EndHostExpirationDays        types.Int64  `tfsdk:"end_host_expiration_days"`
	PingRetries                  types.Int64  `tfsdk:"ping_retries"`
	PingSweepFrequency           types.Int64  `tfsdk:"ping_sweep_frequency"`
	PingTimeout                  types.Int64  `tfsdk:"ping_timeout"`
	PortScanningEnabled          types.Bool   `tfsdk:"port_scanning_enabled"`
	RouteLimit                   types.Int64  `tfsdk:"route_limit"`
	SnmpDiscoveryVersion         types.String `tfsdk:"snmp_discovery_version"`
	Snmpv1DataCollectionFallback types.String `tfsdk:"snmpv1_data_collection_fallback"`
	TcpScanTechnique             types.String `tfsdk:"tcp_scan_technique"`
}

var AdvancedSettingsAttrTypes = map[string]attr.Type{
	"arp_aggregate_limit":             types.Int64Type,
	"arp_cache_refresh_period":        types.Int64Type,
	"device_expiration_days":          types.Int64Type,
	"dns_lookup_options":              types.StringType,
	"dns_lookup_throttle":             types.Int64Type,
	"end_host_expiration_days":        types.Int64Type,
	"ping_retries":                    types.Int64Type,
	"ping_sweep_frequency":            types.Int64Type,
	"ping_timeout":                    types.Int64Type,
	"port_scanning_enabled":           types.BoolType,
	"route_limit":                     types.Int64Type,
	"snmp_discovery_version":          types.StringType,
	"snmpv1_data_collection_fallback": types.StringType,
	"tcp_scan_technique":              types.StringType,
}

var AdvancedSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"arp_aggregate_limit": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Number of MAC addresses.",
	},
	"arp_cache_refresh_period": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "ARP Cache Refresh in seconds or minutes.",
	},
	"device_expiration_days": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Days after which expired device data needs to be purged.",
	},
	"dns_lookup_options": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Specify whether to perform a DNS lookup when discovering the devices. Valid values are: * _disable_   : Turn off DNS lookup when discovering the devices. * _infraonly_ : Collect hostnames for the network devices using DNS lookup. * _all_       : Collect hostnames for network devices (switches and routers) and end host devices using DNS lookup.",
	},
	"dns_lookup_throttle": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "DNS Lookup Throttle. This is the value in percentage that throttles the traffic on the DNS server.",
	},
	"end_host_expiration_days": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Days after which expired assets data needs to be purged.",
	},
	"ping_retries": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Ping Sweep Attempts.",
	},
	"ping_sweep_frequency": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Ping Sweep Frequency in days or hours.",
	},
	"ping_timeout": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Ping Sweep timeout limit in seconds.",
	},
	"port_scanning_enabled": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Controls the Port Scanning for the discovered devices.  Valid values are: * _enabled_: Enables the port scanning for the services specified in the Service Ports. * _disabled_: Disables the port scanning for the services specified in the Service Ports.",
	},
	"route_limit": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "number of routes that can be configured in the discovery.",
	},
	"snmp_discovery_version": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Authenticate and poll using one of these versions of the SNMP protocol. Valid values are: * _snmpv1_  : Use SNMPV1 protocol. * _snmpv2c_ : Use SNMPV2C protocol.",
	},
	"snmpv1_data_collection_fallback": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Version of SNMP protocol for credential discovery and device polling. Valid values are: * _disabled_ : Exclusively use SNMPv2c and up. Do not fallback to SNMPV1. * _enabled_  : Support fallback to SNMPv1, if SNMPV2C and up are not available.",
	},
	"tcp_scan_technique": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Configure the TCP scan technique to be used for the discovery.  Valid values are: * _syn_:     Perform scans on thousands of TCP ports per system, never completing connections across any well-known port.              SYN packets are sent and the poller waits for a response while continuing to scan other ports.              A SYN/ACK response indicates the protocol port is listening while a RST indicates it is not listening.              The SYN option presents less impact on the network. * _connect_: Perform scan on IPv6 networks. Unlike the SYN option, complete connections are attempted on the              scanned system and each successive TCP protocol port being scanned.",
	},
}

func ExpandAdvancedSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.AdvancedSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m AdvancedSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *AdvancedSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.AdvancedSettings {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.AdvancedSettings{
		ArpAggregateLimit:            flex.ExpandInt64Pointer(m.ArpAggregateLimit),
		ArpCacheRefreshPeriod:        flex.ExpandInt64Pointer(m.ArpCacheRefreshPeriod),
		DeviceExpirationDays:         flex.ExpandInt64Pointer(m.DeviceExpirationDays),
		DnsLookupOptions:             flex.ExpandStringPointer(m.DnsLookupOptions),
		DnsLookupThrottle:            flex.ExpandInt64Pointer(m.DnsLookupThrottle),
		EndHostExpirationDays:        flex.ExpandInt64Pointer(m.EndHostExpirationDays),
		PingRetries:                  flex.ExpandInt64Pointer(m.PingRetries),
		PingSweepFrequency:           flex.ExpandInt64Pointer(m.PingSweepFrequency),
		PingTimeout:                  flex.ExpandInt64Pointer(m.PingTimeout),
		PortScanningEnabled:          flex.ExpandBoolPointer(m.PortScanningEnabled),
		RouteLimit:                   flex.ExpandInt64Pointer(m.RouteLimit),
		SnmpDiscoveryVersion:         flex.ExpandStringPointer(m.SnmpDiscoveryVersion),
		Snmpv1DataCollectionFallback: flex.ExpandStringPointer(m.Snmpv1DataCollectionFallback),
		TcpScanTechnique:             flex.ExpandStringPointer(m.TcpScanTechnique),
	}
	return to
}

func FlattenAdvancedSettings(ctx context.Context, from *discoverycontroller.AdvancedSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(AdvancedSettingsAttrTypes)
	}
	m := AdvancedSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, AdvancedSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *AdvancedSettingsModel) Flatten(ctx context.Context, from *discoverycontroller.AdvancedSettings, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = AdvancedSettingsModel{}
	}
	m.ArpAggregateLimit = flex.FlattenInt64Pointer(from.ArpAggregateLimit)
	m.ArpCacheRefreshPeriod = flex.FlattenInt64Pointer(from.ArpCacheRefreshPeriod)
	m.DeviceExpirationDays = flex.FlattenInt64Pointer(from.DeviceExpirationDays)
	m.DnsLookupOptions = flex.FlattenStringPointer(from.DnsLookupOptions)
	m.DnsLookupThrottle = flex.FlattenInt64Pointer(from.DnsLookupThrottle)
	m.EndHostExpirationDays = flex.FlattenInt64Pointer(from.EndHostExpirationDays)
	m.PingRetries = flex.FlattenInt64Pointer(from.PingRetries)
	m.PingSweepFrequency = flex.FlattenInt64Pointer(from.PingSweepFrequency)
	m.PingTimeout = flex.FlattenInt64Pointer(from.PingTimeout)
	m.PortScanningEnabled = types.BoolPointerValue(from.PortScanningEnabled)
	m.RouteLimit = flex.FlattenInt64Pointer(from.RouteLimit)
	m.SnmpDiscoveryVersion = flex.FlattenStringPointer(from.SnmpDiscoveryVersion)
	m.Snmpv1DataCollectionFallback = flex.FlattenStringPointer(from.Snmpv1DataCollectionFallback)
	m.TcpScanTechnique = flex.FlattenStringPointer(from.TcpScanTechnique)
}
