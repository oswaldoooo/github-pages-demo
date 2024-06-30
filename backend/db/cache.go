package db

import (
	"backend/internal/utils"
	"backend/types"
	"context"
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

type RankInfo types.KV[BlogBasic, uint64]

const (
	article_range_set = "article_range_set"
)

type Cache struct {
	db  *redis.Client
	log *log.Logger
}

func NewCache(db *redis.Client, logger *log.Logger) *Cache {
	return &Cache{
		db:  db,
		log: logger,
	}
}
func (c *Cache) TagList(ctx context.Context, author_id uint64) (result []TagInfo) {
	r, err := c.db.ZRevRangeWithScores(tagCountSets+"_"+strconv.FormatUint(author_id, 10), 0, -1).Result()
	if err != nil {
		c.log.Println("get tag count failed " + err.Error())
		return
	}
	utils.SliceConvert(&result, r, func(dst *TagInfo, src redis.Z) {
		dst.Key = src.Member.(string)
		dst.Value = uint64(src.Score)
	})
	return
}

func (c *Cache) BlogView(ctx context.Context, articleId uint64) (views uint64) {
	aid := strconv.FormatUint(articleId, 10)
	res, err := c.db.ZScore(article_range_set, aid).Result()
	if err != nil {
		c.log.Println("find article view error " + err.Error())
		return
	}
	err = c.db.ZIncrBy(article_range_set, 1, aid).Err()
	if err != nil {
		c.log.Println("add view article " + aid + " error " + err.Error())
	}
	views = uint64(res)
	return
}
