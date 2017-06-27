#!/usr/bin/env bash
TESTS_DIR=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd); . $TESTS_DIR/common.sh

run_test() {
  echo "##########"
  echo "# Testing $1"
  $TESTS_DIR/$1
}


run_test sandbox/create_list.sh
