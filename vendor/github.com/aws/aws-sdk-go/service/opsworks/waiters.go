// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package opsworks

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// WaitUntilAppExists uses the AWS OpsWorks API operation
// DescribeApps to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *OpsWorks) WaitUntilAppExists(input *DescribeAppsInput) error {
	return c.WaitUntilAppExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilAppExistsWithContext is an extended version of WaitUntilAppExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilAppExistsWithContext(ctx aws.Context, input *DescribeAppsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilAppExists",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(1 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 200,
			},
			{
				State:    request.FailureWaiterState,
				Matcher:  request.StatusWaiterMatch,
				Expected: 400,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			req, _ := c.DescribeAppsRequest(input)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilDeploymentSuccessful uses the AWS OpsWorks API operation
// DescribeDeployments to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *OpsWorks) WaitUntilDeploymentSuccessful(input *DescribeDeploymentsInput) error {
	return c.WaitUntilDeploymentSuccessfulWithContext(aws.BackgroundContext(), input)
}

// WaitUntilDeploymentSuccessfulWithContext is an extended version of WaitUntilDeploymentSuccessful.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilDeploymentSuccessfulWithContext(ctx aws.Context, input *DescribeDeploymentsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilDeploymentSuccessful",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Deployments[].Status",
				Expected: "successful",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Deployments[].Status",
				Expected: "failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			req, _ := c.DescribeDeploymentsRequest(input)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInstanceOnline uses the AWS OpsWorks API operation
// DescribeInstances to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *OpsWorks) WaitUntilInstanceOnline(input *DescribeInstancesInput) error {
	return c.WaitUntilInstanceOnlineWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceOnlineWithContext is an extended version of WaitUntilInstanceOnline.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilInstanceOnlineWithContext(ctx aws.Context, input *DescribeInstancesInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInstanceOnline",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Instances[].Status",
				Expected: "online",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "setup_failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "shutting_down",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "start_failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stopped",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stopping",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "terminating",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "terminated",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stop_failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			req, _ := c.DescribeInstancesRequest(input)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInstanceRegistered uses the AWS OpsWorks API operation
// DescribeInstances to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *OpsWorks) WaitUntilInstanceRegistered(input *DescribeInstancesInput) error {
	return c.WaitUntilInstanceRegisteredWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceRegisteredWithContext is an extended version of WaitUntilInstanceRegistered.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilInstanceRegisteredWithContext(ctx aws.Context, input *DescribeInstancesInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInstanceRegistered",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Instances[].Status",
				Expected: "registered",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "setup_failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "shutting_down",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stopped",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stopping",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "terminating",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "terminated",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stop_failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			req, _ := c.DescribeInstancesRequest(input)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInstanceStopped uses the AWS OpsWorks API operation
// DescribeInstances to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *OpsWorks) WaitUntilInstanceStopped(input *DescribeInstancesInput) error {
	return c.WaitUntilInstanceStoppedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceStoppedWithContext is an extended version of WaitUntilInstanceStopped.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilInstanceStoppedWithContext(ctx aws.Context, input *DescribeInstancesInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInstanceStopped",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Instances[].Status",
				Expected: "stopped",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "booting",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "online",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "pending",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "rebooting",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "requested",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "running_setup",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "setup_failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "start_failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "stop_failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			req, _ := c.DescribeInstancesRequest(input)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilInstanceTerminated uses the AWS OpsWorks API operation
// DescribeInstances to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *OpsWorks) WaitUntilInstanceTerminated(input *DescribeInstancesInput) error {
	return c.WaitUntilInstanceTerminatedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceTerminatedWithContext is an extended version of WaitUntilInstanceTerminated.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *OpsWorks) WaitUntilInstanceTerminatedWithContext(ctx aws.Context, input *DescribeInstancesInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilInstanceTerminated",
		MaxAttempts: 40,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Instances[].Status",
				Expected: "terminated",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundException",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "booting",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "online",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "pending",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "rebooting",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "requested",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "running_setup",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "setup_failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Instances[].Status",
				Expected: "start_failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			req, _ := c.DescribeInstancesRequest(input)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
