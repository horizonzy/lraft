// Copyright (c) 2020, pole-group. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entity

// Raft Error Code
type RaftErrorCode int

const (
	UNKNOWN             = RaftErrorCode(-1)
	SUCCESS             = RaftErrorCode(0)
	ERaftTimedOut       = RaftErrorCode(10001)
	EStateMachine       = RaftErrorCode(10002)
	ECatchup            = RaftErrorCode(10003)
	ELeaderMoved        = RaftErrorCode(10004)
	EStepEer            = RaftErrorCode(10005)
	ENodeShutdown       = RaftErrorCode(10006)
	EHigherTermRequest  = RaftErrorCode(10007)
	EHigherTermResponse = RaftErrorCode(10008)
	EBadNode            = RaftErrorCode(10009)
	EVoteForCandidate   = RaftErrorCode(10010)
	ENewLeader          = RaftErrorCode(10011)
	ELeaderConflict     = RaftErrorCode(10012)
	ETransferLeaderShip = RaftErrorCode(10013)
	ELogDeleted         = RaftErrorCode(10014)
	ENoMoreUserLog      = RaftErrorCode(10015)
	ERequest            = RaftErrorCode(1000)
	EStop               = RaftErrorCode(1001)
	EAGAIN              = RaftErrorCode(1002)
	EINTR               = RaftErrorCode(1003)
	EInternal           = RaftErrorCode(1004)
	ECANCELED           = RaftErrorCode(1005)
	EHostDown           = RaftErrorCode(1006)
	EShutdown           = RaftErrorCode(1007)
	EPERM               = RaftErrorCode(1008)
	EBUSY               = RaftErrorCode(1009)
	ETIMEDOUT           = RaftErrorCode(1010)
	ESTALE              = RaftErrorCode(1011)
	ENOENT              = RaftErrorCode(1012)
	EExists             = RaftErrorCode(1013)
	EIO                 = RaftErrorCode(1014)
	EINVAL              = RaftErrorCode(1015)
	EACCES              = RaftErrorCode(1016)
)
