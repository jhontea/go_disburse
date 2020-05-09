package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/jhontea/go_disburse/handlers"
	"github.com/jhontea/go_disburse/repositories"
	"github.com/jhontea/go_disburse/services"

	_ "github.com/go-sql-driver/mysql"
)

// database
var db *sql.DB

// repositories
var disburseRepository repositories.DisburseRepositoryInterface

// services
var disburseService services.DisburseServiceInterface

var (
	// CommandDisburse :nodoc:
	CommandDisburse = "disburse"

	// CommandDisburseStatus :nodoc:
	CommandDisburseStatus = "disburse-status"

	// CommandTimeExecution :nodoc:
	CommandTimeExecution = "time"
)

func main() {
	initMysql()
	initRepositories()
	initServices()
	// defer db.Close()

	flag.Parse()
	if len(flag.Args()) > 0 {
		processCommand(flag.Args())
	} else {
		fmt.Println("Please input command")
	}
}

func initMysql() {
	var err error

	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/disburse_db")

	if err != nil {
		log.Fatal(err)
	}
}

func initRepositories() {
	disburseRepository = repositories.NewDisburseRepository(db)
}

func initServices() {
	disburseService = services.NewDisburseService(disburseRepository)
}

func processCommand(commands []string) {
	// handler
	disburseeHandler := handlers.NewDisburseHandler(disburseService)

	switch commands[0] {
	case CommandDisburse:
		{
			fmt.Println("processing disburse")
			disburseeHandler.SendDisburse()

		}
	case CommandDisburseStatus:
		{
			if len(commands) < 2 {
				fmt.Println("Command must be: `disburse-status {transaction_id}`")
				return
			}

			transactionID, _ := strconv.Atoi(commands[1])
			disburseeHandler.GetDisburseStatus(transactionID)
		}
	case CommandTimeExecution:
		{
			start := time.Now()

			if len(commands) < 3 {
				fmt.Println("Command must be: `time {transaction_id} {n}`")
				return
			}

			transactionID, _ := strconv.Atoi(commands[1])
			n, _ := strconv.Atoi(commands[2])
			for i := 1; i <= n; i++ {
				disburseeHandler.CheckTimeExecution(transactionID)
				elapsed := time.Since(start)
				log.Printf("Total execution %d data time: %s", i, elapsed)
			}
		}
	default:
		{
			fmt.Println("command not found")
		}
	}
}
