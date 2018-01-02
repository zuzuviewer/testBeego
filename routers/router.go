package routers

import (
	"github.com/astaxie/beego"
	"testBeego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/query/scanner/result", &controllers.ScannerController{})
}
