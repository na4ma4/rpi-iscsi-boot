package namegen

const maxSerialLength = 8

func Truncate(serial string) string {
	name := serial
	if len(name) > maxSerialLength {
		name = name[len(name)-8:]
	}

	return name
}
