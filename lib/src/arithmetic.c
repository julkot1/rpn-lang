#include "arithmetic.h"

STC_TYPE stc_get_add_type(STC_TYPE type_a, STC_TYPE type_b) {
    if (IS_ALL_SAME_TYPE(type_a, type_b)) {
        return type_a;
    }
    else if (IS_ONE_OF_TYPE(STC_STRING_TYPE, type_a, type_b)) {
        return STC_STRING_TYPE;
    }
    else if (IS_ONE_OF_TYPE(STC_BOOL_TYPE, type_a, type_b)) {
        return STC_BOOL_TYPE;
    }
    else if (IS_ONE_OF_TYPE(STC_FLOAT_TYPE, type_a, type_b)) {
        return STC_FLOAT_TYPE;
    }
    return 0x0;
}

STC_I64 stc_add(STC_I64 a, STC_I64 b, STC_TYPE type_a, STC_TYPE type_b, STC_TYPE result) {
    switch (result) {
        case STC_INT_TYPE:
            if (IS_ALL_SAME_TYPE(type_a, type_b)) {
                return a + b;
            }
            break;
        case STC_FLOAT_TYPE:
            return (STC_I64)((double)a + (double)b);
        case STC_STRING_TYPE:
            return 0x0;
        case STC_BOOL_TYPE:
            return 1;
        default:
            return 0x0;

    }
    return 0x0;
}
