// sectfs - db.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package db

import (
	"log"
	"strings"
	"time"

	"github.com/dgraph-io/badger/v4"
)

var (
	badgerDB *badger.DB
	err      error
)

func Open(path string) (*badger.DB, error) {
	badgerDB, err = badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return badgerDB, nil
}

func Close() error {
	return badgerDB.Close()
}

func Set(key []byte, value []byte) {
	wb := badgerDB.NewWriteBatch()
	defer wb.Cancel()
	err := wb.SetEntry(badger.NewEntry(key, value).WithMeta(0))
	if err != nil {
		log.Println("Failed to write data to cache.", "key", string(key), "value", string(value), "err", err)
	}
	err = wb.Flush()
	if err != nil {
		log.Println("Failed to flush data to cache.", "key", string(key), "value", string(value), "err", err)
	}
}

func SetWithTTL(key []byte, value []byte, ttl int64) {
	wb := badgerDB.NewWriteBatch()
	defer wb.Cancel()
	err := wb.SetEntry(badger.NewEntry(key, value).WithMeta(0).WithTTL(time.Duration(ttl * time.Second.Nanoseconds())))
	if err != nil {
		log.Println("Failed to write data to cache.", "key", string(key), "value", string(value), "err", err)
	}
	err = wb.Flush()
	if err != nil {
		log.Println("Failed to flush data to cache.", "key", string(key), "value", string(value), "err", err)
	}
}

func Get(key []byte) []byte {
	var ival []byte
	err := badgerDB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		ival, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		log.Println("Failed to read data from the cache.", "key", string(key), "error", err)
	}
	return ival
}

func Has(key []byte) (bool, error) {
	var exist = false
	err := badgerDB.View(func(txn *badger.Txn) error {
		_, err := txn.Get(key)
		if err != nil {
			return err
		} else {
			exist = true
		}
		return err
	})
	if strings.HasSuffix(err.Error(), "not found") {
		err = nil
	}
	return exist, err
}

func Delete(key []byte) error {
	wb := badgerDB.NewWriteBatch()
	defer wb.Cancel()
	return wb.Delete(key)
}

func IteratorKeysAndValues() {
	err := badgerDB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				log.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Println("Failed to iterator keys and values from the cache.", "error", err)
	}
}

func IteratorKeys() {
	err := badgerDB.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			log.Printf("key=%s\n", k)
		}
		return nil
	})
	if err != nil {
		log.Println("Failed to iterator keys from the cache.", "error", err)
	}
}

func SeekWithPrefix(prefixStr string) {
	err := badgerDB.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(prefixStr)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				log.Printf("key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Println("Failed to seek prefix from the cache.", "prefix", prefixStr, "error", err)
	}
}
