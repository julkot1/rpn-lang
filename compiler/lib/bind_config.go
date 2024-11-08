package lib

import (
	"errors"
	"strconv"
)

type PreprocessorMacro int

type STCBindConfig struct {
	CanPrevent bool
	Code       int
	Name       string
	Matrix     Matrix
}

type TypesMatch struct {
	Types    []string
	Function string
}

type Matrix struct {
	VtableName   string
	VtableType   string
	TypesMatches []TypesMatch
}

func CreateSTCBindConfig(root PreprocessorToken) (*STCBindConfig, error) {
	if root.Value != MarcosName[MACRO_STC] {
		return nil, errors.New("config must start with macro STC(...)")
	}
	cp := getCanPrevent(root)
	code, er1 := getCode(root)
	name, er1 := getName(root)
	matrix, er1 := getMatrix(root)
	if er1 != nil {
		return nil, er1
	}

	return &STCBindConfig{
		CanPrevent: cp,
		Code:       code,
		Name:       name,
		Matrix:     *matrix,
	}, nil

}

func getMatrix(root PreprocessorToken) (*Matrix, error) {
	var function PreprocessorToken
	for _, arg := range root.Args {
		if arg.Text == MarcosName[MACRO_FUNCTION] {
			err := isDefinitionGood(arg)
			if err != nil {
				return nil, err
			}
			function = arg
		}
	}
	if function.Value == "" {
		return nil, errors.New("no function")
	}

	matrix := &Matrix{VtableType: function.Args[0].Value.(string), VtableName: function.Args[1].Value.(string), TypesMatches: make([]TypesMatch, 0)}
	for _, typ := range function.Args[2:] {
		types := TypesMatch{Types: make([]string, 0)}
		for _, t := range typ.Args {
			types.Types = append(types.Types, t.Value.(string))
		}
		matrix.TypesMatches = append(matrix.TypesMatches, types)
	}

	return matrix, nil

}

func isDefinitionGood(arg PreprocessorToken) error {
	if len(arg.Args) < 3 {
		return errors.New("function must contain at least 3 arguments")
	}
	if arg.Args[0].Typ != PREPROCESSOR_VAR_T {
		return errors.New("function argument[0] must be a variable")
	}
	if arg.Args[1].Typ != PREPROCESSOR_VAR_T {
		return errors.New("function argument[0] must be a variable")
	}
	for idx, a := range arg.Args[2:] {
		if a.Text != MarcosName[MACRO_TYPE] {
			return errors.New("function argument[" + strconv.Itoa(idx) + "] must be a TYPE macro")
		}
	}
	argc := len(arg.Args[2].Args[0].Args)
	for _, x := range arg.Args[2].Args[0].Args {
		for idx, a := range x.Args {
			if len(a.Args) != argc {
				return errors.New("function argument[" + strconv.Itoa(idx) + "] invalid argument length (expected " + strconv.Itoa(argc) + ")")
			}
			for i, t := range arg.Args {
				if t.Typ != PREPROCESSOR_VAR_T {
					return errors.New("function argument[" + strconv.Itoa(idx) + "]: type argument [" + strconv.Itoa(i) + "] must be a variable")
				}
			}
		}
	}

	return nil
}

func getName(root PreprocessorToken) (string, error) {
	for _, arg := range root.Args {
		if arg.Text == MarcosName[MACRO_NAME] {
			if len(arg.Args) != 1 {
				return "", errors.New("code must contain one arg")
			}
			if arg.Args[0].Typ != PREPROCESSOR_STRING {
				return "", errors.New("code must contain string type arg")
			}
			return arg.Args[0].Value.(string), nil
		}
	}
	return "", nil
}

func getCode(root PreprocessorToken) (int, error) {
	for _, arg := range root.Args {
		if arg.Text == MarcosName[MACRO_CODE] {
			if len(arg.Args) != 1 {
				return -1, errors.New("code must contain one arg")
			}
			if arg.Args[0].Typ != PREPROCESSOR_INT_T {
				return -1, errors.New("code must contain int type arg")
			}
			return arg.Args[0].Value.(int), nil
		}
	}
	return -1, errors.New("no code")
}

func getCanPrevent(root PreprocessorToken) bool {
	for _, arg := range root.Args {
		if arg.Text == MarcosName[MACRO_CAN_PREVENT] {
			return true
		}
	}
	return false
}
