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

package stats

import "github.com/uber-go/tally"

const (
	// TagModule is module tag for metrics
	TagModule = "module"
	// TagType is either request or response
	TagType = "type"
)

// HTTPTags creates metrics scope with defined tags
var (
	// TaskTags creates metrics scope with defined tags
	TaskTags = map[string]string{
		TagModule: "task",
	}

	// TaskExecutionCount counts number of executions
	TaskExecutionCount tally.Counter
	// TaskPublishCount counts number of tasks enqueued
	TaskPublishCount tally.Counter
	// TaskExecuteFail counts number of tasks failed to execute
	TaskExecuteFail tally.Counter
	// TaskPublishFail counts number of tasks failed to enqueue
	TaskPublishFail tally.Counter
	// TaskExecutionTime is a turnaround time for execution
	TaskExecutionTime tally.Timer
	// TaskPublishTime is a publish time for tasks
	TaskPublishTime tally.Timer
)

// SetupTaskMetrics allocates counters for necessary setup
func SetupTaskMetrics(scope tally.Scope) {
	taskTagsScope := scope.Tagged(TaskTags)
	TaskExecutionCount = taskTagsScope.Tagged(map[string]string{TagType: "execution"}).Counter("count")
	TaskPublishCount = taskTagsScope.Tagged(map[string]string{TagType: "publish"}).Counter("count")
	TaskExecuteFail = taskTagsScope.Tagged(map[string]string{TagType: "execution"}).Counter("fail")
	TaskPublishFail = taskTagsScope.Tagged(map[string]string{TagType: "publish"}).Counter("fail")

	TaskExecutionTime = taskTagsScope.Tagged(map[string]string{TagType: "execution"}).Timer("time")
	TaskPublishTime = taskTagsScope.Tagged(map[string]string{TagType: "publish"}).Timer("time")
}
