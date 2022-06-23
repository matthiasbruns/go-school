## Build a small program that

- Counts the rows of a table
- Adds a new entry while using transactions
- Counts the rows again

## Table Creation Script

```sql
create table gophers
(
    id   int auto_increment
        primary key,
    name text          not null,
    age  int default 0 not null
);

```

## Run MySQL

`docker run --detach --name=goschool-mysql -p 52000:3306 --platform linux/x86_64  --env="MYSQL_ROOT_PASSWORD=mypassword" mysql`

## Connect to localhost MySQL

```go
db, err := sql.Open("mysql", "root:mypassword@tcp(localhost:52000)/goschool")
```