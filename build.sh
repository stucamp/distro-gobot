#!/bin/sh

go build \
-o "./distrobot" \
cmd/testbot.go \
cmd/rss_parse.go \
cmd/quotes.go \
cmd/randomgen.go \
cmd/distrowatch.go \
cmd/readjson.go