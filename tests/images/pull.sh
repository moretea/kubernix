#!/usr/bin/env bash
TEST_DIR=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd); . $TEST_DIR/../common.sh

kubernix::test_start

hello_name=$(nix-build '<nixpkgs>' -A hello | awk -F'/' '{print $4}')

$CRICTL image pull $hello_name
$CRICTL image list | grep -q $hello_name

kubernix::test_finished
