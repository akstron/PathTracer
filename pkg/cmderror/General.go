package cmderror

type SomethingWentWrong struct{}

func (err *SomethingWentWrong) Error() string {
	return "Something went wrong. Please try again."
}
