package pgclient

import (
	"context"
	"database/sql"
	"sync"

	// "github.com/jackc/pgx/v4"
	// "github.com/jackc/pgx/v4/pgxpool"

	_ "github.com/mattn/go-sqlite3"
)

// ===========================================
// Trasaction Callback
// ===========================================

//TxCallback a callback to running of sql in transaction
type TxCallback func(sql.Tx) error

// ===========================================
// Connection configuration
// ===========================================

//Config configuration for database connection
type Config struct {
	Host string
	Port string
	User string
	Pswd string
	DBNm string
}

// ===========================================
// Connection
// ===========================================

//Conn abstraction of a connection
type Conn struct {
	pool *sql.DB
	ctx  context.Context
}

//Ping test connection
func (c *Conn) Ping() error {
	return c.pool.Ping()
}

func (c *Conn) StartTx() (*sql.Tx, error) {
	return c.pool.Begin()
}

func (c *Conn) StartRTx() (*sql.Tx, error) {
	return c.pool.BeginTx(context.Background(), &sql.TxOptions{
		ReadOnly:  true,
		Isolation: sql.LevelReadCommitted,
	})
}

// ===========================================
// Connection Manager
// ===========================================

// ConnManager a connection manager
type ConnManager struct {
	config *Config
	pool   *sql.DB
	ctx    context.Context
}

// Init initialize a connection pool
func (m *ConnManager) Init() error {
	db, err := sql.Open("sqlite3", "database/foo.db")
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	// store reference to session in manager
	m.pool = db
	// success
	return nil
}

// GetConn create or retrieve a conn object
func (m *ConnManager) GetConn() (*Conn, error) {
	// initialize pool if no initialized
	if m.pool == nil {
		err := m.Init()
		if err != nil {
			return nil, err
		}
	}
	// result conn instance
	return &Conn{
		ctx:  m.ctx,
		pool: m.pool,
	}, nil
}

// ===========================================
// Connection as singleton
// ===========================================

var managerSingletonInst *ConnManager = nil
var managerSingletonOnce = &sync.Once{}

// SetConfiguration set configuration to conn manager
func SetConfiguration(config Config) {
	if managerSingletonInst == nil {
		managerSingletonOnce.Do(func() {
			managerSingletonInst = &ConnManager{
				config: &config,
				ctx:    context.Background(),
			}
		})
	}
}

// GetManager get a singleton instance of conn manager
func GetManager() *ConnManager {
	return managerSingletonInst
}

func CloseAll() error {
	conn, err := managerSingletonInst.GetConn()
	if err != nil {
		return err
	}
	conn.pool.Close()
	return nil
}
