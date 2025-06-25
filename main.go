package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func UnixToTime(timeStamp int) string {
	t := time.Unix(int64(timeStamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func TextLn(str1 string, str2 string) string {
	fmt.Println(str1, str2)

	// 我想要返回 str1 + str2的值要怎么写？
	// 你可以直接返回 str1 + str2
	return str1 + " ---- " + str2
}

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
		"TextLn":     TextLn,
	})
	fmt.Println("Hello World!")
	r.LoadHTMLGlob("templates/**/*.html")

	r.Static("/tmp", "./templates/static")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "There is test page!")
	})
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s", name)
	})
	r.POST("/post", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Post老子干什么？",
		})
	})
	r.PUT("/edit", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Put老子干什么？",
		})
	})
	r.DELETE("/delete", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delete老子干什么？",
		})
	})
	type User struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}

	// r.GET("/user", func(c *gin.Context) {
	// 	user := User{
	// 		Name:  "小明",
	// 		Age:   18,
	// 		Email: "xiaoming@163.com",
	// 	}
	// 	c.JSON(200, user)
	// })

	r.GET("/tom", func(c *gin.Context) {
		tom := User{
			Name:  "Tom",
			Age:   18,
			Email: "tom@163.com",
		}
		c.JSONP(200, tom)
	})

	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"name":  "小明",
			"age":   18,
			"email": "xiaoming@163.com",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "新闻",
		})
	})

	r.GET("/long", func(c *gin.Context) {

		jack := User{
			Name:  "jack",
			Age:   18,
			Email: "jack@163.com",
		}

		c.HTML(http.StatusOK, "long.html", gin.H{
			"user":  jack,
			"title": "新闻",
		})
	})

	r.GET("/admin", func(c *gin.Context) {
		admin := User{
			Name:  "Admin",
			Age:   18,
			Email: "admin@163.com",
		}
		tom := User{
			Name:  "Tom",
			Age:   18,
			Email: "tom@163.com",
		}
		jack := User{
			Name:  "jack",
			Age:   18,
			Email: "jack@163.com",
		}
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "Admin",
			"admin": admin,
			"Tags":  []string{"tag1", "tag2", "tag3"},
			"Items": []interface{}{
				// &User{
				// 	Name:  "Tom",
				// 	Age:   18,
				// 	Email: "tom@163.com",
				// },
				// &User{
				// 	Name:  "jack",
				// 	Age:   18,
				// 	Email: "jack@163.com",
				// },
				admin,
				tom,
				jack,
			},
		})
	})

	r.GET("/user", func(c *gin.Context) {

		elon := User{
			Name:  "Elon",
			Age:   18,
			Email: "Elon@163.com",
		}
		c.HTML(http.StatusOK, "user/index.html", gin.H{
			"title": "User",
			"user":  elon,
			"date":  1744707155,
		})
	})
	r.Run(":8080")
}
