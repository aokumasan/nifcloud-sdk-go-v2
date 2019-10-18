// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type CopyDBSnapshotInput struct {
	_ struct{} `type:"structure"`

	SourceDBSnapshotIdentifier *string `locationName:"SourceDBSnapshotIdentifier" type:"string"`

	TargetDBSnapshotIdentifier *string `locationName:"TargetDBSnapshotIdentifier" type:"string"`
}

// String returns the string representation
func (s CopyDBSnapshotInput) String() string {
	return nifcloudutil.Prettify(s)
}

type CopyDBSnapshotOutput struct {
	_ struct{} `type:"structure"`

	DBSnapshot *DBSnapshot `type:"structure"`
}

// String returns the string representation
func (s CopyDBSnapshotOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opCopyDBSnapshot = "CopyDBSnapshot"

// CopyDBSnapshotRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using CopyDBSnapshotRequest.
//    req := client.CopyDBSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/CopyDBSnapshot
func (c *Client) CopyDBSnapshotRequest(input *CopyDBSnapshotInput) CopyDBSnapshotRequest {
	op := &aws.Operation{
		Name:       opCopyDBSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopyDBSnapshotInput{}
	}

	req := c.newRequest(op, input, &CopyDBSnapshotOutput{})
	return CopyDBSnapshotRequest{Request: req, Input: input, Copy: c.CopyDBSnapshotRequest}
}

// CopyDBSnapshotRequest is the request type for the
// CopyDBSnapshot API operation.
type CopyDBSnapshotRequest struct {
	*aws.Request
	Input *CopyDBSnapshotInput
	Copy  func(*CopyDBSnapshotInput) CopyDBSnapshotRequest
}

// Send marshals and sends the CopyDBSnapshot API request.
func (r CopyDBSnapshotRequest) Send(ctx context.Context) (*CopyDBSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CopyDBSnapshotResponse{
		CopyDBSnapshotOutput: r.Request.Data.(*CopyDBSnapshotOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CopyDBSnapshotResponse is the response type for the
// CopyDBSnapshot API operation.
type CopyDBSnapshotResponse struct {
	*CopyDBSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CopyDBSnapshot request.
func (r *CopyDBSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}