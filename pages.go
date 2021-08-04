package main

import (
	"net/http"

	"github.com/Jarover/reef/readconfig"
	"github.com/gin-gonic/gin"
)

//homePage - отдает страницу
func homePage(c *gin.Context) {

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Reef",
	})

}

// infoPage - интформационная страница
func infoPage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"version": readconfig.Version.VersionStr(),
		"data":    readconfig.Version.BuildTime,
	})
}
