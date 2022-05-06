package helper

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	gMySQL "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DBHelper interface {
	Open() *sql.DB
	Close() error
	Begin() (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
}

type dbHelper struct {
	db *sql.DB
}

func NewDBHelper(host string, port int, username, password, database string, isMigrate bool, argsModel ...interface{}) DBHelper {
	if isMigrate {
		err := initSQLWithGorm(host, port, username, password, database, argsModel...)
		if err != nil {
			panic(err.Error())
		}
	}

	db, err := initSQL(host, port, username, password, database)
	if err != nil {
		zap.S().Errorf("Can't connect to db with err: %v", err)
		panic(err.Error())
	}
	return &dbHelper{
		db: db,
	}
}

func (h *dbHelper) Open() *sql.DB {
	return h.db
}

func (h *dbHelper) Close() error {
	return h.db.Close()
}

func (h *dbHelper) Begin() (*sql.Tx, error) {
	return h.db.Begin()
}

func (h *dbHelper) Commit(tx *sql.Tx) error {
	return tx.Commit()
}
func (h *dbHelper) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func initSQL(host string, port int, username, password, database string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, host, port, database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func ReplaceTableName(tableName string) string {
	fmt.Println(tableName)
	return strings.ToUpper(tableName)
}

func initSQLWithGorm(host string, port int, username, password, database string, argsModel ...interface{}) error {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, host, port, database)
	gdb, err := gorm.Open(gMySQL.Open(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	gdb.AutoMigrate(argsModel...)
	return nil
}
