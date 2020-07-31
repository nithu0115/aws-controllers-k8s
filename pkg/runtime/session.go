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

package runtime

import (
	"github.com/aws/aws-controllers-k8s/pkg/throttle"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pkg/errors"
)

func NewSession(awsThrottleCfg *throttle.ServiceOperationsThrottleConfig) (*session.Session, error) {
	sess := session.Must(session.NewSession(aws.NewConfig()))

	if  &awsThrottleCfg != nil {
		throttler := throttle.NewThrottler(awsThrottleCfg)
		throttler.InjectHandlers(&sess.Handlers)
	}

	metadata := ec2metadata.New(sess)
	region, err := metadata.Region()
	if err != nil {
		return nil, errors.Wrap(err, "failed to introspect region from EC2Metadata, specify --aws-region instead if EC2Metadata is unavailable")
	}

	awsCfg := aws.NewConfig().WithRegion(region).WithSTSRegionalEndpoint(endpoints.RegionalSTSEndpoint)
	sess = sess.Copy(awsCfg)

	return sess, nil
}
