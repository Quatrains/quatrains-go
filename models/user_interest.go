package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserInterest struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	InterestIds string    `json:"interest_ids"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func _createUserInterestPtr(userId int, interestIds []int) *UserInterest {
	interestIdsMar, err := json.Marshal(interestIds)
	if err != nil {
		fmt.Println(err)
	}
	userInterest := UserInterest{
		UserId:      userId,
		InterestIds: string(interestIdsMar),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return &userInterest
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
	o := orm.NewOrm()
	o.Using("default")
	userInterestPtr := _createUserInterestPtr(userId, interestIds)

	id, err := o.Insert(userInterestPtr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}
}
