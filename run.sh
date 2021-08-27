#!/bin/bash

ROOT_DIR=$(pwd)

function build {
    ./build.sh $1
}

function bench {
    hyperfine "$1.exe" --export-json "bench-$2.json"
}

function run_all {
    for FILE in $(ls -d impl*)
    do
        pushd $FILE
        build $FILE
        bench $FILE "$ROOT_DIR/$FILE"
        popd
    done
}

run_all