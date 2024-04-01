#!/bin/bash

cmd_line_args=("$@")

declare -a allowed_cmds=("gen" "clean" "init" "regen")
declare -a folders_to_clean=("node_modules" "build" "dist" ".swc" ".gradle")

# if no command is passed, error
if [ $# -eq 0 ]; then
	echo "No command passed"
	echo "Allowed commands: ${allowed_cmds[@]}"
	exit 1
fi

# if more than one command is passed, error
if [ $# -gt 1 ]; then
	echo "Only one command is allowed"
	echo "Allowed commands: ${allowed_cmds[@]}"
	exit 1
fi

# if command is not allowed, error
cmd_allowed=false
for i in "${allowed_cmds[@]}"
do
	if [ $1 == $i ]; then
		cmd_allowed=true
		break
	fi
done


if ! $cmd_allowed; then
	echo "Invalid command: $1"
	echo "Allowed commands: ${allowed_cmds[@]}"
	exit 1
fi

# function to clean the project
function clean() {
	make remove-gen-proto
	make remove-gen-types

	for i in "${folders_to_clean[@]}"
	do
		find . -name "$i" -type d -prune -exec rm -rf '{}' +
	done
}

# function to generate the project
function generate () {
	make gen-proto
	make gen-types
}

case $1 in 
	"clean") clean;;
	"gen") generate;;
	"init") generate;;
	"regen") clean; generate;;
esac
