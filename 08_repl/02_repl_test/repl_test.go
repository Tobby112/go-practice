package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func main() {
	Start(os.Stdin, os.Stdout)
}

/*
This will behaves like:
>> 1 + 1
#=> 2
>>
*/

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		out.Write([]byte(">> "))

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		// Other logics
		fmt.Printf("%s\n", line)
	}
}

func TestStart(t *testing.T) {

	in := &bytes.Buffer{}
	out := &bytes.Buffer{}

	in.WriteString("1 + 1\n")
	in.WriteString("exit\n") // exit the input mode and continues test

	Start(in, out)

	result := out.String()
	fmt.Println(result)
}

type CHIO struct {
	ch chan []byte
}

func (c *CHIO) Read(p []byte) (n int, err error) {
	readBytes, ok := <-c.ch
	if ok == false {
		return 0, io.EOF
	}
	n = copy(p, readBytes)
	return
}

func (c *CHIO) Write(p []byte) (n int, err error) {
	c.ch <- p
	return len(p), nil
}

func TestStartByChio(t *testing.T) {
	in := &CHIO{ch: make(chan []byte, 1)}
	out := &CHIO{ch: make(chan []byte, 1)}
	go Start(in, out)
	inputList := [][]byte{[]byte("1+1\n"), []byte("exit\n")}
	for _, s := range inputList {
		r, ok := <-out.ch
		if ok == false {
			break
		}
		fmt.Printf("%s", string(r))
		in.ch <- s
	}
	close(in.ch)

	r, _ := <-out.ch
	fmt.Printf("%s", string(r))
	close(out.ch)
	fmt.Printf("\n")
}

func logicOfStart([]byte) []byte {
	return []byte("2")
}

func TestStartLogic(t *testing.T) {
	inputList := [][]byte{[]byte("1+1\n"), []byte("exit\n")}

	r := logicOfStart(inputList[0])

	expectedString := "2"
	c := strings.Compare(expectedString, string(r))
	if c != 0 {
		t.Errorf("c:%d expectedString is %s not %s\n", c, expectedString, r)
	}
}

/*
Question: How to write unit test for Start()?
When testing no matter what io.Reader I use, the scanner closes right away.
I want to write something like:
func TestStart(t *testing.T) {
	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	Start(in, out)
	in.WriteString("1 + 1")
    in.WrtieString("exit") // exit the input mode and continues test
	result := out.String()
	fmt.Println(result)
}
And the result should be:
>> 1 + 1
#=> 2
>>
But I only got:
>>
*/
