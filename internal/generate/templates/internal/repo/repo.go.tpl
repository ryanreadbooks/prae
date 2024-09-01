package repo

import (
  "fmt"

  "{{ .Go.Module }}/internal/config"

  "github.com/zeromicro/go-zero/core/stores/sqlx"
)

func GetDsn(user, pass, addr, dbName string) string {
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, addr, dbName)
}

type Repo struct {
  db sqlx.SqlConn
}

func New(c *config.Config) *Repo {
  db := sqlx.NewMysql(GetDsn(
		c.DB.User,
		c.DB.Pass,
		c.DB.Addr,
		c.DB.DbName,
	))

  rdb, err := db.RawDB()
	if err != nil {
		panic(err)
	}
	if err = rdb.Ping(); err != nil {
		panic(err)
	}

  r := &Repo{
    db: db,
  }

  return r
}

func (d *Repo) DB() sqlx.SqlConn {
	return d.db
}
