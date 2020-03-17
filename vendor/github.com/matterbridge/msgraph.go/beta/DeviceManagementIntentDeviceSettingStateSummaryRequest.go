// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceManagementIntentDeviceSettingStateSummaryRequestBuilder is request builder for DeviceManagementIntentDeviceSettingStateSummary
type DeviceManagementIntentDeviceSettingStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementIntentDeviceSettingStateSummaryRequest
func (b *DeviceManagementIntentDeviceSettingStateSummaryRequestBuilder) Request() *DeviceManagementIntentDeviceSettingStateSummaryRequest {
	return &DeviceManagementIntentDeviceSettingStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementIntentDeviceSettingStateSummaryRequest is request for DeviceManagementIntentDeviceSettingStateSummary
type DeviceManagementIntentDeviceSettingStateSummaryRequest struct{ BaseRequest }

// Get performs GET request for DeviceManagementIntentDeviceSettingStateSummary
func (r *DeviceManagementIntentDeviceSettingStateSummaryRequest) Get(ctx context.Context) (resObj *DeviceManagementIntentDeviceSettingStateSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceManagementIntentDeviceSettingStateSummary
func (r *DeviceManagementIntentDeviceSettingStateSummaryRequest) Update(ctx context.Context, reqObj *DeviceManagementIntentDeviceSettingStateSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceManagementIntentDeviceSettingStateSummary
func (r *DeviceManagementIntentDeviceSettingStateSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}