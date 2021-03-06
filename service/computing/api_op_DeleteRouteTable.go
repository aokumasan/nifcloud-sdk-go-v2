// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DeleteRouteTableInput struct {
	_ struct{} `type:"structure"`

	RouteTableId *string `locationName:"RouteTableId" type:"string"`
}

// String returns the string representation
func (s DeleteRouteTableInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteRouteTableOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DeleteRouteTableOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteRouteTable = "DeleteRouteTable"

// DeleteRouteTableRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DeleteRouteTableRequest.
//    req := client.DeleteRouteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DeleteRouteTable
func (c *Client) DeleteRouteTableRequest(input *DeleteRouteTableInput) DeleteRouteTableRequest {
	op := &aws.Operation{
		Name:       opDeleteRouteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DeleteRouteTableInput{}
	}

	req := c.newRequest(op, input, &DeleteRouteTableOutput{})
	return DeleteRouteTableRequest{Request: req, Input: input, Copy: c.DeleteRouteTableRequest}
}

// DeleteRouteTableRequest is the request type for the
// DeleteRouteTable API operation.
type DeleteRouteTableRequest struct {
	*aws.Request
	Input *DeleteRouteTableInput
	Copy  func(*DeleteRouteTableInput) DeleteRouteTableRequest
}

// Send marshals and sends the DeleteRouteTable API request.
func (r DeleteRouteTableRequest) Send(ctx context.Context) (*DeleteRouteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRouteTableResponse{
		DeleteRouteTableOutput: r.Request.Data.(*DeleteRouteTableOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRouteTableResponse is the response type for the
// DeleteRouteTable API operation.
type DeleteRouteTableResponse struct {
	*DeleteRouteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRouteTable request.
func (r *DeleteRouteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
