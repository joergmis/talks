#!/bin/bash

set -e

basedir=$(pwd)

# compile to wasm
cd ${basedir}/code/wasm/testdata
echo "building wasm with go:"
time GOOS=wasip1 GOARCH=wasm go build -o add-go.wasm add.go
echo "building wasm with tinygo:"
time tinygo build -o add.wasm -target=wasi add.go

# compile the binary which uses the wasm binaries
cd ${basedir}/code/wasm
go build -o wasm-test main.go

# compile different examples with standard go and tinygo for size and time
# comparison
for dir in "code/embedded/c" "code/embedded/stdlib" "code/embedded/println"; do
	cd ${basedir}/${dir}
	name=${PWD##*/}
	echo "building ${dir} with go:"
	time go build -o "build-${name}" main.go
	echo "building ${dir} with tinygo:"
	time tinygo build -o "build-${name}-tinygo" main.go
done

cd ${basedir}

# show the different sizes
ls -lah */**/**/build-*
ls -lah */**/**/*.wasm
