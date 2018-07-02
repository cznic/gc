# Copyright 2016 The GC Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all clean cover cpu editor internalError later mem nuke todo edit fuzz fuzz2 fuzz-more

grep=--include=*.go --include=*.l --include=*.y --include=*.yy
ngrep='TODOOK\|.*_string\.go\|testdata/errchk'

all: editor0 editor
	go vet 2>&1 | grep -v $(ngrep) || true
	golint 2>&1 | grep -v $(ngrep) || true
	make todo
	unused . || true
	misspell *.go
	gosimple || true
	unconvert || true
	maligned || true

clean:
	go clean
	rm -f *~ *.test *.out gc-fuzz.zip y.output

cover:
	t=$(shell tempfile) ; go test -coverprofile $$t && go tool cover -html $$t && unlink $$t

cpu: clean
	go test -run @ -bench Checker/StdCh$$ -cpuprofile cpu.out
	go tool pprof -web *.test cpu.out

edit:
	touch errchk.log log
	@2>/dev/null gvim -p Makefile errchk.log log *.go

editor0:
	go generate 2>&1 | tee -a log
	go test -i
	go test -run Example | fe
	grep -n Invalid example_test.go || true

editor:
	@echo $(shell LC_TIME=c date) | tee log
	gofmt -l -s -w *.go
	go test -i
	go test -noerrchk 2>&1 | tee -a log
	@echo $(shell LC_TIME=c date) | tee errchk.log
	go test -run TestErrchk -chk.summary 2>&1 | tee -a errchk.log
	sed -i -e s\\$(shell pwd)\/\\\\ errchk.log
	go build
	grep -n '^TODO\|[^X]TODO' errchk.log | tee -a errchk.log
	go test -run @ -bench . -benchmem | tee -a errchk.log
	git diff errchk.log

fuzz:
	go-fuzz-build -func FuzzLexer github.com/cznic/gc-priv
	rm -rf testdata/fuzz/lexer/corpus/ testdata/fuzz/lexer/crashers/ testdata/fuzz/lexer/suppressions/
	-go-fuzz -bin gc-fuzz.zip -workdir testdata/fuzz/lexer/

fuzz-more:
	-go-fuzz -bin gc-fuzz.zip -workdir testdata/fuzz/lexer/

fuzz2:
	cat $$(ls testdata/fuzz/lexer/crashers/*.output | head -n 1)

internalError:
	egrep -ho '"internal error.*"' *.go | sort | cat -n

later:
	@grep -n $(grep) LATER * || true
	@grep -n $(grep) MAYBE * || true

mem: clean
	go test -run @ -bench Parser/Std$$ -memprofile mem.out -memprofilerate 1 -timeout 24h
	go tool pprof -lines -web -alloc_space *.test mem.out

nuke: clean
	go clean -i

todo:
	@grep -nr $(grep) ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* * | grep -v $(ngrep) || true
	@grep -nr $(grep) TODO * | grep -v $(ngrep) || true
	@grep -nr $(grep) BUG * | grep -v $(ngrep) || true
	@grep -nr $(grep) [^[:alpha:]]println * | grep -v $(ngrep) || true
