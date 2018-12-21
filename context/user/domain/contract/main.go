package contract

import user "github.com/amiralii/goEchoExample/context/user/domain/model"

type Repository interface {
	Save(u user.User) error
	GetByCredential(username string, password string) (user.User, error)
}
