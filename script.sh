#!/bin/bash

repo_dir="/workspace/axolotl"  # Specify the directory name where you want to clone the repository

# Check if the directory already exists
if [ ! -d "$repo_dir" ]; then
    # Directory does not exist, so clone the repository
    mkdir -p /workspace/axolotl
    cp -r /app/axolotl/* "$repo_dir"
else
    # Directory already exists, print a message or take other action
    echo "Directory '$repo_dir' already exists. Skipping clone."
fi

cp /app/chub "$repo_dir"
cd /workspace/axolotl && ./chub