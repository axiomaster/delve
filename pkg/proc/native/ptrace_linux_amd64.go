package native

import (
	"github.com/go-delve/delve/pkg/proc/linutil"
	sys "golang.org/x/sys/unix"
	"syscall"
	"unsafe"
)

// PtraceGetRegset returns floating point registers of the specified thread
// using PTRACE.
// See amd64_linux_fetch_inferior_registers in gdb/amd64-linux-nat.c.html
// and amd64_supply_xsave in gdb/amd64-tdep.c.html
// and Section 13.1 (and following) of Intel® 64 and IA-32 Architectures Software Developer’s Manual, Volume 1: Basic Architecture
func PtraceGetRegset(tid int) (regset linutil.AMD64Xstate, err error) {
	_, _, err = syscall.Syscall6(syscall.SYS_PTRACE, sys.PTRACE_GETFPREGS, uintptr(tid), uintptr(0), uintptr(unsafe.Pointer(&regset.AMD64PtraceFpRegs)), 0, 0)
	if err == syscall.Errno(0) || err == syscall.ENODEV {
		// ignore ENODEV, it just means this CPU doesn't have X87 registers (??)
		err = nil
	}

	var xstateargs [_X86_XSTATE_MAX_SIZE]byte
	iov := sys.Iovec{Base: &xstateargs[0], Len: _X86_XSTATE_MAX_SIZE}
	_, _, err = syscall.Syscall6(syscall.SYS_PTRACE, sys.PTRACE_GETREGSET, uintptr(tid), _NT_X86_XSTATE, uintptr(unsafe.Pointer(&iov)), 0, 0)
	if err != syscall.Errno(0) {
		if err == syscall.ENODEV || err == syscall.EIO {
			// ignore ENODEV, it just means this CPU or kernel doesn't support XSTATE, see https://github.com/go-delve/delve/issues/1022
			// also ignore EIO, it means that we are running on an old kernel (pre 2.6.34) and PTRACE_GETREGSET is not implemented
			err = nil
		}
		return
	} else {
		err = nil
	}

	regset.Xsave = xstateargs[:iov.Len]
	err = linutil.AMD64XstateRead(regset.Xsave, false, &regset)
	return
}
