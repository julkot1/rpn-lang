package lang

type Type int64

const (
	INT_T Type = iota
	BOOL_T
	CHAR_T
	FLOAT_T
	STRING_T
	ARRAY_T
	REF_T
	Type_T
	Struct_T
	ANY_T
)

func StringToType(s string, program *Program) Type {
	newS := s[1 : len(s)-1]
	switch newS {
	case "I64":
		return INT_T
	case "Bool":
		return BOOL_T
	case "I8":
		return CHAR_T
	case "F64":
		return FLOAT_T
	case "Str":
		return STRING_T
	case "Arr":
		return ARRAY_T
	case "Ref":
		return REF_T
	case "Type":
		return Type_T
	case "Any":
		return ANY_T
	}

	if isStruct(newS, program) {
		return Struct_T
	}

	return 0
}

func isStruct(s string, program *Program) bool {
	for _, x := range program.StcStruct {
		if x.Name == s {
			return true
		}
	}
	return false
}

type PushableToken struct {
	Value interface{}
	Typ   Type
}
