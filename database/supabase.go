package database

import (
	"fmt"
	"splitwise-backend/constants"

	"github.com/supabase-community/supabase-go"
)

func GetAllData(tableName string) {
	client, err := supabase.NewClient(constants.API_URL, constants.API_KEY, nil)
	if err != nil {
		fmt.Println("Error initializing client:", err)
		return
	}

	data, count, err := client.From(tableName).Select("*", "exact", false).Execute()
	if err != nil {
		fmt.Println("Error querying data:", err)
		return
	}

	fmt.Println(string(data), count)
}
