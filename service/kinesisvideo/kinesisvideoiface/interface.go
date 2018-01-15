// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kinesisvideoiface provides an interface to enable mocking the Amazon Kinesis Video Streams service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kinesisvideoiface

import (
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
)

// KinesisVideoAPI provides an interface to enable mocking the
// kinesisvideo.KinesisVideo service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Kinesis Video Streams.
//    func myFunc(svc kinesisvideoiface.KinesisVideoAPI) bool {
//        // Make svc.CreateStream request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := kinesisvideo.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKinesisVideoClient struct {
//        kinesisvideoiface.KinesisVideoAPI
//    }
//    func (m *mockKinesisVideoClient) CreateStream(input *kinesisvideo.CreateStreamInput) (*kinesisvideo.CreateStreamOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKinesisVideoClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type KinesisVideoAPI interface {
	CreateStreamRequest(*kinesisvideo.CreateStreamInput) kinesisvideo.CreateStreamRequest

	DeleteStreamRequest(*kinesisvideo.DeleteStreamInput) kinesisvideo.DeleteStreamRequest

	DescribeStreamRequest(*kinesisvideo.DescribeStreamInput) kinesisvideo.DescribeStreamRequest

	GetDataEndpointRequest(*kinesisvideo.GetDataEndpointInput) kinesisvideo.GetDataEndpointRequest

	ListStreamsRequest(*kinesisvideo.ListStreamsInput) kinesisvideo.ListStreamsRequest

	ListTagsForStreamRequest(*kinesisvideo.ListTagsForStreamInput) kinesisvideo.ListTagsForStreamRequest

	TagStreamRequest(*kinesisvideo.TagStreamInput) kinesisvideo.TagStreamRequest

	UntagStreamRequest(*kinesisvideo.UntagStreamInput) kinesisvideo.UntagStreamRequest

	UpdateDataRetentionRequest(*kinesisvideo.UpdateDataRetentionInput) kinesisvideo.UpdateDataRetentionRequest

	UpdateStreamRequest(*kinesisvideo.UpdateStreamInput) kinesisvideo.UpdateStreamRequest
}

var _ KinesisVideoAPI = (*kinesisvideo.KinesisVideo)(nil)