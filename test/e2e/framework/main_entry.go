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

package framework

import (
	"context"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func MainEntry(m *testing.M) {
	if err := setup(context.Background()); err != nil {
		logrus.Errorf("fail to setup framework: %v", err)
		os.Exit(1)
	}

	code := m.Run()

	if err := teardown(context.Background()); err != nil {
		logrus.Errorf("fail to teardown framework: %v", err)
		os.Exit(1)
	}
	os.Exit(code)
}
