#include "array.h"
#include "stc.h"


void stc_at(STC_I64 idx, STC_I64 array, STC_TYPE idx_typ, STC_TYPE array_typ) 
{
    switch (array_typ) {
        case STC_STRING_TYPE:
        push(TO_INT((TO_STC_STRING(array))[idx]), STC_I8_TYPE);
        break;
    
    }
}


void stc_len(STC_I64 array, STC_TYPE array_typ) 
{
    STC_STR str;
    STC_I64 l=0;
    STC_I8 c;
    switch (array_typ) {
        case STC_STRING_TYPE:
            str = (STC_STR)(array);
            while ((c = *(str+l) )!= '\0') 
               l ++;
            push(l, STC_I64_TYPE);

        break;
    
    }
}