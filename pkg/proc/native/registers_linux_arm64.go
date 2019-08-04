package native

import "github.com/go-delve/delve/pkg/proc"

// SetPC sets RIP to the value specified by 'pc'.
func (thread *Thread) SetPC(pc uint64) error {
	return nil
}

func (thread *Thread) SetSP(sp uint64) (err error) {
	return nil
}

func (thread *Thread) SetDX(dx uint64) (err error) {
	return nil
}

func (t *Thread) Registers(floatingPoint bool) (proc.Registers, error) {
	return nil, nil
}