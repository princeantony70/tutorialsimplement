package main

   import (
       _ "github.com/go-sql-driver/mysql"
       "database/sql"
       "fmt"
   )

   func main() {
       db, err := sql.Open("mysql", "prince:prince@/astaxie")
       checkErr(err)

       // insert
       stmt, err := db.Prepare("INSERT userinfo SET username=antony,departname=it,created=2012-12-09")
       checkErr(err)

       res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
       checkErr(err)

       id, err := res.LastInsertId()
       checkErr(err)

       fmt.Println(id)
       // update
       stmt, err = db.Prepare("update userinfo set username=eric where uid=1")
       checkErr(err)

       res, err = stmt.Exec("astaxieupdate", id)
       checkErr(err)

       affect, err := res.RowsAffected()
       checkErr(err)

       fmt.Println(affect)

       // query
       rows, err := db.Query("SELECT * FROM userinfo")
       checkErr(err)

       for rows.Next() {
           var uid int
           var username string
           var department string
           var created string
           err = rows.Scan(&uid, &username, &department, &created)
           checkErr(err)
           fmt.Println(uid)
           fmt.Println(username)
           fmt.Println(department)
           fmt.Println(created)
       }

       // delete
       stmt, err = db.Prepare("delete from userinfo where uid=1")
       checkErr(err)

       res, err = stmt.Exec(id)
       checkErr(err)

       affect, err = res.RowsAffected()
       checkErr(err)

       fmt.Println(affect)

       db.Close()

   }

   func checkErr(err error) {
       if err != nil {
           panic(err)
       }
   }
