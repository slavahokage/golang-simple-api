package main

import (
	"controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	ac := controllers.СreateArticleController(router)
	initializeRouting(&ac)
	/*db := model.ConnectionToDB()
	a := model.GetArticleModel()
	db.AutoMigrate(a.GetArticle())
	model.CreateNewArticle("test")
	//ourAr := model.GetArticleById(*db,1)
	//fmt.Println(ourAr)
	test2 := a.GetArticle()
	db.First(test2,1)
	fmt.Println(test2.Text)
	defer db.Close()*/
}

func initializeRouting(a *controllers.ArticleController) {
	a.СreateArticle()
	a.DeleteArticleById()
	a.FindAllArticles()
	a.FindArticleById()
	a.Run()
}

