package manager_test

import bMgr "book/manager"
import "testing"
import "fmt"

func Test_Create(t *testing.T) {
	info, err := bMgr.Manager.Create("AAAAAA", "test.txt")
	fmt.Println("return :", info, err)
}

func Test_GetFileInfo(t *testing.T) {
	info, err := bMgr.Manager.GetFileInfo(2)
	fmt.Println("return :", info, err)

	info, err = bMgr.Manager.GetFileInfo(10)
	fmt.Println("return :", info, err)
}

func Test_GetFileInfos(t *testing.T) {
	info, err := bMgr.Manager.GetFileInfos()
	fmt.Println("return :", info, err)
}
