package main

import (
	"log"
	"math"
	"net/http"

	"github.com/Jarover/reef/models"
	"github.com/Jarover/reef/readconfig"
	"github.com/gin-gonic/gin"
)

//homePage - отдает страницу
func homePage(c *gin.Context) {
	var obj = models.Wlevel{}
	err := models.GetDB().Last(&obj, "point_id = ?", 26).Error
	if err != nil {
		log.Println(err)

	}

	olevel := obj.Offset - obj.Level

	deep := int(math.Round(float64(olevel)/10) * 10)
	znak := "+"
	if olevel < 0 {
		znak = ""
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Форт Риф",
		"level": olevel,
		"id":    obj.ID,
		"deep":  deep,
		"znak":  znak,
	})

}

// infoPage - интформационная страница
func infoPage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"version": readconfig.Version.VersionStr(),
		"data":    readconfig.Version.BuildTime,
	})
}

// levels - интформационная страница
func levels(c *gin.Context) {

	var levels = []models.Wlevel{}
	models.GetDB().Where("point_id = ?", 26).Find(&levels)
	type row struct {
		utime string
		level int
	}
	var out []row

	for _, v := range levels {
		var n row
		n.level = v.Level
		n.utime = v.Datetime
		out = append(out, n)
	}

	c.JSON(http.StatusOK, gin.H{
		"out":    out,
		"levels": levels,
	})
}
