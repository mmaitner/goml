package goml

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
)

type Tags struct {
	Href string
}

type Element struct {
	Body []string
	Id   string
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func feildName(a interface{}, n int) string {
	val := reflect.Indirect(reflect.ValueOf(a))
	return val.Type().Field(n).Name
}
func feildByName(v Tags, field string) interface{} {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field).Interface()
	return f
}

func genArgString(args Tags, id string) string {
	id = strings.Split(id, "_")[1]
	length := reflect.TypeOf(args).NumField()
	dec := ""
	for i := 0; i < length; i++ {
		name := feildName(args, i)
		value := feildByName(args, name)
		str := fmt.Sprintf("%v", value)
		if str != "" {

			dec += id + "." + name + "= '" + str + "';\n"
		}
	}
	return dec
}
