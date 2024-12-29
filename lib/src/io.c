#include "io.h"

struct STC_VAR_STRUCT {
    STC_I64 value; 
    STC_I64 type; 
};

typedef struct STC_VAR_STRUCT STC_VAR_STRUCT ;



struct STC_ARRAY_STRUCT {
    STC_I64 len; 
    STC_I64 capacity; 
    STC_VAR_STRUCT *elem;
};

typedef struct STC_ARRAY_STRUCT STC_ARRAY_STRUCT ;


void stc_print(STC_I64 val, STC_TYPE typ){
    char *types[8] = {
        "I64", "Bool", "I8", "F64", "Str", "Array", "Ref", "Type"
    };
    STC_ARRAY_STRUCT *ar;
    STC_VAR_STRUCT el;

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
        case STC_ARRAY_TYPE:
            ar = (STC_ARRAY_STRUCT*)val;
            PRINT_STRING("[");
            for (STC_I64 i = 0; i < ar->len; i++) {
                el = ar->elem[i];
                stc_print(el.value, el.type);
                if (i + 1 != ar->len) 
                    PRINT_STRING(" ");
            }
            PRINT_STRING("]");
        break;

    }
}