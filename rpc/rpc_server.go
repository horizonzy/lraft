// Copyright (c) 2020, pole-group. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"context"

	polerpc "github.com/pole-group/pole-rpc"
)

type RaftRPCServer struct {
	IsReady chan struct{}
	server  polerpc.TransportServer
	Ctx     context.Context
	cancelF context.CancelFunc
}

func NewRaftRPCServer(label string, port int32, openTSL bool) (*RaftRPCServer, error) {
	ctx, cancelF := context.WithCancel(context.Background())

	r := &RaftRPCServer{
		IsReady: make(chan struct{}),
		Ctx:     ctx,
		cancelF: cancelF,
	}

	server, err := polerpc.NewTransportServer(r.Ctx, polerpc.ConnectTypeRSocket, label, port, openTSL)
	if err != nil {
		return nil, err
	}
	r.server = server

	return r, err
}

func (rpcServer *RaftRPCServer) GetRealServer() polerpc.TransportServer {
	return rpcServer.server
}
