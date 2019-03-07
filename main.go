package main

import (
    "io/ioutil"
)

func main() {
    b := []byte{0xDe, 0xaD, 0xBe, 0xeF}

    ioutil.WriteFile("test.bin", b, 0644)
}
