// sectfs - config_test.go
// Copyright (C) 2024 LindSuen <lindsuen@foxmail.com>
//
// Use of this source code is governed by a BSD 2-Clause license that can be
// found in the LICENSE file.

package config

import (
	"fmt"
	"testing"
)

func TestParseIni(t *testing.T) {
	cfg, err := parseInitFile("../../config/sectfs.conf")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The listening address: " + cfg.Address + ":" + cfg.Port)
}
