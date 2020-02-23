package controllers

import (
	"encoding/json"
	"fmt"
	"quatrains-go/models"

	"github.com/astaxie/beego"
)

type UserInterestController struct {
	beego.Controller
}

// @Title CreateUserInterest
// @Description get all Interests
// @Param	body	body models.UserInterestReq true
// @Success 200 {}
// @router / [post]
func (this *UserInterestController) CreateUserInterest() {
	var body map[string][]int
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(body)
	}

	interestIds := body["interest_ids"]
	models.CreateUserInterest(12, interestIds)
	this.Data["json"] = "{}"
	this.ServeJSON()
}
