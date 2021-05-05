package aws

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/takescoop/terraform-provider-awssso/awssso/internal/tfresource"
)

// RetryOnAwsCodes retries AWS error codes for one minute
// Note: This function will be moved out of the aws package in the future.
func RetryOnAwsCodes(codes []string, f func() (interface{}, error)) (interface{}, error) {
	var resp interface{}
	err := resource.Retry(1*time.Minute, func() *resource.RetryError {
		var err error
		resp, err = f()
		if err != nil {
			awsErr, ok := err.(awserr.Error)
			if ok {
				for _, code := range codes {
					if awsErr.Code() == code {
						return resource.RetryableError(err)
					}
				}
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})

	if tfresource.TimedOut(err) {
		resp, err = f()
	}

	return resp, err
}
