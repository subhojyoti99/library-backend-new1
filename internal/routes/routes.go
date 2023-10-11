package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"

	"library-backend-new1/internal/handlers"

	"library-backend-new1/internal/middleware"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	// Middleware to authenticate and authorize users
	r.Use(middleware.AuthMiddleware(db))

	// Library routes
	library := r.Group("/library")
	{
		library.POST("/", handlers.CreateLibrary(db))
		library.PUT("/:id", handlers.UpdateLibrary(db))
		library.GET("/:id", handlers.GetLibraryByID(db))
		library.DELETE("/:id", handlers.DeleteLibrary(db))
		library.GET("/", handlers.ListLibraries(db))
	}

	// User routes
	user := r.Group("/user")
	{
		user.POST("/", handlers.CreateUser(db))
		user.PUT("/:id", handlers.UpdateUser(db))
		user.GET("/:id", handlers.GetUserByID(db))
		user.DELETE("/:id", handlers.DeleteUser(db))
		user.GET("/", handlers.ListUsers(db))
	}

	// Book routes
	book := r.Group("/book")
	{
		book.POST("/", handlers.CreateBook(db))
		book.PUT("/:isbn", handlers.UpdateBook(db))
		book.GET("/:isbn", handlers.GetBookByISBN(db))
		book.DELETE("/:isbn", handlers.DeleteBook(db))
		book.GET("/", handlers.ListBooks(db))
	}

	// Request routes
	request := r.Group("/request")
	{
		request.POST("/", handlers.CreateRequest(db))
		request.PUT("/:id", handlers.UpdateRequest(db))
		request.GET("/:id", handlers.GetRequestByID(db))
		request.DELETE("/:id", handlers.DeleteRequest(db))
		request.GET("/", handlers.ListRequests(db))
	}

	// Issue routes
	issue := r.Group("/issue")
	{
		issue.POST("/", handlers.CreateIssue(db))
		issue.PUT("/:id", handlers.UpdateIssue(db))
		issue.GET("/:id", handlers.GetIssueByID(db))
		issue.DELETE("/:id", handlers.DeleteIssue(db))
		issue.GET("/", handlers.ListIssues(db))
	}
	return r
}
