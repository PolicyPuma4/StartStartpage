#!/bin/bash

docker buildx build --platform linux/386,linux/amd64,linux/arm/v6,linux/arm/v7,linux/arm64/v8,linux/ppc64le,linux/s390x --tag policypuma4/startstartpage . --push
