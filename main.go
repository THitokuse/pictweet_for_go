package main
import (
  "github.com/gin-gonic/gin"
  "net/http"
)
func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Hello World")
  })
  //listen and server on 0.0.0.0:8080
  r.Run()
}
