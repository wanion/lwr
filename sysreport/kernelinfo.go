package sysreport

// getKernelInfo gets information with the `uname` syscall instead of running the `uname`
// utility repeatedly.
func getKernelInfo() (uname struct{ name, release, version, architecture string }) {
	// var un syscall.Utsname
	// err := syscall.Uname(&un)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// uname.architecture = convertFromCString(un.Machine)
	// uname.release = convertFromCString(un.Release)
	// uname.version = convertFromCString(un.Version)
	// uname.name = convertFromCString(un.Sysname)

	return uname
}

func convertFromCString(str [65]int8) string {
	b := make([]byte, 0, 65)
	for _, v := range str {
		if v == 0 {
			break
		}
		b = append(b, byte(v))
	}
	return string(b)
}
