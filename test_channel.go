package main
import "fmt"
import "time"

func s(str string) {
    fmt.Println(time.Now(), str)
}

func caller() chan bool {
    s("calling fn called from caller")
    isDone := make(chan bool)
    go func(isDone chan bool) {
        s("entered go routine")
        called()
        s("called fn called in go routine")
        close(isDone)
        s("closed channel in go routine")
    }(isDone)
    s("called fn called from caller")
    return isDone
}

func called() {
    s("reached fn called")
    time.Sleep(3000*time.Millisecond)
    s("leaving fn called")
}

func test() {
    isDone := caller()
    s("waiting in test")
    <-isDone
    s("waiting completed in test")
}

func main(){
    test()
}
