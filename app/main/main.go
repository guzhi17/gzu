package main

import (
	"log"
	"net/url"
	"strconv"
)

//////////////////////////////////////////////////////////////////////////////////////////////

type RedisConfig struct {
	Host string
	Password string
	Db int
}

func (r RedisConfig)IsValid() bool {
	return len(r.Host)>0
}

//url: "redis://root:bat2019!#$@118.31.126.111:27017/?authSource=admin"
func (r *RedisConfig) UnmarshalCml(text string) error  {
	u, err := url.Parse(text)
	if err != nil {
		return err
	}
	//u.Scheme == "redis"
	r.Host = u.Host
	if pwd, ok := u.User.Password(); ok{
		r.Password = pwd
	}
	qs := u.Query()

	dbs:=qs["db"]
	if len(dbs)>0{
		db, err := strconv.Atoi(dbs[0])
		if err != nil{
			return err
		}
		r.Db = db
	}

	return nil
}
type User struct {
	Id    int    `cml:"i"`
	Name  string `cml:"n,value=Tom,usage=the name of yours"`
	Email string `cml:"e,value=xx@qq.com,usage=the email of yours"`
	Redis *RedisConfig `cml:"r,usage=the redis"`
}

//go run main.go -t=-n -n "Jack" -e "129@qq.com" -r "10.0.1.99:987"
func main()  {
	log.SetFlags(11)

	x, err := url.Parse("file://name:passowrd@D:/hello/file.yml?a=b") //"file://name:passowrd@D:/hello/file.yml?a=b"
	log.Print(x, err)

	log.Print(x.Scheme)
	log.Print(x.Host)
	log.Print(x.RawPath)
	log.Print(x.Path)
	log.Print(x.RequestURI())
	log.Print(x.EscapedPath())
	log.Print(x.Host + x.Path)


	//user := User{
	//	Id:    1,
	//	//Name: &n,
	//	//Name:  "John Doe",
	//	Email: "john@example",
	//
	//	Redis: &RedisConfig{
	//		Db: 2,
	//	},
	//}
	//
	//err = flag.ParseFlags(&user, flag.ContinueOnError)
	//if user.Redis != nil{
	//	log.Println(*user.Redis)
	//}else{
	//	log.Println(user.Redis)
	//}
	//log.Println(user, err)
}

//////////////////////////////////////////////////////////////////////////////////////////////