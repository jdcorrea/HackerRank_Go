package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    
    // Constants
    const delimiter = "\n"

    // Declare second integer, double, and String variables.
    var int2 uint64
    var double2 float64
    var string2 string

    // Read and save an integer, double, and String to your variables.
    fmt.Scanf("%d", &int2) //using pointers to assing the value
    fmt.Scanf("%f", &double2) //using pointers to assing the value

    //using scanner it's mandatory in the exercise, so I use annonymous func to use it
    string2 = func() string{
      scanner.Scan()
      return scanner.Text()
    }()

    fmt.Scanf("%s", &string2) //using pointers to assing the value
    
    // Print the sum of both integer variables on a new line.
    fmt.Printf("%d" + delimiter, i + int2)

    // Print the sum of the double variables on a new line.
    fmt.Printf("%.1f" + delimiter, d + double2)

    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.
    fmt.Println(s + string2)
}
