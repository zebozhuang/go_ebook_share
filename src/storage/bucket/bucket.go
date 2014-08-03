package bucket

import (
	."github.com/qiniu/api/conf"
	"github.com/qiniu/api/rs"
	"github.com/qiniu/api/io"
	"github.com/qiniu/log"
	"conf"
)

func init() {
	ACCESS_KEY = conf.QINIU_ACCESS_KEY
	SECRET_KEY = conf.QINIU_SECRET_KEY
}

type Bucket struct {
	BucketName string
}

func (bucket *Bucket)GetToken() string {
	putPolicy := rs.PutPolicy {
		Scope: bucket.BucketName,
	}
	return putPolicy.Token(nil)
}

func (bucket *Bucket)Upload(fname string, path string, ret interface{}) bool {
	uptoken := bucket.GetToken()	
	var err error	
	var extra = &io.PutExtra{}

	err = io.PutFile(nil, &ret, uptoken, "test.txt", "test.txt",  extra)	
	if err != nil {
		log.Printf("upload file failed:", err)			
		return false
	}
	return true
}

func NewBucket(name string) *Bucket {
	return &Bucket{BucketName: name}
}

var bucket Bucket = *NewBucket(conf.QINIU_BucketName)
