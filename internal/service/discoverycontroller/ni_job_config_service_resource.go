package discoverycontroller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	universalddiclient "github.com/infobloxopen/universal-ddi-go-client/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NIJobConfigServiceResource{}
var _ resource.ResourceWithImportState = &NIJobConfigServiceResource{}

func NewNIJobConfigServiceResource() resource.Resource {
	return &NIJobConfigServiceResource{}
}

// NIJobConfigServiceResource defines the resource implementation.
type NIJobConfigServiceResource struct {
	client *universalddiclient.APIClient
}

func (r *NIJobConfigServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "discovery_ni_job_config"
}

func (r *NIJobConfigServiceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "",
		Attributes:          NIJobConfigResourceSchemaAttributes,
	}
}

func (r *NIJobConfigServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*universalddiclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *universalddiclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NIJobConfigServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NIJobConfigModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DiscoveryControllerAPI.
		NIJobConfigServiceAPI.
		Create(ctx).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create NIJobConfigService, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NIJobConfigServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NIJobConfigModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, httpRes, err := r.client.DiscoveryControllerAPI.
		NIJobConfigServiceAPI.
		Read(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read NIJobConfigService, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NIJobConfigServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NIJobConfigModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiRes, _, err := r.client.DiscoveryControllerAPI.
		NIJobConfigServiceAPI.
		Update(ctx, data.Id.ValueString()).
		Body(*data.Expand(ctx, &resp.Diagnostics)).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update NIJobConfigService, got error: %s", err))
		return
	}

	res := apiRes.GetResult()
	data.Flatten(ctx, &res, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NIJobConfigServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NIJobConfigModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	httpRes, err := r.client.DiscoveryControllerAPI.
		NIJobConfigServiceAPI.
		Delete(ctx, data.Id.ValueString()).
		Execute()
	if err != nil {
		if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete NIJobConfigService, got error: %s", err))
		return
	}
}

func (r *NIJobConfigServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
