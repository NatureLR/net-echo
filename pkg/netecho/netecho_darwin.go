//go:build !windows && !linux

package netecho

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "不支持drawin")
}
