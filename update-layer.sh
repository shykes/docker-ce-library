#!/bin/bash

set -e

list_layers() {
	git config --name-only --get-regexp 'layer\.[^.]*\.url' | sed -e 's/^layer\.//' -e 's/\.url$//'
}

update_sources() {
	echo "Checking layer remote sources for updates"
	list_layers | {
		while read layer; do
			if [ $(git config --get layer.$layer.skip) ]; then
				echo "[$layer] Skipping"
				continue
			fi
			url=$(git config --get layer.$layer.url)
			branch=$(git config --get layer.$layer.branch)
			echo "[$layer] fetching $url $branch"
			git fetch $url $branch:layers/$layer/src
		done
	}
}

update_sources
