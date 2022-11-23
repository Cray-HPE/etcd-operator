// Copyright 2016 The etcd-operator Authors
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

package cluster

import (
	"context"
	"fmt"

	"github.com/coreos/etcd-operator/pkg/util/k8sutil"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func (c *Cluster) upgradeOneMember(memberName string) error {
	c.status.SetUpgradingCondition(c.cluster.Spec.Version)

	ns := c.cluster.Namespace

	cxt := context.TODO()
	pod, err := c.config.KubeCli.CoreV1().Pods(ns).Get(cxt, memberName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("fail to get pod (%s): %v", memberName, err)
	}
	oldpod := pod.DeepCopy()

	c.logger.Infof("upgrading the etcd member %v from %s to %s", memberName, k8sutil.GetEtcdVersion(pod), c.cluster.Spec.Version)
	pod.Spec.Containers[0].Image = k8sutil.ImageName(c.cluster.Spec.Repository, c.cluster.Spec.Version)
	k8sutil.SetEtcdVersion(pod, c.cluster.Spec.Version)

	patchdata, err := k8sutil.CreatePatch(oldpod, pod, v1.Pod{})
	if err != nil {
		return fmt.Errorf("error creating patch: %v", err)
	}

	cxt_2 := context.TODO()
	var opts metav1.PatchOptions
	_, err = c.config.KubeCli.CoreV1().Pods(ns).Patch(cxt_2, pod.GetName(), types.StrategicMergePatchType, patchdata, opts)
	if err != nil {
		return fmt.Errorf("fail to update the etcd member (%s): %v", memberName, err)
	}
	c.logger.Infof("finished upgrading the etcd member %v", memberName)

	cxt_3 := context.TODO()
	var opts_2 metav1.CreateOptions
	_, err = c.eventsCli.Create(cxt_3, k8sutil.MemberUpgradedEvent(memberName, k8sutil.GetEtcdVersion(oldpod), c.cluster.Spec.Version, c.cluster), opts_2)
	if err != nil {
		c.logger.Errorf("failed to create member upgraded event: %v", err)
	}

	return nil
}
