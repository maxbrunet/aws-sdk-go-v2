// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycluster

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycluster/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Set the state of the routing control to reroute traffic. You can set the value
// to be On or Off. When the state is On, traffic flows to a cell. When it's off,
// traffic does not flow. For more information about working with routing controls,
// see Routing control
// (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html) in the
// Route 53 Application Recovery Controller Developer Guide.
func (c *Client) UpdateRoutingControlState(ctx context.Context, params *UpdateRoutingControlStateInput, optFns ...func(*Options)) (*UpdateRoutingControlStateOutput, error) {
	if params == nil {
		params = &UpdateRoutingControlStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRoutingControlState", params, optFns, c.addOperationUpdateRoutingControlStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRoutingControlStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoutingControlStateInput struct {

	// The Amazon Resource Number (ARN) for the routing control that you want to update
	// the state for.
	//
	// This member is required.
	RoutingControlArn *string

	// The state of the routing control. You can set the value to be On or Off.
	//
	// This member is required.
	RoutingControlState types.RoutingControlState

	noSmithyDocumentSerde
}

type UpdateRoutingControlStateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRoutingControlStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateRoutingControlState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateRoutingControlState{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpUpdateRoutingControlStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRoutingControlState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRoutingControlState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-cluster",
		OperationName: "UpdateRoutingControlState",
	}
}
