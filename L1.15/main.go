package main

var justString string

func someFunc() {
  v := createHugeString(1 &lt;&lt; 10)
  buf := make([]byte, 100)      
  copy(buf, v[:100])       
  justString = string(buf)       
} 

func main() {
  someFunc()
}