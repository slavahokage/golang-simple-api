package main

import (
	"controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	ac := controllers.СreateArticleController(router)
	initializeRouting(&ac)
}

func initializeRouting(a *controllers.ArticleController) {
	a.СreateArticle()
	a.DeleteArticleById()
	a.FindAllArticles()
	a.FindArticleById()
	a.Run()
}

