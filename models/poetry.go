package models

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/*
# 推荐系统产出的诗歌id为idx
    idx = pw.IntegerField(unique=True)
    title = pw.CharField(max_length=32)
    title_tr = pw.CharField(max_length=32, default="")
    author = pw.CharField(max_length=32)
    author_tr = pw.CharField(max_length=32, default="")
    content = JSONCharField(max_length=4096, default=list)
    content_tr = JSONCharField(max_length=4096, default=list)
    background = pw.CharField(max_length=2048, default="")
    analysis = pw.TextField(default="")
*/

type Poetry struct {
	Id         int    `json:"id"`
	Idx        int    `json:"idx"`
	Title      string `json:"title"`
	TitleTr    string `json:"title_tr"`
	Author     string `json:"author"`
	AuthorTr   string `json:"author_tr"`
	Content    string `json:"content"`
	ContentTr  string `json:"content_tr"`
	Background string `json:"background"`
	Analysis   string `json:"ananlysis"`
}

type JsonPoetry struct {
	Id         int      `json:"id"`
	Idx        int      `json:"idx"`
	Title      string   `json:"title"`
	TitleTr    string   `json:"title_tr"`
	Author     string   `json:"author"`
	AuthorTr   string   `json:"author_tr"`
	Content    []string `json:"content"`
	ContentTr  []string `json:"content_tr"`
	Background string   `json:"background"`
	Analysis   string   `json:"ananlysis"`
}

type DailyPoetry struct {
	TodayDate string     `json:"today_date"`
	Week      int        `json:"week"`
	Poetry    JsonPoetry `json:"poetry"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	orm.RegisterModel(new(Poetry))
}

func GetDailyPoetry() DailyPoetry {
	o := orm.NewOrm()
	o.Using("default")
	poetry := Poetry{Id: 1}

	err := o.Read(&poetry)
	if err != nil {
		fmt.Println(err)
	}

	jsonPoetry := JsonPoetry{
		Id:         poetry.Id,
		Idx:        poetry.Idx,
		Title:      poetry.Title,
		TitleTr:    poetry.TitleTr,
		Author:     poetry.Author,
		AuthorTr:   poetry.AuthorTr,
		Background: poetry.Background,
		Analysis:   poetry.Analysis,
	}
	json.Unmarshal([]byte(poetry.Content), &jsonPoetry.Content)
	json.Unmarshal([]byte(poetry.ContentTr), &jsonPoetry.ContentTr)

	dp := DailyPoetry{
		TodayDate: "2020-02-22", Week: 1, Poetry: jsonPoetry}

	return dp
}
