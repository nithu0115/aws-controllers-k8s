// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package throttle

import (
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/aws/aws-sdk-go/service/ecr"
	"golang.org/x/time/rate"
	"regexp"
)

// NewDefaultServiceOperationsThrottleConfig returns a ServiceOperationsThrottleConfig with default settings.
func NewDefaultServiceOperationsThrottleConfig() *ServiceOperationsThrottleConfig {
	return &ServiceOperationsThrottleConfig{
		Value: map[string][]Config{
			apigatewayv2.ServiceID: {
				{
					OperationPtn: regexp.MustCompile("^Describe|List"),
					Rate:         rate.Limit(40),
					Burst:        5,
				},
				{
					OperationPtn: regexp.MustCompile("^Create|Update|Delete"),
					Rate:         rate.Limit(8),
					Burst:        5,
				},
			},
			ecr.ServiceID: {
				{
					OperationPtn: regexp.MustCompile("^Describe|List"),
					Rate:         rate.Limit(40),
					Burst:        5,
				},
				{
					OperationPtn: regexp.MustCompile("^Create|Delete"),
					Rate:         rate.Limit(8),
					Burst:        5,
				},
			},
		},
	}
}
