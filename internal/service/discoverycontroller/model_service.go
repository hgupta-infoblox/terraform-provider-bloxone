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

type ServiceModel struct {
	Ophid       types.String `tfsdk:"ophid"`
	ServiceId   types.String `tfsdk:"service_id"`
	ServiceName types.String `tfsdk:"service_name"`
}

var ServiceAttrTypes = map[string]attr.Type{
	"ophid":        types.StringType,
	"service_id":   types.StringType,
	"service_name": types.StringType,
}

var ServiceResourceSchemaAttributes = map[string]schema.Attribute{
	"ophid": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Required. On-Prem Host ID on which the job is to run.",
	},
	"service_id": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Required. Unique Identifier of the discovery service associated with this job config.",
	},
	"service_name": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Read-only. Name of the discovery service.",
	},
}

func ExpandService(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.Service {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ServiceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ServiceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.Service {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.Service{
		Ophid:     flex.ExpandStringPointer(m.Ophid),
		ServiceId: flex.ExpandStringPointer(m.ServiceId),
	}
	return to
}

func FlattenService(ctx context.Context, from *discoverycontroller.Service, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ServiceAttrTypes)
	}
	m := ServiceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ServiceAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ServiceModel) Flatten(ctx context.Context, from *discoverycontroller.Service, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ServiceModel{}
	}
	m.Ophid = flex.FlattenStringPointer(from.Ophid)
	m.ServiceId = flex.FlattenStringPointer(from.ServiceId)
	m.ServiceName = flex.FlattenStringPointer(from.ServiceName)
}
