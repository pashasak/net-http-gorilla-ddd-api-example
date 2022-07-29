package utils

type InternalError struct {
	HumanReadableExplanation string
	ActualError              error
}

func (e InternalError) Error() string {
	if e.HumanReadableExplanation == "" {
		return e.ActualError.Error()
	} else {
		return e.HumanReadableExplanation
	}
}
