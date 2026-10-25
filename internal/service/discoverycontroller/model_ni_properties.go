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

type NIPropertiesModel struct {
	AdvancedSettings  types.Object `tfsdk:"advanced_settings"`
	CliCredentials    types.List   `tfsdk:"cli_credentials"`
	PollingSettings   types.Object `tfsdk:"polling_settings"`
	ServicePorts      types.List   `tfsdk:"service_ports"`
	SnmpCredentials   types.List   `tfsdk:"snmp_credentials"`
	SnmpV3Credentials types.List   `tfsdk:"snmp_v3_credentials"`
}

var NIPropertiesAttrTypes = map[string]attr.Type{
	"advanced_settings":   types.ObjectType{AttrTypes: AdvancedSettingsAttrTypes},
	"cli_credentials":     types.ListType{ElemType: types.ObjectType{AttrTypes: CLICredentialsAttrTypes}},
	"polling_settings":    types.ObjectType{AttrTypes: PollingSettingsAttrTypes},
	"service_ports":       types.ListType{ElemType: types.ObjectType{AttrTypes: ServicePortAttrTypes}},
	"snmp_credentials":    types.ListType{ElemType: types.ObjectType{AttrTypes: SNMPCredentialsAttrTypes}},
	"snmp_v3_credentials": types.ListType{ElemType: types.ObjectType{AttrTypes: SNMPV3CredentialsAttrTypes}},
}

var NIPropertiesResourceSchemaAttributes = map[string]schema.Attribute{
	"advanced_settings": schema.SingleNestedAttribute{
		Attributes:          AdvancedSettingsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "NI advanced settings.",
	},
	"cli_credentials": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: CLICredentialsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "CLI Credential Settings.",
	},
	"polling_settings": schema.SingleNestedAttribute{
		Attributes:          PollingSettingsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "NI polling settings.",
	},
	"service_ports": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ServicePortResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "Service port configuration.",
	},
	"snmp_credentials": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SNMPCredentialsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "SNMP Credentials Settings.",
	},
	"snmp_v3_credentials": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SNMPV3CredentialsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "SNMPv3 Credentials Settings.",
	},
}

func ExpandNIProperties(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.NIProperties {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NIPropertiesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *NIPropertiesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.NIProperties {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.NIProperties{
		AdvancedSettings:  ExpandAdvancedSettings(ctx, m.AdvancedSettings, diags),
		CliCredentials:    flex.ExpandFrameworkListNestedBlock(ctx, m.CliCredentials, diags, ExpandCLICredentials),
		PollingSettings:   ExpandPollingSettings(ctx, m.PollingSettings, diags),
		ServicePorts:      flex.ExpandFrameworkListNestedBlock(ctx, m.ServicePorts, diags, ExpandServicePort),
		SnmpCredentials:   flex.ExpandFrameworkListNestedBlock(ctx, m.SnmpCredentials, diags, ExpandSNMPCredentials),
		SnmpV3Credentials: flex.ExpandFrameworkListNestedBlock(ctx, m.SnmpV3Credentials, diags, ExpandSNMPV3Credentials),
	}
	return to
}

func FlattenNIProperties(ctx context.Context, from *discoverycontroller.NIProperties, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NIPropertiesAttrTypes)
	}
	m := NIPropertiesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NIPropertiesAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *NIPropertiesModel) Flatten(ctx context.Context, from *discoverycontroller.NIProperties, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = NIPropertiesModel{}
	}
	m.AdvancedSettings = FlattenAdvancedSettings(ctx, from.AdvancedSettings, diags)
	m.CliCredentials = flex.FlattenFrameworkListNestedBlock(ctx, from.CliCredentials, CLICredentialsAttrTypes, diags, FlattenCLICredentials)
	m.PollingSettings = FlattenPollingSettings(ctx, from.PollingSettings, diags)
	m.ServicePorts = flex.FlattenFrameworkListNestedBlock(ctx, from.ServicePorts, ServicePortAttrTypes, diags, FlattenServicePort)
	m.SnmpCredentials = flex.FlattenFrameworkListNestedBlock(ctx, from.SnmpCredentials, SNMPCredentialsAttrTypes, diags, FlattenSNMPCredentials)
	m.SnmpV3Credentials = flex.FlattenFrameworkListNestedBlock(ctx, from.SnmpV3Credentials, SNMPV3CredentialsAttrTypes, diags, FlattenSNMPV3Credentials)
}
