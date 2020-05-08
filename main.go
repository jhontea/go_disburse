package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/jhontea/go_disburse/handlers"
	"github.com/jhontea/go_disburse/services"
)

// services
var disburseService services.DisburseServiceInterface

var (
	// CommandDisburse :nodoc:
	CommandDisburse = "disburse"

	// CommandDisburseStatus :nodoc:
	CommandDisburseStatus = "disburse-status"
)

func main() {
	initServices()

	flag.Parse()
	if len(flag.Args()) > 0 {
		processCommand(flag.Args())
	} else {
		fmt.Println("Please input command")
	}
}

func initServices() {
	disburseService = services.NewDisburseService()
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
			fmt.Println("processing get disburse status")
			disburseeHandler.GetDisburseStatus(transactionID)
		}
	default:
		{
			fmt.Println("command not found")
		}
	}
}
