package sqLite_test

import (
	"atp/payment/pkg/adapter/sqLite"
	"log"
	"testing"
)

func Test_Connect(t *testing.T) {
	_, err := sqLite.NewConnection("block.db")
	if err != nil {
		log.Fatalf("<Test_Connect> failed connect to database:" + err.Error())
	}
}
