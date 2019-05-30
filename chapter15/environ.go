package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Executable())

	fmt.Println(os.Environ())
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Getenv("DEBUG"))
	fmt.Println(os.LookupEnv("DEBUG"))

	fmt.Println(os.Setenv("DEBUG", "True"))
	fmt.Println(os.LookupEnv("DEBUG"))

	fmt.Println(os.Unsetenv("DEBUG"))
	fmt.Println(os.LookupEnv("DEBUG"))

	os.Clearenv()
	fmt.Println(os.Environ())

	fmt.Println(os.Hostname())
	fmt.Println(os.TempDir())
	fmt.Println(os.UserHomeDir())
	fmt.Println(os.UserCacheDir())
}
