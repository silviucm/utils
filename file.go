package utils

import "os"
import "io"

type fileUtils struct{}

// single variable acting as the FileUtils "subpackage" inside the legit utils package
var File fileUtils

// Checks if the file with the given path exists, returns true if yes
func (dummyReceiver *fileUtils) Exists(name string) bool {

	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (dummyReceiver *fileUtils) Delete(filePath string) error {

	return os.Remove(filePath)

}

// CheckClose is used to check the return from Close in a defer statement.
// Typical usage would be to pass the pointer of the error that is returned by
// the caller function:
//
// defer checkClose(out, &err)
// io.Copy(out, resp.Body)
// return err
//
// In some scenarios, it is possible the file was closed before the defer statement
// This function insures that an error is still captured in that case
func (dummyReceiver *fileUtils) CheckClose(c io.Closer, err *error) {
	cerr := c.Close()
	if *err == nil {
		*err = cerr
	}
}
