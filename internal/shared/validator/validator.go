package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func CustomErrorMessage(vError error) (err error) {
	if obj, ok := vError.(validator.ValidationErrors); ok {
		for _, vError := range obj {
			switch vError.Tag() {
			case "required":
				return fmt.Errorf("%s harus diisi", vError.Field())
			case "max":
				return fmt.Errorf("%s tidak boleh lebih dari %s ", vError.Field(), vError.Param())
			case "email":
				return fmt.Errorf("%s harus berupa email yang valid", vError.Field())
			case "uuidv4":
				return fmt.Errorf("%s harus berupa uuidV4 yang valid", vError.Field())
			default:
				return vError
			}
		}
	}
	return
}
