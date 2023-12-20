// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package serialization

import (
	"time"

	"github.com/uber/cadence/common/types"
)

var (
	zeroUnix = time.Unix(0, 0)
)

var expectedNil = map[string]map[string]any{
	"*serialization.WorkflowExecutionInfo": {
		"GetAutoResetPoints":                    []uint8(nil),
		"GetAutoResetPointsEncoding":            "",
		"GetCancelRequestID":                    "",
		"GetCancelRequested":                    false,
		"GetClientFeatureVersion":               "",
		"GetClientImpl":                         "",
		"GetClientLibraryVersion":               "",
		"GetCloseStatus":                        int32(0),
		"GetCompletionEvent":                    []uint8(nil),
		"GetCompletionEventEncoding":            "",
		"GetCreateRequestID":                    "",
		"GetCompletionEventBatchID":             int64(0),
		"GetCronSchedule":                       "",
		"GetDecisionAttempt":                    int64(0),
		"GetDecisionOriginalScheduledTimestamp": zeroUnix,
		"GetDecisionRequestID":                  "",
		"GetDecisionScheduleID":                 int64(0),
		"GetDecisionScheduledTimestamp":         zeroUnix,
		"GetDecisionStartedID":                  int64(0),
		"GetDecisionStartedTimestamp":           zeroUnix,
		"GetDecisionTaskTimeout":                time.Duration(0),
		"GetDecisionTimeout":                    time.Duration(0),
		"GetLastFirstEventID":                   int64(0),
		"GetLastProcessedEvent":                 int64(0),
		"GetParentWorkflowID":                   "",
		"GetStartVersion":                       int64(0),
		"GetTaskList":                           "",
		"GetDecisionVersion":                    int64(0),
		"GetEventBranchToken":                   []uint8(nil),
		"GetEventStoreVersion":                  int32(0),
		"GetExecutionContext":                   []uint8(nil),
		"GetFirstExecutionRunID":                []uint8(nil),
		"GetHasRetryPolicy":                     false,
		"GetInitiatedID":                        int64(0),
		"GetIsCron":                             false,
		"GetLastEventTaskID":                    int64(0),
		"GetLastUpdatedTimestamp":               zeroUnix,
		"GetLastWriteEventID":                   int64(0),
		"GetMemo":                               map[string][]uint8(nil),
		"GetParentDomainID":                     []uint8(nil),
		"GetParentRunID":                        []uint8(nil),
		"GetPartitionConfig":                    map[string]string(nil),
		"GetHistorySize":                        int64(0),
		"GetRetryAttempt":                       int64(0),
		"GetRetryBackoffCoefficient":            float64(0),
		"GetRetryExpiration":                    time.Duration(0),
		"GetRetryExpirationTimestamp":           zeroUnix,
		"GetRetryInitialInterval":               time.Duration(0),
		"GetRetryMaximumAttempts":               int32(0),
		"GetRetryMaximumInterval":               time.Duration(0),
		"GetRetryNonRetryableErrors":            []string(nil),
		"GetSearchAttributes":                   map[string][]uint8(nil),
		"GetSignalCount":                        int64(0),
		"GetStartTimestamp":                     zeroUnix,
		"GetState":                              int32(0),
		"GetStickyScheduleToStartTimeout":       time.Duration(0),
		"GetStickyTaskList":                     "",
		"GetVersionHistories":                   []uint8(nil),
		"GetVersionHistoriesEncoding":           "",
		"GetWorkflowTimeout":                    time.Duration(0),
		"GetWorkflowTypeName":                   "",
	},
	"*serialization.TimerTaskInfo": {
		"GetDomainID":        []uint8(nil),
		"GetEventID":         int64(0),
		"GetRunID":           []uint8(nil),
		"GetScheduleAttempt": int64(0),
		"GetTaskType":        int16(0),
		"GetTimeoutType":     int16(0),
		"GetVersion":         int64(0),
		"GetWorkflowID":      "",
	},
	"*serialization.ReplicationTaskInfo": {
		"GetBranchToken":             []uint8(nil),
		"GetCreationTimestamp":       zeroUnix,
		"GetDomainID":                []uint8(nil),
		"GetEventStoreVersion":       int32(0),
		"GetFirstEventID":            int64(0),
		"GetNewRunBranchToken":       []uint8(nil),
		"GetNewRunEventStoreVersion": int32(0),
		"GetNextEventID":             int64(0),
		"GetRunID":                   []uint8(nil),
		"GetScheduledID":             int64(0),
		"GetTaskType":                int16(0),
		"GetVersion":                 int64(0),
		"GetWorkflowID":              "",
	},
	"*serialization.TaskListInfo": {
		"GetAckLevel":        int64(0),
		"GetExpiryTimestamp": zeroUnix,
		"GetKind":            int16(0),
		"GetLastUpdated":     zeroUnix,
	},
	"*serialization.TaskInfo": {
		"GetCreatedTimestamp": zeroUnix,
		"GetExpiryTimestamp":  zeroUnix,
		"GetPartitionConfig":  map[string]string(nil),
		"GetRunID":            []uint8(nil),
		"GetScheduleID":       int64(0),
		"GetWorkflowID":       "",
	},
	"*serialization.TimerInfo": {
		"GetExpiryTimestamp": zeroUnix,
		"GetStartedID":       int64(0),
		"GetTaskID":          int64(0),
		"GetVersion":         int64(0),
	},
	"*serialization.RequestCancelInfo": {
		"GetCancelRequestID":       "",
		"GetInitiatedEventBatchID": int64(0),
		"GetVersion":               int64(0),
	},
	"*serialization.SignalInfo": {
		"GetControl":               []uint8(nil),
		"GetInitiatedEventBatchID": int64(0),
		"GetInput":                 []uint8(nil),
		"GetName":                  "",
		"GetRequestID":             "",
		"GetVersion":               int64(0),
	},
	"*serialization.ChildExecutionInfo": {
		"GetCreateRequestID":        "",
		"GetDomainID":               "",
		"GetDomainNameDEPRECATED":   "",
		"GetInitiatedEvent":         []uint8(nil),
		"GetInitiatedEventBatchID":  int64(0),
		"GetInitiatedEventEncoding": "",
		"GetParentClosePolicy":      int32(0),
		"GetStartedEvent":           []uint8(nil),
		"GetStartedEventEncoding":   "",
		"GetStartedID":              int64(0),
		"GetStartedRunID":           []uint8(nil),
		"GetStartedWorkflowID":      "",
		"GetVersion":                int64(0),
		"GetWorkflowTypeName":       "",
	},
	"*serialization.ActivityInfo": {
		"GetActivityID":               "",
		"GetAttempt":                  int32(0),
		"GetCancelRequestID":          int64(0),
		"GetCancelRequested":          false,
		"GetHasRetryPolicy":           false,
		"GetHeartbeatTimeout":         time.Duration(0),
		"GetRequestID":                "",
		"GetRetryBackoffCoefficient":  float64(0),
		"GetRetryExpirationTimestamp": zeroUnix,
		"GetRetryInitialInterval":     time.Duration(0),
		"GetRetryLastFailureDetails":  []uint8(nil),
		"GetRetryLastFailureReason":   "",
		"GetRetryLastWorkerIdentity":  "",
		"GetRetryMaximumAttempts":     int32(0),
		"GetRetryMaximumInterval":     time.Duration(0),
		"GetRetryNonRetryableErrors":  []string(nil),
		"GetScheduleToCloseTimeout":   time.Duration(0),
		"GetScheduleToStartTimeout":   time.Duration(0),
		"GetScheduledEvent":           []uint8(nil),
		"GetScheduledEventBatchID":    int64(0),
		"GetScheduledEventEncoding":   "",
		"GetScheduledTimestamp":       zeroUnix,
		"GetStartToCloseTimeout":      time.Duration(0),
		"GetStartedEvent":             []uint8(nil),
		"GetStartedEventEncoding":     "",
		"GetStartedID":                int64(0),
		"GetStartedIdentity":          "",
		"GetStartedTimestamp":         zeroUnix,
		"GetTaskList":                 "",
		"GetTimerTaskStatus":          int32(0),
		"GetVersion":                  int64(0),
	},
	"*serialization.HistoryTreeInfo": {
		"GetAncestors":        []*types.HistoryBranchRange(nil),
		"GetCreatedTimestamp": zeroUnix,
		"GetInfo":             "",
	},
	"*serialization.DomainInfo": {
		"GetActiveClusterName":           "",
		"GetArchivalBucket":              "",
		"GetArchivalStatus":              int16(0),
		"GetBadBinaries":                 []uint8(nil),
		"GetBadBinariesEncoding":         "",
		"GetClusters":                    []string(nil),
		"GetConfigVersion":               int64(0),
		"GetData":                        map[string]string(nil),
		"GetDescription":                 "",
		"GetEmitMetric":                  false,
		"GetFailoverEndTimestamp":        zeroUnix,
		"GetFailoverNotificationVersion": int64(0),
		"GetFailoverVersion":             int64(0),
		"GetHistoryArchivalStatus":       int16(0),
		"GetHistoryArchivalURI":          "",
		"GetLastUpdatedTimestamp":        zeroUnix,
		"GetName":                        "",
		"GetNotificationVersion":         int64(0),
		"GetOwner":                       "",
		"GetPreviousFailoverVersion":     int64(0),
		"GetRetention":                   time.Duration(0),
		"GetStatus":                      int32(0),
		"GetVisibilityArchivalStatus":    int16(0),
		"GetVisibilityArchivalURI":       "",
	},
	"*serialization.ShardInfo": {
		"GetClusterReplicationLevel":                   map[string]int64(nil),
		"GetClusterTimerAckLevel":                      map[string]time.Time(nil),
		"GetClusterTransferAckLevel":                   map[string]int64(nil),
		"GetCrossClusterProcessingQueueStates":         []uint8(nil),
		"GetCrossClusterProcessingQueueStatesEncoding": "",
		"GetDomainNotificationVersion":                 int64(0),
		"GetOwner":                                     "",
		"GetPendingFailoverMarkers":                    []uint8(nil),
		"GetPendingFailoverMarkersEncoding":            "",
		"GetReplicationAckLevel":                       int64(0),
		"GetReplicationDlqAckLevel":                    map[string]int64(nil),
		"GetStolenSinceRenew":                          int32(0),
		"GetTimerAckLevel":                             zeroUnix,
		"GetTimerProcessingQueueStates":                []uint8(nil),
		"GetTimerProcessingQueueStatesEncoding":        "",
		"GetTransferAckLevel":                          int64(0),
		"GetTransferProcessingQueueStates":             []uint8(nil),
		"GetTransferProcessingQueueStatesEncoding":     "",
		"GetUpdatedAt":                                 zeroUnix,
	},
}

