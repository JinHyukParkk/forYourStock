package api


type myCloser interface {
	Close() error
}

// closeFile is a helper function which streamlines closing
// with error checking on different file types.
func closeFile(f myCloser) {
	err := f.Close()
	check(err)
}

// check is a helper function which streamlines error checking
func check(e error) {
	if e != nil {
		panic(e)
	}
}