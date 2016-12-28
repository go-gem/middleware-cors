// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

/*
Package corsmidware provides Cross-Origin Resource Sharing middleware for Gem web framework.
*/
package corsmidware

import (
	"net/http"

	"github.com/go-gem/gem"
	"github.com/rs/cors"
)

// New return a CORS middleware instance with the given options.
func New(options cors.Options) *CORS {
	return &CORS{
		cors: cors.New(options),
	}
}

// CORS Cross-Origin Resource Sharing middleware.
type CORS struct {
	cors *cors.Cors
}

// Wrap implements the Middleware interface.
func (c *CORS) Wrap(next gem.Handler) gem.Handler {
	return gem.HandlerFunc(func(ctx *gem.Context) {
		c.cors.ServeHTTP(ctx.Response, ctx.Request, func(w http.ResponseWriter, r *http.Request) {
			next.Handle(ctx)
		})
	})
}
