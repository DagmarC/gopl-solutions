package eval

import (
	"bytes"
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%v", u.op, u.x)

}

func (b binary) String() string {
	return fmt.Sprintf("%v %c %v", b.x, b.op, b.y)
}

func (c call) String() string {
	var buf bytes.Buffer

	buf.Write([]byte(c.fn))
	buf.Write([]byte("("))
	for i, arg := range c.args {
		if i > 0 {
			buf.Write([]byte(", "))
		}
		buf.Write([]byte(arg.String()))
	}
	buf.Write([]byte(")"))

	return buf.String()
}

//!+7.14 min
func (m minimum) String() string {
	var buf bytes.Buffer

	buf.Write([]byte("min"))
	buf.Write([]byte("("))
	for i, arg := range m.operands {
		if i > 0 {
			buf.Write([]byte(", "))
		}
		buf.Write([]byte(arg.String()))
	}
	buf.Write([]byte(")"))

	return buf.String()
}
//!-7.14 min
