// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux || darwin

package main

import (
	"net/http"
	"time"

	"golang.org/x/build/internal/buildgo"
)

// initTryDev registers a mock /try-dev page to make it easier
// to do local development of the trybot status page and its CSS.
func initTryDev(mux *http.ServeMux) {
	ts := &trySet{
		tryKey: tryKey{
			Project:  "go",
			Branch:   "master",
			ChangeID: "I1936e2dbe90634817f1aedabcba3c2b9f94e401b",
			Commit:   "555cfa3ee5e9f3df4b10c96af487424bfde19125",
		},
		tryID: "T4bfde19125",
		trySetState: trySetState{
			failed: []string{"failed-build"},
			remain: 1,
			builds: []*buildStatus{
				{
					BuilderRev: buildgo.BuilderRev{
						Name: "linux-amd64-race",
						Rev:  "555cfa3ee5e9f3df4b10c96af487424bfde19125",
					},
					startTime: time.Now(),
				},
				{
					BuilderRev: buildgo.BuilderRev{
						Name: "darwin-amd64-race",
						Rev:  "555cfa3ee5e9f3df4b10c96af487424bfde19125",
					},
					startTime: time.Now(),
					done:      time.Now().Add(3 * time.Minute),
				},
			},
		},
	}
	mux.HandleFunc("/try-dev", func(w http.ResponseWriter, r *http.Request) {
		tss := ts.trySetState.clone()
		serveTryStatusHTML(w, ts, tss)
	})
}
