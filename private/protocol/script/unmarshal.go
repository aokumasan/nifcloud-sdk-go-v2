package script

import (
	"encoding/xml"
	"io"

	"github.com/aokumasan/nifcloud-sdk-go-v2/private/protocol/xml/xmlutil"
	request "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
)

// UnmarshalHandler is a named request handler for unmarshaling script protocol requests
var UnmarshalHandler = request.NamedHandler{Name: "nifcloudsdk.script.Unmarshal", Fn: Unmarshal}

// UnmarshalMetaHandler is a named request handler for unmarshaling script protocol request metadata
var UnmarshalMetaHandler = request.NamedHandler{Name: "nifcloudsdk.script.UnmarshalMeta", Fn: UnmarshalMeta}

// UnmarshalErrorHandler is a named request handler for unmarshaling script protocol request errors
var UnmarshalErrorHandler = request.NamedHandler{Name: "nifcloudsdk.script.UnmarshalError", Fn: UnmarshalError}

// Unmarshal unmarshals a response body for the script protocol.
func Unmarshal(r *request.Request) {
	defer r.HTTPResponse.Body.Close()
	decoder := xml.NewDecoder(r.HTTPResponse.Body)
	err := xmlutil.UnmarshalXML(r.Data, decoder, r.Operation.Name+"Result")
	if err != nil {
		r.Error = awserr.New("SerializationError", "failed decoding Script response", err)
		return
	}
}

// UnmarshalMeta unmarshals response headers for the script protocol.
func UnmarshalMeta(r *request.Request) {
	// TODO implement unmarshaling of request IDs
}

type xmlErrorResponse struct {
	XMLName   xml.Name `xml:"ErrorResponse"`
	Code      string   `xml:"Error>Code"`
	Message   string   `xml:"Error>Message"`
	RequestID string   `xml:"requestId"`
}

// UnmarshalError unmarshals a response error for the script protocol.
func UnmarshalError(r *request.Request) {
	defer r.HTTPResponse.Body.Close()

	resp := &xmlErrorResponse{}
	err := xml.NewDecoder(r.HTTPResponse.Body).Decode(resp)
	if err != nil && err != io.EOF {
		r.Error = awserr.New("SerializationError", "failed decoding Script error response", err)
	} else {
		r.Error = awserr.NewRequestFailure(
			awserr.New(resp.Code, resp.Message, nil),
			r.HTTPResponse.StatusCode,
			resp.RequestID,
		)
	}
}
