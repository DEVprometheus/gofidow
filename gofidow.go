//#############################################
// gofidow : keyword finder in a web page...
//
// __author__ = 'developer.prometheus@gmail.com (bulent sahin)'
// __version__ = '0.0.1'
//#############################################

package main

import (
     "fmt"
     "os"
     "io/ioutil"
     "net/http"
     "log"
     "strings"
     "regexp"
)

func main() {

     if len(os.Args) < 3 {
         fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
         fmt.Println("Usage:")
         fmt.Println("      go run gofidow.go <FQDN> \"<keyword>\"")
         fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
         os.Exit(1)
        }

     siteo := os.Args[1]

     fido := os.Args[2]; re := regexp.MustCompile(fido)

     resp, err := http.Get(siteo)

     if err != nil {
        fmt.Println("* error: http Get problem...")

        fmt.Printf("%s", err)
        os.Exit(1)
        }

     defer resp.Body.Close()

     body, err := ioutil.ReadAll(resp.Body)

if err == nil {
   lines := strings.Split(string(body), "\n")
   size := len(lines)

   for i:=0; i<size; i++ {
       lino := lines[i]

       resos := (re.FindStringIndex(lino))

       if (len(resos) != 0) {
           fmt.Println(lino)
          }

       }
   } else {
     log.Fatal(err)
          }
//end
}

