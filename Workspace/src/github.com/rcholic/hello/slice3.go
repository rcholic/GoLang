package main
import "fmt"
import "strconv"

func main() {
    var a = make([]string, 5, 5)

    for i := 0; i < 5; i++ {
        //cast int to string using strconv
        a[i] = "hello world " + strconv.Itoa(i)
    }

    for _, val := range a {
        fmt.Println(val)
    }
}
