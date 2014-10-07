net-http-recover
================

[![Build Status](https://drone.io/github.com/bakins/net-http-recover/status.png)](https://drone.io/github.com/bakins/net-http-recover/latest)

See
[![GoDoc](https://godoc.org/github.com/bakins/net-http-recover?status.svg)](https://godoc.org/github.com/bakins/net-http-recover)
for documentation.

Simple Go net/http recovery middleware

`http.Handler` middleware for panic recovery.

I've copy/pasted this code around several times, so I put it into a
proper package.

Inspired by https://github.com/codegangsta/negroni and https://github.com/gorilla/handlers
