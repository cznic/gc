# Copyright 2016 The GC Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean cover cpu editor internalError later mem nuke todo edit

grep=--include=*.go --include=*.l --include=*.y --include=*.yy
ngrep='TODOOK\|parser\.go\|scanner\.go\|.*_string\.go'

all: editor
	go install
	go vet 2>&1 | grep -v $(ngrep) || true
	golint 2>&1 | grep -v $(ngrep) || true
	make todo
	unused . || true
	misspell *.go
	gosimple || true

clean:
	go clean
	rm -f *~ *.test *.out

cover:
	t=$(shell tempfile) ; go test -coverprofile $$t && go tool cover -html $$t && unlink $$t

cpu: clean
	go test -run @ -bench . -cpuprofile cpu.out
	go tool pprof -lines *.test cpu.out

edit:
	gvim -p Makefile scanner.l parser.yy log test.log all_test.go ast2.go builtin.go context.go decl.go etc.go gc.go lexer.go package.go type.go unsafe.go value.go

editor: parser.go scanner.go
	gofmt -l -s -w *.go
	go test -i
	go test 2>&1 | tee log
	grep -n 'extra error' test.log || true
	grep -n 'xc\.Dict' *.go | grep -v global.go || true
	git status test.log

internalError:
	egrep -ho '"internal error.*"' *.go | sort | cat -n

later:
	@grep -n $(grep) LATER * || true
	@grep -n $(grep) MAYBE * || true

mem: clean
	go test -run @ -bench Load -memprofile mem.out -memprofilerate 1 -timeout 24h
	go tool pprof -lines -web -alloc_space *.test mem.out

nuke: clean
	go clean -i

parser.go scanner.go: parser.yy scanner.l xerrors
	go test -i
	go generate 2>&1 | tee log-generate

todo:
	@grep -n $(grep) ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* * | grep -v $(ngrep) || true
	@grep -n $(grep) 'TODO\|todo' * | grep -v $(ngrep) || true
	@grep -n $(grep) BUG * | grep -v $(ngrep) || true
	@grep -n $(grep) [^[:alpha:]]println * | grep -v $(ngrep) || true
