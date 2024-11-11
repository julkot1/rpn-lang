#include "types.h"
#include "stack.h"
#include "decorators.h"

STC(CODE(2), CAN_PREVENT)
void stc_add(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(3), CAN_PREVENT)
void stc_sub(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(4), CAN_PREVENT)
void stc_mul(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(6), CAN_PREVENT)
void stc_div(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(7), CAN_PREVENT)
void stc_mod(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3); 




STC(CODE(255), CAN_PREVENT, NAME("addN"))
void stc_add_I64_I64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(255), CAN_PREVENT, NAME("addF"))
void stc_add_F64_F64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(255), CAN_PREVENT, NAME("subN"))
void stc_sub_I64_I64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(255), CAN_PREVENT, NAME("subF"))
void stc_sub_F64_F64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(255), CAN_PREVENT, NAME("mulN"))
void stc_mul_I64_I64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(255), CAN_PREVENT, NAME("mulF"))
void stc_mul_F64_F64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(255), CAN_PREVENT, NAME("divN"))
void stc_div_I64_I64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(255), CAN_PREVENT, NAME("divF"))
void stc_div_F64_F64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(255), CAN_PREVENT, NAME("modN"))
void stc_mod_I64_I64(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
