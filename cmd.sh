#!/bin/bash

cmd_line_args=("$@")

declare -a allowed_cmds=("gen" "clean" "init")
declare -a folders_to_clean=("node_modules" "build" "dist" ".swc" ".gradle")

for i in "${cmd_line_args[@]}"
do
	if [[ ! " ${allowed_cmds[@]} " =~ " ${i} " ]]; then
		echo "Invalid command: $i"
		echo "Allowed commands: ${allowed_cmds[@]}"
		exit 1
	fi
done

function contains() {
	local n=$#
	local value=${!n}
	for ((i=1; i < $#; i++)) {
		if [ "${!i}" == "${value}" ]; then
			return 0
		fi
	}
	return 1
}

function clean() {
	make remove-gen-proto
	make remove-gen-types

	for i in "${folders_to_clean[@]}"
	do
		find . -name "$i" -type d -prune -exec rm -rf '{}' +
	done
}

function generate () {
	make gen-proto
	make gen-types
}

if contains "$@" "clean"; then
	clean
fi

if contains "$@" "gen"; then
	generate
fi

if contains "$@ " "init"; then
	generate
fi
