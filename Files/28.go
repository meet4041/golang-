package main

import (
	"fmt"
	"os"
)

func main() {

	//Files
	// file, error := os.Open("example.txt")
	// if error != nil {
	// 	panic(error)
	// }

	// fileInfo, error := file.Stat()
	// if error != nil {
	// 	panic(error)
	// }

	// fmt.Println("File Name:", fileInfo.Name())
	// fmt.Println("File Size:", fileInfo.Size())
	// fmt.Println("File or Folder:", fileInfo.IsDir())
	// fmt.Println("File Modified At :", fileInfo.ModTime())

	// defer file.Close()

	// Read
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fileInfo, error := f.Stat()
	// if error != nil {
	// 	panic(error)
	// }
	// defer f.Close()

	// buf := make([]byte, fileInfo.Size())
	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// for i := 0; i < len(buf); i++ {
	// 	fmt.Println("data", d, string(buf[i]))
	// }

	// ReadFile
	// f, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(f))
	// defer f.Close()

	// Read Folder
	// dir, err := os.Open("../")
	// dir, err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }
	// defer dir.Close()

	// fileInfo, err := dir.Readdir(0)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, file := range fileInfo {
	// 	fmt.Println(file.Name(), file.IsDir())
	// }

	// Create a file
	// file, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// file.WriteString("Hello, World!")
	// file.WriteString("Hello, World-1!")

	// bytes := []byte("Hello, World-2!")
	// file.Write(bytes)

	// Read and write to another file streaming fashion
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}
	// }
	// writer.Flush()
	// fmt.Println("Written to new file successfully")

	// Delete a file
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File deleted successfully")
}

// Commmand to run this file -
// go run 28.go
