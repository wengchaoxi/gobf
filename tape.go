package framework

var (
	WRITE      byte = ','
	READ       byte = '.'
	INC        byte = '+'
	DEC        byte = '-'
	INC_PTR    byte = '>'
	DEC_PTR    byte = '<'
	LOOP_BEGIN byte = '['
	LOOP_END   byte = ']'
)

type Tape struct {
	Data  []byte
	Index uint
}

func NewTape(size uint) *Tape {
	return &Tape{
		Data:  make([]byte, size),
		Index: 0,
	}
}

func (t *Tape) IncPtr() {
	t.Index++
}

func (t *Tape) DecPtr() {
	t.Index--
}

func (t *Tape) Inc() {
	t.Data[t.Index]++
}

func (t *Tape) Dec() {
	t.Data[t.Index]--
}

func (t *Tape) Set(v byte) {
	t.Data[t.Index] = v
}

func (t *Tape) Get() byte {
	return t.Data[t.Index]
}
