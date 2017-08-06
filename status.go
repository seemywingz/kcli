package main

import (
	"fmt"
)

var jenkinsURL = "http://prod2.jenkins.lifion.oneadp.com/view/DITC/view/All/job/DitC_Lookup"

// Status : get current ditc status
func Status() {
	fmt.Println("Getting Ditc Status")
	GetJenkinsBuildNumber(jenkinsURL)
}
