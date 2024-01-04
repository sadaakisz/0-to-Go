package main

type Copier interface {
	Copy(sourceFile, destinationFile string) (bytesCopied int)
}

func main() {

}
