HTTP Analyzer
=============

Simple tool written in golang for analyzing incoming http traffic

This is how to build it on non-linux platform for linux os (cross compiling)::

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s"
