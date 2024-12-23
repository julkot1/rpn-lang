#include "io.h"


void stc_print(STC_I64 val, STC_TYPE typ){
    char *types[8] = {
        "I64", "Bool", "I8", "F64", "Str", "Array", "Struct", "Type"
    };

    switch (typ) {
        case STC_I64_TYPE:
            PRINT_I64(val);
        break;
        case STC_I8_TYPE:
            PRINT_I8((STC_I8)val);
        break;
        case STC_F64_TYPE:
            PRINT_F64(*(STC_F64*)(&val));
        break;
        case STC_STRING_TYPE:
            PRINT_STRING((STC_STR)val);
        break;
        case STC_TYPE_TYPE:
            PRINT_STRING(types[val]);
        break;
    }
}