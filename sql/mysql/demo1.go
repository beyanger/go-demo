
package main

import (
    "math/rand"
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


const (
    user_name   = "root"
    password    = "123"
    protocol    = "tcp"
    server      = "127.0.0.1"
    port        = 3306
    database    = "webserver"
)

type User struct {
    Id      int64 `db:"id"`
    Name    sql.NullString `db:"name"`
    Age     int `db:"age"`
}

var (
    db *sql.DB
    insertStmt  *sql.Stmt
)

func init() {
    var err error
    dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user_name, password, protocol, server, port, database)
    // 这里并不会真正的连接数据库
    db, err = sql.Open("mysql", dsn)
    checkError(err)

    insertStmt, err = db.Prepare("INSERT INTO user(name, age) VALUES(?, ?)")
    checkError(err)
}

func deinit() {
    if db != nil {
        db.Close()
    }
    if insertStmt != nil {
        insertStmt.Close()
    }
}


func checkError(err error) {
    if err != nil {
        panic(err.Error())
    }
}


func QueryOne(db *sql.DB) {
    fmt.Println("---------------------------")
    user := new(User)

    row := db.QueryRow("SELECT * FROM user WHERE id=?", 1)

    if err := row.Scan(&user.Id, &user.Name, &user.Age); err != nil {
        panic(err.Error())
    }
    fmt.Println(user)
}

func QueryMulti(db *sql.DB) {
    fmt.Println("---------------------------")
    user := User{}
    rows, err := db.Query("SELECT * FROM user WHERE id > ? LIMIT 10", 1)
    checkError(err)

    // rows 必须scan，不然会导致连接无法关闭而被占用，直到达到最大连接数，会阻塞等待连接关闭
    for rows.Next() {
        if err = rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
            panic(err.Error())
        }
        fmt.Println(user)
    }
    rows.Close()
}


func InsertData(db *sql.DB) {
    fmt.Println("---------------------------")
    r, err := db.Exec("INSERT INTO user(name, age) VALUES(?, ?)", "shabi", 12)

    checkError(err)

    lastid, err := r.LastInsertId()
    checkError(err)
    fmt.Println("last insert id:", lastid)

    rowaffected, err := r.RowsAffected()
    checkError(err)

    fmt.Println("affected: ", rowaffected)
}


func PrepareSQL(db *sql.DB) {

    fmt.Println("---------------------------")

    const (
        name = "axxxdfasdfasdfasdfbcdefghijklmnopoqrst"
    )

    p := rand.Int() % 10 + 1
    fmt.Println(p)

    r, err := insertStmt.Exec(string(name[p:p+p]), p + 10)

    checkError(err)

    lastid, err := r.LastInsertId()
    checkError(err)
    fmt.Println("last insert id:", lastid)

    rowaffected, err := r.RowsAffected()
    checkError(err)

    fmt.Println("affected: ", rowaffected)
}


func UpdateData(db *sql.DB) {
    // 不返回 LastInsertId
}

func DeleteData(db *sql.DB) {
    // 不返回 LastInsertId
}


func main() {

    if err := db.Ping(); err != nil {
        panic(err.Error())
    }
    QueryOne(db)
    QueryMulti(db)
    InsertData(db)
    for i := 0; i < 1000; i++ {
        PrepareSQL(db)
    }
    defer deinit()
}


