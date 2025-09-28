```markdown
# Task Tracker CLI (Go)


Simple CLI task tracker using Go standard library. Data saved to `tasks.json`.


## Build
```
go build -o task-cli
```


## Usage
```
./task-cli add "Buy groceries"
./task-cli list
./task-cli mark-in-progress 1
./task-cli mark-done 1
./task-cli update 1 "New desc"
./task-cli delete 1
```


```