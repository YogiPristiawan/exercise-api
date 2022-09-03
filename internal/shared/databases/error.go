package databases

import (
	"errors"

	"gorm.io/gorm"
)

func CastDatabaseError(err error) int {
	if err == nil {
		return 0
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 404
	}

	return 500
}
