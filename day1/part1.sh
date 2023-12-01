#!/bin/sh

cat input | sed 's/[^0-9]//g' | sed -E 's/(.).*(.)/\1\2/' | awk 'BEGIN { sum = 0 } { if ($1 < 10 ) { sum += $1$1; } else { sum += $1; }} END { print sum }'

