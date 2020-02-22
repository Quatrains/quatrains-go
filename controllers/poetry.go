package controllers

import (
	"quatrains-go/models"

	"github.com/astaxie/beego"
)

// Operations about Poetry
type PoetryController struct {
	beego.Controller
}

// @Title GetDailyPoetry
// @Description get all Interests
// @Success 200 {object} models.DailyPoetry
// @router / [get]
func (i *PoetryController) GetDailyPoetry() {
	dailyPoetry := models.GetDailyPoetry()
	i.Data["json"] = dailyPoetry
	i.ServeJSON()
}
