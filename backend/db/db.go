package db

import (
	"backend/internal/ent"
	"backend/internal/ent/user"
	"context"
	"log"
)

type User struct {
	Id       uint64 `json:"id"`
	Account  string `json:"account"`
	Desc     string `json:"desc"`
	Location string `json:"location"`
	Tags     string `json:"tags"`
	Password string `json:"password,omitempty"`
}
type Blog struct {
	Id      uint64   `json:"id"`
	Title   string   `json:"title"`
	UserId  uint64   `json:"author_id"`
	Tags    []string `json:"tags"`
	Address string   `json:"address"`
	Ctime   int64    `json:"ctime"`
	View    uint64   `json:"view"`
	Desc    string   `json:"desc,omitempty"`
}
type Database struct {
	db  *ent.Client
	log *log.Logger
}

func NewDatabase(db *ent.Client, logger *log.Logger) *Database {
	return &Database{
		db:  db,
		log: logger,
	}
}
func (d *Database) UserBasicInfo(ctx context.Context, account string) *User {
	var u User
	uinfo, err := d.db.User.Query().Select(user.FieldAccount, user.FieldDesc, user.FieldLocation).
		Where(user.Account(account)).Only(ctx)
	if err != nil {
		d.log.Println("find user error " + err.Error())
		return nil
	} else if uinfo == nil {
		return nil
	}
	u.Desc = uinfo.Desc
	u.Account = uinfo.Account
	u.Location = uinfo.Location
	u.Id = uinfo.ID
	return &u
}

// func (d *Database) BlogList(ctx context.Context, author_id uint64, title string) (bloglist []Blog) {
// 	query := d.db.Blog.Query().Where(blog.UserID(author_id))
// 	if len(title) > 0 {
// 		query.Where(blog.TitleContainsFold(title))
// 	}
// 	b, err := query.All(ctx)
// 	if err != nil {
// 		d.log.Println("get blog list error " + err.Error())
// 		return
// 	}
// 	utils.SliceConvert(&bloglist, b, func(dst *Blog, src *ent.Blog) {
// 		dst.Address = src.Address
// 		dst.Ctime = src.Ctime
// 		dst.Id = src.ID
// 		dst.UserId = src.UserID
// 		dst.Title = src.Title
// 		dst.Tags = src.Tags
// 	})
// 	return
// }

// func (d *Database) BlogListWithTag(ctx context.Context, author_id uint64, tagId uint64) (bloglist []Blog) {
// 	query := d.db.Blog.Query().Where(blog.UserID(author_id))
// 	if tagId > 0 {
// 		query.Modify(func(s *sql.Selector) {
// 			tagtable := sql.Table(tags.Table)
// 			s.LeftJoin(tagtable).On(tags.FieldArticleID, blog.FieldID).Where(sql.EQ(tags.FieldID, tagId))
// 		})
// 	}
// 	b, err := query.All(ctx)
// 	if err != nil {
// 		d.log.Println("find blog with tag failed")
// 		return
// 	}
// 	utils.SliceConvert(&bloglist, b, func(dst *Blog, src *ent.Blog) {
// 		dst.Address = src.Address
// 		dst.Ctime = src.Ctime
// 		dst.Id = src.ID
// 		dst.UserId = src.UserID
// 		dst.Title = src.Title
// 		dst.Tags = src.Tags
// 	})
// 	return
// }
