package main

import ( "fmt"; "io"; "strings"; "golang.org/x/tour/reader" )

func main() {
	read()
	infinite()
}

func read() {
	fmt.Println("-- package strings : method NewReader")
	
	r := strings.NewReader("Hello, Reader!")
	fmt.Println("length : ", r.Len())

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF { break }
	}
}

func infinite(){
	fmt.Println("-- Reader type that emits an infinite stream of 'A'")
	reader.Validate(MyReader{})
}

type MyReader struct{}

var n int = 0;
const ncon int = 3000;

// Add Read([]byte) (int, error) method to MyReader
func (m MyReader) Read(b []byte) (int, error){
	for x := range b { b[x] = 'A' }
	fmt.Println("size of b[]: ", len(b))
	if(n + len(b) > ncon) { b[ncon - n] = 'B' } else { n += len(b) }
	return len(b),nil
}
