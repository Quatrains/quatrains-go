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
	var body map[string]interface{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(body)
	}
	var interestIds []int
	// json.Unmarshal set numbers'type is float64
	for _, interestID := range body["interest_ids"].([]interface{}) {
		interestIds = append(interestIds, int(interestID.(float64)))
	}

	models.CreateUserInterest(13, interestIds)
	this.Data["json"] = "{}"
	this.ServeJSON()
}
