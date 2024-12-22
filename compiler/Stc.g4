
grammar Stc;

prog          : (functionDef)+ EOF;

functionDef   : 'fun' ID block; 


subBlock      :( (operation | stackOperation | push | identifier)+ (functionDef | ifBlock | repeatBlock)?  );

block         : '{' (subBlock)+'}';

ifBlock       : 'if' block (elseBlock)?;
elseBlock     : 'else' block;
repeatBlock   : 'repeat' block;

operation     : STACK_PREVENTION? operaor;
operaor       : LOGIC_OPERATOR 
              | ARITHMETIC_OPERATOR 
              | BUILD_IN_OPERATOR
              ;
stackOperation: 'dup' 
              | 'rot' 
              | 'swap' 
              | 'pop'
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
BUILD_IN_OPERATOR  : 'typeof' 
                   | ':=' 
                   ;

ID: [a-zA-Z_][a-zA-Z_0-9]*;  

identifier: ID ((':' ID)+)?;



WS: [ \t\r\n]+ -> skip;  // Ignore whitespace