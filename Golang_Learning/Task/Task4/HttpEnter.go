package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var g_secret = []byte("KpcL4pmBaYkxx")

func startHttp() {
	router := gin.Default()
	router.Use(authenticate())
	router.POST("/register", register)
	router.POST("/login", login)
	router.POST("/getUserName", getUserName)
	router.POST("/createPost", createPost)
	router.POST("/queryPost", queryPost)
	router.POST("/queryAllPost", queryAllPost)
	router.POST("/updatePost", updatePost)
	router.POST("/deletePost", deletePost)
	router.POST("/createComment", createComment)
	router.POST("/queryAllComment", queryAllComment)

	router.Run()
}

// 注册
func register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取用户信息失败"})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
		return
	}
	user.Password = string(hashedPassword)

	if err := insertRecordUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

// 登录
func login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storedUser, err := queryRecordUserByName(user.UserName)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码不正确"})
		return
	}

	// 生成 JWT
	expirationTime := time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        storedUser.ID,
		"user_name": storedUser.UserName,
		"exp":       expirationTime,
	})

	tokenString, err := token.SignedString(g_secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成Token失败"})
		return
	}

	// 返回token给客户端
	c.JSON(http.StatusOK, gin.H{
		"token":      tokenString,
		"expires_at": expirationTime,
		"user": gin.H{
			"id":        storedUser.ID,
			"user_name": storedUser.UserName,
		},
	})
}

// 获取用户名
func getUserName(c *gin.Context) {
	user := c.MustGet("currentUser").(User)

	c.JSON(http.StatusOK, gin.H{
		"user_name": user.UserName,
	})
}

// 获取token信息
func getTokenInfo(c *gin.Context) (claims jwt.MapClaims, err error) {
	// 从Header获取token
	tokenString := c.GetHeader("token")
	if tokenString == "" {
		err = errors.New("缺少认证token")
		return
	}

	// 解析和验证token
	token, jwtErr := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return g_secret, nil
	})
	if jwtErr != nil || !token.Valid {
		err = errors.New("无效的token")
		return
	}

	return
}

// 鉴权中间件
func authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := getTokenInfo(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		user, err := queryRecordUserByID(uint(claims["id"].(float64)))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户不存在"})
			c.Abort()
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

// 创建文章
func createPost(c *gin.Context) {
	user := c.MustGet("currentUser").(User)
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取文章信息失败"})
		return
	}

	value, exists := jsonData["title"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少标题"})
		return
	}

	title, ok := value.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "标题类型错误"})
		return
	}

	value, exists = jsonData["content"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少内容"})
		return
	}

	content, ok := value.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "内容类型错误"})
		return
	}

	post := Post{Title: title,
		Content: content,
		UserID:  user.ID}
	if err := insertRecordPost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "创建成功"})
}

// 查询文章信息
func queryPost(c *gin.Context) {
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取参数失败"})
		return
	}

	value, exists := jsonData["id"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少文章ID"})
		return
	}

	id, ok := value.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章ID类型错误"})
		return
	}

	post, err := queryRecordPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      post.ID,
		"title":   post.Title,
		"content": post.Content,
	})
}

// 查询用户的所有文章
func queryAllPost(c *gin.Context) {
	user := c.MustGet("currentUser").(User)

	posts, err := queryRecordPostByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
		return
	}

	results := []map[string]interface{}{}
	for _, value := range posts {
		results = append(results, map[string]interface{}{
			"id":      value.ID,
			"title":   value.Title,
			"content": value.Content,
		})
	}

	c.JSON(http.StatusOK, results)
}

// 查询文章信息
func updatePost(c *gin.Context) {
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取参数失败"})
		return
	}

	value, exists := jsonData["id"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少文章ID"})
		return
	}

	id, ok := value.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章ID类型错误"})
		return
	}

	updates := make(map[string]interface{})
	value, exists = jsonData["title"]
	if exists {
		text, ok := value.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "文章标题类型错误"})
			return
		}
		updates["title"] = text
	}

	value, exists = jsonData["content"]
	if exists {
		text, ok := value.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "文章内容类型错误"})
			return
		}
		updates["content"] = text
	}

	if len(updates) == 0 {
		c.JSON(http.StatusCreated, gin.H{"message": "更新成功"})
		return
	}

	err := updateRecordPost(uint(id), updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "更新成功"})
}

// 删除文章信息
func deletePost(c *gin.Context) {
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取参数失败"})
		return
	}

	value, exists := jsonData["id"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少文章ID"})
		return
	}

	id, ok := value.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章ID类型错误"})
		return
	}

	err := deleteRecordPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "删除成功"})
}

// 创建评论
func createComment(c *gin.Context) {
	user := c.MustGet("currentUser").(User)
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取文章信息失败"})
		return
	}

	value, exists := jsonData["post_id"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少文章ID"})
		return
	}

	postId, ok := value.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章ID类型错误"})
		return
	}

	value, exists = jsonData["content"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少内容"})
		return
	}

	content, ok := value.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "内容类型错误"})
		return
	}

	comment := Comment{PostID: uint(postId),
		Content: content,
		UserID:  user.ID}
	if err := insertRecordComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "创建成功"})
}

// 查询文章的所有评论
func queryAllComment(c *gin.Context) {
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "获取文章信息失败"})
		return
	}

	value, exists := jsonData["post_id"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少文章ID"})
		return
	}

	postId, ok := value.(float64)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文章ID类型错误"})
		return
	}

	comments, err := queryRecordCommentByPostID(uint(postId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
		return
	}

	results := []map[string]interface{}{}
	for _, value := range comments {
		results = append(results, map[string]interface{}{
			"id":      value.ID,
			"content": value.Content,
		})
	}

	c.JSON(http.StatusOK, results)
}
