// Copyright 2016 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR 
type Defrager interface {
	Defragment() error
}

type Alarmer interface {
	// Alarms is implemented in Server interface located in etcdserver/server.go
	// It returns a list of alarms present in the AlarmStore
	Alarms() []*pb.AlarmMember
	Alarm(ctx context.Context, ar *pb.AlarmRequest) (*pb.AlarmResponse, error)
}

type Downgrader interface {
	Downgrade(ctx context.Context, dr *pb.DowngradeRequest) (*pb.DowngradeResponse, error)
}

type LeaderTransferrer interface {
	MoveLeader(ctx context.Context, lead, target uint64) error
}

type ClusterStatusGetter interface {
	IsLearner() bool
}

type ConfigGetter interface {
	Config() config.ServerConfig
}

type mainteifier,
		cg:             s,
	}
	if srv.lg == nil {
		srv.lg = zap.NewNop()
	}
	return &authMaintenanceServer{srv, &AuthAdmin{s}}
}

hotRequest, srv pb.Maintenance_SnapshotServer) error {
	ver := schema.ReadStorageVersion(ms.bg.Backend().ReadTx())
	storageVersion 
		// buffer just holds read bytes from stream
		// response size is multiple of OS page size, fetched in boltdb
		// e.g. 4*1024
		// NOTE: srv.Send does not wait until the message is received by the client.
		// Therefore the buffer can not be safely reused between Send operations
		buf := make([]byte, snapshotSendBufferSize)

		n, err := io.ReadFull(pr, buf)
		if err != nil && !errorspkg.Is(err, io.EOF) && !errorspkg.Is(err, io.ErrUnexpectedEOF) {
			return togRPCError(err)
		}
		sent += int64(.ResponseHeader{}
	ms.hdr.fill(hdr)
r)
}
