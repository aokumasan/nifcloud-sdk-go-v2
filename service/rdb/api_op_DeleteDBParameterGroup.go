// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteDBParameterGroupInput struct {
	_ struct{} `type:"structure"`

	DBParameterGroupName *string `locationName:"DBParameterGroupName" type:"string"`
}

// String returns the string representation
func (s DeleteDBParameterGroupInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteDBParameterGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDBParameterGroupOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteDBParameterGroup = "DeleteDBParameterGroup"

// DeleteDBParameterGroupRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using DeleteDBParameterGroupRequest.
//    req := client.DeleteDBParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/DeleteDBParameterGroup
func (c *Client) DeleteDBParameterGroupRequest(input *DeleteDBParameterGroupInput) DeleteDBParameterGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteDBParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBParameterGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteDBParameterGroupOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteDBParameterGroupRequest{Request: req, Input: input, Copy: c.DeleteDBParameterGroupRequest}
}

// DeleteDBParameterGroupRequest is the request type for the
// DeleteDBParameterGroup API operation.
type DeleteDBParameterGroupRequest struct {
	*aws.Request
	Input *DeleteDBParameterGroupInput
	Copy  func(*DeleteDBParameterGroupInput) DeleteDBParameterGroupRequest
}

// Send marshals and sends the DeleteDBParameterGroup API request.
func (r DeleteDBParameterGroupRequest) Send(ctx context.Context) (*DeleteDBParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBParameterGroupResponse{
		DeleteDBParameterGroupOutput: r.Request.Data.(*DeleteDBParameterGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBParameterGroupResponse is the response type for the
// DeleteDBParameterGroup API operation.
type DeleteDBParameterGroupResponse struct {
	*DeleteDBParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBParameterGroup request.
func (r *DeleteDBParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
