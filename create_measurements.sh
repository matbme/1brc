#!/bin/sh

samples=1000000000

if [ "$1" = "small" ]; then
    samples=10000000
fi

cd create || exit
go build
./create $samples
cd ..
