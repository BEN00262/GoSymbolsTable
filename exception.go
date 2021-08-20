package main

type Exception struct {
	message string
}

func ThrowException(message string) *Exception {
	return &Exception{
		message: message,
	}
}

func (m *Exception) Error() string {
	return m.message
}
