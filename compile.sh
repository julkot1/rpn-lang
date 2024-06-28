#!/bin/bash

opt out.ll -O2 > out_opt.ll
llc out_opt.ll -filetype=obj -o out.o
clang out.o -o out  -lm
rm out.o out_opt.ll