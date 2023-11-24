package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func NewRouter() *Router {
	router := gin.Default()

	public := router.Group("/api")

	public.POST("/register", func(c *gin.Context) { authCtrl.Register(c) })
	public.POST("/login", func(c *gin.Context) { authCtrl.GetToken(c) })

	protected := router.Group("/api/admin")
	protected.Use(IsValidToken())
	protected.GET("/user", func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")
		authCtrl.GetUserInfo(c, bearToken)
	})
	protected.POST("/enable", func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")
		authCtrl.Enable(c, bearToken)
	})
	protected.POST("/disable", func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")
		authCtrl.Disable(c, bearToken)
	})

	return &Router{router: router}
}

func (r *Router) Run() {
	err := r.router.Run(":9999")
	if err != nil {
		panic("Error while running server")
	}
}

func IsValidToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearToken := c.Request.Header.Get("Authorization")

		userId, err := tkn.ExtractUserId(bearToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		foundUser, err := orm.FindUserByID(userId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if foundUser.Disabled == true {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Target user is disabled"})
			c.Abort()
			return
		}

		c.Next()
	}
}
