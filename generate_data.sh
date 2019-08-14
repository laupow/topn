#!/bin/bash

ITERATIONS=$(( $1 ))
OUTPUT_PATH=$2

echo "Generating numbers in $OUTPUT_PATH"

x=1
while [ $x -le $ITERATIONS ]
do
  echo "Iteration: $x/$ITERATIONS..."
  x=$(( $x + 1 ))
  shuf -i 10-1000000000000 -n 500000 >> $OUTPUT_PATH
done