#include "array.h"
#include "stc.h"


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

void stc_at(STC_I64 idx, STC_I64 array, STC_TYPE idx_typ, STC_TYPE array_typ) 
{
    STC_ARRAY_STRUCT *arr;
    STC_VAR_STRUCT var;
    switch (array_typ) {
        case STC_STRING_TYPE:
            push(TO_INT((TO_STC_STRING(array))[idx]), STC_I8_TYPE);
        break;
        case STC_ARRAY_TYPE: 
            arr = (STC_ARRAY_STRUCT*)array;
            var = (arr->elem[idx]);
            push(var.value, var.type);

        break;
    }
}


void stc_len(STC_I64 array, STC_TYPE array_typ) 
{
    STC_STR str;
    STC_I64 l=0;
    STC_I8 c;
    STC_ARRAY_STRUCT *arr;
    switch (array_typ) {
        case STC_STRING_TYPE:
            str = (STC_STR)(array);
            while ((c = *(str+l) )!= '\0') 
               l ++;
            push(l, STC_I64_TYPE);
        break;
        case STC_ARRAY_TYPE:
            arr = (STC_ARRAY_STRUCT *)array;
            push((*arr).len, STC_I64_TYPE);
        break;
    }
}