#!/usr/bin/env bash
currentDir=$(cd "$(dirname "$0")" && pwd)
cd "$currentDir/proto" && {
  protoc --go_out="$currentDir/" *.proto
}