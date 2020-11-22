//+build os OMIT

package main

// INTERFACE_S OMIT
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

// INTERFACE_E OMIT

func main() {

}

/* USG_S OMIT
// File implementa Reader y Writer
File, err := os.Open("archivito.txt")

contents, err := ioutil.ReadAll(File)
bytesWritten, err := File.Write([]byte("Hola Mundo!"))
*/ // USG_E OMIT
