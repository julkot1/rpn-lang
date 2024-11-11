#include "arithmetic.h"
#include <stdio.h>

typedef void (*STC_bin_function)(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
STC_bin_function add_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_add_I64_I64, 0, stc_add_I64_I64, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{stc_add_I64_I64, 0, stc_add_I64_I64, 0, 0, 0, 0, 0},
	{0, 0, 0, stc_add_F64_F64, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};
STC_bin_function mul_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_mul_I64_I64, 0, stc_mul_I64_I64, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{stc_mul_I64_I64, 0, stc_mul_I64_I64, 0, 0, 0, 0, 0},
	{0, 0, 0, stc_mul_F64_F64, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};
STC_bin_function sub_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_sub_I64_I64, 0, stc_sub_I64_I64, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{stc_sub_I64_I64, 0, stc_sub_I64_I64, 0, 0, 0, 0, 0},
	{0, 0, 0, stc_sub_F64_F64, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};
STC_bin_function mod_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_mod_I64_I64, 0, stc_mod_I64_I64, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{stc_mod_I64_I64, 0, stc_mod_I64_I64, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};
STC_bin_function div_funcs[STC_TYPES_SIZE][STC_TYPES_SIZE] = {
	{stc_div_I64_I64, 0, stc_div_I64_I64, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{stc_div_I64_I64, 0, stc_div_I64_I64, 0, 0, 0, 0, 0},
	{0, 0, 0, stc_div_F64_F64, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0}
};




void stc_add(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_bin_function f = add_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}

void stc_sub(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_bin_function f = sub_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}

void stc_mul(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_bin_function f = mul_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}


void stc_div(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_bin_function f = div_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}

void stc_mod(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_bin_function f = mod_funcs[arg2][arg3];
	if(f != 0)
		f(arg0,arg1,arg2,arg3);
}



void stc_add_I64_I64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg0+arg1, STC_I64_TYPE);
}

void stc_sub_I64_I64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1-arg0, STC_I64_TYPE);
}

void stc_mul_I64_I64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1*arg0, STC_I64_TYPE);
}

void stc_div_I64_I64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1/arg0, STC_I64_TYPE);
}

void stc_mod_I64_I64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	push(arg1%arg0, STC_I64_TYPE);
}

void stc_add_F64_F64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_F64 x = *(STC_F64 *)(&arg0)+*(STC_F64 *)(&arg1);
	push(*(STC_I64 *)&x, STC_F64_TYPE);
}
void stc_sub_F64_F64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_F64 x = *(STC_F64 *)(&arg1)-*(STC_F64 *)(&arg0);
	push(*(STC_I64 *)&x, STC_F64_TYPE);
}
void stc_div_F64_F64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_F64 x = (*(STC_F64 *)(&arg1)) / (*(STC_F64 *)(&arg0));
	push(*(STC_I64 *)&x, STC_F64_TYPE);
}
void stc_mul_F64_F64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3)
{
	STC_F64 x = (*(STC_F64 *)(&arg1)) * (*(STC_F64 *)(&arg0));
	push(*(STC_I64 *)&x, STC_F64_TYPE);
}
