#!/usr/bin/env bash

# Build the source file
#make dev-docker

# Run it
docker run --rm -it $(docker build -q .)
