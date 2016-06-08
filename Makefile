# Copyright 2016 The GC Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean cover cpu editor internalError later mem nuke todo edit

grep=--include=*.go --include=*.l --include=*.y --include=*.yy

all: editor
	go install
	go vet || true
	golint || true
	make todo
	unused . || true
	misspell *.go
	gosimple

clean:
	go clean
	rm -f *~ cpu.test mem.test

cover:
	t=$(shell tempfile) ; go test -coverprofile $$t && go tool cover -html $$t && unlink $$t

cpu:
	go test -c -o cpu.test
	./cpu.test -noerr -test.cpuprofile cpu.out
	go tool pprof --lines cpu.test cpu.out

edit:
	gvim -p Makefile *.l parser.yy log *.go

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

mem:
	go test -c -o mem.test
	./mem.test -test.bench . -test.memprofile mem.out
	go tool pprof --lines --web --alloc_space mem.test mem.out

nuke: clean
	go clean -i

parser.go scanner.go: parser.yy scanner.l xerrors
	go test -i
	go generate 2>&1 | tee log-generate

todo:
	@grep -n $(grep) ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* * || true
	@grep -n $(grep) TODO * || true
	@grep -n $(grep) BUG * || true
	@grep -n $(grep) [^[:alpha:]]println * || true
