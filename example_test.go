// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package corsmidware

import (
	"fmt"
	"net/http/httptest"

	"github.com/go-gem/gem"
	"github.com/rs/cors"
)

func Example() {
	corsMidware := New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
	})
	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	ctx := &gem.Context{Request: req, Response: resp}

	var pass bool
	handler := corsMidware.Wrap(gem.HandlerFunc(func(ctx *gem.Context) {
		pass = true
	}))
	handler.Handle(ctx)

	fmt.Println(pass)
	// Output:
	// true
}
