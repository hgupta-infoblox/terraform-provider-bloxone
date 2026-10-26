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

type ScanNetworkModel struct {
	DestinationIpSpaceId    types.String `tfsdk:"destination_ip_space_id"`
	DestinationIpSpaceName  types.String `tfsdk:"destination_ip_space_name"`
	ExclusionIpAddressRange types.String `tfsdk:"exclusion_ip_address_range"`
	ExclusionMode           types.String `tfsdk:"exclusion_mode"`
	IpAddressRange          types.String `tfsdk:"ip_address_range"`
	IpSpaceAction           types.String `tfsdk:"ip_space_action"`
	PingSweep               types.Bool   `tfsdk:"ping_sweep"`
}

var ScanNetworkAttrTypes = map[string]attr.Type{
	"destination_ip_space_id":    types.StringType,
	"destination_ip_space_name":  types.StringType,
	"exclusion_ip_address_range": types.StringType,
	"exclusion_mode":             types.StringType,
	"ip_address_range":           types.StringType,
	"ip_space_action":            types.StringType,
	"ping_sweep":                 types.BoolType,
}

var ScanNetworkResourceSchemaAttributes = map[string]schema.Attribute{
	"destination_ip_space_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"destination_ip_space_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Read-only. Name of the Read-Write IPSpace where discovered objects would be created.",
	},
	"exclusion_ip_address_range": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Represents the ipaddress or range that has to be excluded from the discovery scan.",
	},
	"exclusion_mode": schema.StringAttribute{
		Optional: true,
	},
	"ip_address_range": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Represents the network to be scanned in the format of cidr/mask or range of IPs.",
	},
	"ip_space_action": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The inheritance setting for IP Space Valid values are: * _inherit_: Use the inherited value from parent IPSpace * _override_: Use the value set in the object.  Defaults to _inherit_.",
	},
	"ping_sweep": schema.BoolAttribute{
		Optional: true,
	},
}

func ExpandScanNetwork(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.ScanNetwork {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ScanNetworkModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ScanNetworkModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.ScanNetwork {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.ScanNetwork{
		DestinationIpSpaceId:    flex.ExpandStringPointer(m.DestinationIpSpaceId),
		DestinationIpSpaceName:  flex.ExpandStringPointer(m.DestinationIpSpaceName),
		ExclusionIpAddressRange: flex.ExpandStringPointer(m.ExclusionIpAddressRange),
		ExclusionMode:           flex.ExpandStringPointer(m.ExclusionMode),
		IpAddressRange:          flex.ExpandStringPointer(m.IpAddressRange),
		IpSpaceAction:           flex.ExpandStringPointer(m.IpSpaceAction),
		PingSweep:               flex.ExpandBoolPointer(m.PingSweep),
	}
	return to
}

func FlattenScanNetwork(ctx context.Context, from *discoverycontroller.ScanNetwork, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ScanNetworkAttrTypes)
	}
	m := ScanNetworkModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ScanNetworkAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ScanNetworkModel) Flatten(ctx context.Context, from *discoverycontroller.ScanNetwork, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ScanNetworkModel{}
	}
	m.DestinationIpSpaceId = flex.FlattenStringPointer(from.DestinationIpSpaceId)
	m.DestinationIpSpaceName = flex.FlattenStringPointer(from.DestinationIpSpaceName)
	m.ExclusionIpAddressRange = flex.FlattenStringPointer(from.ExclusionIpAddressRange)
	m.ExclusionMode = flex.FlattenStringPointer(from.ExclusionMode)
	m.IpAddressRange = flex.FlattenStringPointer(from.IpAddressRange)
	m.IpSpaceAction = flex.FlattenStringPointer(from.IpSpaceAction)
	m.PingSweep = types.BoolPointerValue(from.PingSweep)
}
