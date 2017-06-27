#!/usr/bin/env bash
TEST_DIR=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd); . $TEST_DIR/../common.sh

kubernix::test_start

$CRICTL sandbox run $TEST_DIR/sandbox.json > /dev/null
$CRICTL sandbox ls | grep -q hdishd83djaidwnduwk28bcsb

kubernix::test_finished
