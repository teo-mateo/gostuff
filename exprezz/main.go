package main

import "github.com/teo-mateo/gostuff/exprezz/impl"
import "fmt"

func main(){
	expression, err := exprezz.Parse("1+1")
	if err != nil{
		panic(err.Error())
	}
	result := expression.Eval(nil)

	fmt.Println(result)
}