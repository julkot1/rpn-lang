//
// Created by julian on 10/2/24.
//

#ifndef STCLIBC_ARITHMETIC_H
#define STCLIBC_ARITHMETIC_H
#include "types.h"


STC_TYPE stc_add_type(STC_TYPE type_a, STC_TYPE type_b);
void stc_add(STC_I64 a, STC_I64 b, STC_TYPE type_a, STC_TYPE type_b);



#endif //STCLIBC_ARITHMETIC_H
