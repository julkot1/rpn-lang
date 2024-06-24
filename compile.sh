#!/bin/bash

llc out.ll -filetype=obj -o out.o
clang out.o -o out
rm out.o