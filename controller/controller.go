package controller

import (
	"github.com/spf13/viper"
	"hot/router"
	"hot/utils"
	"net/http"
)

func Router()  {
	router.GetHandleFunc("/hello", helloGet)
	router.PostHandleFunc("/hello", helloPost)
	router.PostHandleFunc("/login", login)
	err := http.ListenAndServe(":"+viper.GetString("port"), router.DefaultRouter)
	if err != nil {
		utils.Error.Panic(err)
	}
}
