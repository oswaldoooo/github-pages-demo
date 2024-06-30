package route

import (
	"backend/db"
	"backend/internal/ent"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type BlogV1 struct {
	db    *db.Database
	cache *db.Cache
	cb    *db.CacheDatabase
}

func NewBlogV1(d *ent.Client, cache *redis.Client, errorlog *log.Logger) *BlogV1 {
	return &BlogV1{
		db:    db.NewDatabase(d, errorlog),
		cache: db.NewCache(cache, errorlog),
		cb:    db.NewCacheDatabase(d, cache, errorlog),
	}
}
func (b *BlogV1) RegisterRouter(r gin.IRoutes) {
	r.GET("/blog/list", b.BlogList)
	r.POST("/blog/list", b.BlogList)
	r.GET("/blog/rank", b.BlogRank)
	r.GET("/blog/tag/list", b.BlogTagList)
	r.GET("/blog/view", b.BlogView)
	r.GET("/user/:account", b.UserInfo)
}
func (b *BlogV1) UserInfo(ctx *gin.Context) {
	defer ctx.Next()
	account := ctx.Param("account")
	uinfo := b.db.UserBasicInfo(ctx, account)
	if uinfo == nil {
		ctx.JSON(http.StatusOK, BaseResponse{Status: NotFound})
		return
	}
	ctx.JSON(http.StatusOK, BaseResponse{Status: Ok, Data: *uinfo})
}
func (b *BlogV1) BlogList(ctx *gin.Context) {
	var req struct {
		Tag      string `json:"tag"`
		Page     uint64 `json:"page"`
		PageSize uint64 `json:"page_size"`
		Title    string `json:"title"`
		AuthorId uint64 `json:"author_id"`
	}
	defer ctx.Next()
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, BaseResponse{
			Status: ParamInvalid,
		})
		return
	}
	if len(req.Tag) > 0 {

		return
	}
	var data []db.Blog
	if len(req.Tag) > 0 {
		data = b.cb.BlogListWithTag(ctx, req.AuthorId, req.Tag)
	} else {
		data = b.cb.BlogList(ctx, req.AuthorId, req.Title)
	}

	ctx.JSON(http.StatusOK, BaseResponse{
		Status: Ok,
		Data:   data,
	})
}
func (b *BlogV1) BlogRank(ctx *gin.Context) {
	pagesize, _ := strconv.ParseUint(ctx.Query("page_size"), 10, 64)
	page, _ := strconv.ParseUint(ctx.Query("page"), 10, 64)
	rinfo := b.cb.GetRank(ctx, page, pagesize)
	ctx.JSON(http.StatusOK, BaseResponse{Status: Ok, Data: rinfo})
	ctx.Next()
}
func (b *BlogV1) BlogTagList(ctx *gin.Context) {
	authorId, _ := strconv.ParseUint(ctx.Query("author_id"), 10, 64)
	taglist := b.cache.TagList(ctx, authorId)
	ctx.JSON(http.StatusOK, BaseResponse{Status: Ok, Data: taglist})
	ctx.Next()
}

func (b *BlogV1) BlogView(ctx *gin.Context) {
	authorId, _ := strconv.ParseUint(ctx.Query("author_id"), 10, 64)
	ctx.JSON(http.StatusOK, BaseResponse{Status: Ok, Data: b.cache.BlogView(ctx, authorId)})
	ctx.Next()
}
