package db

import (
	"backend/internal/ent"
	"backend/internal/ent/blog"
	"backend/internal/ent/tags"
	"backend/internal/ent/user"
	"backend/internal/utils"
	"backend/types"
	"context"
	"log"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"github.com/go-redis/redis"
)

const (
	tagCountSets = "tag_count_sets"
)

type BlogBasic struct {
	Title   string `json:"title"`
	Author  string `json:"author"`
	Address string `json:"address"`
}
type TagInfo types.KV[string, uint64]
type CacheDatabase struct {
	db    *ent.Client
	cache *redis.Client
	log   *log.Logger
}
type pointerStruct struct {
	title   *string
	address *string
	account *string
}

func NewCacheDatabase(d *ent.Client, c *redis.Client, logger *log.Logger) *CacheDatabase {
	return &CacheDatabase{
		db:    d,
		cache: c,
		log:   logger,
	}
}
func (c *CacheDatabase) GetRank(ctx context.Context, start, limit uint64) (results []RankInfo) {
	val, err := c.cache.ZRevRangeWithScores(article_range_set, int64(start), int64(limit)).Result()
	if err != nil {
		c.log.Println("find rank failed")
		return
	}
	var (
		keys  = make([]uint64, 0, len(val))
		links = make(map[uint64]pointerStruct)
	)
	utils.SliceConvert(&results, val, func(dst *RankInfo, src redis.Z) {
		dst.Key.Title = src.Member.(string)
		uval, err := strconv.ParseUint(dst.Key.Title, 10, 64)
		if err != nil {
			dst.Key.Title = "not found blog"
			c.log.Println("not found blog", dst.Key, "key invalid", err)
			return
		}
		keys = append(keys, uval)
		links[uval] = pointerStruct{
			title:   &dst.Key.Title,
			address: &dst.Key.Address,
			account: &dst.Key.Author,
		}
		dst.Value = uint64(src.Score)
	})
	if len(val) > 0 {
		var blogs []struct {
			ID      uint64 `json:"id"`
			Address string `json:"address"`
			Title   string `json:"title"`
			Account string `json:"account"`
		}
		err := c.db.Blog.Query().Select(blog.FieldTitle, blog.FieldID,
			blog.FieldAddress, user.FieldAccount).Where(blog.IDIn(keys...)).Modify(func(s *sql.Selector) {
			s.LeftJoin(sql.Table(user.Table)).On(user.FieldID, blog.FieldUserID)
		}).Scan(ctx, &blogs)
		if err != nil {
			c.log.Println("found blogs error " + err.Error())
			return
		}
		for _, v := range blogs {
			ps := links[v.ID]
			*ps.address = v.Address
			*ps.title = v.Title
			*ps.account = v.Account
		}
	}
	return
}

func (cb *CacheDatabase) BlogList(ctx context.Context, author_id uint64, title string) (bloglist []Blog) {
	query := cb.db.Blog.Query().Where(blog.UserID(author_id))
	if len(title) > 0 {
		query.Where(blog.TitleContainsFold(title))
	}
	b, err := query.All(ctx)
	if err != nil {
		cb.log.Println("get blog list error " + err.Error())
		return
	}
	var (
		idlist []uint64 = make([]uint64, 0, len(b))
		links           = make(map[uint64]*[]string)
	)
	utils.SliceConvert(&bloglist, b, func(dst *Blog, src *ent.Blog) {
		dst.Address = src.Address
		dst.Ctime = src.Ctime
		dst.Id = src.ID
		dst.UserId = src.UserID
		dst.Title = src.Title
		vi, err := cb.cache.ZScore(article_range_set, strconv.FormatUint(dst.Id, 10)).Result()
		if err == nil {
			dst.View = uint64(vi)
		}
		idlist = append(idlist, dst.Id)
		links[dst.Id] = &dst.Tags
	})
	//get tag name
	if len(idlist) > 0 {
		taginfo, err := cb.db.Tags.Query().Where(tags.ArticleIDIn(idlist...)).All(ctx)
		if err != nil {
			cb.log.Println("get tags info error " + err.Error())
			return
		}
		for _, tv := range taginfo {
			l := links[tv.ArticleID]
			(*l) = append((*l), tv.Name)
		}
	}

	return
}
func (cb *CacheDatabase) BlogListWithTag(ctx context.Context, author_id uint64, tagName string) (bloglist []Blog) {
	query := cb.db.Blog.Query().Where(blog.UserID(author_id))
	if len(tagName) > 0 {
		query.Modify(func(s *sql.Selector) {
			tagtable := sql.Table(tags.Table)
			s.LeftJoin(tagtable).On(tags.FieldArticleID, blog.FieldID).Where(sql.EQ(tags.FieldName, tagName))
		})
	}
	b, err := query.All(ctx)
	if err != nil {
		cb.log.Println("find blog with tag failed")
		return
	}
	var (
		idlist []uint64 = make([]uint64, 0, len(b))
		links           = make(map[uint64]*[]string)
	)
	utils.SliceConvert(&bloglist, b, func(dst *Blog, src *ent.Blog) {
		dst.Address = src.Address
		dst.Ctime = src.Ctime
		dst.Id = src.ID
		dst.UserId = src.UserID
		dst.Title = src.Title
		vi, err := cb.cache.ZScore(article_range_set, strconv.FormatUint(dst.Id, 10)).Result()
		if err == nil {
			dst.View = uint64(vi)
		}
		idlist = append(idlist, dst.Id)
		links[dst.Id] = &dst.Tags
	})
	//get tag name
	if len(idlist) > 0 {
		taginfo, err := cb.db.Tags.Query().Where(tags.ArticleIDIn(idlist...)).All(ctx)
		if err != nil {
			cb.log.Println("get tags info error " + err.Error())
			return
		}
		for _, tv := range taginfo {
			l := links[tv.ArticleID]
			(*l) = append((*l), tv.Name)
		}
	}
	return
}

// func (cb *CacheDatabase) TagByName(ctx context.Context, tagname string, author_id uint64) (id uint64) {
// 	var err error
// 	id, err = cb.cache.HGet(tagHashTable, tagname).Uint64()
// 	if err == nil {
// 		return
// 	}
// 	return
// }
// func (cb *CacheDatabase) TagById(ctx context.Context, tagid uint64) (name string) {

// 	return
// }
