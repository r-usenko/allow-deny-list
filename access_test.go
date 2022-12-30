package access_test

import (
	"testing"

	access "github.com/r-usenko/allow-deny-list"
)

const (
	One = iota
	Two
)

func TestEmptyList(t *testing.T) {
	//Allow: NONE, Deny: ALL
	if access.Allowed().IsAllowed(One) == nil {
		t.Error("License allowed but allowed list are empty")
	}

	//Deny: NONE, Allow: ALL
	if access.Denied().IsAllowed(One) != nil {
		t.Error("License allowed but denied list are empty")
	}
}

func TestExists(t *testing.T) {
	//exist in denied
	if access.Denied(One, Two).IsAllowed(One) == nil {
		t.Error("License allowed but contains in denied list")
	}

	//exist allowed
	if access.Allowed(One, Two).IsAllowed(One) != nil {
		t.Error("License denied but contains in allowed list")
	}

	//missing in denied, then allowed
	if access.Denied(Two).IsAllowed(One) != nil {
		t.Error("License denied but missing in denied list")
	}

	//missing in allowed, then denied
	if access.Allowed(Two).IsAllowed(One) == nil {
		t.Error("License allowed but missing in allowed list")
	}
}
