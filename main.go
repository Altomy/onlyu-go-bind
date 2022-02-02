package onlyugobind

import "errors"

var result string
var err error

func Execute(action string, payload string) (string, error) {

	switch action {
	case "openDatabase":
		result, err = OpenDatabase(payload)
	case "storeUser":
		result, err = StoreUser(payload)
	case "checkUser":
		result, err = CheckUser(payload)
	default:
		result = "NotExecute"
		err = errors.New("not execute")
	}

	return result, err
}
