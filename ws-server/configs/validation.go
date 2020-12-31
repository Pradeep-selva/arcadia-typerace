package configs

import (
	"fmt"
)

var validationTypes = struct {
	JOIN string
	CREATE string
}{"join", "create"}

func GetIfRequiredUsersPresent(valType string, users int) error {
	message := ""

	if valType == validationTypes.JOIN || valType == validationTypes.CREATE {
		message = "Unable to "+valType+" room. Please go to home."
	} else {
		message = "Invalid operation encountered. Please go to home."
	}

	if valType == validationTypes.JOIN &&  users == 1 {
		return nil
	} else if valType == validationTypes.CREATE && users == 0 {
		return nil
	} else {
		return fmt.Errorf(message)
	}
}