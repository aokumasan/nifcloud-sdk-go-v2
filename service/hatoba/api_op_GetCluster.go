// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetClusterInput struct {
	_ struct{} `type:"structure"`

	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"ClusterName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetClusterInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetClusterInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetClusterInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ClusterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetClusterOutput struct {
	_ struct{} `type:"structure"`

	Cluster *Cluster `locationName:"cluster" type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s GetClusterOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetClusterOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Cluster != nil {
		v := s.Cluster

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cluster", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetCluster = "GetCluster"

// GetClusterRequest returns a request value for making API operation for
// NIFCLOUD Hatoba beta.
//
//    // Example sending a request using GetClusterRequest.
//    req := client.GetClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/hatoba-2019-03-27/GetCluster
func (c *Client) GetClusterRequest(input *GetClusterInput) GetClusterRequest {
	op := &aws.Operation{
		Name:       opGetCluster,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/clusters/{ClusterName}",
	}

	if input == nil {
		input = &GetClusterInput{}
	}

	req := c.newRequest(op, input, &GetClusterOutput{})
	return GetClusterRequest{Request: req, Input: input, Copy: c.GetClusterRequest}
}

// GetClusterRequest is the request type for the
// GetCluster API operation.
type GetClusterRequest struct {
	*aws.Request
	Input *GetClusterInput
	Copy  func(*GetClusterInput) GetClusterRequest
}

// Send marshals and sends the GetCluster API request.
func (r GetClusterRequest) Send(ctx context.Context) (*GetClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetClusterResponse{
		GetClusterOutput: r.Request.Data.(*GetClusterOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetClusterResponse is the response type for the
// GetCluster API operation.
type GetClusterResponse struct {
	*GetClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCluster request.
func (r *GetClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}