package model_test

import "book/model"
import "testing"
import "fmt"

func Test_GetTablet(t *testing.T) {
	fmt.Println("table name: ", model.Table)
	fmt.Println("table struct: ", model.TableStruct)
}
