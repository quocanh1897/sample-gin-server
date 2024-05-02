package constant

import (
	"strings"
)

type Scope string

const (
	AgencyAnyRead   Scope = "org:read"
	AgencyAnyManage Scope = "org:edit"
	UsersSelfRead   Scope = "profile:read"
	UsersSelfManage Scope = "profile:edit"
)

type Role string

const (
	SuperAdmin Role = "SuperAdmin"
)

func MockScopeMapping(email string) []Scope {
	normalizedEmail := strings.ToLower(email)

	if strings.Contains(normalizedEmail, "+customersupport") {
		return []Scope{
			AgencyAnyRead,
			AgencyAnyManage,
			UsersSelfRead,
			UsersSelfManage,
		}
	}
	// otherwise just do SuperAdmin
	return []Scope{
		AgencyAnyRead,
		AgencyAnyManage,
		UsersSelfRead,
		UsersSelfManage,
	}
}
