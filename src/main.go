package main

import "plugin" 

func main() {
    plug, err   := plugin.Open("./build/tests/test.so")
    if err != nil {
        panic(err)
    }

    fn, err     := plug.Lookup("Fooandbar")
    if err != nil {
        panic(err)
    }

    fn.(func())()
}

