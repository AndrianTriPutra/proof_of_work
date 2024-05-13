package transaction_test

import (
	"atp/payment/pkg/adapter/sqLite"
	"atp/payment/pkg/repository/transaction"
	"atp/payment/pkg/utils/echos/util"
	"context"
	"errors"
	"log"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func Test_FindALL(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	log.Printf("basepath:%s", basepath)
	base := basepath[0:strings.Index(basepath, "pkg")]
	path := base + "database/block.db"
	log.Printf("path:%s", path)

	db, err := sqLite.NewConnection(path)
	if err != nil {
		log.Fatalf("FAILED connect to database:" + err.Error())
	}

	ctx := context.Background()
	repo := transaction.NewRepository(db)
	data, err := repo.GetALLTrans(ctx)
	if errors.Is(err, util.ErrorNotFound) {
		log.Fatalf("error-1:%s", err.Error())
	} else if err != nil {
		log.Fatalf("error-2:%s", err.Error())
	}

	for _, value := range data {
		log.Printf("id:%v", value.ID)
	}
}
