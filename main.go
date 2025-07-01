package main
import(
	"cart/routes"
	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()
	routes.UserRoutes(router)
	router.Run("localhost:9090")
	
}