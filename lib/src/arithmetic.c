#include "types.h"
#include "stack.h"
typedef void (*STC_bin_function)(STC_I64 arg0, STC_I64 arg1);
STC_bin_function add_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{add_I64_I64, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};

void stc_add(STC_TYPE arg0, STC_TYPE arg1, STC_I64 arg2, STC_I64 arg3)
{
	STC_bin_function f = add_funcs[arg0][arg1];
	if(f != 0)
		f(arg2,arg3);
}

void add_I64_I64(STC_I64 arg0, STC_I64 arg1)
{
	push(arg0+arg1, STC_I64_TYPE);
}

void add_F64_F64(STC_I64 arg0, STC_I64 arg1)
{
	push(*(double *)&arg0+*(double *)&arg1, STC_FLOAT_TYPE);
}
