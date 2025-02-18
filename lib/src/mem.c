#include "mem.h"
#include "stc.h"


void stc_memcpy(STC_I64 dest, STC_I64 src, STC_I64 n){
    __asm__ volatile (
        "rep movsb" 
        : "+D" (dest), "+S" (src), "+c" (n) 
        :                                  
        : "memory"
    );
}