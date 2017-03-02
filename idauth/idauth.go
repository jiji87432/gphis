// gphis/idauth - Identity Authentication 身份认证模块
//
package idauth

type IAuth interface {
	Authenticate() bool
}

func Auth(ia IAuth) bool {
	return ia.Authenticate()
}
