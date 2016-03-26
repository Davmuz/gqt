// Copyright 2016 Davide Muzzarelli. All right reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gqt

import (
	"os"
	"path/filepath"
	"testing"
)

var testDir string

func init() {
	var err error
	testDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	testDir = filepath.Join(testDir, "test")
}

func Test(t *testing.T) {
	for _, dir := range []string{"pkg1", "pkg2"} {
		err := Add(filepath.Join(testDir, dir), "*.sql")
		if err != nil {
			t.Error(err)
		}
	}

	if Get("test1") != "test1" {
		t.Error("test1")
	}

	if Exec("test2", true) != "Yes" {
		t.Error("test2")
	}

	if Get("sub/test3") != "test3" {
		t.Error("test3")
	}

	if Get("test4") != "test4" {
		t.Error("test4")
	}

	if Get("test5") != "test5" {
		t.Error("test5")
	}
}
