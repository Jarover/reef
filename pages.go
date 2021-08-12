package main

import (
	"log"
	"math"
	"net/http"
	"time"

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
	now := time.Now()
	preday := now.AddDate(0, 0, -1)
	var levels = []models.Wlevel{}
	var queue []int64
	//preday = preday.Add(-3 * time.Hour)

	models.GetDB().Order("datetime desc").Where("point_id = ? and channel_id = ? and datetime > ?", 26, 1, preday.Format("2006-01-02 15:04:05")).Find(&levels)
	type row struct {
		Utime string
		Level int64
		Unix  int64
	}
	var out []row
	var n row
	ii := 0
	for _, v := range levels {

		queue = append(queue, v.Offset-v.Level)
		if ii > 1 {
			queue = queue[1:]
		}
		n.Level = Sum(queue) / int64(len(queue))
		n.Utime = v.Datetime
		t, _ := time.Parse(time.RFC3339, v.Datetime)
		n.Unix = (t.Unix() + 3*3600) * 1000
		out = append(out, n)
		ii = ii + 1
	}

	c.JSON(http.StatusOK, gin.H{
		"out": out,
	})
}

// chanels - информация
func chanels(c *gin.Context) {
	now := time.Now()
	preday := now.AddDate(0, 0, -1)

	ch := c.Param("ch")
	var datas = []models.Wlevel{}
	models.GetDB().Order("datetime").Where("point_id = ? and channel_id = ? and datetime > ?", 28, ch, preday.Format("2006-01-02 15:04:05")).Find(&datas)
	type row struct {
		Utime string
		Level int64
		Unix  int64
	}
	var out []row
	var n row

	for _, v := range datas {

		n.Level = v.Level
		n.Utime = v.Datetime
		t, _ := time.Parse(time.RFC3339, v.Datetime)
		n.Unix = (t.Unix() + 3*3600) * 1000
		out = append(out, n)
	}

	c.JSON(http.StatusOK, gin.H{
		"out": out,
	})
}
