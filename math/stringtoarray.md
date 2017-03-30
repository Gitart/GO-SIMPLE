# How to iterate over a []string(array)

```golang
package main

 import (
         "fmt"
 )

 func main() {

         // taken from http://www.asitis.com/16/

         divineValues := []string{
                 "fearlessness",
                 "purification of one's existence",
                 "cultivation of spritual knowledge",
                 "charity",
                 "self-control",
                 "performance of sacrifice",
                 "study of the Vedas",
                 "austerity and simplicity",
                 "non-violence",
                 "truthfulness",
                 "freedom from anger",
                 "renunciation",
                 "tranquility",
                 "aversion to faultfinding",
                 "compassion and freedom from covetousness",
                 "gentleness",
                 "modesty and steady determination",
                 "vigor",
                 "forgiveness",
                 "fortitude",
                 "cleanliness",
                 "freedom from envy and the passion for honor",
         }

         for index, each := range divineValues {
                 fmt.Printf("Divine value [%d] is [%s]\n", index, each)
         }
 }
 ```
