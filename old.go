package main

// import (
// 	"embed"
// 	"fmt"
// 	"io"
// 	"os"
// 	"strings"
// )

// //go:embed licenses/*.txt
// var licenseFiles embed.FS

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("usage: licensr list/print/make [license-name]")
// 		os.Exit(0)
// 	}

// 	files, err := licenseFiles.ReadDir("licenses")
// 	if err != nil {
// 		panic(err)
// 	}

// 	switch os.Args[1] {
// 	case "list":
// 		for _, file := range files {
// 			fmt.Println(strings.TrimSuffix(file.Name(), ".txt"))
// 		}

// 	case "print":
// 		if len(os.Args) != 3 {
// 			goto noLicense
// 		}

// 		want := strings.ToLower(os.Args[2])
// 		for _, file := range files {
// 			if strings.ToLower(strings.TrimSuffix(file.Name(), ".txt")) == want {
// 				contents, err := licenseFiles.ReadFile("licenses/" + file.Name())
// 				if err != nil {
// 					panic(err)
// 				}
// 				fmt.Println(string(contents))
// 			}
// 		}

// 	case "make":
// 		if len(os.Args) != 3 {
// 			goto noLicense
// 		}

// 		var licenseTxt string

// 		want := strings.ToLower(os.Args[2])
// 		for _, file := range files {
// 			if strings.ToLower(strings.TrimSuffix(file.Name(), ".txt")) == want {
// 				contents, err := licenseFiles.ReadFile("licenses/" + file.Name())
// 				if err != nil {
// 					panic(err)
// 				}
// 				licenseTxt = string(contents)
// 			}
// 		}
// 		if licenseTxt == "" {
// 			fmt.Printf("No license with name %s found.\n", os.Args[2])
// 			os.Exit(1)
// 		}

// 		dst, err := os.Create("LICENSE")
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer dst.Close()

// 		_, err = io.Copy(dst, strings.NewReader(licenseTxt))
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Printf("Success! You are now licensed under %s.\n", want)

// 	default:
// 		fmt.Printf("Error: Not understood: %s\n", os.Args[1])
// 		os.Exit(1)
// 	}

// 	os.Exit(0)

// noLicense:
// 	fmt.Println("Error: You must specify a license.")
// 	os.Exit(1)

// }
