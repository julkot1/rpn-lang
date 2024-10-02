#!/bin/bash

opt output.ll -O2 > out_opt.ll
llc out_opt.ll -filetype=obj -relocation-model=pic -o out.o
clang out.o -o out  -lm -fPIE -pie
rm out.o out_opt.ll