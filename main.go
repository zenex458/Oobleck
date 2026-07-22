package main

import (
    "fmt"
    "io/fs"
    "path/filepath"
    "plugin" 
)

func main() {
    // TODO: make "./build/scripts/" part of this cwd-agnostic
    filepath.WalkDir("./build/scripts", func(path string, d fs.DirEntry, err error) error {
        if d.IsDir() {
            return nil
        }

        plug, err   := plugin.Open(path)
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (plugin couldn't be opened)\n", path)
            return nil;
        }

        name, err     := plug.Lookup("Name")
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (`Test` symbol not found)\n", path)
            return nil;
        }

        fmt.Printf("Found: %s\n", name.(func() string)())
        return nil
    })
}

