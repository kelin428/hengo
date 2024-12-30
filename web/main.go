package main

import (
	"hengo"
)

type Name struct {
	Id string `json:"id"`
}

func main() {
	hen := hengo.New()
	hen.Get("/hello", func(c *hengo.Context) {
		var name Name
		c.BindJSON(&name)
		//c.JsonResponse(hengo.StatusOK, "Hello World", hengo.H{
		//	"name":    "Hen",
		//	"version": "0.1",
		//})
		c.JsonErrorResponse(hengo.StatusBadRequest, "Hello World")
	})

	hen.Run(":8080")
}
