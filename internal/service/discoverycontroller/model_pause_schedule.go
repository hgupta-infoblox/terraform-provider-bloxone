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

type PauseScheduleModel struct {
	EndTime   types.Int64 `tfsdk:"end_time"`
	StartTime types.Int64 `tfsdk:"start_time"`
	Weekday   types.Int64 `tfsdk:"weekday"`
}

var PauseScheduleAttrTypes = map[string]attr.Type{
	"end_time":   types.Int64Type,
	"start_time": types.Int64Type,
	"weekday":    types.Int64Type,
}

var PauseScheduleResourceSchemaAttributes = map[string]schema.Attribute{
	"end_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time of the day in minutes that indicates the end of the blackout window.",
	},
	"start_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The time of the day in minutes that indicates the start of the blackout window.",
	},
	"weekday": schema.Int64Attribute{
		Optional: true,
	},
}

func ExpandPauseSchedule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *discoverycontroller.PauseSchedule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m PauseScheduleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

func (m *PauseScheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *discoverycontroller.PauseSchedule {
	if m == nil {
		return nil
	}
	to := &discoverycontroller.PauseSchedule{
		EndTime:   flex.ExpandInt64Pointer(m.EndTime),
		StartTime: flex.ExpandInt64Pointer(m.StartTime),
		Weekday:   flex.ExpandInt64Pointer(m.Weekday),
	}
	return to
}

func FlattenPauseSchedule(ctx context.Context, from *discoverycontroller.PauseSchedule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(PauseScheduleAttrTypes)
	}
	m := PauseScheduleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, PauseScheduleAttrTypes, m)
	diags.Append(d...)
	return t
}

func (m *PauseScheduleModel) Flatten(ctx context.Context, from *discoverycontroller.PauseSchedule, diags *diag.Diagnostics) {
	if from == nil {
		return
	}
	if m == nil {
		*m = PauseScheduleModel{}
	}
	m.EndTime = flex.FlattenInt64Pointer(from.EndTime)
	m.StartTime = flex.FlattenInt64Pointer(from.StartTime)
	m.Weekday = flex.FlattenInt64Pointer(from.Weekday)
}
