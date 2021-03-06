// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReleaseHostsRequest
type ReleaseHostsInput struct {
	_ struct{} `type:"structure"`

	// The IDs of the Dedicated Hosts to release.
	//
	// HostIds is a required field
	HostIds []string `locationName:"hostId" locationNameList:"item" type:"list" required:"true"`
}

// String returns the string representation
func (s ReleaseHostsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReleaseHostsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ReleaseHostsInput"}

	if s.HostIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("HostIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReleaseHostsResult
type ReleaseHostsOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the Dedicated Hosts that were successfully released.
	Successful []string `locationName:"successful" locationNameList:"item" type:"list"`

	// The IDs of the Dedicated Hosts that could not be released, including an error
	// message.
	Unsuccessful []UnsuccessfulItem `locationName:"unsuccessful" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s ReleaseHostsOutput) String() string {
	return awsutil.Prettify(s)
}

const opReleaseHosts = "ReleaseHosts"

// ReleaseHostsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// When you no longer want to use an On-Demand Dedicated Host it can be released.
// On-Demand billing is stopped and the host goes into released state. The host
// ID of Dedicated Hosts that have been released can no longer be specified
// in another request, for example, to modify the host. You must stop or terminate
// all instances on a host before it can be released.
//
// When Dedicated Hosts are released, it may take some time for them to stop
// counting toward your limit and you may receive capacity errors when trying
// to allocate new Dedicated Hosts. Wait a few minutes and then try again.
//
// Released hosts still appear in a DescribeHosts response.
//
//    // Example sending a request using ReleaseHostsRequest.
//    req := client.ReleaseHostsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ReleaseHosts
func (c *Client) ReleaseHostsRequest(input *ReleaseHostsInput) ReleaseHostsRequest {
	op := &aws.Operation{
		Name:       opReleaseHosts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReleaseHostsInput{}
	}

	req := c.newRequest(op, input, &ReleaseHostsOutput{})
	return ReleaseHostsRequest{Request: req, Input: input, Copy: c.ReleaseHostsRequest}
}

// ReleaseHostsRequest is the request type for the
// ReleaseHosts API operation.
type ReleaseHostsRequest struct {
	*aws.Request
	Input *ReleaseHostsInput
	Copy  func(*ReleaseHostsInput) ReleaseHostsRequest
}

// Send marshals and sends the ReleaseHosts API request.
func (r ReleaseHostsRequest) Send(ctx context.Context) (*ReleaseHostsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReleaseHostsResponse{
		ReleaseHostsOutput: r.Request.Data.(*ReleaseHostsOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReleaseHostsResponse is the response type for the
// ReleaseHosts API operation.
type ReleaseHostsResponse struct {
	*ReleaseHostsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReleaseHosts request.
func (r *ReleaseHostsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
