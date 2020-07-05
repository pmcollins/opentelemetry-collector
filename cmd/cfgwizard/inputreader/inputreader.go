package inputreader

import (
	"bufio"
	"io"
	"strconv"
)

type InputReader struct {
	bufioReader *bufio.Reader
}

func New(r io.Reader) InputReader {
	return InputReader{
		bufioReader: bufio.NewReader(r),
	}
}

func (r InputReader) ReadInt() *int {
	str := r.ReadString()
	if str == "" {
		return nil
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return nil
	}
	return &num
}

func (r InputReader) ReadStringDefault(defaultVal string) string {
	str := r.ReadString()
	if str == "" {
		return defaultVal
	}
	return str
}

func (r InputReader) ReadString() string {
	str, err := r.bufioReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return str[:len(str)-1]
}
