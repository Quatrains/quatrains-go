package models

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserInterest struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	InterestIds string `json:"interest_ids"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	orm.RegisterModel(new(UserInterest))
}

type UserInterestReq struct {
	InterestIds []int `json:"interest_ids"`
}

func CreateUserInterest(userId int, interestIds []int) {
	interestIdsMar, err := json.Marshal(interestIds)
	if err != nil {
		fmt.Println(err)
	}
	o := orm.NewOrm()
	o.Using("default")
	userInterest := UserInterest{UserId: userId, InterestIds: string(interestIdsMar)}
	id, err := o.Insert(&userInterest)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}
}
