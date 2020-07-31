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
	flag "github.com/spf13/pflag"
	ctrlrt "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

const (
	flagBindPort             = "bind-port"
	flagEnableLeaderElection = "enable-leader-election"
	flagMetricAddr           = "metrics-addr"
	flagEnableDevLogging     = "enable-development-logging"
)

type Config struct {
	BindPort                 int
	MetricsAddr              string
	EnableLeaderElection     bool
	EnableDevelopmentLogging bool
}

func (c *Config) BindFlags() {
	flag.IntVar(
		&c.BindPort, flagBindPort,
		9443,
		"The port the service controller binds to.",
	)
	flag.StringVar(
		&c.MetricsAddr, flagMetricAddr,
		"0.0.0.0:8080",
		"The address the metric endpoint binds to.",
	)
	flag.BoolVar(
		&c.EnableLeaderElection, flagEnableLeaderElection,
		false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.",
	)
	flag.BoolVar(
		&c.EnableDevelopmentLogging, flagEnableDevLogging,
		false,
		"Configures the logger to use a Zap development config (encoder=consoleEncoder,logLevel=Debug,stackTraceLevel=Warn, no sampling), " +
			"otherwise a Zap production config will be used (encoder=jsonEncoder,logLevel=Info,stackTraceLevel=Error), sampling).",
	)
}

func (c *Config) SetupLogger()() {
	zapOptions := zap.Options{
		Development:    c.EnableDevelopmentLogging,
	}
	ctrlrt.SetLogger(zap.New(zap.UseFlagOptions(&zapOptions)))
}
