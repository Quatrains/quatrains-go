package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["quatrains-go/controllers:InterestController"] = append(beego.GlobalControllerRouter["quatrains-go/controllers:InterestController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quatrains-go/controllers:PoetryController"] = append(beego.GlobalControllerRouter["quatrains-go/controllers:PoetryController"],
        beego.ControllerComments{
            Method: "GetDailyPoetry",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["quatrains-go/controllers:UserInterestController"] = append(beego.GlobalControllerRouter["quatrains-go/controllers:UserInterestController"],
        beego.ControllerComments{
            Method: "CreateUserInterest",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
