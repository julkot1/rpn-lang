
grammar Stc;

prog          : (functionDef)+ EOF;

functionDef   : 'fun' ID arguments? block; 


subBlock      :( (arrayIndex| arrayNew | array | operation | stackOperation | push | varReference | varAssign | identifier)+ (functionDef | ifBlock | repeatBlock)?  );

block         : '{' (subBlock)+'}';

ifBlock       : 'if' block (elseBlock)?;
elseBlock     : 'else' block;
repeatBlock   : 'repeat' arguments? block;

arguments : '(' argument+ ')' ;

operation     : STACK_PREVENTION? operaor;
operaor       : LOGIC_OPERATOR 
              | ARITHMETIC_OPERATOR 
              | BUILD_IN_OPERATOR
              | ASSIGN_OPERATOR
              ;
stackOperation: 'dup' 
              | 'rot' 
              | 'swap' 
              | 'pop'
              | 'over'
              ;
push          : SIGNED_NUMBER 
              | NUMBER
              | SIGNED_FLOAT 
              | FLOAT 
              | SIGNED_FLOAT_E 
              | FLOAT_E 
              | BIN_NUMBER 
              | HEX_NUMBER 
              | CHAR 
              | STRING 
              | BOOL 
              | SIMPLE_TYPE
              ;
arrayElement  : SIGNED_NUMBER 
              | NUMBER
              | SIGNED_FLOAT 
              | FLOAT 
              | SIGNED_FLOAT_E 
              | FLOAT_E 
              | BIN_NUMBER 
              | HEX_NUMBER 
              | CHAR 
              | STRING 
              | BOOL 
              | SIMPLE_TYPE
              | arrayIndex
              | varIdentifier
              | array
              ;


arrayIndex: varIdentifier '@' (NUMBER|varIdentifier );



array: ARRAY_OPEN (arrayElement)+ ARRAY_CLOSE;
arrayDescriber: ARRAY_OPEN NUMBER ARRAY_CLOSE;
arrayNew: ARRAY_OPERATOR arrayDescriber arrayDescriber?;


argument    : ID;
varAssign   : (varIdentifier|arrayIndex) ASSIGN_OPERATOR;
varReference: REFERENCE_OPERATOR varIdentifier;
varIdentifier: ID ((':' ID)+)?;
identifier: (STACK_PREVENTION)? ID ((':' ID)+)?;




STACK_PREVENTION: '!';

NUMBER: DIGIT+; // Unsigned number
SIGNED_NUMBER: [+-] DIGIT+; // Signed number
FLOAT: DIGIT+ '.' DIGIT+; // Unsigned float
SIGNED_FLOAT: [+-] DIGIT+ '.' DIGIT+; // Signed float
FLOAT_E: DIGIT+ ('.' DIGIT+)? [eE] [+-]? DIGIT+; // Scientific notation
SIGNED_FLOAT_E: [+-] DIGIT+ ('.' DIGIT+)? [eE] [+-]? DIGIT+; // Signed scientific notation

fragment DIGIT: [0-9];

BIN_NUMBER: '0' [bB] [01]+;
HEX_NUMBER:  '0' [xX] [0-9a-fA-F]+;
STRING: '"' (ESC | ~["\\])* '"';     
CHAR: '\'' (ESC | ~['\\]) '\'';
BOOL: 'true' | 'false';
fragment ESC: '\\' [btnfr"'\\];
SIMPLE_TYPE        : 'I64' 
                   | 'I8' 
                   | 'Str' 
                   | 'Float' 
                   | 'Bool'
                   ;
LOGIC_OPERATOR     : '<=' 
                   | '>='
                   | '!='
                   | '==' 
                   | '<' 
                   | '>' 
                   | 'and' 
                   | 'or' 
                   | 'not'
                   ;
ARITHMETIC_OPERATOR: '+' 
                   | '-'
                   | '*' 
                   | '/' 
                   | '%' 
                   ;  
ASSIGN_OPERATOR    : ':=';
REFERENCE_OPERATOR : '&';
BUILD_IN_OPERATOR  : 'typeof'
                   | 'at'
                   | 'len'
                   | 'call'
                   ;


ARRAY_OPEN: '[';
ARRAY_CLOSE: ']';
ARRAY_OPERATOR: 'arr';

ID: [a-zA-Z_][a-zA-Z_0-9]*;  





WS: [ \t\r\n]+ -> skip;  // Ignore whitespace