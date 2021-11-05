package service

import (
	"github.com/beego/beego/v2/client/orm"
	."metal_ty/models"
)

type ArticleService interface {
	Save(article *Article) (int64,error)
}

type articleService struct {

}

func NewService() ArticleService {
	return &articleService{}
}

func (as *articleService) Save(article *Article)(int64, error){
	o := orm.NewOrm()
	return o.Insert(article)
}

