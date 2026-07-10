// test.go
// Type information and utilities for tests



// Possible categories for a test
type TestCategory int

const (
    OpsecTest TestCategory = itoa
)

// A utility function that converts from a [TestCategory] type to a string
func category_name(category TestCategory) {
    // TODO
}



// Underlying test struct - includes all information about a test that will be
// displayed to the user
type Test struct {
    name string
    description string
    categories []string

    fix_description string
    fix_auto func()
}
