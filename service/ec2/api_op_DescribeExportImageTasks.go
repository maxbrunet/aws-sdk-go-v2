// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified export image tasks or all of your export image tasks.
func (c *Client) DescribeExportImageTasks(ctx context.Context, params *DescribeExportImageTasksInput, optFns ...func(*Options)) (*DescribeExportImageTasksOutput, error) {
	if params == nil {
		params = &DescribeExportImageTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExportImageTasks", params, optFns, addOperationDescribeExportImageTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExportImageTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExportImageTasksInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The IDs of the export image tasks.
	ExportImageTaskIds []string

	// Filter tasks using the task-state filter and one of the following values:
	// active, completed, deleting, or deleted.
	Filters []types.Filter

	// The maximum number of results to return in a single call.
	MaxResults int32

	// A token that indicates the next page of results.
	NextToken *string
}

type DescribeExportImageTasksOutput struct {

	// Information about the export image tasks.
	ExportImageTasks []types.ExportImageTask

	// The token to use to get the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeExportImageTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeExportImageTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeExportImageTasks{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExportImageTasks(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeExportImageTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeExportImageTasks",
	}
}