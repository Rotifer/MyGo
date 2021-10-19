package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	tmpDir := path.Join(os.TempDir())
	fmt.Printf("Temp Dir: %v\n", os.TempDir())
	fmt.Printf("The temp dir on this machine is %s\n", tmpDir)
}
