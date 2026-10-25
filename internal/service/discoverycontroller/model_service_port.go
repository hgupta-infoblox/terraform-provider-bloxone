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

type ServicePortModel struct {
	Port     types.Int64  `tfsdk:"port"`
	Protocol types.String `tfsdk:"protocol"`
	Service  types.String `tfsdk:"service"`
}

var ServicePortAttrTypes = map[string]attr.Type{
	"port":     types.Int64Type,
	"protocol": types.StringType,
	"service":  types.StringType,
}

var ServicePortResourceSchemaAttributes = map[string]schema.Attribute{
	"port": schema.Int64Attribute{
		Optional: true,
	},
	"protocol": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Service protocol correspond to one of TCP or UDP.",
	},
	"service": schema.StringAttribute{
		Optional: true,
	},
}

func ExpandServicePort(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.ServicePort {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ServicePortModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ServicePortModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.ServicePort {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.ServicePort{
		Port:     flex.ExpandInt64Pointer(m.Port),
		Protocol: flex.ExpandStringPointer(m.Protocol),
		Service:  flex.ExpandStringPointer(m.Service),
	}
	return to
}

func FlattenServicePort(ctx context.Context, from *discoverycontroller.ServicePort, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ServicePortAttrTypes)
	}
	m := ServicePortModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ServicePortAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ServicePortModel) Flatten(ctx context.Context, from *discoverycontroller.ServicePort, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ServicePortModel{}
	}
	m.Port = flex.FlattenInt64Pointer(from.Port)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.Service = flex.FlattenStringPointer(from.Service)
}
