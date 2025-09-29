package utils

type HeaderWriter interface {
	WriteHeader(int)
}

func WriteException(hw HeaderWriter) {
	hw.WriteHeader(500)
}
