// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyModifyAddressAttributeInput struct {
	_ struct{} `type:"structure"`

	Attribute *string `locationName:"Attribute" type:"string"`

	PrivateIpAddress *string `locationName:"PrivateIpAddress" type:"string"`

	PublicIp *string `locationName:"PublicIp" type:"string"`

	Value *string `locationName:"Value" type:"string"`
}

// String returns the string representation
func (s NiftyModifyAddressAttributeInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyModifyAddressAttributeOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyModifyAddressAttributeOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyModifyAddressAttribute = "NiftyModifyAddressAttribute"

// NiftyModifyAddressAttributeRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyModifyAddressAttributeRequest.
//    req := client.NiftyModifyAddressAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyModifyAddressAttribute
func (c *Client) NiftyModifyAddressAttributeRequest(input *NiftyModifyAddressAttributeInput) NiftyModifyAddressAttributeRequest {
	op := &aws.Operation{
		Name:       opNiftyModifyAddressAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyModifyAddressAttributeInput{}
	}

	req := c.newRequest(op, input, &NiftyModifyAddressAttributeOutput{})
	return NiftyModifyAddressAttributeRequest{Request: req, Input: input, Copy: c.NiftyModifyAddressAttributeRequest}
}

// NiftyModifyAddressAttributeRequest is the request type for the
// NiftyModifyAddressAttribute API operation.
type NiftyModifyAddressAttributeRequest struct {
	*aws.Request
	Input *NiftyModifyAddressAttributeInput
	Copy  func(*NiftyModifyAddressAttributeInput) NiftyModifyAddressAttributeRequest
}

// Send marshals and sends the NiftyModifyAddressAttribute API request.
func (r NiftyModifyAddressAttributeRequest) Send(ctx context.Context) (*NiftyModifyAddressAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyModifyAddressAttributeResponse{
		NiftyModifyAddressAttributeOutput: r.Request.Data.(*NiftyModifyAddressAttributeOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyModifyAddressAttributeResponse is the response type for the
// NiftyModifyAddressAttribute API operation.
type NiftyModifyAddressAttributeResponse struct {
	*NiftyModifyAddressAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyModifyAddressAttribute request.
func (r *NiftyModifyAddressAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}