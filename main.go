package splitwise_backend

import (
	"fmt"

	"github.com/supabase-community/supabase-go"
)

func main() {

	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, count, err := client.From("todos").Select("*", "exact", false).Execute()
	if err == nil {
		fmt.Println(string(data), err, count)
	}
}
