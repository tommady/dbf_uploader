package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

type doer struct {
	fpath        string
	ftype        string
	uploader     *s3manager.Uploader
	uploadFolder string
	uploadBucket string
	uploadRegion string
}

func newDoer(folderPath, filesType, accessKeyID, secretKey, bucket, region string) (*doer, error) {
	if !filepath.IsAbs(folderPath) {
		return nil, errors.Errorf("doer[newDoer] file path:%s is not absolute", folderPath)
	}
	if _, err := os.Stat(folderPath); err != nil {
		return nil, errors.Wrapf(err, "doer[newDoer] file path:%s detect failed", folderPath)
	}

	return &doer{
		fpath:        folderPath,
		ftype:        filesType,
		uploadFolder: filepath.Base(folderPath),
		uploadBucket: bucket,
		uploader: s3manager.NewUploader(session.New(&aws.Config{
			CredentialsChainVerboseErrors: aws.Bool(true),
			Region:                        aws.String(region),
			Credentials:                   credentials.NewStaticCredentials(accessKeyID, secretKey, ""),
		})),
	}, nil
}

func (d *doer) do() error {
	fs, err := ioutil.ReadDir(d.fpath)
	if err != nil {
		return errors.Wrapf(err, "doer[do] read dir failed")
	}

	for _, f := range fs {
		if !f.IsDir() && filepath.Ext(f.Name()) == d.ftype {
			sf, err := ioutil.ReadFile(f.Name())
			if err != nil {
				return errors.Wrapf(err, "doer[do] failed on file:%s", f.Name())
			}

			_, err = d.uploader.Upload(&s3manager.UploadInput{
				Bucket: aws.String(d.uploadBucket),
				Key:    aws.String(filepath.Join(d.uploadFolder, f.Name())),
				Body:   bytes.NewReader(sf),
			})
			if err != nil {
				return errors.Wrapf(err, "doer[do] upload file:%s failed", f.Name())
			}
		}
	}

	return nil
}
