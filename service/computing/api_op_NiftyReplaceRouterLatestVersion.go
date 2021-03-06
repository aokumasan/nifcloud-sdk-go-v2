// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyReplaceRouterLatestVersionInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	RouterId *string `locationName:"RouterId" type:"string"`

	RouterName *string `locationName:"RouterName" type:"string"`
}

// String returns the string representation
func (s NiftyReplaceRouterLatestVersionInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyReplaceRouterLatestVersionOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyReplaceRouterLatestVersionOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyReplaceRouterLatestVersion = "NiftyReplaceRouterLatestVersion"

// NiftyReplaceRouterLatestVersionRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyReplaceRouterLatestVersionRequest.
//    req := client.NiftyReplaceRouterLatestVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyReplaceRouterLatestVersion
func (c *Client) NiftyReplaceRouterLatestVersionRequest(input *NiftyReplaceRouterLatestVersionInput) NiftyReplaceRouterLatestVersionRequest {
	op := &aws.Operation{
		Name:       opNiftyReplaceRouterLatestVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyReplaceRouterLatestVersionInput{}
	}

	req := c.newRequest(op, input, &NiftyReplaceRouterLatestVersionOutput{})
	return NiftyReplaceRouterLatestVersionRequest{Request: req, Input: input, Copy: c.NiftyReplaceRouterLatestVersionRequest}
}

// NiftyReplaceRouterLatestVersionRequest is the request type for the
// NiftyReplaceRouterLatestVersion API operation.
type NiftyReplaceRouterLatestVersionRequest struct {
	*aws.Request
	Input *NiftyReplaceRouterLatestVersionInput
	Copy  func(*NiftyReplaceRouterLatestVersionInput) NiftyReplaceRouterLatestVersionRequest
}

// Send marshals and sends the NiftyReplaceRouterLatestVersion API request.
func (r NiftyReplaceRouterLatestVersionRequest) Send(ctx context.Context) (*NiftyReplaceRouterLatestVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyReplaceRouterLatestVersionResponse{
		NiftyReplaceRouterLatestVersionOutput: r.Request.Data.(*NiftyReplaceRouterLatestVersionOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyReplaceRouterLatestVersionResponse is the response type for the
// NiftyReplaceRouterLatestVersion API operation.
type NiftyReplaceRouterLatestVersionResponse struct {
	*NiftyReplaceRouterLatestVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyReplaceRouterLatestVersion request.
func (r *NiftyReplaceRouterLatestVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
