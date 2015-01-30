package main
import "fmt"

type Pill struct {
    name string
    age int
}

//construct a stringer for outputing the sring format
func (p *Pill) toString() string {

    return fmt.Sprintf("name = %s, age = %d\n", p.name, p.age)
}

func main() {
    //aPill is a pointer to Pill struct
    var aPill = &Pill{"Aspirin", 20}
    fmt.Printf("aPill's age is %d\n", (*aPill).age)
    fmt.Printf(aPill.toString())
}
