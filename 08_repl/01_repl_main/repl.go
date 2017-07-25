package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func main() {
	Start(os.Stdin, os.Stdout, nil)
}

/*
This will behaves like:
>> 1 + 1
#=> 2
>>
*/

func Start(in io.Reader, out io.Writer, wg *sync.WaitGroup) {
	scanner := bufio.NewScanner(in)

	for {
		out.Write([]byte(">> "))
		scanned := scanner.Scan()

		if !scanned {
			fmt.Printf("scan:%v\n", scanner.Err())
			break
			//continue
		}

		line := scanner.Text()
		// Other logics
		fmt.Printf("line:%s\n", line)
		time.Sleep(time.Second * 1)
	}
	fmt.Printf("[Start] done\n")
	wg.Done()
}
