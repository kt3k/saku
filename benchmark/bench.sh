#!/bin/sh

# info action
hyperfine --export-json result_info.json \
	"node_modules/.bin/saku -i" \
	"$GOPATH/bin/saku -i"

# single run
hyperfine --export-json result_run.json \
	"node_modules/.bin/saku h" \
	"$GOPATH/bin/saku h"

H10="h h h h h h h h h h"

# 10 in serial
hyperfine --export-json result_run_s_10.json \
	"node_modules/.bin/saku $H10" \
	"$GOPATH/bin/saku $H10"

# 10 in parallel
hyperfine --export-json result_run_p_10.json \
	"node_modules/.bin/saku -p $H10" \
	"$GOPATH/bin/saku -p $H10"
