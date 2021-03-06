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

package tracing

import (
	"testing"

	"go.uber.org/fx/metrics"
	"go.uber.org/fx/testutils"
	"go.uber.org/fx/ulog"

	"github.com/stretchr/testify/assert"
	"github.com/uber-go/zap"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

var (
	_serviceName            = "serviceName"
	_statsReporter          = metrics.NopCachedStatsReporter
	_emptyJaegerConfig      = &config.Configuration{}
	_disabledJaegerConfig   = &config.Configuration{Disabled: true}
	_jaegerConfigWithLogger = &config.Configuration{Logger: jaeger.NullLogger}
)

func getLogger() ulog.Log {
	return ulog.Logger()
}

func TestInitGlobalTracer_Simple(t *testing.T) {
	tracer, closer, err := InitGlobalTracer(
		_emptyJaegerConfig, _serviceName, getLogger(), _statsReporter,
	)
	defer closer.Close()
	assert.NotNil(t, tracer)
	assert.NotNil(t, closer)
	assert.NoError(t, err)
}

func TestInitGlobalTracer_Disabled(t *testing.T) {
	tracer, closer, err := InitGlobalTracer(
		_disabledJaegerConfig, _serviceName, getLogger(), _statsReporter,
	)
	defer closer.Close()
	assert.NotNil(t, tracer)
	assert.NotNil(t, closer)
	assert.NoError(t, err)
}

func TestInitGlobalTracer_NoServiceName(t *testing.T) {
	tracer, closer, err := InitGlobalTracer(_emptyJaegerConfig, "", getLogger(), _statsReporter)
	assert.Error(t, err)
	assert.Nil(t, tracer)
	assert.Nil(t, closer)
}

func TestLoadAppConfig(t *testing.T) {
	jConfig := loadAppConfig(_emptyJaegerConfig, getLogger())
	assert.NotNil(t, jConfig)
	assert.NotNil(t, jConfig.Logger)
}

func TestLoadAppConfig_JaegerConfigWithLogger(t *testing.T) {
	jConfig := loadAppConfig(_jaegerConfigWithLogger, getLogger())
	assert.NotNil(t, jConfig)
	assert.Equal(t, jaeger.NullLogger, jConfig.Logger)
}

func TestLoadAppConfig_NilJaegerConfig(t *testing.T) {
	jConfig := loadAppConfig(nil, getLogger())
	assert.NotNil(t, jConfig)
	assert.NotNil(t, jConfig.Logger)
}

func TestJaegerLogger(t *testing.T) {
	testutils.WithInMemoryLogger(t, nil, func(zapLogger zap.Logger, buf *testutils.TestBuffer) {
		loggerWithZap := ulog.Builder().SetLogger(zapLogger).Build()
		jLogger := jaegerLogger{log: loggerWithZap}
		jLogger.Infof("info message")
		jLogger.Infof("info message: %s", "oddArg")
		jLogger.Infof("info message: %s %s", "value1", "value2")
		jLogger.Error("error message")
		assert.Equal(t, []string{
			`{"level":"info","msg":"info message"}`,
			`{"level":"info","msg":"info message: oddArg"}`,
			`{"level":"info","msg":"info message: value1 value2"}`,
			`{"level":"error","msg":"error message"}`,
		}, buf.Lines(), "Incorrect output from logger")
	})
}
