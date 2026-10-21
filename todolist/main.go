package main

import (
	"go-learning/todolist/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	const cachePath = "data.json"
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    "Hello, Go!",
		})
	})

	r.GET("/list", func(c *gin.Context) {
		todolist, err := utils.InitTodoList(cachePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"msg":     "服务器错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    todolist,
		})
	})

	r.POST("/todo", func(c *gin.Context) {
		var req struct {
			Content string `json:"content" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		todo, err2 := utils.AddTodo(cachePath, req.Content)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"msg":     err2.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "添加成功",
			"data":    todo,
		})
	})

	r.PATCH("/todo/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"msg":     "id格式错误",
			})
			return
		}
		todo, err2 := utils.ToggleTodo(cachePath, id)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"msg":     err2.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "任务状态修改成功",
			"data":    todo,
		})
	})

	r.DELETE("/todo/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"msg":     "id格式错误",
			})
			return
		}
		todo, err2 := utils.RemoveTodo(cachePath, id)
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"msg":     err2.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "任务状态修改成功",
			"data":    todo,
		})
	})

	r.DELETE("/completed", func(c *gin.Context) {
		todolist, err := utils.ClearCompletedTodos(cachePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"msg":     err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "已成功清除所有已完成任务",
			"data":    todolist,
		})
	})

	r.Run(":3000")
}
