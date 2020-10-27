package transport

import (
	"todo/handler"
	"todo/storage"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//InitRoutes create noteroutes and userroutes
func InitRoutes(engine *gin.Engine, db *gorm.DB) {
	initNoteRoutes(engine, db)
	initUserRoutes(engine, db)
}

func initUserRoutes(engine *gin.Engine, db *gorm.DB) {
	groupRouter := engine.Group("/v1/users")
	groupRouter.POST("/signup", func(c *gin.Context) {
		userDB := &storage.UserDB{
			DB: db,
		}
		result, err := handler.UserSignUp(c, userDB)
		resultsHandler(c, err, result)
	})
	groupRouter.POST("/auth", func(c *gin.Context) {
		userDB := &storage.UserDB{
			DB: db,
		}
		result, err := handler.UserLogin(c, userDB)
		resultsHandler(c, err, result)
	})
}

func initNoteRoutes(engine *gin.Engine, db *gorm.DB) {
	groupRouter := engine.Group("/v1/notes")
	groupRouter.GET("/:id", func(c *gin.Context) {
		noteDB := &storage.NoteDB{
			DB: db,
		}
		result, err := handler.NoteList(c, noteDB)
		resultsHandler(c, err, result)
	})
	groupRouter.POST("", func(c *gin.Context) {
		noteDB := &storage.NoteDB{
			DB: db,
		}
		result, err := handler.NoteCreate(c, noteDB)
		resultsHandler(c, err, result)
	})
	groupRouter.PUT("/:id", func(c *gin.Context) {
		noteDB := &storage.NoteDB{
			DB: db,
		}
		err := handler.NoteUpdate(c, noteDB)
		resultsHandler(c, err, nil)
	})
	groupRouter.DELETE("/:id", func(c *gin.Context) {
		noteDB := &storage.NoteDB{
			DB: db,
		}
		err := handler.NoteDelete(c, noteDB)
		resultsHandler(c, err, nil)
	})

}

func resultsHandler(c *gin.Context, err error, result interface{}) {
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, result)
}
