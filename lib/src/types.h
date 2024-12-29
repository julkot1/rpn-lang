
#include <math.h>
#if defined(__gnu_linux__) || defined(__linux__) || defined(linux) || defined(__linux)
#include <sys/types.h>
#include <stdint.h>

typedef int64_t STC_I64;
typedef int64_t STC_TYPE;
typedef int64_t STC_SIZE;
typedef double_t STC_F64;
typedef char STC_I8;
typedef char *STC_STR;

#define STC_I64_TYPE 0x0
#define STC_BOOL_TYPE 0x1
#define STC_I8_TYPE 0x2
#define STC_F64_TYPE 0x3
#define STC_STRING_TYPE 0x4
#define STC_ARRAY_TYPE 0x5
#define STC_REF_TYPE 0x6
#define STC_TYPE_TYPE 0x7

#define STC_TYPES_SIZE 8


#define IS_ONE_OF_TYPE(type, type_a, type_b) (type_a == type || type_b == type)
#define IS_ALL_TYPE(type, type_a, type_b) (type_a == type && type_b == type)
#define IS_ALL_SAME_TYPE(type_a, type_b) (type_a == type_b)

#define TO_STC_STRING(x) (STC_STR)(x)
#define TO_INT(x) (STC_I64)(x)

#endif
