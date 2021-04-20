package controller

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func GetVisits(c *gin.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// pong, err := client.Ping(ctx).Result()
	// fmt.Println(pong, err)
	val, getErr := client.Get(ctx, "visits").Result()
	if getErr == redis.Nil {
		if err := client.Set(ctx, "visits", 0, 0).Err(); err != nil {
			c.JSON(http.StatusInternalServerError, errors.New("error when setting keys"))
			return
		}
	}
	if getErr != nil {
		c.JSON(http.StatusInternalServerError, errors.New("error when getting val"))
		return
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	valInt = valInt + 1
	if err := client.Set(ctx, "visits", valInt, 0).Err(); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, val)

}
