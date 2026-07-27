package main

import (
    "fmt"
    "io/fs"
    "path/filepath"
    "plugin" 
)

// if the symbol with field of `symbolname' has no default value then the function will return `defaultvalue`
func setdefaultvalue[strorslice string | []string](defaultvalue strorslice, symbolname string, plug *plugin.Plugin) strorslice {
		lookup, err := plug.Lookup(symbolname)
    if err == nil {
		    return *lookup.(*strorslice)
		}
		return defaultvalue
}

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
        name := setdefaultvalue(filepath.Base(path), "Name", plug)
        description := setdefaultvalue("None provided.", "Description", plug)
        categories := setdefaultvalue([]string{"uncategorised"}, "Categories", plug)

        check, err := plug.Lookup("Check")
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (symbol `Check` not found)\n", path)
            return nil
        }

        fixDescription, err := plug.Lookup("FixDescription")
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (symbol `FixDescription` not found)\n", path)
            return nil
        }

        /*
        fixauto, err := plug.Lookup("FixAuto")
        if err != nil {
            fmt.Printf("ERROR:\tSkipping %s (symbol `FixAuto` not found)\n", path)
            return nil
        }
        */

        // Running tests
        if check.(func() bool)() {
            fmt.Printf("---- %s ----\nNAME: %s\nDESCRIPTION: %s\nCATEGORIES: %s\nFIX DESCRIPTIONS: %s\n\n", path, name, description, categories, fixDescription.(func() string)())
        }

        return nil
    })
}

