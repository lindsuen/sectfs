// sectfs - db_test.go
// Copyright (C) 2025 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package db

import (
	"encoding/base64"
	"fmt"
	"testing"
)

var (
	dataDirectoryPath = "../../data"
	testKey           = []byte("")
	testKeyPrefix     = ""
)

func TestGet(t *testing.T) {
	_, err := Open(dataDirectoryPath)
	if err != nil {
		fmt.Println(err)
	}
	value, _ := base64.RawURLEncoding.DecodeString(string(Get(testKey)))
	fmt.Println("value=" + string(value))
}

func TestIteratorKeys(t *testing.T) {
	_, err := Open(dataDirectoryPath)
	if err != nil {
		fmt.Println(err)
	}
	IteratorKeys()
}

func TestIteratorKeysAndValues(t *testing.T) {
	_, err := Open(dataDirectoryPath)
	if err != nil {
		fmt.Println(err)
	}
	IteratorKeysAndValues()
}

func TestSeekWithPrefix(t *testing.T) {
	_, err := Open(dataDirectoryPath)
	if err != nil {
		fmt.Println(err)
	}
	SeekWithPrefix(testKeyPrefix)
}
