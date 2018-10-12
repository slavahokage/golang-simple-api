package controllers

import (
	"github.com/gin-gonic/gin"
	"model"
	"strconv"
)

type ArticleController struct {
	router *gin.Engine
}

func СreateArticleController(router *gin.Engine) ArticleController {
	a := model.GetArticleModel()
	db := model.ConnectionToDB()
	db.AutoMigrate(a.GetArticle())
	ac := ArticleController{router: router}
	return ac
}

func (a *ArticleController) СreateArticle() {
	a.router.POST("/create", func(c *gin.Context) {
		a := c.Query("article")
		h := c.Query("headline")
		db := model.ConnectionToDB()
		ar := model.GetArticleModel().GetArticle()
		ar.Text = a
		ar.Headline = h
		db.Create(ar)
		c.JSON(200, gin.H{
			"article": a,
			"headline": h,
		})
	})

}

func (a *ArticleController) FindArticleById() {
	a.router.GET("/find", func(c *gin.Context) {
		id := c.Query("id")
		db := model.ConnectionToDB()
		article := model.GetArticleModel().GetArticle();
		i, err := strconv.Atoi(id)

		if err != nil {
			panic("something go wrong")
		}

		db.First(article, i)

		c.JSON(200, gin.H{
			"headline": article.Headline,
			"article": article.Text,

		})
	})
}

func (a *ArticleController) FindAllArticles() {
	a.router.GET("/findAll", func(c *gin.Context) {
		db := model.ConnectionToDB()
		all := model.GetAllArticles()
		db.Find(&all)
		c.JSON(200, gin.H{
			"article": all,
		})
	})
}

func (a *ArticleController) DeleteArticleById() {
	a.router.DELETE("/delete", func(c *gin.Context) {
		id := c.Query("id")
		db := model.ConnectionToDB()
		a := model.GetArticleModel().GetArticle()
		db.Where("id = ?", id).Delete(a)
		c.JSON(200, gin.H{"id #" + id: "deleted"})
	})
}

func (a *ArticleController) Run() {
	a.router.Run()
}