#include "decorators.h"
#include "types.h"
#include "stack.h"

#if defined(__gnu_linux__) || defined(__linux__) || defined(linux) || defined(__linux)
#include <stdio.h>
#define PRINT_I64(x) printf("%ld\n", x)
#define PRINT_F64(x) printf("%f\n", x)
#define PRINT_I8(x) printf("%c\n", x)
#define PRINT_STRING(x) printf("%s\n", x)
#endif


STC(CODE(5), CAN_PREVENT)
void stc_print(STC_I64 val, STC_TYPE typ);
