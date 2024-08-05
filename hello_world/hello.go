package main

import "fmt"

var languagePrefixMap = map[string]string {
    "English": "Hello",
    "Spanish": "Hola",
    "Uzbek": "Salom",
}

func Hello(name, language string) string {
    if name == ""{
        name = "World"
    }
    if language == ""{
        language = "Uzbek"
    }

    return languagePrefixMap[language] + ", " + name
}

func main() {
    fmt.Println(Hello("World",""))
}

