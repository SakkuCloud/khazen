package model

type Auth struct {
	AccessKey string
	SecretKey string
}

func (a *Auth) IsAuthorized(accessKey string, secretKey string) bool {
	if accessKey == a.AccessKey && secretKey == a.SecretKey {
		return true
	}
	return false
}
