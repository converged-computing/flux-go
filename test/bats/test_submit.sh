#!/bin/bash

HERE=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd $HERE
../../bin/fluxgo-submit
flux jobs -a | grep sleep

# Give time to finish
sleep 10