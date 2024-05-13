package sqLite

import (
	"atp/payment/pkg/adapter/model"
	"context"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	ormlog "gorm.io/gorm/logger"
)

type sqliteDB struct {
	db *gorm.DB
}

type DatabaseI interface {
	Db(ctx context.Context) interface{}
	WithTransaction(ctx context.Context, fn func(ctxWithTx context.Context, dbt *gorm.DB) error) error
}

func NewConnection(dsn string) (DatabaseI, error) {
	var err error

	newLogger := ormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		ormlog.Config{
			SlowThreshold:             5 * time.Second, // Slow SQL threshold
			Colorful:                  true,            // Disable color
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			LogLevel:                  ormlog.Error,    // Log level
		},
	)

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return sqliteDB{db: db}, err
	}

	db.AutoMigrate(&model.Transaction{})

	log.Printf("successfully connected to database: %v", dsn)

	return &sqliteDB{db: db}, nil
}

func (p sqliteDB) Db(ctx context.Context) interface{} {
	tx := ctx.Value("txContext")
	if tx == nil {
		return p.db
	}
	return tx.(*gorm.DB)
}

func (d sqliteDB) WithTransaction(ctx context.Context, fn func(ctxWithTx context.Context, dbt *gorm.DB) error) error {
	tx := d.db.Begin()
	ctxWithTx := context.WithValue(ctx, tx, "txContext")
	errFn := fn(ctxWithTx, tx)
	if errFn != nil {
		errRlbck := tx.Rollback().Error
		if errRlbck != nil {
			log.Printf("<WithTransaction> failed on rollback:%s", errRlbck.Error())
		}
		return errFn
	}

	errCmmt := tx.Commit().Error
	if errCmmt != nil {
		log.Printf("<WithTransaction> failed on commit:%s", errCmmt.Error())
	}
	return errFn
}
