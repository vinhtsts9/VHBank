package service

type (
	IUserLogin interface{}
)

var (
	localUserLogin IUserLogin
)

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}

func NewUserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement localUserLogin notfound")
	}

	return localUserLogin
}
