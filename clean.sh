#!/bin/bash

make remove-gen-proto
make remove-gen-types

declare -a folders=( "node_modules" "build" "dist" ".swc" ".gradle")

for i in "${folders[@]}"
do
  find . -name "$i" -type d -prune -exec rm -rf '{}' +
done

make gen-proto
make gen-types
