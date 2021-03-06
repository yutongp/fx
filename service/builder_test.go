// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package service

import (
	"errors"
	"testing"

	. "go.uber.org/fx/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewBuilder_NoConfig(t *testing.T) {
	_, err := NewBuilder().Build()
	assert.Error(t, err)
}

func TestNewBuilder_WithConfig(t *testing.T) {
	b := NewBuilder(
		WithConfiguration(StaticAppData(nil)),
	)

	svc, err := b.Build()
	require.NoError(t, err)
	assert.NotEmpty(t, svc.Name())
}

func TestBuilder_WithModules(t *testing.T) {
	_, err := NewBuilder(
		WithConfiguration(StaticAppData(nil)),
	).WithModules(nopModule).Build()
	assert.NoError(t, err)
}

func TestBuilder_WithErrModule(t *testing.T) {
	_, err := NewBuilder(
		WithConfiguration(StaticAppData(nil)),
	).WithModules(errModule).Build()
	assert.Error(t, err)
}

func TestBuilder_SkipsModulesBadInit(t *testing.T) {
	empty := ""
	_, err := NewBuilder(
		WithConfiguration(StaticAppData(&empty)),
	).WithModules(nopModule).Build()
	assert.Error(t, err, "Expected service name to be provided")
}

func TestWithModules_OK(t *testing.T) {
	_, err := WithModules(nopModule).WithOptions(
		WithConfiguration(StaticAppData(nil)),
	).Build()
	assert.NoError(t, err)
}

func nopModule(_ ModuleCreateInfo) ([]Module, error) {
	return nil, nil
}

func errModule(_ ModuleCreateInfo) ([]Module, error) {
	return nil, errors.New("intentional module creation failure")
}
