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
			left, right := right, left
			for range 2 {
				select {
				case left <- val:
					left = nil
				case right <- val:
					right = nil
				}
			}
		}
	}()

	return left, right
}
