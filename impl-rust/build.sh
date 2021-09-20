#!/bin/bash
if [ ! -d "bin" ]
then
    mkdir bin
fi
cargo build --release
cp target/release/*.exe bin/$1