// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceAppManagementTaskUpdateStatusRequestParameter undocumented
type DeviceAppManagementTaskUpdateStatusRequestParameter struct {
	// Status undocumented
	Status *DeviceAppManagementTaskStatus `json:"status,omitempty"`
	// Note undocumented
	Note *string `json:"note,omitempty"`
}

//
type DeviceAppManagementTaskUpdateStatusRequestBuilder struct{ BaseRequestBuilder }

// UpdateStatus action undocumented
func (b *DeviceAppManagementTaskRequestBuilder) UpdateStatus(reqObj *DeviceAppManagementTaskUpdateStatusRequestParameter) *DeviceAppManagementTaskUpdateStatusRequestBuilder {
	bb := &DeviceAppManagementTaskUpdateStatusRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updateStatus"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceAppManagementTaskUpdateStatusRequest struct{ BaseRequest }

//
func (b *DeviceAppManagementTaskUpdateStatusRequestBuilder) Request() *DeviceAppManagementTaskUpdateStatusRequest {
	return &DeviceAppManagementTaskUpdateStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceAppManagementTaskUpdateStatusRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}