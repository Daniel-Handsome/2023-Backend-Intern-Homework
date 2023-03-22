package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Daniel-Handsome/2023-Backend-intern-Homework/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(cfg Config) *gorm.DB {
	connector := &pqsqlConn{cfg: cfg}
	return connector.Open()
}

type pqsqlConn struct {
	cfg Config
}

func (p *pqsqlConn) Open() *gorm.DB {
	dsn := p.cfg.BuildConnString()
	return p.Connect(dsn)
}

func (p *pqsqlConn) Connect(dsn string) *gorm.DB {
	db, err := sql.Open(string(pq), dsn)
	if err != nil {
		panic(err)
	}

	p.cfg.Conn.apply(db)

	ok := make(chan bool)
	go func() {
		var name string
		for {
			query := "SELECT current_database()"
			err := db.QueryRow(query).Scan(name)
			if err != nil && name == "" {
				continue
			}

			ok <- true
			return
		}
	}()

	select {
	case <-time.After(10 * time.Second):
		panic("db timeout for 10 seconds")
	case <-ok:
		orm, err := gorm.Open(postgres.New(postgres.Config{
			Conn: db,
		}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
		if err != nil {
			panic(err)
		}
		return orm
	}
}

type Config struct {
	DBName   string
	Host     string
	Port     int
	User     string
	Password string
	SSL      SSL
	Conn     Conn
}

type SSL struct {
	Mode string
}

func (s *SSL) toOptions() []option {
	return []option{
		{"sslmode", s.Mode},
	}
}

type Conn struct {
	Max int
}

func (c *Conn) apply(db *sql.DB) {
	db.SetMaxOpenConns(c.Max)
}

type option struct {
	key   string
	value string
}

func (o *option) toString() string {
	return fmt.Sprintf("%s=%s", o.key, o.value)
}

func (c *Config) BuildConnString() string {
	options := c.toOptions()
	return c.build(options)
}

func (c *Config) toOptions() (options []option) {
	if c.SSL.Mode != "" {
		options = append(options, c.SSL.toOptions()...)
	} else {
		options = append(options, option{
			key: "sslmode", value: "disable",
		})
	}

	return append(options, c.defautOptions()...)
}

func (c *Config) defautOptions() (options []option) {
	return []option{
		{"user", c.User},
		{"password", c.Password},
		{"host", c.Host},
		{"dbname", c.DBName},
	}
}

func (c *Config) build(options []option) string {
	params := []string{}
	for _, option := range options {
		params = append(params, option.toString())
	}

	return strings.Join(params, " ")
}

func DefaultCfg(dbnName string) Config {
	return Config{
		DBName: dbnName,
		Host:   utils.GetConfigToString("Host"),
		Port:   int(utils.GetConfigToInt("Port")),
		SSL: SSL{
			Mode: "",
		},
		Conn: Conn{
			Max: 100,
		},
	}
}
