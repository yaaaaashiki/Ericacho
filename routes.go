package routers

import (
	"github.com/astaxie/beego"
	controllersFacebook "github.com/yaaaaashiki/Ericacho/controller/facebook"
)

func init() {
	beego.Router("/facebook/oauth2", &controllersFacebook.Oauth2Controller{})
	beego.Router("/facebook/callback", &controllersFacebook.CallbackController{})

}
