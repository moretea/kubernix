set -euo pipefail

KUBERNIX_DIR=$(cd $(dirname "${BASH_SOURCE[0]}")/.. && pwd)
KUBERNIX_BIN=${KUBERNIX_BIN:-$KUBERNIX_DIR/kubernix}

kubernix::test::ensure_bin() {
  if [ ! -f $KUBERNIX_BIN ]; then
    echo "kubernix is not build yet!"
    echo "build it with:"
    echo "  go build github.com/moretea/kubernix/cmd/kubernix"
    exit 1
  fi
}

kubernix::test_start() {
  kubernix::test::ensure_bin

  kubernix_test_scratch=$(mktemp -d -t tmp.XXXXXXXXXX)
  cd $kubernix_test_scratch
  $KUBERNIX_BIN --socket ./kubernix.sock &
  KUBERNIX_PID=$!

  finish() {
    kill -KILL $KUBERNIX_PID
    rm -rf "$kubernix_test_scratch"
    echo "TEST FAILED"
    exit 1
  }

  trap finish EXIT

  while [ ! -S ./kubernix.sock ]; do
    echo "waiting for kubernix socket .."
    sleep 0.1
  done
  CRICTL="crictl -r ./kubernix.sock"
}

kubernix::test_finished() {
  echo "Test passed!"
  trap '' EXIT
  set +x
  kill -INT $KUBERNIX_PID
  rm -rf "$kubernix_test_scratch"
  exit 0
}
