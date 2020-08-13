package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

var Client *elastic.Client

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func Init() (err error) {
	Client, err = elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
		return err
	}
	fmt.Println("connect to es success", Client)
	return
}
func Run() {
	p1 := &Person{Name: "luolita", Age: 22, Married: false}
	Index(p1)
}

func Index(p1 *Person) {
	//put1, err := client.Index().Index("student").Type("go").BodyJson(p1).Do(context.Background())
	put1, err := Client.Index().Index("student").Type("go").BodyJson(p1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed student %s to index %s,type %s\n", put1.Id, put1.Index, put1.Type)
}
