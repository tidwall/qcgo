//go:build !windows && !slowcgo

#include "textflag.h"

TEXT ·morestack(SB), NOSPLIT|NOFRAME, $0
    JMP runtime·morestack_noctxt(SB)  // add more stack

TEXT ·availstack(SB), NOSPLIT|NOFRAME, $0-8
    MOVQ    SP, AX         // read current stack pointer
    MOVQ    0(g), BX       // read stack lo boundary
    SUBQ    BX, AX         // avail = current - lo
    MOVD    AX, ret+0(FP)  // set avail return value
    RET                    // return

TEXT ·Call(SB), NOSPLIT|NOPTR, $0-16
    MOVD    cfunc+0(FP), AX  // cFunc unsafe.Pointer
    MOVD    arg+8(FP), DI    // arg unsafe.Pointer
    MOVQ    SP, BX           // save SP
    ANDQ    $~15, SP         // align stack to 16 bytes
    CALL    AX               // call function
    MOVQ    BX, SP           // restore SP
    RET                      // return
