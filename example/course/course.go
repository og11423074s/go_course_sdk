package main

import (
	"errors"
	"fmt"
	courseSdk "github.com/og11423074s/go_course_sdk/course"
	"os"
)

func main() {
	courseTrans := courseSdk.NewHttpClient("http://localhost:8082/", "")

	course, err := courseTrans.Get("ce856191-ab87-4f88-a688-65a9d7a152ba")

	if err != nil {
		if errors.As(err, &courseSdk.ErrNotFound{}) {
			fmt.Println("Not found", err.Error())
			os.Exit(1)
		}

		fmt.Println("Internal server error", err.Error())
		os.Exit(1)
	}

	fmt.Println("Course", course)
	os.Exit(0)

}
