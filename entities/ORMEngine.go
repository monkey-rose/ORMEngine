package entities

import (
    "database/sql"
    "strings"
    "reflect"
    "fmt"
    "time"
    "strconv"
    _ "github.com/go-sql-driver/mysql"
)

type ORMExec interface {
	Insert(o interface{})
	Find(o interface{})
}


type ORMEngine struct {
	ORMExec
	db *sql.DB
}

func NewORMEngine() *ORMEngine {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
    engine := ORMEngine{}
    engine.db = db
    return &engine
}

func (engine *ORMEngine) Insert(o interface{}) error {
    s := reflect.ValueOf(o).Elem()
    tableName := strings.ToLower(s.Type().Name())
    typeOfT := s.Type()
    col := make([]string, 0)
    value := make([]string, 0)
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
		if f.Type().String() == "int" {
			continue
		}
		col = append(col, typeOfT.Field(i).Name)
		if f.Type().String() == "*time.Time" {
				value = append(value, fmt.Sprintf("%v", f.Interface().(*time.Time).Format("2006-01-02 15:04:05")))
			} else {
				value = append(value, fmt.Sprintf("%v", f.Interface()))
			}
		
    }
	stmt, err := engine.db.Prepare("INSERT " + tableName + "(" + strings.Join(col, ",") + ")" + "values (?,?,?)")
	checkErr(err)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(value[0], value[1], value[2])
	checkErr(err)
	if err != nil {
		return err
	}
	return nil
}

func (engine *ORMEngine) Find(o interface{}) error {
	rows, err := engine.db.Query("SELECT * FROM userinfo")
	checkErr(err)
	if err != nil {
		return err
	}
	columns, _ := rows.Columns()
    count := len(columns)
    values := make([]interface{}, count)
    valuePtrs := make([]interface{}, count)

    s := reflect.ValueOf(o).Elem()
	typeUser := reflect.TypeOf(s.Interface()).Elem()

    for rows.Next() {
        for i, _ := range columns {
            valuePtrs[i] = &values[i]
        }

        rows.Scan(valuePtrs...)
        item := reflect.New(typeUser).Elem()
        for i, col := range columns {
            val := values[i]
			if col == "createat" {
				b, _ := val.([]byte)
                v := string(b)
				tm2, err := time.Parse("2006-01-02 15:04:05", v)
				checkErr(err)
				if err != nil {
					return err
				}
				item.Field(i).Set(reflect.ValueOf(&tm2))
			} else if col == "uid" {
				b, _ := val.([]byte)
                v := string(b)
				id , err:= strconv.Atoi(v)
				if err != nil {
					return err
				}
				id_64 := int64(id)
				item.Field(i).SetInt(id_64)
			} else {
				item.Field(i).Set(reflect.ValueOf(val).Convert(item.Field(i).Type()))
			}
		}
		s.Set(reflect.Append(s, item))
    }
	return nil
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
