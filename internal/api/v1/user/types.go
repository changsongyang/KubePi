package user

import v1User "github.com/KubeOperator/kubepi/internal/model/v1/user"

type User struct {
	v1User.User
	Roles []string `json:"roles"`
	Password string `json:"password"`
}

