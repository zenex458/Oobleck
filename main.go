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
        // Skipping directories
        if d.IsDir() {
            return nil
        }

        // Opening shared library
        plug, err   := plugin.Open(path)
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (plugin couldn't be opened)\n", path)
            return nil
        }

        // Running prerequisite check if it exists
        prereq, err := plug.Lookup("Prereq");
        if err != nil {
            fmt.Printf("WARN:\t Symbol `Prereq` cannot be found in %s\n", path)
        } else {
            if !prereq.(func() bool)() {
                fmt.Printf("INFO:\t Skipping %s (`Prereq` returned false)\n", path)
                return nil;
            }
        }

        // Unwrapping other script symbols
        name, err := plug.Lookup("Name");
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (symbol `Name` not found)\n", path)
            return nil
        }

        description, err := plug.Lookup("Description");
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (symbol `Description` not found)\n", path)
            return nil
        }

        categories, err := plug.Lookup("Categories");
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (symbol `Categories` not found)\n", path)
            return nil
        }

        check, err := plug.Lookup("Check");
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (symbol `Check` not found)\n", path)
            return nil
        }

        fixdescription, err := plug.Lookup("FixDescription");
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (symbol `FixDescription` not found)\n", path)
            return nil
        }

        /*
        fixauto, err := plug.Lookup("FixAuto");
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (symbol `FixAuto` not found)\n", path)
            return nil
        }
        */

        // Running tests
        if check.(func() bool)() {
            fmt.Printf("---- %s ----\nNAME: %s\nDESCRIPTION: %s\nCATEGORIES: %s\nFIX DESCRIPTIONS: %s\n", path, *name.(*string), *description.(*string), *categories.(*[]string), fixdescription.(func() string)())
        }

        return nil
    })
}

