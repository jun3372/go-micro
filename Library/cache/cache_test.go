package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	fmt.Println(myCache.Put("xxx", 1, 10*time.Second))
	fmt.Println(GetInt(myCache.Get("xxx")))
	fmt.Println(myCache.IsExist("xxx"))
	// fmt.Println(myCache.Delete("astaxie"))
}
