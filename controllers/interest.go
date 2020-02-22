package controllers

import (
	"quatrains-go/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about Interest
type InterestController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all Interests
// @Param	page	query	integer	1	false	"page"
// @Param	ipp		query	integer	10	false	"ipp"
// @Success 200 {object} models.InterestList
// @router / [get]
func (i *InterestController) GetAll() {
	page, _ := strconv.Atoi(i.Input().Get("page"))
	if page == 0 {
		page = 1
	}
	ipp, _ := strconv.Atoi(i.Input().Get("ipp"))
	if ipp == 0 {
		ipp = 10
	}
	interests := models.GetAllInterests(page, ipp)
	i.Data["json"] = interests
	i.ServeJSON()
}
