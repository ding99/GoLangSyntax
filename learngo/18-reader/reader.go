package main

import ( "fmt"; "io"; "os"; "strings"; "golang.org/x/tour/reader" )

func main() {
	read()
	infinite()
	rot13()
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
	fmt.Println("-- Reader type that emits an 'infinite' stream of 'A'")
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

type rot13Reader struct { r io.Reader }

func rot13(){
	fmt.Println("-- rot13Reader");
	s := strings.NewReader("Lbh abcd efgh afdasfasd")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func (rot *rot13Reader) Read(b []byte) (read int, err error){
	read, err = rot.r.Read(b)
	for i := 0; i < read; i++ { b[i] = rot13byte(b[i]) }
	return
}

func rot13byte(b byte) byte {
    s := rune(b)
    if s >= 'a' && s < 'n' || s >= 'A' && s < 'N' { b += 13 }
    if s > 'm' && s <= 'z' || s > 'M' && s <= 'Z' { b -= 13 }
    return b
}
