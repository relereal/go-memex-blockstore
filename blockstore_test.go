package blockstore

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/relereal/go-sqlite-datastore"
)

func getDatastore() *datastore.Datastore {
	os.Mkdir("test", 0777)
	os.Remove("test/testdb.db")
	ds := datastore.NewDatastore("test/testdb.db", "keystore")
	ds.Connect()
	return ds
}

func getBlockstore() *Blockstore {
	ds := getDatastore()
	return NewBlockstore(ds)
}

func clearDatastore(ds *datastore.Datastore) {
	ds.CloseDb()
	os.RemoveAll("test")
}

func TestBlockstore(t *testing.T) {
	bs := getBlockstore()
	defer clearDatastore(bs.Store)

	key := "testkey"
	value := []byte("testvalue")

	has, err := bs.Has(context.Background(), key)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}
	if has {
		t.Errorf("Expected has=false but got has=%t", has)
	}

	err = bs.Put(context.Background(), key, value)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	has, err = bs.Has(context.Background(), key)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}
	if !has {
		t.Errorf("Expected has=true but got has=%t", has)
	}

	content, err := bs.Get(context.Background(), key)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	same := bytes.Compare(content, value)
	if same != 0 {
		t.Errorf("Expected same %s=%s", content, value)
	}
}
