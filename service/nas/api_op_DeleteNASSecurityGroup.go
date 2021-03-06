// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package nas

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteNASSecurityGroupInput struct {
	_ struct{} `type:"structure"`

	NASSecurityGroupName *string `locationName:"NASSecurityGroupName" type:"string"`
}

// String returns the string representation
func (s DeleteNASSecurityGroupInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteNASSecurityGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteNASSecurityGroupOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteNASSecurityGroup = "DeleteNASSecurityGroup"

// DeleteNASSecurityGroupRequest returns a request value for making API operation for
// NIFCLOUD NAS.
//
//    // Example sending a request using DeleteNASSecurityGroupRequest.
//    req := client.DeleteNASSecurityGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/nas-N2016-02-24/DeleteNASSecurityGroup
func (c *Client) DeleteNASSecurityGroupRequest(input *DeleteNASSecurityGroupInput) DeleteNASSecurityGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteNASSecurityGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNASSecurityGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteNASSecurityGroupOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteNASSecurityGroupRequest{Request: req, Input: input, Copy: c.DeleteNASSecurityGroupRequest}
}

// DeleteNASSecurityGroupRequest is the request type for the
// DeleteNASSecurityGroup API operation.
type DeleteNASSecurityGroupRequest struct {
	*aws.Request
	Input *DeleteNASSecurityGroupInput
	Copy  func(*DeleteNASSecurityGroupInput) DeleteNASSecurityGroupRequest
}

// Send marshals and sends the DeleteNASSecurityGroup API request.
func (r DeleteNASSecurityGroupRequest) Send(ctx context.Context) (*DeleteNASSecurityGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNASSecurityGroupResponse{
		DeleteNASSecurityGroupOutput: r.Request.Data.(*DeleteNASSecurityGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNASSecurityGroupResponse is the response type for the
// DeleteNASSecurityGroup API operation.
type DeleteNASSecurityGroupResponse struct {
	*DeleteNASSecurityGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNASSecurityGroup request.
func (r *DeleteNASSecurityGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