var expectedEmpty = map[string]map[string]any{
	"*serialization.WorkflowExecutionInfo": {
		"GetAutoResetPoints":                    []uint8(nil),
		"GetAutoResetPointsEncoding":            "",
		"GetCancelRequestID":                    "",
		"GetCancelRequested":                    false,
		"GetClientFeatureVersion":               "",
		"GetClientImpl":                         "",
		"GetClientLibraryVersion":               "",
		"GetCloseStatus":                        int32(0),
		"GetCompletionEvent":                    []uint8(nil),
		"GetCompletionEventEncoding":            "",
		"GetCreateRequestID":                    "",
		"GetCompletionEventBatchID":             int64(0),
		"GetCronSchedule":                       "",
		"GetDecisionAttempt":                    int64(0),
		"GetDecisionOriginalScheduledTimestamp": time.Time{},
		"GetDecisionRequestID":                  "",
		"GetDecisionScheduleID":                 int64(0),
		"GetDecisionScheduledTimestamp":         time.Time{},
		"GetDecisionStartedID":                  int64(0),
		"GetDecisionStartedTimestamp":           time.Time{},
		"GetDecisionTaskTimeout":                time.Duration(0),
		"GetDecisionTimeout":                    time.Duration(0),
		"GetLastFirstEventID":                   int64(0),
		"GetLastProcessedEvent":                 int64(0),
		"GetParentWorkflowID":                   "",
		"GetStartVersion":                       int64(0),
		"GetTaskList":                           "",
		"GetDecisionVersion":                    int64(0),
		"GetEventBranchToken":                   []uint8(nil),
		"GetEventStoreVersion":                  int32(0),
		"GetExecutionContext":                   []uint8(nil),
		"GetFirstExecutionRunID":                []uint8(nil),
		"GetHasRetryPolicy":                     false,
		"GetInitiatedID":                        int64(0),
		"GetIsCron":                             false,
		"GetLastEventTaskID":                    int64(0),
		"GetLastUpdatedTimestamp":               time.Time{},
		"GetLastWriteEventID":                   int64(0),
		"GetMemo":                               map[string][]uint8(nil),
		"GetParentDomainID":                     []uint8(nil),
		"GetParentRunID":                        []uint8(nil),
		"GetPartitionConfig":                    map[string]string(nil),
		"GetHistorySize":                        int64(0),
		"GetRetryAttempt":                       int64(0),
		"GetRetryBackoffCoefficient":            float64(0),
		"GetRetryExpiration":                    time.Duration(0),
		"GetRetryExpirationTimestamp":           time.Time{},
		"GetRetryInitialInterval":               time.Duration(0),
		"GetRetryMaximumAttempts":               int32(0),
		"GetRetryMaximumInterval":               time.Duration(0),
		"GetRetryNonRetryableErrors":            []string(nil),
		"GetSearchAttributes":                   map[string][]uint8(nil),
		"GetSignalCount":                        int64(0),
		"GetStartTimestamp":                     time.Time{},
		"GetState":                              int32(0),
		"GetStickyScheduleToStartTimeout":       time.Duration(0),
		"GetStickyTaskList":                     "",
		"GetVersionHistories":                   []uint8(nil),
		"GetVersionHistoriesEncoding":           "",
		"GetWorkflowTimeout":                    time.Duration(0),
		"GetWorkflowTypeName":                   "",
	},
	"*serialization.TimerTaskInfo": {
		"GetDomainID":        []uint8(nil),
		"GetEventID":         int64(0),
		"GetRunID":           []uint8(nil),
		"GetScheduleAttempt": int64(0),
		"GetTaskType":        int16(0),
		"GetTimeoutType":     int16(0),
		"GetVersion":         int64(0),
		"GetWorkflowID":      "",
	},
	"*serialization.ReplicationTaskInfo": {
		"GetBranchToken":             []uint8(nil),
		"GetCreationTimestamp":       time.Time{},
		"GetDomainID":                []uint8(nil),
		"GetEventStoreVersion":       int32(0),
		"GetFirstEventID":            int64(0),
		"GetNewRunBranchToken":       []uint8(nil),
		"GetNewRunEventStoreVersion": int32(0),
		"GetNextEventID":             int64(0),
		"GetRunID":                   []uint8(nil),
		"GetScheduledID":             int64(0),
		"GetTaskType":                int16(0),
		"GetVersion":                 int64(0),
		"GetWorkflowID":              "",
	},
	"*serialization.TaskListInfo": {
		"GetAckLevel":        int64(0),
		"GetExpiryTimestamp": time.Time{},
		"GetKind":            int16(0),
		"GetLastUpdated":     time.Time{},
	},
	"*serialization.TaskInfo": {
		"GetCreatedTimestamp": time.Time{},
		"GetExpiryTimestamp":  time.Time{},
		"GetPartitionConfig":  map[string]string(nil),
		"GetRunID":            []uint8(nil),
		"GetScheduleID":       int64(0),
		"GetWorkflowID":       "",
	},
	"*serialization.TimerInfo": {
		"GetExpiryTimestamp": time.Time{},
		"GetStartedID":       int64(0),
		"GetTaskID":          int64(0),
		"GetVersion":         int64(0),
	},
	"*serialization.RequestCancelInfo": {
		"GetCancelRequestID":       "",
		"GetInitiatedEventBatchID": int64(0),
		"GetVersion":               int64(0),
	},
	"*serialization.SignalInfo": {
		"GetControl":               []uint8(nil),
		"GetInitiatedEventBatchID": int64(0),
		"GetInput":                 []uint8(nil),
		"GetName":                  "",
		"GetRequestID":             "",
		"GetVersion":               int64(0),
	},
	"*serialization.ChildExecutionInfo": {
		"GetCreateRequestID":        "",
		"GetDomainID":               "",
		"GetDomainNameDEPRECATED":   "",
		"GetInitiatedEvent":         []uint8(nil),
		"GetInitiatedEventBatchID":  int64(0),
		"GetInitiatedEventEncoding": "",
		"GetParentClosePolicy":      int32(0),
		"GetStartedEvent":           []uint8(nil),
		"GetStartedEventEncoding":   "",
		"GetStartedID":              int64(0),
		"GetStartedRunID":           []uint8(nil),
		"GetStartedWorkflowID":      "",
		"GetVersion":                int64(0),
		"GetWorkflowTypeName":       "",
	},
	"*serialization.ActivityInfo": {
		"GetActivityID":               "",
		"GetAttempt":                  int32(0),
		"GetCancelRequestID":          int64(0),
		"GetCancelRequested":          false,
		"GetHasRetryPolicy":           false,
		"GetHeartbeatTimeout":         time.Duration(0),
		"GetRequestID":                "",
		"GetRetryBackoffCoefficient":  float64(0),
		"GetRetryExpirationTimestamp": time.Time{},
		"GetRetryInitialInterval":     time.Duration(0),
		"GetRetryLastFailureDetails":  []uint8(nil),
		"GetRetryLastFailureReason":   "",
		"GetRetryLastWorkerIdentity":  "",
		"GetRetryMaximumAttempts":     int32(0),
		"GetRetryMaximumInterval":     time.Duration(0),
		"GetRetryNonRetryableErrors":  []string(nil),
		"GetScheduleToCloseTimeout":   time.Duration(0),
		"GetScheduleToStartTimeout":   time.Duration(0),
		"GetScheduledEvent":           []uint8(nil),
		"GetScheduledEventBatchID":    int64(0),
		"GetScheduledEventEncoding":   "",
		"GetScheduledTimestamp":       time.Time{},
		"GetStartToCloseTimeout":      time.Duration(0),
		"GetStartedEvent":             []uint8(nil),
		"GetStartedEventEncoding":     "",
		"GetStartedID":                int64(0),
		"GetStartedIdentity":          "",
		"GetStartedTimestamp":         time.Time{},
		"GetTaskList":                 "",
		"GetTimerTaskStatus":          int32(0),
		"GetVersion":                  int64(0),
	},
	"*serialization.HistoryTreeInfo": {
		"GetAncestors":        []*types.HistoryBranchRange(nil),
		"GetCreatedTimestamp": time.Time{},
		"GetInfo":             "",
	},
	"*serialization.DomainInfo": {
		"GetActiveClusterName":           "",
		"GetArchivalBucket":              "",
		"GetArchivalStatus":              int16(0),
		"GetBadBinaries":                 []uint8(nil),
		"GetBadBinariesEncoding":         "",
		"GetClusters":                    []string(nil),
		"GetConfigVersion":               int64(0),
		"GetData":                        map[string]string(nil),
		"GetDescription":                 "",
		"GetEmitMetric":                  false,
		"GetFailoverEndTimestamp":        zeroUnix,
		"GetFailoverNotificationVersion": int64(0),
		"GetFailoverVersion":             int64(0),
		"GetHistoryArchivalStatus":       int16(0),
		"GetHistoryArchivalURI":          "",
		"GetLastUpdatedTimestamp":        time.Time{},
		"GetName":                        "",
		"GetNotificationVersion":         int64(0),
		"GetOwner":                       "",
		"GetPreviousFailoverVersion":     int64(0),
		"GetRetention":                   time.Duration(0),
		"GetStatus":                      int32(0),
		"GetVisibilityArchivalStatus":    int16(0),
		"GetVisibilityArchivalURI":       "",
	},
	"*serialization.ShardInfo": {
		"GetClusterReplicationLevel":                   map[string]int64(nil),
		"GetClusterTimerAckLevel":                      map[string]time.Time(nil),
		"GetClusterTransferAckLevel":                   map[string]int64(nil),
		"GetCrossClusterProcessingQueueStates":         []uint8(nil),
		"GetCrossClusterProcessingQueueStatesEncoding": "",
		"GetDomainNotificationVersion":                 int64(0),
		"GetOwner":                                     "",
		"GetPendingFailoverMarkers":                    []uint8(nil),
		"GetPendingFailoverMarkersEncoding":            "",
		"GetReplicationAckLevel":                       int64(0),
		"GetReplicationDlqAckLevel":                    map[string]int64(nil),
		"GetStolenSinceRenew":                          int32(0),
		"GetTimerAckLevel":                             time.Time{},
		"GetTimerProcessingQueueStates":                []uint8(nil),
		"GetTimerProcessingQueueStatesEncoding":        "",
		"GetTransferAckLevel":                          int64(0),
		"GetTransferProcessingQueueStates":             []uint8(nil),
		"GetTransferProcessingQueueStatesEncoding":     "",
		"GetUpdatedAt":                                 time.Time{},
	},
}

