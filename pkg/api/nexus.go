package api

import (
	context "context"
	"errors"

	internal_raft "github.com/flipkart-incubator/nexus/internal/raft"
	"github.com/flipkart-incubator/nexus/pkg/db"
	"github.com/flipkart-incubator/nexus/pkg/raft"
)

type RaftReplicator interface {
	Start()
	Save(context.Context, []byte) ([]byte, error)
	AddMember(context.Context, int, string) error
	RemoveMember(context.Context, int) error
	Stop()
}

func NewRaftReplicator(store db.Store, opts ...raft.Option) (RaftReplicator, error) {
	if store == nil {
		return nil, errors.New("Store must be given")
	}
	if options, err := raft.NewOptions(opts...); err != nil {
		return nil, err
	} else {
		return internal_raft.NewReplicator(store, options), nil
	}
}
