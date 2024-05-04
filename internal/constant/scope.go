package constant

type Scope string

const (
	AdminAnyRead    Scope = "org:read"
	AdminAnyManage  Scope = "org:edit"
	UsersSelfRead   Scope = "profile:read"
	UsersSelfManage Scope = "profile:edit"
)

type Role string

const (
	SuperAdmin Role = "SuperAdmin"
)