var expectedNonEmpty = map[string]map[string]any{
	"*serialization.WorkflowExecutionInfo": {
		"GetAutoResetPoints":                    []byte("resetpoints"),
		"GetAutoResetPointsEncoding":            "",
		"GetCancelRequestID":                    "",
		"GetCancelRequested":                    false,
		"GetClientFeatureVersion":               "",
		"GetClientImpl":                         "",
		"GetClientLibraryVersion":               "",
		"GetCloseStatus":                        int32(6),
		"GetCompletionEvent":                    []byte("completionEvent"),
		"GetCompletionEventEncoding":            "completionEventEncoding",
		"GetCreateRequestID":                    "",
		"GetCompletionEventBatchID":             int64(2),
		"GetCronSchedule":                       "",
		"GetDecisionAttempt":                    int64(0),
		"GetDecisionOriginalScheduledTimestamp": time.Time{},
		"GetDecisionRequestID":                  "",
		"GetDecisionScheduleID":                 int64(0),
		"GetDecisionScheduledTimestamp":         time.Time{},
		"GetDecisionStartedID":                  int64(0),
		"GetDecisionStartedTimestamp":           time.Time{},
		"GetDecisionTaskTimeout":                time.Duration(0),
		"GetDecisionTimeout":                    time.Duration(4),
		"GetLastFirstEventID":                   int64(7),
		"GetLastProcessedEvent":                 int64(0),
		"GetParentWorkflowID":                   "parentWorkflowID",
		"GetStartVersion":                       int64(0),
		"GetTaskList":                           "taskList",
		"GetDecisionVersion":                    int64(0),
		"GetEventBranchToken":                   []uint8(nil),
		"GetEventStoreVersion":                  int32(0),
		"GetExecutionContext":                   []byte("executionContext"),
		"GetFirstExecutionRunID":                []uint8(nil),
		"GetHasRetryPolicy":                     false,
		"GetInitiatedID":                        int64(1),
		"GetIsCron":                             false,
		"GetLastEventTaskID":                    int64(0),
		"GetLastUpdatedTimestamp":               time.Time{},
		"GetLastWriteEventID":                   int64(0),
		"GetMemo":                               map[string][]uint8(nil),
		"GetParentDomainID":                     []byte(parentDomainID),
		"GetParentRunID":                        []byte(parentRunID),
		"GetPartitionConfig":                    map[string]string(nil),
		"GetHistorySize":                        int64(0),
		"GetRetryAttempt":                       int64(0),
		"GetRetryBackoffCoefficient":            float64(0),
		"GetRetryExpiration":                    time.Duration(0),
		"GetRetryExpirationTimestamp":           time.Time{},
		"GetRetryInitialInterval":               time.Duration(0),
		"GetRetryMaximumAttempts":               int32(0),
		"GetRetryMaximumInterval":               time.Duration(0),
		"GetRetryNonRetryableErrors":            []string(nil),
		"GetSearchAttributes": map[string][]uint8{
			"key": []byte("value"),
		},
		"GetSignalCount":                  int64(0),
		"GetStartTimestamp":               time.Time{},
		"GetState":                        int32(5),
		"GetStickyScheduleToStartTimeout": time.Duration(0),
		"GetStickyTaskList":               "",
		"GetVersionHistories":             []uint8(nil),
		"GetVersionHistoriesEncoding":     "",
		"GetWorkflowTimeout":              time.Duration(3),
		"GetWorkflowTypeName":             "workflowTypeName",
	},
	"*serialization.TimerTaskInfo": {
		"GetDomainID":        []byte(taskDomainID),
		"GetEventID":         int64(5),
		"GetRunID":           []byte(taskRunID),
		"GetScheduleAttempt": int64(4),
		"GetTaskType":        int16(1),
		"GetTimeoutType":     int16(2),
		"GetVersion":         int64(3),
		"GetWorkflowID":      "workflowID",
	},
	"*serialization.ReplicationTaskInfo": {
		"GetBranchToken":             []byte("branchToken"),
		"GetCreationTimestamp":       replicationCreationTimestamp,
		"GetDomainID":                []uint8(replicationTaskDomainID),
		"GetEventStoreVersion":       int32(6),
		"GetFirstEventID":            int64(3),
		"GetNewRunBranchToken":       []byte("newRunBranchToken"),
		"GetNewRunEventStoreVersion": int32(7),
		"GetNextEventID":             int64(4),
		"GetRunID":                   []byte(replicationTaskRunID),
		"GetScheduledID":             int64(5),
		"GetTaskType":                int16(1),
		"GetVersion":                 int64(2),
		"GetWorkflowID":              "workflowID",
	},
	"*serialization.TaskListInfo": {
		"GetAckLevel":        int64(2),
		"GetExpiryTimestamp": taskListInfoExpireTime,
		"GetKind":            int16(1),
		"GetLastUpdated":     taskListInfoLastUpdateTime,
	},
	"*serialization.TaskInfo": {
		"GetCreatedTimestamp": taskInfoCreateTime,
		"GetExpiryTimestamp":  taskInfoExpiryTime,
		"GetPartitionConfig":  map[string]string{"key": "value"},
		"GetRunID":            []byte(taskInfoRunID),
		"GetScheduleID":       int64(1),
		"GetWorkflowID":       "workflowID",
	},
	"*serialization.TimerInfo": {
		"GetExpiryTimestamp": timerInfoExpireTime,
		"GetStartedID":       int64(2),
		"GetTaskID":          int64(3),
		"GetVersion":         int64(1),
	},
	"*serialization.RequestCancelInfo": {
		"GetCancelRequestID":       "cancelRequestID",
		"GetInitiatedEventBatchID": int64(2),
		"GetVersion":               int64(1),
	},
	"*serialization.SignalInfo": {
		"GetControl":               []byte("signalControl"),
		"GetInitiatedEventBatchID": int64(2),
		"GetInput":                 []byte("signalInput"),
		"GetName":                  "signalName",
		"GetRequestID":             "signalRequestID",
		"GetVersion":               int64(1),
	},
	"*serialization.ChildExecutionInfo": {
		"GetCreateRequestID":        "createRequestID",
		"GetDomainID":               "domainID",
		"GetDomainNameDEPRECATED":   "",
		"GetInitiatedEvent":         []byte("initiatedEvent"),
		"GetInitiatedEventBatchID":  int64(2),
		"GetInitiatedEventEncoding": "initiatedEventEncoding",
		"GetParentClosePolicy":      int32(1),
		"GetStartedEvent":           []byte("startedEvent"),
		"GetStartedEventEncoding":   "startedEventEncoding",
		"GetStartedID":              int64(3),
		"GetStartedRunID":           []byte(childExecutionInfoStartedRunID),
		"GetStartedWorkflowID":      "startedWorkflowID",
		"GetVersion":                int64(1),
		"GetWorkflowTypeName":       "workflowTypeName",
	},
	"*serialization.ActivityInfo": {
		"GetActivityID":               "activityID",
		"GetAttempt":                  int32(6),
		"GetCancelRequestID":          int64(4),
		"GetCancelRequested":          true,
		"GetHasRetryPolicy":           true,
		"GetHeartbeatTimeout":         time.Duration(4),
		"GetRequestID":                "requestID",
		"GetRetryBackoffCoefficient":  float64(8),
		"GetRetryExpirationTimestamp": activeInfoRetryExpirationTime,
		"GetRetryInitialInterval":     time.Duration(5),
		"GetRetryLastFailureDetails":  []byte("retryLastFailureDetails"),
		"GetRetryLastFailureReason":   "retryLastFailureReason",
		"GetRetryLastWorkerIdentity":  "retryLastWorkerIdentity",
		"GetRetryMaximumAttempts":     int32(7),
		"GetRetryMaximumInterval":     time.Duration(6),
		"GetRetryNonRetryableErrors":  []string{"error1", "error2"},
		"GetScheduleToCloseTimeout":   time.Duration(1),
		"GetScheduleToStartTimeout":   time.Duration(2),
		"GetScheduledEvent":           []byte("scheduledEvent"),
		"GetScheduledEventBatchID":    int64(2),
		"GetScheduledEventEncoding":   "scheduledEventEncoding",
		"GetScheduledTimestamp":       activityInfoScheduledTime,
		"GetStartToCloseTimeout":      time.Duration(3),
		"GetStartedEvent":             []byte("startedEvent"),
		"GetStartedEventEncoding":     "startedEventEncoding",
		"GetStartedID":                int64(3),
		"GetStartedIdentity":          "startedIdentity",
		"GetStartedTimestamp":         activeInfoStartedTime,
		"GetTaskList":                 "taskList",
		"GetTimerTaskStatus":          int32(5),
		"GetVersion":                  int64(1),
	},
	"*serialization.HistoryTreeInfo": {
		"GetAncestors": []*types.HistoryBranchRange{
			{
				BranchID: "branchID1",
			},
			{
				BranchID: "branchID2",
			},
		},
		"GetCreatedTimestamp": historyTreeEventCreatedTime,
		"GetInfo":             "historyTreeInfo",
	},
	"*serialization.DomainInfo": {
		"GetActiveClusterName":           "cluster1",
		"GetArchivalBucket":              "archivalBucket",
		"GetArchivalStatus":              int16(2),
		"GetBadBinaries":                 []byte("badBinaries"),
		"GetBadBinariesEncoding":         "badBinariesEncoding",
		"GetClusters":                    []string{"cluster1", "cluster2"},
		"GetConfigVersion":               int64(3),
		"GetData":                        map[string]string{"datakey": "datavalue"},
		"GetDescription":                 "description",
		"GetEmitMetric":                  true,
		"GetFailoverEndTimestamp":        domainInfoFailoverEndTimestamp,
		"GetFailoverNotificationVersion": int64(5),
		"GetFailoverVersion":             int64(6),
		"GetHistoryArchivalStatus":       int16(7),
		"GetHistoryArchivalURI":          "historyArchivalURI",
		"GetLastUpdatedTimestamp":        domainInfoLastUpdatedTimestamp,
		"GetName":                        "name",
		"GetNotificationVersion":         int64(4),
		"GetOwner":                       "owner",
		"GetPreviousFailoverVersion":     int64(9),
		"GetRetention":                   time.Duration(1),
		"GetStatus":                      int32(1),
		"GetVisibilityArchivalStatus":    int16(8),
		"GetVisibilityArchivalURI":       "visibilityArchivalURI",
	},
	"*serialization.ShardInfo": {
		"GetClusterReplicationLevel": map[string]int64{
			"cluster1": 6,
		},
		"GetClusterTimerAckLevel": map[string]time.Time{
			"cluster1": shardInfoTimerAckLevel,
		},
		"GetClusterTransferAckLevel": map[string]int64{
			"cluster1": 5,
		},
		"GetCrossClusterProcessingQueueStates":         []byte("crossClusterProcessingQueueStates"),
		"GetCrossClusterProcessingQueueStatesEncoding": "crossClusterProcessingQueueStatesEncoding",
		"GetDomainNotificationVersion":                 int64(4),
		"GetOwner":                                     "owner",
		"GetPendingFailoverMarkers":                    []byte("pendingFailoverMarkers"),
		"GetPendingFailoverMarkersEncoding":            "pendingFailoverMarkersEncoding",
		"GetReplicationAckLevel":                       int64(2),
		"GetReplicationDlqAckLevel": map[string]int64{
			"cluster1": 7,
		},
		"GetStolenSinceRenew":                      int32(1),
		"GetTimerAckLevel":                         shardInfoTimerAckLevel,
		"GetTimerProcessingQueueStates":            []byte("timerProcessingQueueStates"),
		"GetTimerProcessingQueueStatesEncoding":    "timerProcessingQueueStatesEncoding",
		"GetTransferAckLevel":                      int64(3),
		"GetTransferProcessingQueueStates":         []byte("transferProcessingQueueStates"),
		"GetTransferProcessingQueueStatesEncoding": "transferProcessingQueueStatesEncoding",
		"GetUpdatedAt":                             shardInfoUpdatedTime,
	},
}