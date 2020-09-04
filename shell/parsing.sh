#!/bin/sh

cat ../data/CORPCODE.xml | grep '<stock_code>' | sed 's/[^0-9]//g' | grep . > list.txt