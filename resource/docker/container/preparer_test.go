// Copyright © 2016 Asteris, LLC
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

package container_test

import (
	"testing"

	"github.com/asteris-llc/converge/helpers/fakerenderer"
	"github.com/asteris-llc/converge/resource"
	"github.com/asteris-llc/converge/resource/docker/container"
	"github.com/stretchr/testify/assert"
)

// TestPreparerInterface tests that the Preparer interface is properly
// implemented
func TestPreparerInterface(t *testing.T) {
	t.Parallel()
	assert.Implements(t, (*resource.Resource)(nil), new(container.Preparer))
}

// TestPreparerInvalidStatus tests preparer validation
func TestPreparerInvalidStatus(t *testing.T) {
	t.Run("status is invalid", func(t *testing.T) {
		p := &container.Preparer{Name: "test", Image: "nginx", Status: "exited"}
		_, err := p.Prepare(fakerenderer.New())
		if assert.Error(t, err) {
			assert.EqualError(t, err, "status must be 'running' or 'created'")
		}
	})

	t.Run("name is invalid", func(t *testing.T) {
		p := &container.Preparer{Name: "", Image: "nginx"}
		_, err := p.Prepare(fakerenderer.New())
		if assert.Error(t, err) {
			assert.EqualError(t, err, "name must be provided")
		}
	})

	t.Run("image is invalid", func(t *testing.T) {
		p := &container.Preparer{Name: "nginx", Image: ""}
		_, err := p.Prepare(fakerenderer.New())
		if assert.Error(t, err) {
			assert.EqualError(t, err, "image must be provided")
		}
	})
}
