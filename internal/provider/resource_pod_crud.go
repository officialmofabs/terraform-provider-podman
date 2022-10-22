package provider

import (
	"context"
	"fmt"

	"github.com/containers/podman/v4/pkg/bindings/pods"
	"github.com/containers/podman/v4/pkg/domain/entities"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/project0/terraform-provider-podman/internal/utils"
)

func (r podResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data podResourceData

	client := r.initClientData(ctx, &data, req.Config.Get, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	podSpec := toPodmanPodSpecGenerator(ctx, data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create
	podCreateResponse, errCreate := pods.CreatePodFromSpec(client, &entities.PodSpec{
		PodSpecGen: *podSpec,
	})
	if errCreate != nil {
		resp.Diagnostics.AddError("Podman client error", fmt.Sprintf("Failed to create pod resource: %s", errCreate.Error()))
		return
	}

	podResponse, err := pods.Inspect(client, podCreateResponse.Id, nil)
	if err != nil {
		resp.Diagnostics.AddError("Podman client error", fmt.Sprintf("Failed to read pod resource after creation: %s", err.Error()))
		return
	}

	// Set state
	resp.Diagnostics.Append(
		resp.State.Set(ctx, fromPodResponse(podResponse, &resp.Diagnostics))...,
	)
}

func (r podResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data podResourceData

	client := r.initClientData(ctx, &data, req.State.Get, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	podResponse, err := pods.Inspect(client, data.Name.Value, nil)
	if err != nil {
		resp.Diagnostics.AddError("Podman client error", fmt.Sprintf("Failed to read pod resource: %s", err.Error()))
		return
	}

	// Set state
	resp.Diagnostics.Append(
		resp.State.Set(ctx, fromPodResponse(podResponse, &resp.Diagnostics))...,
	)
}

// Update is not implemented
func (r podResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	utils.AddUnexpectedError(
		&resp.Diagnostics,
		"Update triggered for a pod resource",
		"pods are immutable resources and cannot be updated, it always needs to be replaced.",
	)
}

func (r podResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data podResourceData

	client := r.initClientData(ctx, &data, req.State.Get, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: Allow force ?
	// TODO: handle report messages
	_, err := pods.Remove(client, data.Name.Value, nil)
	if err != nil {
		resp.Diagnostics.AddError("Podman client error", fmt.Sprintf("Failed to delete pod resource: %s", err.Error()))
	}

	resp.State.RemoveResource(ctx)
}

func (r podResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}
