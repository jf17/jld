package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	fileUrl := "http://central.maven.org/maven2/com/google/code/gson/gson/2.8.5/gson-2.8.5.jar"

	fileName := "gson-2.8.5.jar"

	err := DownloadFile(fileName, fileUrl)
	if err != nil {
		panic(err)
	}

}


func DownloadFile(fileName string, url string) error {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	oneFolder := "JAR"
	twoFolder := "build"
	threeFolder := "lib"

	dirPath := filepath.Join(pwd, oneFolder, twoFolder, threeFolder)

	fullPath := filepath.Join(dirPath, fileName)

	_, err = os.Stat(dirPath)
	if err != nil {
		fmt.Println("create dir ...")
		os.MkdirAll(dirPath, os.ModePerm)
	}

	fmt.Println("Downloading file...")

	// Create the file
	out, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	size, err := io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	if (size / 1024) > 0 {
		if (size / 1048576) > 0 {
			result := size / 1048576

			fmt.Printf("%s with %v Mbytes downloaded \n", fileName, result)
		} else {
			result := size / 1024
			fmt.Printf("%s with %v Kbytes downloaded \n", fileName, result)
		}
	} else {
		fmt.Printf("%s with %v bytes downloaded \n", fileName, size)
	}

	return nil
}
