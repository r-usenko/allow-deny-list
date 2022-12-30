package access

type IAccess interface {
	//IsAllowed allowed from list and denied all others
	IsAllowed(value interface{}) error
}

type module struct {
	list     []interface{}
	isDenied bool
}

// Allowed list, other denied
func Allowed(list ...interface{}) IAccess {
	return &module{
		list:     list,
		isDenied: false,
	}
}

// Denied list, other allowed
func Denied(list ...interface{}) IAccess {
	return &module{
		list:     list,
		isDenied: true,
	}
}

// IsAllowed check access
func (m *module) IsAllowed(value interface{}) error {
	if m.checkAccess(value) {
		return nil
	}

	return ErrAccess{
		IsDenied: m.isDenied,
		List:     m.list,
		Value:    value,
	}
}

func (m *module) checkAccess(value interface{}) bool {
	for _, l := range m.list {
		if l == value {
			return !m.isDenied
		}
	}

	return m.isDenied
}
