# Cross-Origin Resource Sharing Middleware

[![Build Status](https://travis-ci.org/go-gem/middleware-cors.svg?branch=master)](https://travis-ci.org/go-gem/middleware-cors)
[![GoDoc](https://godoc.org/github.com/go-gem/middleware-cors?status.svg)](https://godoc.org/github.com/go-gem/middleware-cors)
[![Coverage Status](https://coveralls.io/repos/github/go-gem/middleware-cors/badge.svg?branch=master)](https://coveralls.io/github/go-gem/middleware-cors?branch=master)

Cross-Origin Resource Sharing Middleware for [Gem](https://github.com/go-gem/gem) Web framework.

## Install

```
$ go get -u github.com/go-gem/middleware-cors
```

## Example

```
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
```