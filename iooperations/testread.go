package iooperations

import (
	"fmt"
	"os"
	"bufio"
	"io"

)

/**
 * Created by John Tsantilis 
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 24/8/2017.
 */

func readFileWithReadString(fn string) (err error) {
	fmt.Println("readFileWithReadString")

	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return err
	}

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')

		fmt.Printf(" > Read %d characters\n", len(line))

		// Process the line here.
		fmt.Println(" > > " + limitLength(line, 50))

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return
}

func readFileWithScanner(fn string) (err error) {
	fmt.Println("readFileWithScanner - this will fail!")

	// Don't use this, it doesn't work with long lines...

	file, err := os.Open(fn)
	defer file.Close()

	if err != nil {
		return err
	}

	// Start reading from the file using a scanner.
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Printf(" > Read %d characters\n", len(line))

		// Process the line here.
		fmt.Println(" > > " + limitLength(line, 50))
	}

	if scanner.Err() != nil {
		fmt.Printf(" > Failed!: %v\n", scanner.Err())
	}

	return
}

func main() {
	testLongLines()
	testLinesThatDoNotFinishWithALinebreak()

}

func testLongLines() {
	fmt.Println("Long lines")
	fmt.Println()

	createFileWithLongLine("longline.txt")
	readFileWithReadString("longline.txt")
	fmt.Println()
	readFileWithScanner("longline.txt")
	fmt.Println()
	readFileWithReadLine("longline.txt")
	fmt.Println()
}

func testLinesThatDoNotFinishWithALinebreak() {
	fmt.Println("No linebreak")
	fmt.Println()

	createFileThatDoesNotEndWithALineBreak("nolinebreak.txt")
	readFileWithReadString("nolinebreak.txt")
	fmt.Println()
	readFileWithScanner("nolinebreak.txt")
	fmt.Println()
	readFileWithReadLine("nolinebreak.txt")
	fmt.Println()
}

func createFileThatDoesNotEndWithALineBreak(fn string) (err error) {
	file, err := os.Create(fn)
	defer file.Close()

	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)
	w.WriteString("Does not end with linebreak.")
	w.Flush()

	return
}

func createFileWithLongLine(fn string) (err error) {
	file, err := os.Create(fn)
	defer file.Close()

	if err != nil {
		return err
	}

	w := bufio.NewWriter(file)

	fs := 1024 * 1024 * 4 // 4MB

	// Create a 4MB long line consisting of the letter a.
	for i := 0; i < fs; i++ {
		w.WriteRune('a')
	}

	// Terminate the line with a break.
	w.WriteRune('\n')

	// Put in a second line, which doesn't have a linebreak.
	w.WriteString("Second line.")

	w.Flush()

	return
}

func limitLength(s string, length int) string {
	if len(s) < length {
		return s
	}

	return s[:length]
}
