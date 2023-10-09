// Package util
package util

import "os"

// PathExist check the path directory if exist
func PathExist(p string) bool {
	if stat, err := os.Stat(p); err == nil && stat.IsDir() {
		return true
	}
	return false
}

// AnyError returns first non-nil error
func AnyError(errs ...error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}

	return nil
}
