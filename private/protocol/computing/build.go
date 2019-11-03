// Package computing provides serialization of NIFCLOUD Computing requests and responses.
package computing

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/aokumasan/nifcloud-sdk-go-v2/private/protocol/query/queryutil"
	request "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
)

// BuildHandler is a named request handler for building computing protocol requests
var BuildHandler = request.NamedHandler{Name: "awssdk.computing.Build", Fn: Build}

// Build builds a request for the EC2 protocol.
func Build(r *request.Request) {
	body := url.Values{
		"Action":  {r.Operation.Name},
		"Version": {r.Metadata.APIVersion},
	}
	if err := queryutil.Parse(body, r.Params, true); err != nil {
		r.Error = awserr.New("SerializationError", "failed encoding NIFCLOUD Computing Query request", err)
	}

	// Fix request parameters of DescribeLoadBalancers for NIFCLOUD
	if r.Operation.Name == "DescribeLoadBalancers" {
		parameterLength := reflect.ValueOf(r.Params).Elem().FieldByName("LoadBalancerNames").Len()
		prefix := "LoadBalancerNames"
		for i := 1; i <= parameterLength; i++ {
			body[fmt.Sprintf("%s.LoadBalancerPort.%d", prefix, i)] = body[fmt.Sprintf("%s.member.%d.LoadBalancerPort", prefix, i)]
			body[fmt.Sprintf("%s.InstancePort.%d", prefix, i)] = body[fmt.Sprintf("%s.member.%d.InstancePort", prefix, i)]
			body[fmt.Sprintf("%s.member.%d", prefix, i)] = body[fmt.Sprintf("%s.member.%d.LoadBalancerName", prefix, i)]
			delete(body, fmt.Sprintf("%s.member.%d.LoadBalancerPort", prefix, i))
			delete(body, fmt.Sprintf("%s.member.%d.InstancePort", prefix, i))
			delete(body, fmt.Sprintf("%s.member.%d.LoadBalancerName", prefix, i))
		}
	}

	if r.ExpireTime == 0 {
		r.HTTPRequest.Method = "POST"
		r.HTTPRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
		r.SetBufferBody([]byte(body.Encode()))
	} else { // This is a pre-signed request
		r.HTTPRequest.Method = "GET"
		r.HTTPRequest.URL.RawQuery = body.Encode()
	}
}
