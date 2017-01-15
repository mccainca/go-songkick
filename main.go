package songkick

import (
	"fmt"
	"os"
)

//APIUrl ... base url
const APIUrl = "http://api.songkick.com/api/3.0/"

//APIKey ... key for API.  Set environment variable
var APIKey = os.Getenv("SK_API_KEY")

//Hello ... test your env vars
func Hello() {
	fmt.Printf("Hello.  Your API Key is %s", APIKey)
}
