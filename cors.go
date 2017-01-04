// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

/*
Package corsmidware provides Cross-Origin Resource Sharing middleware for Gem web framework.

Example

	package main

	import (
		"github.com/go-gem/gem"
		"github.com/go-gem/middleware-cors"
		"github.com/rs/cors"
	)

	var corsMiddleware = corsmidware.New(cors.Options{})

	func main() {
		router := gem.NewRouter()
		router.Use(corsMiddleware)
		router.GET("/", func(ctx *gem.Context) {
			ctx.HTML(200, "foo")
		})

		gem.ListenAndServe(":8080", router.Handler())
	}
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
		c.cors.ServeHTTP(ctx.Response, ctx.Request, func(_ http.ResponseWriter, _ *http.Request) {
			next.Handle(ctx)
		})
	})
}
