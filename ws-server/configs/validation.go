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

	if valType == validationTypes.CREATE {
		message = "Unable to create room. Please go to home and try again."
	} else if valType == validationTypes.JOIN {
		message = "This room is already full. Please join/create a different room."
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