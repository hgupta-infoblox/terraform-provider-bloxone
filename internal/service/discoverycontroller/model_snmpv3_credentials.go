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

type SNMPV3CredentialsModel struct {
	Creds    types.Object `tfsdk:"creds"`
	Priority types.Int64  `tfsdk:"priority"`
}

var SNMPV3CredentialsAttrTypes = map[string]attr.Type{
	"creds":    types.ObjectType{AttrTypes: KeyStoreCredentialAttrTypes},
	"priority": types.Int64Type,
}

var SNMPV3CredentialsResourceSchemaAttributes = map[string]schema.Attribute{
	"creds": schema.SingleNestedAttribute{
		Attributes: KeyStoreCredentialResourceSchemaAttributes,
		Optional:   true,
	},
	"priority": schema.Int64Attribute{
		Optional: true,
	},
}

func ExpandSNMPV3Credentials(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.SNMPV3Credentials {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SNMPV3CredentialsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SNMPV3CredentialsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.SNMPV3Credentials {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.SNMPV3Credentials{
		Creds:    ExpandKeyStoreCredential(ctx, m.Creds, diags),
		Priority: flex.ExpandInt64Pointer(m.Priority),
	}
	return to
}

func FlattenSNMPV3Credentials(ctx context.Context, from *discoverycontroller.SNMPV3Credentials, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SNMPV3CredentialsAttrTypes)
	}
	m := SNMPV3CredentialsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SNMPV3CredentialsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SNMPV3CredentialsModel) Flatten(ctx context.Context, from *discoverycontroller.SNMPV3Credentials, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SNMPV3CredentialsModel{}
	}
	m.Creds = FlattenKeyStoreCredential(ctx, from.Creds, diags)
	m.Priority = flex.FlattenInt64Pointer(from.Priority)
}
