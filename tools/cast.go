package tools

func CheckMapInterfaceInterface(interfaceParam interface{}) bool {
	if interfaceParam == nil {
		return false
	}
	switch interfaceParam.(type) {
	case map[interface{}]interface{}:
		return true
	default:
		return false
	}
}

func CheckMapStringInterface(interfaceParam interface{}) bool {
	if interfaceParam == nil {
		return false
	}
	switch interfaceParam.(type) {
	case map[string]interface{}:
		return true
	default:
		return false
	}
}

func CheckArrayInterface(interfaceParam interface{}) bool {
	if interfaceParam == nil {
		return false
	}
	switch interfaceParam.(type) {
	case []interface{}:
		return true
	default:
		return false
	}
}

func CheckString(interfaceParam interface{}) bool {
	if interfaceParam == nil {
		return false
	}
	switch interfaceParam.(type) {
	case string:
		return true
	default:
		return false
	}
}

func CheckArrayString(interfaceParam interface{}) bool {
	if interfaceParam == nil {
		return false
	}
	switch interfaceParam.(type) {
	case []string:
		return true
	default:
		return false
	}
}
