package main

import "fmt"

type person struct {
    name string
    age int
}

func Older(p1, p2 person) (person, int) {
    if p1.age > p2.age {
        return p1, p1.age - p2.age;
    }

    return p2, p2.age - p1.age
}

func main() {
    var rob person
    rob.name = "rob"
    rob.age = 12

    rock := person{"rock", 13}
    lob := person{name:"lob", age:32}

    rr_Older,rr_Diff := Older(rob, rock)
    rl_Older,rl_Diff := Older(rock, lob)
    lr_Older,lr_Diff := Older(lob, rob)

    fmt.Printf("Of %s and %s, %s is older by %d years\n",
    			rob.name, rock.name, rr_Older.name, rr_Diff)
    fmt.Printf("Of %s and %s, %s is older by %d years\n",
    			rock.name, lob.name, rl_Older.name, rl_Diff)
    fmt.Printf("Of %s and %s, %s is older by %d years\n",
    			lob.name, rob.name, lr_Older.name, lr_Diff)
}