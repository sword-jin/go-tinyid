#!/bin/bash

set -e
set -x

DIR=$(cd `dirname $0` && pwd -P)
cd $DIR

go run .
