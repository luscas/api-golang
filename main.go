package main

import (
	"github.com/hoisie/web"
	"github.com/luscas/api-golang/controllers"
)

func FindUser(id string) string {
	return "User ID: " + id
}

// func Radio() string {
// 	response, erro := http.Get("https://painel.audiovox.pw/api-json/Njc2OCsw")
// 	if erro != nil {
// 		panic(erro)
// 	}

// 	responseJson, erro := ioutil.ReadAll(response.Body)

// 	if erro != nil {
// 		panic(erro)
// 	}

// 	return string(responseJson)
// }

func main() {
	web.Get("/", controllers.HelloWorld)
	web.Get("/radio", Radio)
	web.Get("/users/(.*)", FindUser)
	web.Run("0.0.0.0:9999")
}
