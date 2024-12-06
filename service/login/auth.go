package login

type AuthRepo interface {
	Login(username, password string) (int64, error)
}
