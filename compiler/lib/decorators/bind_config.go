package decorators

import (
	"errors"
)

type PreprocessorMacro int

type STCBindConfig struct {
	CanPrevent bool
	Code       int
	Name       string
}

type TypesMatch struct {
	Types    []string
	Function string
}

func CreateSTCBindConfig(root PreprocessorToken) (*STCBindConfig, error) {
	if root.Value != MarcosName[MACRO_STC] {
		return nil, errors.New("config must start with macro STC(...)")
	}
	cp := getCanPrevent(root)
	code, er1 := getCode(root)
	name, er1 := getName(root)
	if er1 != nil {
		return nil, er1
	}

	return &STCBindConfig{
		CanPrevent: cp,
		Code:       code,
		Name:       name,
	}, nil

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
