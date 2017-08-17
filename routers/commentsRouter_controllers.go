package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["RESTv1/controllers:TUser2Controller"] = append(beego.GlobalControllerRouter["RESTv1/controllers:TUser2Controller"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["RESTv1/controllers:TUser2Controller"] = append(beego.GlobalControllerRouter["RESTv1/controllers:TUser2Controller"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["RESTv1/controllers:TUser2Controller"] = append(beego.GlobalControllerRouter["RESTv1/controllers:TUser2Controller"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["RESTv1/controllers:TUser2Controller"] = append(beego.GlobalControllerRouter["RESTv1/controllers:TUser2Controller"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["RESTv1/controllers:TUser2Controller"] = append(beego.GlobalControllerRouter["RESTv1/controllers:TUser2Controller"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["RESTv1/controllers:TUserController"] = append(beego.GlobalControllerRouter["RESTv1/controllers:TUserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["RESTv1/controllers:TUserController"] = append(beego.GlobalControllerRouter["RESTv1/controllers:TUserController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["RESTv1/controllers:TUserController"] = append(beego.GlobalControllerRouter["RESTv1/controllers:TUserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["RESTv1/controllers:TUserController"] = append(beego.GlobalControllerRouter["RESTv1/controllers:TUserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["RESTv1/controllers:TUserController"] = append(beego.GlobalControllerRouter["RESTv1/controllers:TUserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
