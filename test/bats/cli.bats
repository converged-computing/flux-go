#!/usr/bin/env bats

# This is run relative to the root
# We also assume the binaries are built
BIN_DIR=./bin
TEST_DIR=./test/bats
SLEEP_TIME=10

@test "keygen test" {

  ${BIN_DIR}/fluxgo-keygen | grep public-key
  cat ./curve.cert | grep ZeroMQ
  cat ./curve.cert | grep curve-cert
  cat ./curve.cert | grep keygen.hostname
  cat ./curve.cert | grep public-key
  cat ./curve.cert | grep secret-key
}

@test "submit test" {
  flux start ${TEST_DIR}/test_submit.sh
}
