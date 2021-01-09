package a07builder

import (
	"fmt"
	"testing"
)

func Test_Func(t *testing.T) {
	var Director DirectorStu
	var BuilderText BuilderTextStu
	BuilderText.Init()
	Director.Init(&BuilderText)
	Director.Construct()
	fmt.Println(BuilderText.GetResult())

}

func DisFunc() {

}
