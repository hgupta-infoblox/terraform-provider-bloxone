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

type ScanDomainModel struct {
	Id                    types.String `tfsdk:"id"`
	InterfaceName         types.String `tfsdk:"interface_name"`
	IpSpaceId             types.String `tfsdk:"ip_space_id"`
	IpSpaceName           types.String `tfsdk:"ip_space_name"`
	Name                  types.String `tfsdk:"name"`
	Networks              types.List   `tfsdk:"networks"`
	SeedRouterIpAddresses types.List   `tfsdk:"seed_router_ip_addresses"`
	VrfHint               types.String `tfsdk:"vrf_hint"`
}

var ScanDomainAttrTypes = map[string]attr.Type{
	"id":                       types.StringType,
	"interface_name":           types.StringType,
	"ip_space_id":              types.StringType,
	"ip_space_name":            types.StringType,
	"name":                     types.StringType,
	"networks":                 types.ListType{ElemType: types.ObjectType{AttrTypes: ScanNetworkAttrTypes}},
	"seed_router_ip_addresses": types.ListType{ElemType: types.StringType},
	"vrf_hint":                 types.StringType,
}

var ScanDomainResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Read-only. Unique identifier of this scan domain.",
	},
	"interface_name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Host interface associated with this scan domain (e.g., \"eth0\"). Defaults to \"eth0\" if empty.",
	},
	"ip_space_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"ip_space_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Read-only. Name of the IP space, resolved from IPAM.",
	},
	"name": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "User-given name for this scan domain (e.g., \"default\"). Must be unique within a service config.",
	},
	"networks": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ScanNetworkResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Scan networks belonging to this scan domain.",
	},
	"seed_router_ip_addresses": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "Seed router IP addresses belonging to this scan domain.",
	},
	"vrf_hint": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Regex to match VRF discovered over the interface. Defaults to \".*\" (match all).",
	},
}

func ExpandScanDomain(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.ScanDomain {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ScanDomainModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ScanDomainModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.ScanDomain {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.ScanDomain{
		InterfaceName:         flex.ExpandStringPointer(m.InterfaceName),
		IpSpaceId:             flex.ExpandStringPointer(m.IpSpaceId),
		Name:                  flex.ExpandStringPointer(m.Name),
		Networks:              flex.ExpandFrameworkListNestedBlock(ctx, m.Networks, diags, ExpandScanNetwork),
		SeedRouterIpAddresses: flex.ExpandFrameworkListString(ctx, m.SeedRouterIpAddresses, diags),
		VrfHint:               flex.ExpandStringPointer(m.VrfHint),
	}
	return to
}

func FlattenScanDomain(ctx context.Context, from *discoverycontroller.ScanDomain, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ScanDomainAttrTypes)
	}
	m := ScanDomainModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ScanDomainAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ScanDomainModel) Flatten(ctx context.Context, from *discoverycontroller.ScanDomain, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ScanDomainModel{}
	}
	m.Id = flex.FlattenStringPointer(from.Id)
	m.InterfaceName = flex.FlattenStringPointer(from.InterfaceName)
	m.IpSpaceId = flex.FlattenStringPointer(from.IpSpaceId)
	m.IpSpaceName = flex.FlattenStringPointer(from.IpSpaceName)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Networks = flex.FlattenFrameworkListNestedBlock(ctx, from.Networks, ScanNetworkAttrTypes, diags, FlattenScanNetwork)
	m.SeedRouterIpAddresses = flex.FlattenFrameworkListString(ctx, from.SeedRouterIpAddresses, diags)
	m.VrfHint = flex.FlattenStringPointer(from.VrfHint)
}
