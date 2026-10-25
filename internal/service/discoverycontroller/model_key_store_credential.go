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

type KeyStoreCredentialModel struct {
	Id     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Notes  types.String `tfsdk:"notes"`
	Values types.List   `tfsdk:"values"`
}

var KeyStoreCredentialAttrTypes = map[string]attr.Type{
	"id":     types.StringType,
	"name":   types.StringType,
	"notes":  types.StringType,
	"values": types.ListType{ElemType: types.ObjectType{AttrTypes: KeyStoreValueAttrTypes}},
}

var KeyStoreCredentialResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed: true,
	},
	"name": schema.StringAttribute{
		Computed: true,
	},
	"notes": schema.StringAttribute{
		Computed: true,
	},
	"values": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: KeyStoreValueResourceSchemaAttributes,
		},
		Computed: true,
	},
}

func ExpandKeyStoreCredential(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.KeyStoreCredential {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m KeyStoreCredentialModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *KeyStoreCredentialModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.KeyStoreCredential {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.KeyStoreCredential{}
	return to
}

func FlattenKeyStoreCredential(ctx context.Context, from *discoverycontroller.KeyStoreCredential, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(KeyStoreCredentialAttrTypes)
	}
	m := KeyStoreCredentialModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, KeyStoreCredentialAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *KeyStoreCredentialModel) Flatten(ctx context.Context, from *discoverycontroller.KeyStoreCredential, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = KeyStoreCredentialModel{}
	}
	m.Id = flex.FlattenStringPointer(from.Id)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Notes = flex.FlattenStringPointer(from.Notes)
	m.Values = flex.FlattenFrameworkListNestedBlock(ctx, from.Values, KeyStoreValueAttrTypes, diags, FlattenKeyStoreValue)
}
