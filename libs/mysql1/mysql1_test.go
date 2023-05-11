package mysql1

// # DSN

// > [https://github.com/go-sql-driver/mysql]()

// ```go
// DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Asia%%2FShanghai", uid, pwd, host, port, dbName)
// ```

// 1. If you do not want to preselect a database, leave dbname empty:

// ```go
// uid.pwd@tcp()/?
// ```

// 2. charset

// 3. parseTime，default false, 返回结果是[]byte,还是time.Time

// 4. loc, 时区，默认UTC0 , loc=Asia%%2FShanghai

// > datetime类型，存储为无意义字面值[]byte，loc影响time.Time转换字面值的方式

// insert的数据为： t.Loc.String() t.UTC.String()

// select得到的Time为： time.Parse(字面值), 或者 time.ParseLocation(字面值)
