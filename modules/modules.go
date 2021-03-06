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

package modules

import (
	"go.uber.org/fx/service"

	"github.com/opentracing/opentracing-go"
	"github.com/uber-go/tally"
)

// A ModuleConfig holds configuration for a mobule
type ModuleConfig struct {
	Roles []string `yaml:"roles"`
}

// ModuleBase is an embeddable base for all UberFx modules
type ModuleBase struct {
	name      string
	host      service.Host
	isRunning bool
	roles     []string
	scope     tally.Scope
	tracer    opentracing.Tracer
}

// NewModuleBase configures a new ModuleBase
func NewModuleBase(
	name string,
	service service.Host,
	roles []string,
) *ModuleBase {
	return &ModuleBase{
		name:   name,
		host:   service,
		roles:  roles,
		scope:  service.Metrics().SubScope(name),
		tracer: service.Tracer(),
	}
}

// Host returns the module's service host
func (mb ModuleBase) Host() service.Host {
	return mb.host
}

// Roles returns the module's roles
func (mb ModuleBase) Roles() []string {
	return mb.roles
}

// Name returns the module's name
func (mb ModuleBase) Name() string {
	return mb.name
}

// Tracer returns the module's service tracer
func (mb ModuleBase) Tracer() opentracing.Tracer {
	return mb.tracer
}
