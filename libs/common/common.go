package common

import "log"

type myCloser interface {
	Close() error
}

// closeFile is a helper function which streamlines closing
// with error checking on different file types.
func CloseFile(f myCloser) {
	err := f.Close()
	Check(err)
}

// check is a helper function which streamlines error checking
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
