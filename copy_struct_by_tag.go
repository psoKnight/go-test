package main

import (
	"fmt"
	"github.com/ybzhanghx/copier"
	"time"
)

type structA struct {
	Items   string `mson:"Item_string"`
	UserId  string `mson:"UserId_string"`
	PubTime int64  `mson:"PubTime_int64"`
}

type structB struct {
	Item    string    `mson:"Item_string"`
	UserId  int64     `mson:"UserId_int64"`
	PubTime time.Time `mson:"PubTime_time_Time"`
}

func main() {
	obj1 := structA{Items: "233", UserId: "5433", PubTime: time.Now().Unix()}
	obj2 := &structB{}

	fmt.Println(obj2)

	err := copier.CopyByTag(obj2, &obj1, "mson")
	if err != nil {
		fmt.Println("Should not raise error")
	}

	fmt.Println(obj2)
}
