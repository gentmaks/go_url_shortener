package main 

import "fmt"
import "github.com/gin-gonic/gin"
import "github.com/gentmaks/go_url_shortener/handler"
import "github.com/gentmaks/go_url_shortener/store"

func main() {

  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "URL shortener...",
    })
  })
  
  r.POST("/create-short-url", func(c *gin.Context) {
    handler.CreateShortUrl(c)
  })
  
  r.GET("/:shortUrl", func(c *gin.Context) {
    handler.HandleShortUrlRedirect(c)
  })

  store.InitializeStore()

  err := r.Run(":9808")
  if err != nil {
    panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
  }
}
