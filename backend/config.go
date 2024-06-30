package main

import (
	"backend/internal/ent"
	"io"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Database   Database `yaml:"database"`
	Cache      Cache    `yaml:"cache"`
	ListenPort int      `yaml:"listen_port"`
	LogOut     Stream   `yaml:"log_out"`
	BackEndOut Stream   `yaml:"backend_out"`
}
type Database struct {
	Address  string `yaml:"address"`
	User     string `yaml:"user"`
	Password int    `yaml:"password"`
	Db       string `yaml:"db"`
}
type Cache struct {
	Address  string `yaml:"address"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Db       string `yaml:"db"`
}

// stream use like: protocol://address||token
type Stream struct {
	Protocol string
	Address  string
	Token    string
}

type Configer interface {
	ReadConfig() (cnf Config, err error)
}

type File string

func (f File) ReadConfig() (cnf Config, err error) {
	var fi *os.File
	fi, err = os.OpenFile(string(f), os.O_RDONLY, 0644)
	if err != nil {
		return
	}
	defer fi.Close()
	err = yaml.NewDecoder(fi).Decode(&cnf)
	return
}

type RuntimeInfo struct {
	db               *ent.Client
	cache            *redis.Client
	loggerout        io.WriteCloser
	backendLoggerOut io.WriteCloser
}

func BuildRuntimeInfo(c Configer) (rif RuntimeInfo, err error) {
	var cnf Config
	cnf, err = c.ReadConfig()
	if err != nil {
		return
	}
	rif.db, err = buildDb(&cnf.Database)
	if err != nil {
		return
	}
	rif.cache, err = buildCache(&cnf.Cache)
	if err != nil {
		return
	}
	rif.loggerout, err = buildLogger(&cnf.LogOut)
	if err != nil {
		return
	}
	rif.backendLoggerOut, err = buildLogger(&cnf.BackEndOut)
	return
}

func buildDb(c *Database) (d *ent.Client, err error) {
	d, err = ent.Open("sqlite3", c.Address)
	return
}
func buildCache(c *Cache) (d *redis.Client, err error) {
	var db int
	if len(c.Db) > 0 {
		db, err = strconv.Atoi(c.Db)
		if err != nil {
			return
		}
	}
	d = redis.NewClient(&redis.Options{
		Addr:     c.Address,
		DB:       db,
		Password: c.Password,
	})
	err = d.Ping().Err()
	return
}
