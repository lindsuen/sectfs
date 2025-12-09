// sectfs - file.go
// Copyright (C) 2025 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package core

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID          string
	Name        string
	Size        int64
	Path        string
	CreatedTime int64  // timestamp
	Hash        string // sha1
}

func (f *File) SetFileID() {
	f.ID = uuid.New().String()
}

func (f *File) SetFileName(s string) {
	f.Name = s
}

func (f *File) SetFileSize(i int64) {
	f.Size = i
}

func (f *File) SetFilePath(s string) {
	f.Path = s
}

func (f *File) SetFileCreatedTime() {
	f.CreatedTime = time.Now().UnixMilli()
}

func (f *File) SetFileHash(s string) {
	f.Hash = s
}
