package main

import (
	"log"
	"unsafe"
	"zcache"
)

type User struct {
	name string
	age int
}

func (u *User) Len() int {
	return len(u.name) + int(unsafe.Sizeof(u.age))
}

func main() {
	z := zcache.NewGroup("scores",2<<10, zcache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("数据库读取"+key)
			return []byte("val"),nil
		}))

	z.Get("xxx")
	z.Get("xxx")
	z.Get("xxx")
	z.Get("yyy")

}
