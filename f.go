package fib

func Fib(n int64) int64 {
    if n < 2{return 1}
    return Fib(n-2) + Fib(n-1)
}



