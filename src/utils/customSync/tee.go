package customSync

func Tee[T any](in <-chan T) (_, _ <-chan T) {

	left := make(chan T)
	right := make(chan T)

	go func() {
		defer func() {
			close(left)
			close(right)
		}()
		for val := range in {
			leftScope, rightScope := right, left
			for range 2 {
				select {
				case leftScope <- val:
					leftScope = nil
				case rightScope <- val:
					rightScope = nil
				}
			}
		}
	}()

	return left, right
}
