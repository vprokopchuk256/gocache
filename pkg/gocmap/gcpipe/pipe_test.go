package gcpipe_test

// func TestPipeSuccess(t *testing.T) {
// 	m := gcmap.New()
// 	s := &sockettest.Socket{ReadData: "ins key 10"}
// 	defer s.Close()

// 	gcpipe.Pipe(m, s)

// 	log := <-s.input

// 	// 	if log != "key := (integer) 10" {
// 	// 		t.Errorf("expected log to be correct, got: %v", log)
// 	// 	}
// }

// func TestPipeParsingError(t *testing.T) {
// 	m := gcmap.New()
// 	s := newSocket("some undefined command")
// 	defer s.close()

// 	gcpipe.Pipe(m, s)

// 	err := <-s.errors

// 	if err == nil {
// 		t.Errorf("expected error to be returned")
// 	}
// }

// func TestPipeOperationError(t *testing.T) {
// 	m := gcmap.New()
// 	s := newSocket("inc key")
// 	defer s.close()

// 	gcpipe.Pipe(m, s)

// 	err := <-s.errors

// 	if err == nil {
// 		t.Errorf("expected error to be returned")
// 	}
// }
