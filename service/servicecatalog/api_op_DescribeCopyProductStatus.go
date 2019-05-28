// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeCopyProductStatusInput
type DescribeCopyProductStatusInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The token for the copy product operation. This token is returned by CopyProduct.
	//
	// CopyProductToken is a required field
	CopyProductToken *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCopyProductStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCopyProductStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCopyProductStatusInput"}

	if s.CopyProductToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("CopyProductToken"))
	}
	if s.CopyProductToken != nil && len(*s.CopyProductToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CopyProductToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeCopyProductStatusOutput
type DescribeCopyProductStatusOutput struct {
	_ struct{} `type:"structure"`

	// The status of the copy product operation.
	CopyProductStatus CopyProductStatus `type:"string" enum:"true"`

	// The status message.
	StatusDetail *string `type:"string"`

	// The identifier of the copied product.
	TargetProductId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeCopyProductStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCopyProductStatus = "DescribeCopyProductStatus"

// DescribeCopyProductStatusRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets the status of the specified copy product operation.
//
//    // Example sending a request using DescribeCopyProductStatusRequest.
//    req := client.DescribeCopyProductStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeCopyProductStatus
func (c *Client) DescribeCopyProductStatusRequest(input *DescribeCopyProductStatusInput) DescribeCopyProductStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeCopyProductStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCopyProductStatusInput{}
	}

	req := c.newRequest(op, input, &DescribeCopyProductStatusOutput{})
	return DescribeCopyProductStatusRequest{Request: req, Input: input, Copy: c.DescribeCopyProductStatusRequest}
}

// DescribeCopyProductStatusRequest is the request type for the
// DescribeCopyProductStatus API operation.
type DescribeCopyProductStatusRequest struct {
	*aws.Request
	Input *DescribeCopyProductStatusInput
	Copy  func(*DescribeCopyProductStatusInput) DescribeCopyProductStatusRequest
}

// Send marshals and sends the DescribeCopyProductStatus API request.
func (r DescribeCopyProductStatusRequest) Send(ctx context.Context) (*DescribeCopyProductStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCopyProductStatusResponse{
		DescribeCopyProductStatusOutput: r.Request.Data.(*DescribeCopyProductStatusOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCopyProductStatusResponse is the response type for the
// DescribeCopyProductStatus API operation.
type DescribeCopyProductStatusResponse struct {
	*DescribeCopyProductStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCopyProductStatus request.
func (r *DescribeCopyProductStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}