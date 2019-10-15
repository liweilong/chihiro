package routers

import (
	"github.com/astaxie/beego"
	"github.com/liweilong/chihiro/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
