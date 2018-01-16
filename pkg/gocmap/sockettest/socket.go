package sockettest

type Socket struct {
	WriteData string
	ReadData  string
	ErrData   error
	ReadErr   error
	Open      bool
}

func (s *Socket) Read() (string, error) {
	if s.ReadErr != nil {
		return "", s.ReadErr
	}

	return s.ReadData, nil
}

func (s *Socket) Write(w string) error {
	s.WriteData = w

	return nil
}

func (s *Socket) Error(err error) error {
	s.ErrData = err

	return nil
}

func (s *Socket) Close() {
	s.Open = false
}
