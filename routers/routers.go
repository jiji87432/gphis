package routers

import (
	"github.com/astaxie/beego"
	"gphis/controllers"
)

func init() {
	beego.Router("/", &controllers.DefaultController{})
	beego.Router("/auth/:type", &controllers.AuthController{})
}
