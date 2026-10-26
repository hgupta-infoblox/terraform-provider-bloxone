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

type ScheduleTimeModel struct {
	Dom      types.Int64  `tfsdk:"dom"`
	Dow      types.String `tfsdk:"dow"`
	Duration types.String `tfsdk:"duration"`
	Enabled  types.Bool   `tfsdk:"enabled"`
	Every    types.Int64  `tfsdk:"every"`
	Time     types.String `tfsdk:"time"`
	Type     types.String `tfsdk:"type"`
}

var ScheduleTimeAttrTypes = map[string]attr.Type{
	"dom":      types.Int64Type,
	"dow":      types.StringType,
	"duration": types.StringType,
	"enabled":  types.BoolType,
	"every":    types.Int64Type,
	"time":     types.StringType,
	"type":     types.StringType,
}

var ScheduleTimeResourceSchemaAttributes = map[string]schema.Attribute{
	"dom": schema.Int64Attribute{
		Optional: true,
	},
	"dow": schema.StringAttribute{
		Optional: true,
	},
	"duration": schema.StringAttribute{
		Optional: true,
	},
	"enabled": schema.BoolAttribute{
		Optional: true,
	},
	"every": schema.Int64Attribute{
		Optional: true,
	},
	"time": schema.StringAttribute{
		Optional: true,
	},
	"type": schema.StringAttribute{
		Optional: true,
	},
}

func ExpandScheduleTime(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.ScheduleTime {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ScheduleTimeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *ScheduleTimeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.ScheduleTime {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.ScheduleTime{
		Dom:      flex.ExpandInt64Pointer(m.Dom),
		Dow:      flex.ExpandStringPointer(m.Dow),
		Duration: flex.ExpandStringPointer(m.Duration),
		Enabled:  flex.ExpandBoolPointer(m.Enabled),
		Every:    flex.ExpandInt64Pointer(m.Every),
		Time:     flex.ExpandStringPointer(m.Time),
		Type:     flex.ExpandStringPointer(m.Type),
	}
	return to
}

func FlattenScheduleTime(ctx context.Context, from *discoverycontroller.ScheduleTime, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ScheduleTimeAttrTypes)
	}
	m := ScheduleTimeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ScheduleTimeAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *ScheduleTimeModel) Flatten(ctx context.Context, from *discoverycontroller.ScheduleTime, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = ScheduleTimeModel{}
	}
	m.Dom = flex.FlattenInt64Pointer(from.Dom)
	m.Dow = flex.FlattenStringPointer(from.Dow)
	m.Duration = flex.FlattenStringPointer(from.Duration)
	m.Enabled = types.BoolPointerValue(from.Enabled)
	m.Every = flex.FlattenInt64Pointer(from.Every)
	m.Time = flex.FlattenStringPointer(from.Time)
	m.Type = flex.FlattenStringPointer(from.Type)
}
