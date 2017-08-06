package main

import (
	"fmt"

	gt "github.com/seemywingz/gtills"
)

var jenkinsURL = "http://prod2.jenkins.lifion.oneadp.com/view/DITC/view/All/job/DitC_Lookup/"

// Status : get current ditc status
func Status() {
	fmt.Println("Getting Ditc Status")
	res := gt.SendRequest("GET", jenkinsURL)
	fmt.Println(res)
}
