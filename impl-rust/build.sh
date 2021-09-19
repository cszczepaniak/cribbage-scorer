#!/bin/bash
cargo build --release
cp target/release/*.exe $1.exe