package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    //return map[string]int{"x": 1}
    words := make([]string,0)
    words = append(words,strings.Fields(s)...)
    
    count := make(map[string]int)
    
    for _,word := range words {
        if cnt,present := count[word] ;present {
            count[word] = cnt+1
        } else {
            count[word] = 1
        }
    }
    
    return count
        
}

func main() {
    wc.Test(WordCount)
}

