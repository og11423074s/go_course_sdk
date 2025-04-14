package main

import (
	"errors"
	"fmt"
	userSdk "github.com/og11423074s/go_course_sdk/user"
	"os"
)

func main() {
	userTrans := userSdk.NeHttpClient("http://localhost:8081/", "")

	user, err := userTrans.Get("13ae2f6b-4050-46f1-94d4-afb13492f96c")

	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found", err.Error())
			os.Exit(1)
		}

		fmt.Println("Internal server error", err.Error())
		os.Exit(1)
	}

	fmt.Println("User", user)
	os.Exit(0)

}
