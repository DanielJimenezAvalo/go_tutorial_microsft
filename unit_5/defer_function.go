/*A typical use case for the defer function is to close a file after you finish using it. 
Here's an example:

After you create or open a file, you defer the .Close() function to avoid forgetting to close the file after you're done.

*/

package main

import (
	"io"
	"os"
	"fmt"
)

func main(){

	newfile, error:=os.Create("learnGo.txt")

	if error != nil {

		fmt.Println("error: could not create file")

		return 

	}

	defer newfile.Close()

	if _, error =io.WriteString(newfile, "learnign Go!"); error != nil {

		fmt.Println("error: could not write to file")

		return
	}
	
	newfile.Sync()

}