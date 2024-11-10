#include "types.h"
#include "stack.h"
#include "decorators.h"

STC(CODE(22), CAN_PREVENT)
void stc_equals(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(28), CAN_PREVENT)
void stc_not_equals(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(24), CAN_PREVENT)
void stc_greater(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(25), CAN_PREVENT)
void stc_less(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(26), CAN_PREVENT)
void stc_greater_or_eq(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(27), CAN_PREVENT)
void stc_less_or_eq(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(29), CAN_PREVENT)
void stc_or(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

STC(CODE(30), CAN_PREVENT)
void stc_and(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);

void stc_equals_N(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
void stc_greater_I(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
void stc_less_I(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
void stc_less_eq_I(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
void stc_greater_eq_I(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
void stc_not_eq_N(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
void stc_or_N(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
void stc_and_N(STC_I64 arg0, STC_I64 arg1, STC_TYPE arg2, STC_TYPE arg3);
