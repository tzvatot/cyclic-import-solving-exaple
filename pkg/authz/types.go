package authz

type Action string
type Resource string

const (
	CreateAction Action = "create"
	GetAction    Action = "get"
	UpdateAction Action = "update"
	DeleteAction Action = "delete"

	UserResource Resource = "user"
)
