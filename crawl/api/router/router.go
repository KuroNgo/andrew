package router

import (
	"crawl/initialization"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func SetUp(env *initialization.Database, timeout time.Duration, db *mongo.Database, client *mongo.Client, gin *gin.Engine, cacheTTL time.Duration) {

	// Đếm các route
	routeCount := countRoutes(gin)
	fmt.Printf("The number of API endpoints: %d\n", routeCount)
}

func countRoutes(r *gin.Engine) int {
	count := 0
	routes := r.Routes()
	for range routes {
		count++
	}
	return count
}
