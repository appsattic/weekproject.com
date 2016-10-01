all:
	echo 'Provide a target: weekproject clean'

vendor:
	gb vendor fetch github.com/boltdb/bolt

fmt:
	find src/ -name '*.go' -exec go fmt {} ';'

build: fmt
	gb build all

weekproject: build
	./bin/weekproject

test:
	gb test -v

clean:
	rm -rf bin/ pkg/

.PHONY: weekproject
