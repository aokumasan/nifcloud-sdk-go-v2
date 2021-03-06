// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type RebootDBInstanceInput struct {
	_ struct{} `type:"structure"`

	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string"`

	ForceFailover *bool `locationName:"ForceFailover" type:"boolean"`

	NiftyRebootType *string `locationName:"NiftyRebootType" type:"string"`
}

// String returns the string representation
func (s RebootDBInstanceInput) String() string {
	return nifcloudutil.Prettify(s)
}

type RebootDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s RebootDBInstanceOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opRebootDBInstance = "RebootDBInstance"

// RebootDBInstanceRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using RebootDBInstanceRequest.
//    req := client.RebootDBInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/RebootDBInstance
func (c *Client) RebootDBInstanceRequest(input *RebootDBInstanceInput) RebootDBInstanceRequest {
	op := &aws.Operation{
		Name:       opRebootDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RebootDBInstanceInput{}
	}

	req := c.newRequest(op, input, &RebootDBInstanceOutput{})
	return RebootDBInstanceRequest{Request: req, Input: input, Copy: c.RebootDBInstanceRequest}
}

// RebootDBInstanceRequest is the request type for the
// RebootDBInstance API operation.
type RebootDBInstanceRequest struct {
	*aws.Request
	Input *RebootDBInstanceInput
	Copy  func(*RebootDBInstanceInput) RebootDBInstanceRequest
}

// Send marshals and sends the RebootDBInstance API request.
func (r RebootDBInstanceRequest) Send(ctx context.Context) (*RebootDBInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RebootDBInstanceResponse{
		RebootDBInstanceOutput: r.Request.Data.(*RebootDBInstanceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RebootDBInstanceResponse is the response type for the
// RebootDBInstance API operation.
type RebootDBInstanceResponse struct {
	*RebootDBInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RebootDBInstance request.
func (r *RebootDBInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
