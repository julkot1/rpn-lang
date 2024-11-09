#include "logic.h"
typedef void (*STC_logic_bin_function)(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
STC_logic_bin_function equals_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_equals_N, 0, stc_equals_N, stc_equals_N, 0, 0, 0, 0},
	{0, stc_equals_N, 0, 0, 0, 0, 0, 0},
	{stc_equals_N, 0, stc_equals_N, 0, 0, 0, 0, 0},
	{stc_equals_N, 0, 0, stc_equals_N, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};

void stc_equals(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_logic_bin_function f = equals_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}
typedef void (*STC_logic_bin_function)(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
STC_logic_bin_function not_equals_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_not_eq_N, 0, stc_not_eq_N, stc_not_eq_N, 0, 0, 0, 0},
	{0, stc_not_eq_N, 0, 0, 0, 0, 0, 0},
	{stc_not_eq_N, 0, stc_not_eq_N, 0, 0, 0, 0, 0},
	{stc_not_eq_N, 0, 0, stc_not_eq_N, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};

void stc_not_equals(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_logic_bin_function f = not_equals_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}
STC_logic_bin_function greater_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_greater_I, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};

void stc_greater(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_logic_bin_function f = greater_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}
STC_logic_bin_function less_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_less_I, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};

void stc_less(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_logic_bin_function f = less_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}
STC_logic_bin_function greater_or_eq_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_greater_eq_I, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};

void stc_greater_or_eq(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_logic_bin_function f = greater_or_eq_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}
STC_logic_bin_function less_or_eq_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_less_eq_I, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};

void stc_less_or_eq(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_logic_bin_function f = less_or_eq_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}
STC_logic_bin_function or_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_or_N, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};

void stc_or(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_logic_bin_function f = or_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}
STC_logic_bin_function and_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_and_N, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};

void stc_and(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_logic_bin_function f = and_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}

void stc_equals_N(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1==arg0, STC_BOOL_TYPE);
}

void stc_greater_I(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1>arg0, STC_BOOL_TYPE);
}

void stc_less_I(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1<arg0, STC_BOOL_TYPE);
}

void stc_less_eq_I(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1<=arg0, STC_BOOL_TYPE);
}

void stc_greater_eq_I(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1>=arg0, STC_BOOL_TYPE);
}

void stc_not_eq_N(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1!=arg0, STC_BOOL_TYPE);
}

void stc_or_N(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1||arg0, STC_BOOL_TYPE);
}

void stc_and_N(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1&&arg0, STC_BOOL_TYPE);
}
