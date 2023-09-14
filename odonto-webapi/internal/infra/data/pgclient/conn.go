package pgclient

// import (
// 	"context"
// 	"fmt"
// 	"net/url"
// 	"sync"
// 	"time"

// 	"github.com/jackc/pgx/v4"
// 	"github.com/jackc/pgx/v4/pgxpool"
// 	"odonto/internal/infra/config"
// )

// // ===========================================
// // Trasaction Callback
// // ===========================================

// //TxCallback a callback to running of sql in transaction
// type TxCallback func(pgx.Tx) error

// // ===========================================
// // Connection configuration
// // ===========================================

// //Config configuration for database connection
// type Config struct {
// 	Host string
// 	Port string
// 	User string
// 	Pswd string
// 	DBNm string
// }

// // ===========================================
// // Connection
// // ===========================================

// //Conn abstraction of a connection
// type Conn struct {
// 	pool *pgxpool.Pool
// 	ctx  context.Context
// }

// //Ping test connection
// func (c *Conn) Ping() error {
// 	return c.pool.Ping(c.ctx)
// }

// func (c *Conn) StartTx() (pgx.Tx, error) {
// 	return c.pool.Begin(c.ctx)
// }

// func (c *Conn) StartRTx() (pgx.Tx, error) {
// 	return c.pool.BeginTx(context.Background(), pgx.TxOptions{
// 		IsoLevel:   pgx.ReadCommitted,
// 		AccessMode: pgx.ReadOnly,
// 	})
// }

// // ===========================================
// // Connection Manager
// // ===========================================

// // ConnManager a connection manager
// type ConnManager struct {
// 	config *Config
// 	pool   *pgxpool.Pool
// 	ctx    context.Context
// }

// // Init initialize a connection pool
// func (m *ConnManager) Init() error {
// 	// create db url
// 	dburl := url.URL{
// 		Scheme: "postgres",
// 		Host:   fmt.Sprintf("%s:%s", m.config.Host, m.config.Port),
// 		User:   url.UserPassword(m.config.User, m.config.Pswd),
// 		Path:   fmt.Sprintf("/%s", m.config.DBNm),
// 	}
// 	cfg, err := pgxpool.ParseConfig(dburl.String())
// 	if err != nil {
// 		return err
// 	}
// 	//get settings
// 	settings := config.GetSettings()
// 	// get config connection
// 	conn := settings.Connections
// 	// max open connections
// 	cfg.MaxConns = int32(conn.MaxOpenConns)
// 	// max idle connections
// 	cfg.MinConns = int32(conn.MaxIdleConns)
// 	// max idle time
// 	cfg.MaxConnIdleTime = time.Duration(conn.ConnMaxIdleTime) * time.Second
// 	// max life time
// 	cfg.MaxConnLifetime = time.Duration(conn.ConnMaxLifetime) * time.Second
// 	// create session connection
// 	pool, err := pgxpool.ConnectConfig(m.ctx, cfg)
// 	if err != nil {
// 		return err
// 	}
// 	err = pool.Ping(m.ctx)
// 	if err != nil {
// 		return err
// 	}
// 	// store reference to session in manager
// 	m.pool = pool
// 	// success
// 	return nil
// }

// // GetConn create or retrieve a conn object
// func (m *ConnManager) GetConn() (*Conn, error) {
// 	// initialize pool if no initialized
// 	if m.pool == nil {
// 		err := m.Init()
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	// result conn instance
// 	return &Conn{
// 		ctx:  m.ctx,
// 		pool: m.pool,
// 	}, nil
// }

// // ===========================================
// // Connection as singleton
// // ===========================================

// var managerSingletonInst *ConnManager = nil
// var managerSingletonOnce = &sync.Once{}

// // SetConfiguration set configuration to conn manager
// func SetConfiguration(config Config) {
// 	if managerSingletonInst == nil {
// 		managerSingletonOnce.Do(func() {
// 			managerSingletonInst = &ConnManager{
// 				config: &config,
// 				ctx:    context.Background(),
// 			}
// 		})
// 	}
// }

// // GetManager get a singleton instance of conn manager
// func GetManager() *ConnManager {
// 	return managerSingletonInst
// }

// func CloseAll() error {
// 	conn, err := managerSingletonInst.GetConn()
// 	if err != nil {
// 		return err
// 	}
// 	conn.pool.Close()
// 	return nil
// }
