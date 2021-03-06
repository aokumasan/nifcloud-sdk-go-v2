// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyEnableDhcpInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	DhcpConfigId *string `locationName:"DhcpConfigId" type:"string"`

	DhcpOptionsId *string `locationName:"DhcpOptionsId" type:"string"`

	NetworkId *string `locationName:"NetworkId" type:"string"`

	NetworkName *string `locationName:"NetworkName" type:"string"`

	RouterId *string `locationName:"RouterId" type:"string"`

	RouterName *string `locationName:"RouterName" type:"string"`
}

// String returns the string representation
func (s NiftyEnableDhcpInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyEnableDhcpOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyEnableDhcpOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyEnableDhcp = "NiftyEnableDhcp"

// NiftyEnableDhcpRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyEnableDhcpRequest.
//    req := client.NiftyEnableDhcpRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyEnableDhcp
func (c *Client) NiftyEnableDhcpRequest(input *NiftyEnableDhcpInput) NiftyEnableDhcpRequest {
	op := &aws.Operation{
		Name:       opNiftyEnableDhcp,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyEnableDhcpInput{}
	}

	req := c.newRequest(op, input, &NiftyEnableDhcpOutput{})
	return NiftyEnableDhcpRequest{Request: req, Input: input, Copy: c.NiftyEnableDhcpRequest}
}

// NiftyEnableDhcpRequest is the request type for the
// NiftyEnableDhcp API operation.
type NiftyEnableDhcpRequest struct {
	*aws.Request
	Input *NiftyEnableDhcpInput
	Copy  func(*NiftyEnableDhcpInput) NiftyEnableDhcpRequest
}

// Send marshals and sends the NiftyEnableDhcp API request.
func (r NiftyEnableDhcpRequest) Send(ctx context.Context) (*NiftyEnableDhcpResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyEnableDhcpResponse{
		NiftyEnableDhcpOutput: r.Request.Data.(*NiftyEnableDhcpOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyEnableDhcpResponse is the response type for the
// NiftyEnableDhcp API operation.
type NiftyEnableDhcpResponse struct {
	*NiftyEnableDhcpOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyEnableDhcp request.
func (r *NiftyEnableDhcpResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
