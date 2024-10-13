#!/bin/bash
cd ..
mkdir build
cd build
opt output.ll -O2  -S > out_opt.ll
llc out_opt.ll -filetype=obj -relocation-model=pic -o out.o
/usr/lib/llvm14/bin/clang out.o -o out  -lm -fPIE -pie
#rm out.o out_opt.ll14.0wl;