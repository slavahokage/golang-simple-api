package model


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type articleModel struct {
	db *gorm.DB
	a article
}

type article struct {
	gorm.Model
	Headline string
	Text string
}

func ConnectionToDB() *gorm.DB{
	db, err := gorm.Open("mysql", "root:@/articleDB?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func CreateNewArticle(text string) {
	ConnectionToDB().Create(&article{Text: text})
}

func GetAllArticles() *[]article  {
	return &[]article{}
}

func GetArticleModel() *articleModel {
	return &articleModel{a: article{}}
}

func (articleModel) GetArticle() *article {
	return &article{}
}

func (a article) GetArticleText() string {
	return a.Text
}

func (a article) GetArticleHead() string {
	return a.Headline
}

