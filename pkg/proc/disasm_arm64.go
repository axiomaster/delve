package proc

import (
	"golang.org/x/arch/arm64/arm64asm"
)

var maxInstructionLength uint64 = 4

type archInst struct {
	arm64asm.Inst
	Len int
}

// lism
func asmDecode(mem []byte, pc uint64) (*archInst, error) {

	return nil, nil
}

func (inst *archInst) Size() int {
	return inst.Len
}

func (inst *AsmInstruction) Text(flavour AssemblyFlavour, bi *BinaryInfo) string {
	return "arm64 instrument"
}

// IsCall returns true if the instruction is a CALL or LCALL instruction.
func (inst *AsmInstruction) IsCall() bool {
	if inst.Inst == nil {
		return false
	}
	return false
}

// IsRet returns true if the instruction is a RET or LRET instruction.
func (inst *AsmInstruction) IsRet() bool {
	return false
}

// lism
func resolveCallArg(inst *archInst, currentGoroutine bool, regs Registers, mem MemoryReadWriter, bininfo *BinaryInfo) *Location {
	return nil
}

// firstPCAfterPrologueDisassembly returns the address of the first
// instruction after the prologue for function fn by disassembling fn and
// matching the instructions against known split-stack prologue patterns.
// If sameline is set firstPCAfterPrologueDisassembly will always return an
// address associated with the same line as fn.Entry
func firstPCAfterPrologueDisassembly(p Process, fn *Function, sameline bool) (uint64, error) {
	return fn.Entry, nil
}