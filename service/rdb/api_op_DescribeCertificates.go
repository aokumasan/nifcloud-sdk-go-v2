// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeCertificatesInput struct {
	_ struct{} `type:"structure"`

	CertificateIdentifier *string `locationName:"CertificateIdentifier" type:"string"`

	Marker *string `locationName:"Marker" type:"string"`

	MaxRecords *int64 `locationName:"MaxRecords" type:"integer"`
}

// String returns the string representation
func (s DescribeCertificatesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeCertificatesOutput struct {
	_ struct{} `type:"structure"`

	Certificates []Certificate `locationNameList:"Certificate" type:"list"`

	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeCertificatesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeCertificates = "DescribeCertificates"

// DescribeCertificatesRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using DescribeCertificatesRequest.
//    req := client.DescribeCertificatesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/DescribeCertificates
func (c *Client) DescribeCertificatesRequest(input *DescribeCertificatesInput) DescribeCertificatesRequest {
	op := &aws.Operation{
		Name:       opDescribeCertificates,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCertificatesInput{}
	}

	req := c.newRequest(op, input, &DescribeCertificatesOutput{})
	return DescribeCertificatesRequest{Request: req, Input: input, Copy: c.DescribeCertificatesRequest}
}

// DescribeCertificatesRequest is the request type for the
// DescribeCertificates API operation.
type DescribeCertificatesRequest struct {
	*aws.Request
	Input *DescribeCertificatesInput
	Copy  func(*DescribeCertificatesInput) DescribeCertificatesRequest
}

// Send marshals and sends the DescribeCertificates API request.
func (r DescribeCertificatesRequest) Send(ctx context.Context) (*DescribeCertificatesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCertificatesResponse{
		DescribeCertificatesOutput: r.Request.Data.(*DescribeCertificatesOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCertificatesResponse is the response type for the
// DescribeCertificates API operation.
type DescribeCertificatesResponse struct {
	*DescribeCertificatesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCertificates request.
func (r *DescribeCertificatesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
