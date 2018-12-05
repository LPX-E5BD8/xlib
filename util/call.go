/*
Copyright 2018 liipx(lipengxiang)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import "runtime"

// Caller will return the function caller
func Caller() string {
	// get the callers as []uintptr
	uintPointerList := make([]uintptr, 1)

	// skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(3, uintPointerList)
	if n == 0 {
		return "n/a"
	}

	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(uintPointerList[0] - 1)
	if fun == nil {
		return "n/a"
	}

	// return caller name
	return fun.Name()
}
