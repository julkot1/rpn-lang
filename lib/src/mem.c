#include "mem.h"
#include "stc.h"
#include "string.h"
#include "stdio.h"

struct point {
    STC_I64 x, y,z;
};

void stc_memcpy(STC_I64 dest,register STC_I64 src, STC_I64 n, STC_TYPE destT,  STC_TYPE srcT,  STC_TYPE nT){
    printf("%ld %ld %ld\n", (*(struct point*)src).x, (*(struct point*)src).y, (*(struct point*)src).z);
    while (n--) {
        ((char*)dest)[n-1] = ((char*)src)[n-1];
    }
}