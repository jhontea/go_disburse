# How to use it
1. Run migration
go too root project
`cd migration`
`go run migration.go`

2. Run command
go to root project
`go run *.go` + command

# Command
CommandDisburse = command to send disburse
```
go run *.go disburse
```

CommandDisburseStatus = command to get and update disburse status
```
go run *.go disburse-status {transaction_id}
```

CommandTimeExecution = command to check time execution get and update disburse status
```
go run *.go time {transaction_id} {loop}
```

# Time Execution Result
Flow 
1. Hit endpoint get status from url https://nextar.flip.id/disburse/{transaction_id}
2. Update to database

```
2020/05/09 13:52:18 Total execution 1 data time: 272.676937ms
2020/05/09 13:52:18 Total execution 2 data time: 279.903582ms
2020/05/09 13:52:18 Total execution 3 data time: 287.465483ms
2020/05/09 13:52:18 Total execution 4 data time: 296.55344ms
2020/05/09 13:52:18 Total execution 5 data time: 303.070336ms
2020/05/09 13:52:18 Total execution 6 data time: 312.435731ms
2020/05/09 13:52:18 Total execution 7 data time: 320.736942ms
2020/05/09 13:52:18 Total execution 8 data time: 328.142377ms
2020/05/09 13:52:18 Total execution 9 data time: 334.802556ms
2020/05/09 13:52:18 Total execution 10 data time: 342.754148ms
```