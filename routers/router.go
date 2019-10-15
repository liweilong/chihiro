package routers

import (
	"github.com/liweilong/chihiro/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
