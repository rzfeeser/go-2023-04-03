/* Alta3 Research | RZFeeser@alta3.com
   Understanding Go Receiver Functions (i.e. methods)

   func(receiver_name Type) method_name(parameter_list)(return_type){
   // Code
   }

*/

package main

import (
    "fmt"
    "github.com/username/supermario/models"
)


func main() {
    //mario := Player{3, 1, []string{"mushroom"}}
    mario := models.Player{3, 1, []string{"mushroom"}}

    // display mario's current lives
    fmt.Println(mario.Lives)
    // mario just touched a greenmushroom!
    mario.Greenmushroom()
    // display mario's current lives have changed
    fmt.Println(mario.Lives)


    // display mario's current inventory
    fmt.Println(mario.Inventory)
    // mario just picked up a flower!
    mario.Pickup("flower")
    // display mario's current inventory has changed
    fmt.Println(mario.Inventory)


    // below stage 5 mario cannot use the warp whistle
    fmt.Println(mario.CanWhistle())
    // move to a stage where mario can use the warp whistle
    mario.Stage = 7
    // mario can only use the warp whistle on or above stage 5
    fmt.Println(mario.CanWhistle())

}

