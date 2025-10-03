#!/usr/bin/env bash

rm -rf completions
mkdir completions
go run . completion bash > completions/heyFrame-cli.bash
go run . completion zsh > completions/heyFrame-cli.zsh
go run . completion fish > completions/heyFrame-cli.fish