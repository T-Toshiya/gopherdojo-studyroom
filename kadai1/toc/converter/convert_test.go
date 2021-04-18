package converter

import (
	"testing"
)

func TestConverter_Convert(t *testing.T) {
	c := &Converter{
		BeforeFmt: "jpg",
		AfterFmt:  "png",
		Directory: "./testCase",
		FilePath:  "test.jpg",
	}
	err := c.Convert()
	if err != nil {
		t.Errorf("変換処理でエラーが発生")
	}
}
