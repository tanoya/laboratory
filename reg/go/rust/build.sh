#! /bin/bash
CGO_LDFLAGS="-L$(pwd)/lib" go build ./main.go