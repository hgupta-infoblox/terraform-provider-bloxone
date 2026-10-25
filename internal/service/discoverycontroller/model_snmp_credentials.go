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

type SNMPCredentialsModel struct {
	Creds    types.Object `tfsdk:"creds"`
	Priority types.Int64  `tfsdk:"priority"`
}

var SNMPCredentialsAttrTypes = map[string]attr.Type{
	"creds":    types.ObjectType{AttrTypes: KeyStoreCredentialAttrTypes},
	"priority": types.Int64Type,
}

var SNMPCredentialsResourceSchemaAttributes = map[string]schema.Attribute{
	"creds": schema.SingleNestedAttribute{
		Attributes: KeyStoreCredentialResourceSchemaAttributes,
		Optional:   true,
	},
	"priority": schema.Int64Attribute{
		Optional: true,
	},
}

func ExpandSNMPCredentials(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.SNMPCredentials {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SNMPCredentialsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *SNMPCredentialsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.SNMPCredentials {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.SNMPCredentials{
		Creds:    ExpandKeyStoreCredential(ctx, m.Creds, diags),
		Priority: flex.ExpandInt64Pointer(m.Priority),
	}
	return to
}

func FlattenSNMPCredentials(ctx context.Context, from *discoverycontroller.SNMPCredentials, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SNMPCredentialsAttrTypes)
	}
	m := SNMPCredentialsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SNMPCredentialsAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *SNMPCredentialsModel) Flatten(ctx context.Context, from *discoverycontroller.SNMPCredentials, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = SNMPCredentialsModel{}
	}
	m.Creds = FlattenKeyStoreCredential(ctx, from.Creds, diags)
	m.Priority = flex.FlattenInt64Pointer(from.Priority)
}
