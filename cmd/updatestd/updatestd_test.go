// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.14
// +build go1.14

package main

import (
	"runtime"
	"testing"
)

func TestBuildList(t *testing.T) {
	// This test is in the golang.org/x/module.
	// Check that buildList(".") returns sensible results.
	main, deps := buildList(".")
	if want := (module{
		Path:     "golang.org/x/build",
		Main:     true,
		Indirect: false,
	}); main != want {
		t.Errorf("got main = %+v, want %+v", main, want)
	}
	for i, m := range deps {
		if m.Path == "" {
			t.Errorf("deps[%d]: module path is empty", i)
		}
		if m.Main {
			t.Errorf("deps[%d]: unexpectedly a main module", i)
		}
	}
	if len(deps) < 10 {
		t.Errorf("buildList returned %d (less than 10) non-main modules in build list of x/build; "+
			"either buildList is broken or TestBuildList needs to be updated (that'll be the day)",
			len(deps))
	}
}

func TestGOROOTVersion(t *testing.T) {
	// This test requires Go 1.14 or higher to run.
	// Verify gorootVersion(runtime.GOROOT()) returns a fitting version.
	v, err := gorootVersion(runtime.GOROOT())
	if err != nil {
		t.Fatal("gorootVersion returned non-nil error:", err)
	}
	if v < 14 {
		t.Errorf("gorootVersion returned unexpectedly low version %d", v)
	}
}
