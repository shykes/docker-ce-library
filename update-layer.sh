#!/bin/bash

set -x
set -e

list_layers() {
	git config --name-only --get-regexp 'layer\.[^.]*\.url' | sed -e 's/^layer\.//' -e 's/\.url$//'
}

update_sources() {
	list_layers | {
		while read layer; do
			if [ $(git config --get layer.$layer.skip) ]; then
				echo "Skipping $layer"
				continue
			fi
			url=$(git config --get layer.$layer.url)
			branch=$(git config --get layer.$layer.branch)
			git fetch $url $branch:layers/$layer/src
		done
	}
}

update_sources
