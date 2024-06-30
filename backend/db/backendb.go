package db

import (
	"backend/internal/ent"
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type BackendDb struct {
	db    *ent.Client
	cache *redis.Client
	log   *log.Logger
}

func NewBackendDb(d *ent.Client, c *redis.Client, logger *log.Logger) *BackendDb {
	return &BackendDb{
		db:    d,
		cache: c,
		log:   logger,
	}
}
func (b *BackendDb) AddBlog(ctx context.Context, title string, tags []string, address string, authorId uint64) (err error) {
	var db *ent.Tx
	db, err = b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			err = db.Commit()
		}
	}()
	var binfo *ent.Blog
	binfo, err = db.Blog.Create().SetCtime(time.Now().Unix()).SetAddress(address).SetTitle(title).SetUserID(authorId).Save(ctx)
	if err != nil {
		return
	}
	if len(tags) > 0 {
		var taglist []*ent.TagsCreate = make([]*ent.TagsCreate, len(tags))
		for i, tname := range tags {
			taglist[i] = db.Tags.Create().SetArticleID(binfo.ID).SetUserID(authorId).SetName(tname)
		}
		err = db.Tags.CreateBulk(taglist...).Exec(ctx)
	}
	return
}
