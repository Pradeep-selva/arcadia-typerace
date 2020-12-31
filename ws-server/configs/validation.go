package configs

import (
	"fmt"
)

var validationTypes = struct {
	JOIN string
	CREATE string
}{"join", "create"}

func GetIfRequiredUsersPresent(valType string, users int) error {
	if valType == validationTypes.JOIN &&  users == 1 {
		return nil
	} else if valType == validationTypes.CREATE && users == 0 {
		return nil
	} else {
		return fmt.Errorf("Unable to "+valType+" room.")
	}
}