// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// IOSCertificateProfileBaseRequestBuilder is request builder for IOSCertificateProfileBase
type IOSCertificateProfileBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSCertificateProfileBaseRequest
func (b *IOSCertificateProfileBaseRequestBuilder) Request() *IOSCertificateProfileBaseRequest {
	return &IOSCertificateProfileBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSCertificateProfileBaseRequest is request for IOSCertificateProfileBase
type IOSCertificateProfileBaseRequest struct{ BaseRequest }

// Get performs GET request for IOSCertificateProfileBase
func (r *IOSCertificateProfileBaseRequest) Get(ctx context.Context) (resObj *IOSCertificateProfileBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSCertificateProfileBase
func (r *IOSCertificateProfileBaseRequest) Update(ctx context.Context, reqObj *IOSCertificateProfileBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSCertificateProfileBase
func (r *IOSCertificateProfileBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}