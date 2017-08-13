package service

/*CaptureTest creates a unit test (appended to file if it already exists),
 * by saving the input and output structs.
 * It's called like:
 *   CaptureTest("filename", "testname", "function", {inputkey1: inputvalue, ...}, {outputkey1: outputvalue, ...})
 */
func CaptureTest(filename, testname, function string, inhash, outhash interface{}) (err error) {
	//check if filename exists
	//if not, create it with standard test boilerplate
	//else open to append
	//iterate over input hash, print out all values
	//iterate over output hash, print out all values
	// make a stub that calls the function and compares output
	return nil
}
