package main

import (
	"JH_2024_MJJ/internal/middleware"
	database "JH_2024_MJJ/internal/pkg/databse"
	"JH_2024_MJJ/internal/router"
	"JH_2024_MJJ/internal/service"
	"JH_2024_MJJ/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	utils.InitLogger()
	db := database.Init()
	service.ServiceInit(db)
	r := gin.Default()
	r.NoMethod(middleware.HandleNotFond)
	r.NoRoute(middleware.HandleNotFond)
	router.Init(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
