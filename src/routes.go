package src

import (
	"io/ioutil"
	"os"
	"goql/src/utils"
	"github.com/gin-gonic/gin"
)

type Routes struct{}

func (c Routes) StartGin() {
	// Release Mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	}

	r := gin.Default()
	
	r.GET("/graphql", utils.TimeoutNew(gin.WrapH(GraphqlSetting())))
	r.POST("/graphql", utils.TimeoutNew(gin.WrapH(GraphqlSetting())))

	// Server Start
	port := os.Getenv("PORT")
	r.Run("0.0.0.0:" + port)
}
