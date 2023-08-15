#!/bin/bash

set -e

for script in .devcontainer/postCreate.d/* ; do
    $script
done