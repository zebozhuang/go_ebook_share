package bucket_test

import "testing"
import ."storage/bucket"
import "github.com/qiniu/log"
import "github.com/qiniu/api/io"

var (
	bucketName string
)
 
func init() {
	bucketName = "bookshare"
}

func Test_GetToken(t *testing.T) {
	b := Bucket{bucketName}	
	uptoken := b.GetToken()
	if(uptoken != ""){
		log.Printf(uptoken)
	} else {
		t.Error("fail to get token")
	}
}

func Test_Upload(t *testing.T) {
	ret := new(io.PutRet)
	b := Bucket{bucketName}

	result := b.Upload("test.txt", "", &ret)
	if (result) {
		log.Printf("return:", ret)
	} else {
		t.Error("fail to upload:", ret)
	}
}
