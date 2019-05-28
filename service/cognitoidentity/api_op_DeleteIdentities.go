// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentity

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input to the DeleteIdentities action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/DeleteIdentitiesInput
type DeleteIdentitiesInput struct {
	_ struct{} `type:"structure"`

	// A list of 1-60 identities that you want to delete.
	//
	// IdentityIdsToDelete is a required field
	IdentityIdsToDelete []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteIdentitiesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteIdentitiesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteIdentitiesInput"}

	if s.IdentityIdsToDelete == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdentityIdsToDelete"))
	}
	if s.IdentityIdsToDelete != nil && len(s.IdentityIdsToDelete) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdentityIdsToDelete", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Returned in response to a successful DeleteIdentities operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/DeleteIdentitiesResponse
type DeleteIdentitiesOutput struct {
	_ struct{} `type:"structure"`

	// An array of UnprocessedIdentityId objects, each of which contains an ErrorCode
	// and IdentityId.
	UnprocessedIdentityIds []UnprocessedIdentityId `type:"list"`
}

// String returns the string representation
func (s DeleteIdentitiesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteIdentities = "DeleteIdentities"

// DeleteIdentitiesRequest returns a request value for making API operation for
// Amazon Cognito Identity.
//
// Deletes identities from an identity pool. You can specify a list of 1-60
// identities that you want to delete.
//
// You must use AWS Developer credentials to call this API.
//
//    // Example sending a request using DeleteIdentitiesRequest.
//    req := client.DeleteIdentitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-identity-2014-06-30/DeleteIdentities
func (c *Client) DeleteIdentitiesRequest(input *DeleteIdentitiesInput) DeleteIdentitiesRequest {
	op := &aws.Operation{
		Name:       opDeleteIdentities,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteIdentitiesInput{}
	}

	req := c.newRequest(op, input, &DeleteIdentitiesOutput{})
	return DeleteIdentitiesRequest{Request: req, Input: input, Copy: c.DeleteIdentitiesRequest}
}

// DeleteIdentitiesRequest is the request type for the
// DeleteIdentities API operation.
type DeleteIdentitiesRequest struct {
	*aws.Request
	Input *DeleteIdentitiesInput
	Copy  func(*DeleteIdentitiesInput) DeleteIdentitiesRequest
}

// Send marshals and sends the DeleteIdentities API request.
func (r DeleteIdentitiesRequest) Send(ctx context.Context) (*DeleteIdentitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteIdentitiesResponse{
		DeleteIdentitiesOutput: r.Request.Data.(*DeleteIdentitiesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteIdentitiesResponse is the response type for the
// DeleteIdentities API operation.
type DeleteIdentitiesResponse struct {
	*DeleteIdentitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteIdentities request.
func (r *DeleteIdentitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}