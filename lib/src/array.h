#include "stc.h"


STC(CODE(37), CAN_PREVENT)
void stc_set(STC_I64 idx, STC_I64 array, STC_I64 val, STC_TYPE idx_typ, STC_TYPE array_typ, STC_TYPE val_typ);

STC(CODE(33), CAN_PREVENT)
void stc_at(STC_I64 idx, STC_I64 array, STC_TYPE idx_typ, STC_TYPE array_typ);

STC(CODE(34), CAN_PREVENT)
void stc_len(STC_I64 array, STC_TYPE array_typ);
