package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
	"word-list/config"
	"word-list/dao"
	"word-list/dto"
)

func InitialAPI() {
	log.Info("---------- initial api start ----------")
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Use(HandlerInterceptorAdapter())
	router.Use(gin.LoggerWithFormatter(Logger))
	router.Use(gin.Recovery())
	router.LoadHTMLGlob("templates/*")
	moduleGroup := router.Group("/word-list") //Module
	{
		moduleGroup.GET("/", WordGroup)
		moduleGroup.GET("/:id/:page/:pageSize", WordList)
	}
	_ = router.Run(config.Conf.Http.Listen)
}

func HandlerInterceptorAdapter() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		// Set example variable
		context.Set("example", "12345")
		// before request
		context.Next()
		// after request
		latency := time.Since(t)
		log.Print(latency)
		// access the status we are sending
		status := context.Writer.Status()
		log.Println(status)
	}
}

func Logger(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func WordGroup(context *gin.Context) {
	var wordGroupArray []dto.WordGroup
	wordGroupArray = append(wordGroupArray, dto.WordGroup{Id: int64(1), Name: "The Longman 2000"})
	wordGroupArray = append(wordGroupArray, dto.WordGroup{Id: int64(2), Name: "The Oxford 3000"})
	context.HTML(http.StatusOK, "index.html", gin.H{
		"CTX":  config.Conf.Environment.CTX,
		"data": wordGroupArray,
	})
}

func WordList(context *gin.Context) {
	idQuery := context.Param("id")
	pageQuery := context.Param("page")
	pageSizeQuery := context.Param("pageSize")
	id, _ := strconv.ParseInt(idQuery, 10, 64)
	page, _ := strconv.ParseInt(pageQuery, 10, 64)
	pageSize, _ := strconv.ParseInt(pageSizeQuery, 10, 64)
	if page < 0 {
		page = 0
	}
	wordArray, _ := dao.FindWordListPageByParams(id, page, pageSize)
	context.HTML(http.StatusOK, "detail.html", gin.H{
		"CTX":       config.Conf.Environment.CTX,
		"data":      wordArray,
		"id":        idQuery,
		"page_size": pageSizeQuery,
		"previous":  page - 1,
		"next":      page + 1,
	})
}
