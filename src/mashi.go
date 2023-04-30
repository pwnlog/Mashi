package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "path/filepath"
    "strings"
    "unicode"
)

func showHelp(programName string) {
	fmt.Println(`
  __  __           _     _ 
 |  \/  |         | |   (_)
 | \  / | __ _ ___| |__  _ 
 | |\/| |/ _| / __| '_ \| |
 | |  | | (_| \__ \ | | | |
 |_|  |_|\__,_|___/_| |_|_|
          
 	Author: pwnlog
	`)
	fmt.Printf("Usage: %s names.txt\n", filepath.Base(programName))
	fmt.Println("Description: Reads names from a text file and generates possible usernames based on those names.")
	fmt.Println("Option:")
	fmt.Println("\t-h, --help  Show this help message and exit.")
}

func main() {
    if len(os.Args) != 2 {
        showHelp(os.Args[0])
        os.Exit(0)
    }
    
    if os.Args[1] == "-h" || os.Args[1] == "--help" {
        showHelp(os.Args[0])
        os.Exit(0)
    }

    filename := os.Args[1]
    if _, err := os.Stat(filename); os.IsNotExist(err) {
	    fmt.Printf("Error: %s not found\n", filename)
        os.Exit(0)
    }

    data, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
        os.Exit(0)
    }

    lines := strings.Split(string(data), "\n")
    for _, line := range lines {
        name := strings.Join(strings.FieldsFunc(line, func(c rune) bool {
            return !unicode.IsLetter(c) && !unicode.IsSpace(c)
        }), "")

        tokens := strings.Fields(strings.ToLower(name))
        if len(tokens) < 1 {
            continue
        }

        fname := tokens[0]
        lname := ""
        if len(tokens) == 2 {
            lname = tokens[1]
        } else if len(tokens) > 2 {
            lname = strings.Join(tokens[1:], "")
        }

        if len(fname) > 0 && len(lname) > 0 {
            fmt.Println("\n[+] Minimalized")
            fmt.Println(fname + lname)
            fmt.Println(lname + fname)
            fmt.Println(fname + "." + lname)
            fmt.Println(lname + "." + fname)
            fmt.Println(lname + fname[:1])
            fmt.Println(fname[:1] + lname)
            fmt.Println(lname[:1] + fname)
            fmt.Println(fname[:1] + "." + lname)
            fmt.Println(lname[:1] + "." + fname)
            fmt.Println("\n[+] Capitalized")
            fmt.Println(strings.Title(fname) + strings.Title(lname))
            fmt.Println(strings.Title(lname) + strings.Title(fname))
            fmt.Println(strings.Title(fname) + "." + strings.Title(lname))
            fmt.Println(strings.Title(lname) + "." + strings.Title(fname))
            fmt.Println(strings.Title(lname) + strings.Title(fname[:1]))
            fmt.Println(strings.Title(fname[:1]) + strings.Title(lname))
            fmt.Println(strings.Title(lname[:1]) + strings.Title(fname))
            fmt.Println(strings.Title(fname[:1]) + "." + strings.Title(lname))
            fmt.Println(strings.Title(lname[:1]) + "." + strings.Title(fname))
        }
        if len(fname) > 0 {
            fmt.Println("\n[+] Single Detected")
            fmt.Println(fname)
            fmt.Println(strings.Title(fname))
        }
        if len(lname) > 0 {
            fmt.Println("\n[+] Single Detected")
            fmt.Println(lname)
            fmt.Println(strings.Title(lname))
        }
    }
}

