package models

import (
	"fmt"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Interest struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type InterestList struct {
	Total   int         `json:"total"`
	Page    int         `json:"page"`
	Ipp     int         `json:"ipp"`
	Objects []*Interest `json:"objects"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	orm.RegisterModel(new(Interest))
}

// GetAllInterests - return all interests by paginate
func GetAllInterests(page, ipp int) InterestList {
	o := orm.NewOrm()
	o.Using("default")
	var interests []*Interest

	_, err := o.QueryTable("interest").Limit(ipp, page-1).All(&interests)
	if err != nil {
		fmt.Println(err)
	}

	il := InterestList{Total: 10, Page: page, Ipp: ipp, Objects: interests}

	return il
}
