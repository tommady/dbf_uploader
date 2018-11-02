package main

import (
	"flag"
	"log"
)

func main() {
	folderPath := flag.String("folder_path", "", "the bdf folder exact path")
	filesType := flag.String("files_type", "*.dbf", "the files type want to upload")
	checkPeriodSec := flag.Int("check_period_sec", 60, "checking files status period second")
	flag.Parse()

	exec, err := newAutoExec(*checkPeriodSec)
	if err != nil {
		log.Fatalf("new auto exec failed:%v", err)
	}

	newDoer(*folderPath, *filesType)

	exec.close()
}
