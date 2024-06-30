package route

import (
	"backend/db"
	"backend/internal/ent"
	"backend/internal/temptor"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type PublishReq struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags,omitempty"`
}
type Backend struct {
	t   temptor.Temptor
	log *log.Logger
	db  *db.BackendDb
}

func NewBackend(d *ent.Client, c *redis.Client, l *log.Logger) *Backend {
	return &Backend{}
}
func (b *Backend) RegisterRouter(r gin.IRoutes) {
	r.POST("/publish", b.Publish)
}

func (b *Backend) ParseToken(ctx *gin.Context) {

}

// publish blog
func (b *Backend) Publish(ctx *gin.Context) {
	defer ctx.Next()
	var req PublishReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusOK, BaseResponse{Status: ParamInvalid})
		return
	}
	uid := ctx.Value("user_id").(uint64)
	address, err := b.t.Generate(ctx, uid, req.Body)
	if err != nil {
		b.log.Println("generate html file failed " + err.Error())
		ctx.JSON(http.StatusOK, BaseResponse{Status: ServerError})
		return
	}
	err = b.db.AddBlog(ctx, req.Title, req.Tags, address, uid)
	if err != nil {
		b.log.Println("add blog to db error " + err.Error())
		ctx.JSON(http.StatusOK, BaseResponse{Status: ServerError})
		return
	}
	ctx.JSON(http.StatusOK, BaseResponse{Status: Ok})
}
