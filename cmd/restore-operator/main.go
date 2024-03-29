/*
 *
 *  MIT License
 *
 *  (C) Copyright 2022 Hewlett Packard Enterprise Development LP
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a
 *  copy of this software and associated documentation files (the "Software"),
 *  to deal in the Software without restriction, including without limitation
 *  the rights to use, copy, modify, merge, publish, distribute, sublicense,
 *  and/or sell copies of the Software, and to permit persons to whom the
 *  Software is furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included
 *  in all copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
 *  THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR
 *  OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
 *  ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
 *  OTHER DEALINGS IN THE SOFTWARE.
 *
 */
// Copyright 2017 The etcd-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	controller "github.com/coreos/etcd-operator/pkg/controller/restore-operator"
	"github.com/coreos/etcd-operator/pkg/util/constants"
	"github.com/coreos/etcd-operator/pkg/util/k8sutil"
	version "github.com/coreos/etcd-operator/version"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/client-go/tools/record"
)

var (
	namespace   string
	createCRD   bool
	clusterWide bool
	curlImage   string
)

const (
	serviceNameForMyself = "etcd-restore-operator"
	servicePortForMyself = 19999
)

func init() {
	flag.BoolVar(&createCRD, "create-crd", true, "The restore operator will not create the EtcdRestore CRD when this flag is set to false.")
	flag.BoolVar(&clusterWide, "cluster-wide", false, "Enable operator to watch clusters in all namespaces")
	flag.StringVar(&curlImage, "curl-image", "curlimages/curl:latest", "The curl docker image to use during restores")
	flag.Parse()
}

func main() {
	namespace = os.Getenv(constants.EnvOperatorPodNamespace)
	if len(namespace) == 0 {
		logrus.Fatalf("must set env %s", constants.EnvOperatorPodNamespace)
	}
	name := os.Getenv(constants.EnvOperatorPodName)
	if len(name) == 0 {
		logrus.Fatalf("must set env %s", constants.EnvOperatorPodName)
	}
	id, err := os.Hostname()
	if err != nil {
		logrus.Fatalf("failed to get hostname: %v", err)
	}

	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("Go OS/Arch: %s/%s", runtime.GOOS, runtime.GOARCH)
	logrus.Infof("etcd-restore-operator Version: %v", version.Version)
	logrus.Infof("Git SHA: %s", version.GitSHA)
	logrus.Infof("Curl image version: %s", curlImage)

	kubecli := k8sutil.MustNewKubeClient()

	err = createServiceForMyself(context.TODO(), kubecli, name, namespace)
	if err != nil {
		logrus.Fatalf("create service failed: %+v", err)
	}

	rl, err := resourcelock.New(
		resourcelock.EndpointsResourceLock,
		namespace,
		"etcd-restore-operator",
		kubecli.CoreV1(),
		kubecli.CoordinationV1(),
		resourcelock.ResourceLockConfig{
			Identity:      id,
			EventRecorder: createRecorder(kubecli, name, namespace),
		},
	)
	if err != nil {
		logrus.Fatalf("error creating lock: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:          rl,
		LeaseDuration: 15 * time.Second,
		RenewDeadline: 10 * time.Second,
		RetryPeriod:   2 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: run,
			OnStoppedLeading: func() {
				logrus.Fatalf("leader election lost")
			},
		},
	})
}

func createRecorder(kubecli kubernetes.Interface, name, namespace string) record.EventRecorder {
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(logrus.Infof)
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: v1core.New(kubecli.CoreV1().RESTClient()).Events(namespace)})
	return eventBroadcaster.NewRecorder(scheme.Scheme, v1.EventSource{Component: name})
}

func run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	c := controller.New(newControllerConfig())
	err := c.Start(ctx)
	if err != nil {
		logrus.Fatalf("etcd restore operator stopped with error: %v", err)
	}
}

func newControllerConfig() controller.Config {
	cfg := controller.Config{
		Namespace:   namespace,
		ClusterWide: clusterWide,
		CreateCRD:   createCRD,
		MySvcAddr:   fmt.Sprintf("%s.%s:%d", serviceNameForMyself, namespace, servicePortForMyself),
		CurlImage:   curlImage,
	}

	return cfg
}
