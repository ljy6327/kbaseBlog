package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"kbase/common/datebase"
	"kbase/common/router"
	"log"
	"os"
)

func main() {
	//init database
	issucc := datebase.GetInstance().InitDataPool()
	if !issucc {
		log.Println("init database pool failure...")
		os.Exit(1)
	}

	// init router
	router.Init()

}
