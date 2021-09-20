#!/bin/bash

ROOT_DIR=$(pwd)
RESULT_FILE="$ROOT_DIR/results.txt"

function build {
    ./build.sh $1
}

function bench {
    echo "===========================================================================" >> "$RESULT_FILE"
    echo "$1" >> "$RESULT_FILE"
    { time "./bin/$1" ; } 2>> "$RESULT_FILE"
    echo "" >> "$RESULT_FILE"
}

function run_all {
    rm -f "$RESULT_FILE"
    for FILE in $(ls -d impl*)
    do
        pushd $FILE
        build $FILE
        bench $FILE
        popd
    done
}

run_all