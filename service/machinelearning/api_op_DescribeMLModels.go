// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of MLModel that match the search criteria in the request.
func (c *Client) DescribeMLModels(ctx context.Context, params *DescribeMLModelsInput, optFns ...func(*Options)) (*DescribeMLModelsOutput, error) {
	if params == nil {
		params = &DescribeMLModelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMLModels", params, optFns, addOperationDescribeMLModelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMLModelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMLModelsInput struct {

	// The equal to operator. The MLModel results will have FilterVariable values that
	// exactly match the value specified with EQ.
	EQ *string

	// Use one of the following variables to filter a list of MLModel:
	//
	// * CreatedAt -
	// Sets the search criteria to MLModel creation date.
	//
	// * Status - Sets the search
	// criteria to MLModel status.
	//
	// * Name - Sets the search criteria to the contents
	// of MLModelName.
	//
	// * IAMUser - Sets the search criteria to the user account that
	// invoked the MLModel creation.
	//
	// * TrainingDataSourceId - Sets the search criteria
	// to the DataSource used to train one or more MLModel.
	//
	// * RealtimeEndpointStatus -
	// Sets the search criteria to the MLModel real-time endpoint status.
	//
	// *
	// MLModelType - Sets the search criteria to MLModel type: binary, regression, or
	// multi-class.
	//
	// * Algorithm - Sets the search criteria to the algorithm that the
	// MLModel uses.
	//
	// * TrainingDataURI - Sets the search criteria to the data file(s)
	// used in training a MLModel. The URL can identify either a file or an Amazon
	// Simple Storage Service (Amazon S3) bucket or directory.
	FilterVariable types.MLModelFilterVariable

	// The greater than or equal to operator. The MLModel results will have
	// FilterVariable values that are greater than or equal to the value specified with
	// GE.
	GE *string

	// The greater than operator. The MLModel results will have FilterVariable values
	// that are greater than the value specified with GT.
	GT *string

	// The less than or equal to operator. The MLModel results will have FilterVariable
	// values that are less than or equal to the value specified with LE.
	LE *string

	// The less than operator. The MLModel results will have FilterVariable values that
	// are less than the value specified with LT.
	LT *string

	// The number of pages of information to include in the result. The range of
	// acceptable values is 1 through 100. The default value is 100.
	Limit *int32

	// The not equal to operator. The MLModel results will have FilterVariable values
	// not equal to the value specified with NE.
	NE *string

	// The ID of the page in the paginated results.
	NextToken *string

	// A string that is found at the beginning of a variable, such as Name or Id. For
	// example, an MLModel could have the Name2014-09-09-HolidayGiftMailer. To search
	// for this MLModel, select Name for the FilterVariable and any of the following
	// strings for the Prefix:
	//
	// * 2014-09
	//
	// * 2014-09-09
	//
	// * 2014-09-09-Holiday
	Prefix *string

	// A two-value parameter that determines the sequence of the resulting list of
	// MLModel.
	//
	// * asc - Arranges the list in ascending order (A-Z, 0-9).
	//
	// * dsc -
	// Arranges the list in descending order (Z-A, 9-0).
	//
	// Results are sorted by
	// FilterVariable.
	SortOrder types.SortOrder
}

// Represents the output of a DescribeMLModels operation. The content is
// essentially a list of MLModel.
type DescribeMLModelsOutput struct {

	// The ID of the next page in the paginated results that indicates at least one
	// more page follows.
	NextToken *string

	// A list of MLModel that meet the search criteria.
	Results []types.MLModel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMLModelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMLModels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMLModels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMLModels(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMLModels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "DescribeMLModels",
	}
}