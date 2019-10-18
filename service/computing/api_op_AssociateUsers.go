// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type AssociateUsersInput struct {
	_ struct{} `type:"structure"`

	FunctionName *string `locationName:"FunctionName" type:"string"`

	Users []RequestUsersStruct `locationName:"Users" locationNameList:"member" type:"list"`
}

// String returns the string representation
func (s AssociateUsersInput) String() string {
	return nifcloudutil.Prettify(s)
}

type AssociateUsersOutput struct {
	_ struct{} `type:"structure"`

	AssociateUsersResult *AssociateUsersResult `locationName:"AssociateUsersResult" type:"structure"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`

	Users []UsersMemberItem `locationName:"Users" locationNameList:"member" type:"list"`
}

// String returns the string representation
func (s AssociateUsersOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opAssociateUsers = "AssociateUsers"

// AssociateUsersRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using AssociateUsersRequest.
//    req := client.AssociateUsersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/AssociateUsers
func (c *Client) AssociateUsersRequest(input *AssociateUsersInput) AssociateUsersRequest {
	op := &aws.Operation{
		Name:       opAssociateUsers,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &AssociateUsersInput{}
	}

	req := c.newRequest(op, input, &AssociateUsersOutput{})
	return AssociateUsersRequest{Request: req, Input: input, Copy: c.AssociateUsersRequest}
}

// AssociateUsersRequest is the request type for the
// AssociateUsers API operation.
type AssociateUsersRequest struct {
	*aws.Request
	Input *AssociateUsersInput
	Copy  func(*AssociateUsersInput) AssociateUsersRequest
}

// Send marshals and sends the AssociateUsers API request.
func (r AssociateUsersRequest) Send(ctx context.Context) (*AssociateUsersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateUsersResponse{
		AssociateUsersOutput: r.Request.Data.(*AssociateUsersOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateUsersResponse is the response type for the
// AssociateUsers API operation.
type AssociateUsersResponse struct {
	*AssociateUsersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateUsers request.
func (r *AssociateUsersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}