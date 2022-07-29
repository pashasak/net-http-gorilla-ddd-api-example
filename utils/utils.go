package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func AppendUniqueInts(a, b []int) []int {
	for i := range b {
		a = AppendUniqueInt(a, b[i])
	}

	return a
}

func AppendUniqueInt(a []int, b int) []int {
	for i := range a {
		if b == a[i] {
			return a
		}
	}

	return append(a, b)
}

func ReplaceNumbersWithZero(s string) string {
	out := make([]rune, len(s))

	i, added := 0, false
	for _, r := range s {
		if r >= '0' && r <= '9' {
			if added {
				continue
			}
			added, out[i] = true, '0'
		} else {
			added, out[i] = false, r
		}
		i++
	}
	return string(out[:i])
}

func JsonPresenter(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(v)
}

func JsonDecoder(r io.ReadCloser, v interface{}) error {
	defer r.Close()

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, v)
}

func ErrorPresenter(w http.ResponseWriter, err error, status int) {
	AddActualError(w, err)
	http.Error(w, err.Error(), status)
}

func AddActualError(w http.ResponseWriter, err error) {
	if actualError, ok := err.(InternalError); ok {
		if actualError.ActualError != nil {
			w.Header().Set("actual-error", ReplaceNumbersWithZero(actualError.ActualError.Error()))
		}
	}
}
