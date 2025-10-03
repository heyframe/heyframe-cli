#!/usr/bin/env bash

rm -rf completions
mkdir completions
go run . completion bash > completions/heyframe-cli.bash
go run . completion zsh > completions/heyframe-cli.zsh
go run . completion fish > completions/heyframe-cli.fish