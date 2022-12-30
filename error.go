package access

import (
	"fmt"
	"strings"
)

var _ error = (*ErrAccess)(nil)

type ErrAccess struct {
	IsDenied bool
	List     []interface{}
	Value    interface{}
}

func (m ErrAccess) Error() string {
	var listStrings = "*"
	var licLen = len(m.List)

	if licLen != 0 {
		var tmp = make([]string, 0, licLen)
		for _, l := range m.List {
			tmp = append(tmp, fmt.Sprint(l))
		}

		listStrings = strings.Join(tmp, ",")
	}

	var tpl = "Available only on [%s], got %s"
	if m.IsDenied {
		tpl = "Unavailable on [%s], got %s"
	}

	return fmt.Sprintf(tpl, listStrings, m.Value)
}
