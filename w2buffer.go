package w2buffer

type WriterToBuffer struct {
	buffer []byte
}

func (w *WriterToBuffer) Write(p []byte) (n int, err error) {
	if len(p) == 0 {
		return 0, nil
	}
	w.buffer = append(w.buffer, p...)
	return len(p), nil
}

func (w *WriterToBuffer) GetBuffer() []byte {
	return w.buffer
}

func (w *WriterToBuffer) IsEmpty() bool {
	return len(w.buffer) > 0
}
