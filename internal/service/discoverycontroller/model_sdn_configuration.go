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

type SDNConfigurationModel struct {
	Address        types.String `tfsdk:"address"`
	CaCert         types.String `tfsdk:"ca_cert"`
	Creds          types.Object `tfsdk:"creds"`
	Handle         types.String `tfsdk:"handle"`
	OnPrem         types.Bool   `tfsdk:"on_prem"`
	Protocol       types.String `tfsdk:"protocol"`
	SdnType        types.String `tfsdk:"sdn_type"`
	UseGlobalProxy types.Bool   `tfsdk:"use_global_proxy"`
}

var SDNConfigurationAttrTypes = map[string]attr.Type{
	"address":          types.StringType,
	"ca_cert":          types.StringType,
	"creds":            types.ObjectType{AttrTypes: KeyStoreCredentialAttrTypes},
	"handle":           types.StringType,
	"on_prem":          types.BoolType,
	"protocol":         types.StringType,
	"sdn_type":         types.StringType,
	"use_global_proxy": types.BoolType,
}

var SDNConfigurationResourceSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
	},
	"ca_cert": schema.StringAttribute{
		Optional: true,
	},
	"creds": schema.SingleNestedAttribute{
		Attributes: KeyStoreCredentialResourceSchemaAttributes,
		Optional:   true,
	},
	"handle": schema.StringAttribute{
		Optional: true,
	},
	"on_prem": schema.BoolAttribute{
		Optional: true,
	},
	"protocol": schema.StringAttribute{
		Optional: true,
	},
	"sdn_type": schema.StringAttribute{
		Optional: true,
	},
	"use_global_proxy": schema.BoolAttribute{
		Optional: true,
	},
}

func ExpandSDNConfiguration(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.SDNConfiguration {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SDNConfigurationModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SDNConfigurationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.SDNConfiguration {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.SDNConfiguration{
		Address:        flex.ExpandStringPointer(m.Address),
		CaCert:         flex.ExpandStringPointer(m.CaCert),
		Creds:          ExpandKeyStoreCredential(ctx, m.Creds, diags),
		Handle:         flex.ExpandStringPointer(m.Handle),
		OnPrem:         flex.ExpandBoolPointer(m.OnPrem),
		Protocol:       flex.ExpandStringPointer(m.Protocol),
		SdnType:        flex.ExpandStringPointer(m.SdnType),
		UseGlobalProxy: flex.ExpandBoolPointer(m.UseGlobalProxy),
	}
	return to
}

func FlattenSDNConfiguration(ctx context.Context, from *discoverycontroller.SDNConfiguration, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SDNConfigurationAttrTypes)
	}
	m := SDNConfigurationModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SDNConfigurationAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SDNConfigurationModel) Flatten(ctx context.Context, from *discoverycontroller.SDNConfiguration, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SDNConfigurationModel{}
	}
	m.Address = flex.FlattenStringPointer(from.Address)
	m.CaCert = flex.FlattenStringPointer(from.CaCert)
	m.Creds = FlattenKeyStoreCredential(ctx, from.Creds, diags)
	m.Handle = flex.FlattenStringPointer(from.Handle)
	m.OnPrem = types.BoolPointerValue(from.OnPrem)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.SdnType = flex.FlattenStringPointer(from.SdnType)
	m.UseGlobalProxy = types.BoolPointerValue(from.UseGlobalProxy)
}
