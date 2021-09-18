package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "errors"
)

func ReadFile(fileNames []string) (string, error){
    var contents string
    for _, file := range fileNames {
        //Get absolute file path
        fileName, err := filepath.Abs(file)
        if (err != nil) {
            errMsg := fmt.Sprintf ("File not found: %s\n", err)
            fmt.Println (errMsg)
            os.Exit(0)
        }
        //Read file
        //s := fmt.Sprintf("Reading file : %s\n", fileName)
        //fmt.Println(s)
        data, err := ioutil.ReadFile(fileName)
        if (err != nil) {
            return "", errors.New(fmt.Sprintf("Could not read file : %s\n", err))
        } else {
            contents += "\n" + "[+] Reading file : "+fileName + "\n" + string(data)
        }
    }
    return contents, nil
}

func main() {
    if (len(os.Args) == 1) {
        fmt.Println("Usage : cat <file_name1> [<file_name2> ...]")
        os.Exit(0)
    }
    
    //Read the file
    contents, err := ReadFile(os.Args[1:])
    if (err != nil) {
        errMsg := fmt.Sprintf ("Error occurred when reading file: %s\n", err)
        fmt.Println (errMsg)
        os.Exit(0)
    }
    io.WriteString(os.Stdout, contents)    
}