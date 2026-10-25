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

type KeyStoreValueModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

var KeyStoreValueAttrTypes = map[string]attr.Type{
	"name":  types.StringType,
	"value": types.StringType,
}

var KeyStoreValueResourceSchemaAttributes = map[string]schema.Attribute{
	"name": schema.StringAttribute{
		Computed: true,
	},
	"value": schema.StringAttribute{
		Computed: true,
	},
}

func ExpandKeyStoreValue(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.KeyStoreValue {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m KeyStoreValueModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *KeyStoreValueModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.KeyStoreValue {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.KeyStoreValue{}
	return to
}

func FlattenKeyStoreValue(ctx context.Context, from *discoverycontroller.KeyStoreValue, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(KeyStoreValueAttrTypes)
	}
	m := KeyStoreValueModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, KeyStoreValueAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *KeyStoreValueModel) Flatten(ctx context.Context, from *discoverycontroller.KeyStoreValue, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = KeyStoreValueModel{}
	}
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Value = flex.FlattenStringPointer(from.Value)
}
