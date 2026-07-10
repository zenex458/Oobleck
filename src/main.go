package main

import "plugin" 

func main() {
    plug, err   := plugin.Open("./build/scripts/example_script.so")
    if err != nil {
        panic(err)
    }

    fn, err     := plug.Lookup("Fooandbar")
    if err != nil {
        panic(err)
    }

    fn.(func())()
}

