package pkg

import (
	"fmt"
	"os"
)

func Create_database() {
	if _, err := os.Stat("data.db"); os.IsNotExist(err) {
		_, err := os.Create("data.db")
		if err != nil {
			fmt.Println("Create data.db file failed:", err)
		}
		fmt.Println("Create data.db file succeed!")
	} else {
		fmt.Println("data.db file already exists!")
	}
}
