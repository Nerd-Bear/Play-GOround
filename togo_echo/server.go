package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"net/http"
)

var ToGoList []ToGo

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}
type ToGo struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Coordinate Coordinate `json:"coordinate"`
}

func AddToGo(c echo.Context) error {
	toGo := new(ToGo)
	if err := c.Bind(&toGo); err != nil {
		return err
	}

	toGo.Id = uuid.New().String()
	ToGoList = append(ToGoList, *toGo)
	return c.NoContent(http.StatusCreated)
}

func GetToGO(c echo.Context) error {
	if len(ToGoList) == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, ToGoList)
}

func FindIndexById(t []ToGo, id string) int{
	for i, toGo := range t {
		if toGo.Id == id {
			return i
		}
	}
	return -1
}

func DeleteToGo(c echo.Context) error {
	id := c.Param("id")
	deleteIndex := FindIndexById(ToGoList, id)
	if deleteIndex == -1 {
		return c.NoContent(http.StatusNotFound)
	}
	ToGoList = append(ToGoList[:deleteIndex], ToGoList[deleteIndex+1:]...)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	e.GET("/", GetToGO)
	e.POST("/", AddToGo)
	e.DELETE("/:id", DeleteToGo)


	e.Logger.Fatal(e.Start(":1234"))
}
