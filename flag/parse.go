package flag

import (
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	tagName = "cml"
	valueName = "value" //default value
	usageName = "usage"
)
func (cml *FlagSet)ParseFlags(v interface{}) (err error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	rv = reflect.Indirect(rv)
	rt := reflect.TypeOf(v).Elem()
	// Iterate over all available fields and read the tag value
	//cml := NewFlagSet(rt.Name(), errorHandling)
	//
	//CommandLine.String()
	var cmlvs = map[string]interface{}{}
	var cmlptrs = map[string]reflect.Value{}
	for i := 0; i < rt.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := rt.Field(i)
		name, ms := parseTag(field)
		if name == "-"{
			continue
		}
		value := rv.Field(i)
		if value.Kind() == reflect.Ptr && value.IsNil(){
			if !value.CanSet(){
				continue
			}
			//value.Set(reflect.New(value.Type().Elem()))
			value = reflect.New(value.Type().Elem())
			cmlptrs[field.Name] = value
		}
		value = reflect.Indirect(value)

		switch value.Kind() {
		case reflect.String:
			v := cml.String(name, ms[valueName], ms[usageName])
			cmlvs[field.Name] = v
		case reflect.Int64, reflect.Int, reflect.Int32:
			v := cml.Int64(name, toInt(ms[valueName]), ms[usageName])
			cmlvs[field.Name] = v
		case reflect.Bool:
			v := cml.Bool(name, toBool(ms[valueName]), ms[usageName])
			cmlvs[field.Name] = v
		case reflect.Float64, reflect.Float32:
			v := cml.Float64(name, toFloat64(ms[valueName]), ms[usageName])
			cmlvs[field.Name] = v
		case reflect.Struct, reflect.Slice:
			//has interface
			valraw := reflect.Indirect(value).Addr()

			if valraw.Type().NumMethod() > 0 && valraw.CanInterface() {
				if _, ok := valraw.Interface().(CmlUnmarshaler); ok {

					v := cml.String(name, ms[valueName], ms[usageName])
					cmlvs[field.Name] = v
				}
			}
		}
	}

	if len(cmlvs) < 1{
		return nil
	}

	err = cml.Parse(os.Args[1:])
	if err != nil{
		if !strings.HasPrefix(err.Error(), "flag provided but not defined"){
			return err
		}
	}

	var value reflect.Value
	for ki, vi := range cmlvs{
		vr := rv.FieldByName(ki)

		var newaddr = false
		if vr.Kind() == reflect.Ptr && vr.IsNil(){
			newaddr = true
			value = cmlptrs[ki]
		}else{
			value = vr
		}

		value = reflect.Indirect(value)
		//log.Println(value.Kind())
		switch value.Kind() {
		case reflect.String:
			v := *(vi.(*string))
			if len(v)>0{
				value.SetString(v)
			}
		case reflect.Int64, reflect.Int, reflect.Int32:
			value.SetInt(*(vi.(*int64)))
		case reflect.Bool:
			value.SetBool(*(vi.(*bool)))
		case reflect.Float64, reflect.Float32:
			value.SetFloat(*(vi.(*float64)))
		case reflect.Struct, reflect.Slice:
			var valraw  = value.Addr()
			if valraw.Type().NumMethod() > 0 && valraw.CanInterface() {
				v := *(vi.(*string))
				if len(v) > 0{
					if u, ok := valraw.Interface().(CmlUnmarshaler); ok {
						err = u.UnmarshalCml(v)
						if err != nil{
							return err
						}
					}
					if newaddr{
						vr.Set(cmlptrs[ki])
					}
				}
			}
		}
	}
	return nil
}

//todo add support for struct, mushalltext
func ParseFlags(v interface{}, errorHandlings... ErrorHandling) (err error) {
	var errorHandling = ContinueOnError
	if len(errorHandlings) > 0{
		errorHandling = errorHandlings[0]
	}
	rt := reflect.TypeOf(v).Elem()
	cml := NewFlagSet(rt.Name(), errorHandling)
	return cml.ParseFlags(v)
}

type CmlUnmarshaler interface {
	UnmarshalCml(text string) error
}

func toInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil{
		return 0
	}
	return i
}

func toFloat64(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil{
		return 0
	}
	return i
}

func toBool(s string)bool  {
	if s == "true" || s == "TRUE" || s == "1"{
		return true
	}
	return false
}

func parseTag(field reflect.StructField) (n string, m map[string]string) {
	s := field.Tag.Get(tagName)
	ts := strings.Split(s, ",")
	n = ts[0]
	m = map[string]string{}
	for i:=1; i< len(ts); i++{
		kv := strings.SplitN(ts[i], "=", 2)
		if len(kv) < 2{
			m[kv[0]] = ""
		}else{
			m[kv[0]] = kv[1]
		}
	}
	return
}

type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "json: Unmarshal(nil " + e.Type.String() + ")"
}



//////////////////////////////////////////////////////////////////////////////////////////////
type RedisConfig struct {
	Host string
	Password string
	Port int
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
	Redis RedisConfig `cml:"r,usage=the redis"`
}

//go run main.go -t=-n -n "Jack" -e "129@qq.com" -r "10.0.1.99:987"
func main()  {
	user := User{
		Id:    1,
		//Name: &n,
		//Name:  "John Doe",
		Email: "john@example",
	}
	ParseFlags(&user, ContinueOnError)
}

//////////////////////////////////////////////////////////////////////////////////////////////