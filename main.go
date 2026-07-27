package main

import (
    "fmt"
    "io/fs"
    "path/filepath"
    "plugin" 
)

// if `symbol` has no default value then the function will return `defaultvalue`
func setdefaultvalue[strorslice string | []string](defaultvalue strorslice, symbol plugin.Symbol, er error) strorslice {
    if er == nil {
		    return *symbol.(*strorslice)
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
        nameSymbol, err := plug.Lookup("Name")
        name := setdefaultvalue(filepath.Base(path), nameSymbol, err)
        descriptionSymbol, err := plug.Lookup("Description")
        description := setdefaultvalue("None provided.", descriptionSymbol, err)
        categoriesSymbol, err := plug.Lookup("Categories")
        categories := setdefaultvalue([]string{"uncategorised"}, categoriesSymbol, err)

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

