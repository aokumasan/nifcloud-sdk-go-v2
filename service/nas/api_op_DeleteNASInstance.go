// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package nas

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DeleteNASInstanceInput struct {
	_ struct{} `type:"structure"`

	DirectoryServiceAdministratorName *string `locationName:"DirectoryServiceAdministratorName" type:"string"`

	DirectoryServiceAdministratorPassword *string `locationName:"DirectoryServiceAdministratorPassword" type:"string"`

	NASInstanceIdentifier *string `locationName:"NASInstanceIdentifier" type:"string"`
}

// String returns the string representation
func (s DeleteNASInstanceInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteNASInstanceOutput struct {
	_ struct{} `type:"structure"`

	NASInstance *NASInstance `type:"structure"`
}

// String returns the string representation
func (s DeleteNASInstanceOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteNASInstance = "DeleteNASInstance"

// DeleteNASInstanceRequest returns a request value for making API operation for
// NIFCLOUD NAS.
//
//    // Example sending a request using DeleteNASInstanceRequest.
//    req := client.DeleteNASInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/nas-N2016-02-24/DeleteNASInstance
func (c *Client) DeleteNASInstanceRequest(input *DeleteNASInstanceInput) DeleteNASInstanceRequest {
	op := &aws.Operation{
		Name:       opDeleteNASInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNASInstanceInput{}
	}

	req := c.newRequest(op, input, &DeleteNASInstanceOutput{})
	return DeleteNASInstanceRequest{Request: req, Input: input, Copy: c.DeleteNASInstanceRequest}
}

// DeleteNASInstanceRequest is the request type for the
// DeleteNASInstance API operation.
type DeleteNASInstanceRequest struct {
	*aws.Request
	Input *DeleteNASInstanceInput
	Copy  func(*DeleteNASInstanceInput) DeleteNASInstanceRequest
}

// Send marshals and sends the DeleteNASInstance API request.
func (r DeleteNASInstanceRequest) Send(ctx context.Context) (*DeleteNASInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteNASInstanceResponse{
		DeleteNASInstanceOutput: r.Request.Data.(*DeleteNASInstanceOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteNASInstanceResponse is the response type for the
// DeleteNASInstance API operation.
type DeleteNASInstanceResponse struct {
	*DeleteNASInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteNASInstance request.
func (r *DeleteNASInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
