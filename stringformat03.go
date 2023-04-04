/* Alta3 Research | RZFeeser
   Using fmt.Printf - String formatting structs and PirateTreasure*/

package main

import (
    "fmt"
)

// a structure would be analagous to a "class" in other languages
// like a "class" a structure works like a blueprint
type PirateTreasure struct {
    x, y int            // x and y are both "int" {x y}
}

func main() {

    // create a "struct" called, 'p' of the type "PirateTreasure"
    p := PirateTreasure{1, 2}                // here we create a struct "p"
    
    
    // This prints an instance of our PirateTreasure struct.
    fmt.Printf("struct1: %v\n", p)

    // If the value is a struct, the %+v variant will include the struct’s field names
    fmt.Printf("struct2: %+v\n", p)

    // The %#v variant prints a Go syntax representation of the value
    //   i.e. the source code snippet that would produce that value
    fmt.Printf("struct3: %#v\n", p)

    // To print the type of a value, use %T
    fmt.Printf("type: %T\n", p)

    // To print a representation of a PirateTreasure, use %p
    fmt.Printf("PirateTreasure: %p\n", &p)     // structures are a topic we'll dive deeper into in subsequent labs
}

