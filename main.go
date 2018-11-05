package main

import (
	"flag"
	"log"
)

func main() {
	folderPath := flag.String("folder_path", "", "the bdf folder exact path")
	filesType := flag.String("files_type", "*.dbf", "the files type want to upload")
	accessKeyID := flag.String("access_key_id", "", "s3 access key id")
	secretKey := flag.String("secret_key", "", "s3 secret key")
	bucket := flag.String("bucket", "", "s3 bucket to upload")
	region := flag.String("region", "", "s3 upload region")
	checkPeriodSec := flag.Int("check_period_sec", 60, "checking files status period second")
	flag.Parse()

	dr, err := newDoer(*folderPath, *filesType, *accessKeyID, *secretKey, *bucket, *region)
	if err != nil {
		log.Fatalf("new doer failed:%v", err)
	}

	exec, err := newAutoExec(*checkPeriodSec, dr)
	if err != nil {
		log.Fatalf("new auto exec failed:%v", err)
	}

	exec.close()
}
