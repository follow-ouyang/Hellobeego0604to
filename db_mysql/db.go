package db_mysql

import (
	"HelloBeego0604to/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_"go-sql-driver/mysql"
)

var Db *sql.DB
//连接数据库
func Connect()  {
	config := beego.AppConfig

	////获取配置选项
	dbDriver := config.String("db_driver")
	dbUser := config.String("db_user")
	dbpassword := config.String("db_password")
	dbIP := config.String("db_ip")
	dbName := config.String("db_name")
	//fmt.Println(dbDriver,dbUser,dbpassword)

	connUrl := dbUser+":"+dbpassword+"@tcp("+dbIP+")/"+dbName+"?charset=utf8"
	db,err := sql.Open(dbDriver,connUrl)
	if err != nil {//err不为nil，表示连接数据库时出现错误，程序在此中断就行，不用再执行
		panic("数据库连接失败，请重试")
	}
	Db = db
	fmt.Println(db)

}


//将用户信息保存到数据库中去的函数

func AddUser(u models.User) (int64,error) {
	//1、将密码进行hash计算，得到密码hash值
	md5Hash := md5.New()
	md5Hash.Write([]byte(u.Password))
	passwordBytes := md5Hash.Sum(nil)
	u.Password = hex.EncodeToString(passwordBytes)

	result,err := Db.Exec("insert into text(name,birthday,address,password)" +
		"values(?,?,?,?)",u.Name,u.Birthday,u.Address,u.Password)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("遇到错误了")
		return -1,err
	}
	row,err := result.RowsAffected()
	if err != nil {
		fmt.Println("遇到错误啦")
		return -1,err
	}
	fmt.Println(row)
	return row, nil
}

func QueryUser() ([]models.User,error) {
	rows,err := Db.Query("select *from text")
	if err != nil {
		return nil,err
	}
	users := make([]models.User,0)
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Name)
		if err != nil {
			return nil,err
		}
		users = append(users,user)
	}
	return users,err
}