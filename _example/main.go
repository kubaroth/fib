// fib(40)
// time: 4.364014633s
// in gomacro 2min16s
package main

import "fmt"
import "time"
import "github.com/kuba/fib"

func main(){
	start := time.Now()
    //fmt.Printf("%v\n", fib(40));
	fmt.Printf("%v\n", fib.Fib(20));
	fmt.Println(time.Since(start))
}

