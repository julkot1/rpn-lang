#define STC(...)
#define CODE(x)
#define NAME(x)
#define CAN_PREVENT
#define FUNCTION(type, vtable, ...)
#define TYPE(...)



/*
    STC - is a base macro; can contain: CODE, NAME, CAN_PREVENT, FUNCTION
    CODE - has int64 code for stc token 
    NAME - has name for function in stc (optional)
    FUNCTION - contains type of vtable, vtable name and n TYPEs
    TYPE - has n values of arguments type for matching function
*/