package core

type CoreValidator struct {
	Bool   BoolValidator
	Date   DateValidator
	Int    IntValidator
	Slice  SliceValidator
	String StringValidator
}

func NewCore() CoreValidator {
	return CoreValidator{
		Bool:   NewBool(),
		Date:   NewDate(),
		Int:    NewInt(),
		Slice:  NewSlice(),
		String: NewString(),
	}
}
