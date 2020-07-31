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
	"github.com/aws/aws-sdk-go/aws/request"
	"regexp"
)

type Condition func(r *request.Request) bool

func matchService(serviceID string) Condition {
	return func(r *request.Request) bool {
		return r.ClientInfo.ServiceID == serviceID
	}
}

func matchServiceOperation(serviceID string, operation string) Condition {
	return func(r *request.Request) bool {
		if r.Operation == nil {
			return false
		}
		return r.ClientInfo.ServiceID == serviceID && r.Operation.Name == operation
	}
}

func matchServiceOperationPattern(serviceID string, operationPtn *regexp.Regexp) Condition {
	return func(r *request.Request) bool {
		if r.Operation == nil {
			return false
		}
		return r.ClientInfo.ServiceID == serviceID && operationPtn.Match([]byte(r.Operation.Name))
	}
}
