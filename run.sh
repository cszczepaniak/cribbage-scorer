#!/bin/bash

ROOT_DIR=$(pwd)

function build {
    ./build.sh $1
}

function bench {
    hyperfine $1 --export-json "$1.json"
}

function run_all {
    for FILE in $(ls -d impl*)
    do
        pushd $FILE
        build $FILE
        bench $FILE
        popd
    done
}

run_all