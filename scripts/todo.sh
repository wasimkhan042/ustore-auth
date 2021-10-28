#!/bin/bash

grep -rni -E "TODO|FIXME" *.go **/*.go > TODO.md
sed -i '1s/^/# TODOs\n/' TODO.md