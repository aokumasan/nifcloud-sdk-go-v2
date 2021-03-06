// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyReplaceRouteTableAssociationWithVpnGatewayInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	AssociationId *string `locationName:"AssociationId" type:"string"`

	RouteTableId *string `locationName:"RouteTableId" type:"string"`
}

// String returns the string representation
func (s NiftyReplaceRouteTableAssociationWithVpnGatewayInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyReplaceRouteTableAssociationWithVpnGatewayOutput struct {
	_ struct{} `type:"structure"`

	NewAssociationId *string `locationName:"newAssociationId" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s NiftyReplaceRouteTableAssociationWithVpnGatewayOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyReplaceRouteTableAssociationWithVpnGateway = "NiftyReplaceRouteTableAssociationWithVpnGateway"

// NiftyReplaceRouteTableAssociationWithVpnGatewayRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyReplaceRouteTableAssociationWithVpnGatewayRequest.
//    req := client.NiftyReplaceRouteTableAssociationWithVpnGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyReplaceRouteTableAssociationWithVpnGateway
func (c *Client) NiftyReplaceRouteTableAssociationWithVpnGatewayRequest(input *NiftyReplaceRouteTableAssociationWithVpnGatewayInput) NiftyReplaceRouteTableAssociationWithVpnGatewayRequest {
	op := &aws.Operation{
		Name:       opNiftyReplaceRouteTableAssociationWithVpnGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyReplaceRouteTableAssociationWithVpnGatewayInput{}
	}

	req := c.newRequest(op, input, &NiftyReplaceRouteTableAssociationWithVpnGatewayOutput{})
	return NiftyReplaceRouteTableAssociationWithVpnGatewayRequest{Request: req, Input: input, Copy: c.NiftyReplaceRouteTableAssociationWithVpnGatewayRequest}
}

// NiftyReplaceRouteTableAssociationWithVpnGatewayRequest is the request type for the
// NiftyReplaceRouteTableAssociationWithVpnGateway API operation.
type NiftyReplaceRouteTableAssociationWithVpnGatewayRequest struct {
	*aws.Request
	Input *NiftyReplaceRouteTableAssociationWithVpnGatewayInput
	Copy  func(*NiftyReplaceRouteTableAssociationWithVpnGatewayInput) NiftyReplaceRouteTableAssociationWithVpnGatewayRequest
}

// Send marshals and sends the NiftyReplaceRouteTableAssociationWithVpnGateway API request.
func (r NiftyReplaceRouteTableAssociationWithVpnGatewayRequest) Send(ctx context.Context) (*NiftyReplaceRouteTableAssociationWithVpnGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyReplaceRouteTableAssociationWithVpnGatewayResponse{
		NiftyReplaceRouteTableAssociationWithVpnGatewayOutput: r.Request.Data.(*NiftyReplaceRouteTableAssociationWithVpnGatewayOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyReplaceRouteTableAssociationWithVpnGatewayResponse is the response type for the
// NiftyReplaceRouteTableAssociationWithVpnGateway API operation.
type NiftyReplaceRouteTableAssociationWithVpnGatewayResponse struct {
	*NiftyReplaceRouteTableAssociationWithVpnGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyReplaceRouteTableAssociationWithVpnGateway request.
func (r *NiftyReplaceRouteTableAssociationWithVpnGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
