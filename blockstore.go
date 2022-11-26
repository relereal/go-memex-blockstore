// Right now this module only passes all the ipld storage
// interface function into the underlying datastore
// implemented using sqlite3.
//
// Eventually, it will contain additional functionality,
// equivalent to ipfs blockstore, like caching or ???
package blockstore

import (
	"context"

	"github.com/relereal/go-sqlite-datastore"
)

type Blockstore struct {
	Store *datastore.Datastore
}

func NewBlockstore(ds *datastore.Datastore) *Blockstore {
	return &Blockstore{ds}
}

func (bs *Blockstore) Has(ctx context.Context, key string) (bool, error) {
	return bs.Store.Has(ctx, key)
}

func (bs *Blockstore) Get(ctx context.Context, key string) ([]byte, error) {
	return bs.Store.Get(ctx, key)
}

func (bs *Blockstore) Put(ctx context.Context, key string, content []byte) error {
	return bs.Store.Put(ctx, key, content)
}
