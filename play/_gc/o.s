"".main t=1 size=1275 args=0x0 locals=0x98
	0x0000 00000 (1.go:27)	TEXT	"".main(SB), $152-0
	0x0000 00000 (1.go:27)	MOVQ	TLS, CX
	0x0009 00009 (1.go:27)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:27)	LEAQ	-24(SP), AX
	0x0015 00021 (1.go:27)	CMPQ	AX, 16(CX)
	0x0019 00025 (1.go:27)	JLS	1265
	0x001f 00031 (1.go:27)	SUBQ	$152, SP
	0x0026 00038 (1.go:27)	MOVQ	BP, 144(SP)
	0x002e 00046 (1.go:27)	LEAQ	144(SP), BP
	0x0036 00054 (1.go:27)	FUNCDATA	$0, gclocals·c89758d07e85e5e9784036c1fc126388(SB)
	0x0036 00054 (1.go:27)	FUNCDATA	$1, gclocals·6dbd0f82d210f8c1253a4f13f77840be(SB)
	0x0036 00054 (1.go:28)	MOVQ	$500, (SP)
	0x003e 00062 (1.go:28)	PCDATA	$0, $0
	0x003e 00062 (1.go:28)	CALL	runtime/debug.SetGCPercent(SB)
	0x0043 00067 (1.go:30)	LEAQ	type."".MM(SB), AX
	0x004a 00074 (1.go:30)	MOVQ	AX, (SP)
	0x004e 00078 (1.go:30)	PCDATA	$0, $0
	0x004e 00078 (1.go:30)	CALL	runtime.newobject(SB)
	0x0053 00083 (1.go:30)	MOVQ	8(SP), DI
	0x0058 00088 (1.go:30)	MOVQ	DI, "".&mm+104(SP)
	0x005d 00093 (1.go:30)	MOVL	runtime.writeBarrier(SB), AX
	0x0063 00099 (1.go:30)	TESTL	AX, AX
	0x0065 00101 (1.go:30)	JNE	$0, 1227
	0x006b 00107 (1.go:30)	LEAQ	"".statictmp_10(SB), SI
	0x0072 00114 (1.go:30)	DUFFCOPY	$868
	0x0085 00133 (1.go:31)	LEAQ	type.map[int]"".User(SB), AX
	0x008c 00140 (1.go:31)	MOVQ	AX, (SP)
	0x0090 00144 (1.go:31)	MOVQ	$0, 8(SP)
	0x0099 00153 (1.go:31)	MOVQ	$0, 16(SP)
	0x00a2 00162 (1.go:31)	MOVQ	$0, 24(SP)
	0x00ab 00171 (1.go:31)	PCDATA	$0, $1
	0x00ab 00171 (1.go:31)	CALL	runtime.makemap(SB)
	0x00b0 00176 (1.go:31)	MOVQ	32(SP), AX
	0x00b5 00181 (1.go:31)	MOVL	runtime.writeBarrier(SB), CX
	0x00bb 00187 (1.go:31)	MOVQ	"".&mm+104(SP), DX
	0x00c0 00192 (1.go:31)	LEAQ	24(DX), BX
	0x00c4 00196 (1.go:31)	TESTL	CX, CX
	0x00c6 00198 (1.go:31)	JNE	$0, 1203
	0x00cc 00204 (1.go:31)	MOVQ	AX, 24(DX)
	0x00d0 00208 (1.go:33)	LEAQ	type.int(SB), AX
	0x00d7 00215 (1.go:33)	MOVQ	AX, (SP)
	0x00db 00219 (1.go:33)	PCDATA	$0, $1
	0x00db 00219 (1.go:33)	CALL	runtime.newobject(SB)
	0x00e0 00224 (1.go:33)	MOVQ	8(SP), AX
	0x00e5 00229 (1.go:33)	MOVQ	AX, "".&index+112(SP)
	0x00ea 00234 (1.go:33)	MOVQ	$0, (AX)
	0x00f1 00241 (1.go:34)	LEAQ	type."".User(SB), CX
	0x00f8 00248 (1.go:34)	MOVQ	CX, (SP)
	0x00fc 00252 (1.go:34)	PCDATA	$0, $2
	0x00fc 00252 (1.go:34)	CALL	runtime.newobject(SB)
	0x0101 00257 (1.go:34)	MOVQ	8(SP), AX
	0x0106 00262 (1.go:34)	MOVQ	AX, "".&deUser+136(SP)
	0x010e 00270 (1.go:34)	MOVQ	"".statictmp_11+16(SB), CX
	0x0115 00277 (1.go:34)	MOVQ	"".statictmp_11+8(SB), DX
	0x011c 00284 (1.go:34)	MOVQ	"".statictmp_11(SB), BX
	0x0123 00291 (1.go:34)	MOVQ	"".statictmp_11+24(SB), SI
	0x012a 00298 (1.go:34)	MOVQ	SI, ""..autotmp_21+96(SP)
	0x012f 00303 (1.go:34)	MOVQ	BX, (AX)
	0x0132 00306 (1.go:34)	MOVQ	CX, 16(AX)
	0x0136 00310 (1.go:34)	MOVL	runtime.writeBarrier(SB), CX
	0x013c 00316 (1.go:34)	LEAQ	8(AX), BX
	0x0140 00320 (1.go:34)	LEAQ	24(AX), DI
	0x0144 00324 (1.go:34)	MOVQ	DI, ""..autotmp_22+88(SP)
	0x0149 00329 (1.go:34)	TESTL	CX, CX
	0x014b 00331 (1.go:34)	JNE	$0, 1147
	0x0151 00337 (1.go:34)	MOVQ	DX, 8(AX)
	0x0155 00341 (1.go:34)	MOVQ	SI, 24(AX)
	0x0159 00345 (1.go:35)	LEAQ	type.*"".User(SB), CX
	0x0160 00352 (1.go:35)	MOVQ	CX, (SP)
	0x0164 00356 (1.go:35)	PCDATA	$0, $4
	0x0164 00356 (1.go:35)	CALL	runtime.newobject(SB)
	0x0169 00361 (1.go:35)	MOVQ	8(SP), AX
	0x016e 00366 (1.go:35)	MOVQ	AX, "".&deUser2+128(SP)
	0x0176 00374 (1.go:34)	LEAQ	type."".User(SB), CX
	0x017d 00381 (1.go:35)	MOVQ	CX, (SP)
	0x0181 00385 (1.go:35)	PCDATA	$0, $5
	0x0181 00385 (1.go:35)	CALL	runtime.newobject(SB)
	0x0186 00390 (1.go:35)	MOVQ	8(SP), AX
	0x018b 00395 (1.go:35)	MOVQ	$0, (AX)
	0x0192 00402 (1.go:35)	MOVQ	$0, 8(AX)
	0x019a 00410 (1.go:35)	MOVQ	$0, 16(AX)
	0x01a2 00418 (1.go:35)	MOVQ	$0, 24(AX)
	0x01aa 00426 (1.go:35)	MOVL	runtime.writeBarrier(SB), CX
	0x01b0 00432 (1.go:35)	TESTL	CX, CX
	0x01b2 00434 (1.go:35)	JNE	$0, 1112
	0x01b8 00440 (1.go:35)	MOVQ	"".&deUser2+128(SP), CX
	0x01c0 00448 (1.go:35)	MOVQ	AX, (CX)
	0x01c3 00451 (1.go:36)	MOVQ	(CX), AX
	0x01c6 00454 (1.go:36)	MOVL	runtime.writeBarrier(SB), DX
	0x01cc 00460 (1.go:36)	TESTL	DX, DX
	0x01ce 00462 (1.go:36)	JNE	$0, 1072
	0x01d4 00468 (1.go:36)	MOVQ	"".&deUser+136(SP), DX
	0x01dc 00476 (1.go:36)	MOVQ	AX, 24(DX)
	0x01e0 00480 (1.go:39)	LEAQ	type.struct { F uintptr; "".mm *"".MM }(SB), AX
	0x01e7 00487 (1.go:39)	MOVQ	AX, (SP)
	0x01eb 00491 (1.go:39)	PCDATA	$0, $6
	0x01eb 00491 (1.go:39)	CALL	runtime.newobject(SB)
	0x01f0 00496 (1.go:39)	MOVQ	8(SP), AX
	0x01f5 00501 (1.go:39)	MOVQ	AX, ""..autotmp_23+80(SP)
	0x01fa 00506 (1.go:39)	LEAQ	"".main.func1(SB), CX
	0x0201 00513 (1.go:39)	MOVQ	CX, (AX)
	0x0204 00516 (1.go:39)	TESTB	AL, (AX)
	0x0206 00518 (1.go:39)	MOVL	runtime.writeBarrier(SB), CX
	0x020c 00524 (1.go:39)	LEAQ	8(AX), DX
	0x0210 00528 (1.go:39)	TESTL	CX, CX
	0x0212 00530 (1.go:39)	JNE	$0, 1038
	0x0218 00536 (1.go:39)	MOVQ	"".&mm+104(SP), CX
	0x021d 00541 (1.go:39)	MOVQ	CX, 8(AX)
	0x0221 00545 (1.go:46)	LEAQ	type.chan bool(SB), DX
	0x0228 00552 (1.go:46)	MOVQ	DX, (SP)
	0x022c 00556 (1.go:46)	MOVQ	$0, 8(SP)
	0x0235 00565 (1.go:46)	PCDATA	$0, $7
	0x0235 00565 (1.go:46)	CALL	runtime.makechan(SB)
	0x023a 00570 (1.go:46)	MOVQ	16(SP), AX
	0x023f 00575 (1.go:46)	MOVQ	AX, "".doned+72(SP)
	0x0244 00580 (1.go:79)	MOVQ	"".&deUser2+128(SP), CX
	0x024c 00588 (1.go:79)	MOVQ	CX, 56(SP)
	0x0251 00593 (1.go:79)	MOVQ	$100000, 16(SP)
	0x025a 00602 (1.go:79)	MOVQ	"".&index+112(SP), CX
	0x025f 00607 (1.go:79)	MOVQ	CX, 24(SP)
	0x0264 00612 (1.go:79)	MOVQ	"".&mm+104(SP), DX
	0x0269 00617 (1.go:79)	MOVQ	DX, 32(SP)
	0x026e 00622 (1.go:79)	MOVQ	"".&deUser+136(SP), BX
	0x0276 00630 (1.go:79)	MOVQ	BX, 40(SP)
	0x027b 00635 (1.go:79)	MOVQ	""..autotmp_23+80(SP), SI
	0x0280 00640 (1.go:79)	MOVQ	SI, 48(SP)
	0x0285 00645 (1.go:79)	MOVL	$48, (SP)
	0x028c 00652 (1.go:79)	LEAQ	"".main.func2·f(SB), DI
	0x0293 00659 (1.go:79)	MOVQ	DI, 8(SP)
	0x0298 00664 (1.go:79)	PCDATA	$0, $8
	0x0298 00664 (1.go:79)	CALL	runtime.newproc(SB)
	0x029d 00669 (1.go:92)	MOVQ	"".&deUser+136(SP), AX
	0x02a5 00677 (1.go:92)	MOVQ	AX, 40(SP)
	0x02aa 00682 (1.go:92)	MOVQ	"".doned+72(SP), AX
	0x02af 00687 (1.go:92)	MOVQ	AX, 16(SP)
	0x02b4 00692 (1.go:92)	MOVQ	"".&index+112(SP), AX
	0x02b9 00697 (1.go:92)	MOVQ	AX, 24(SP)
	0x02be 00702 (1.go:92)	MOVQ	"".&mm+104(SP), CX
	0x02c3 00707 (1.go:92)	MOVQ	CX, 32(SP)
	0x02c8 00712 (1.go:92)	MOVL	$32, (SP)
	0x02cf 00719 (1.go:92)	LEAQ	"".main.func3·f(SB), DX
	0x02d6 00726 (1.go:92)	MOVQ	DX, 8(SP)
	0x02db 00731 (1.go:92)	PCDATA	$0, $9
	0x02db 00731 (1.go:92)	CALL	runtime.newproc(SB)
	0x02e0 00736 (1.go:105)	MOVQ	"".&mm+104(SP), AX
	0x02e5 00741 (1.go:105)	MOVQ	AX, 24(SP)
	0x02ea 00746 (1.go:105)	MOVQ	"".&index+112(SP), AX
	0x02ef 00751 (1.go:105)	MOVQ	AX, 16(SP)
	0x02f4 00756 (1.go:105)	MOVL	$16, (SP)
	0x02fb 00763 (1.go:105)	LEAQ	"".main.func4·f(SB), AX
	0x0302 00770 (1.go:105)	MOVQ	AX, 8(SP)
	0x0307 00775 (1.go:105)	PCDATA	$0, $10
	0x0307 00775 (1.go:105)	CALL	runtime.newproc(SB)
	0x030c 00780 (1.go:33)	LEAQ	type.int(SB), AX
	0x0313 00787 (1.go:109)	MOVQ	AX, (SP)
	0x0317 00791 (1.go:109)	PCDATA	$0, $10
	0x0317 00791 (1.go:109)	CALL	runtime.newobject(SB)
	0x031c 00796 (1.go:109)	MOVQ	8(SP), AX
	0x0321 00801 (1.go:109)	MOVQ	AX, "".&g+120(SP)
	0x0326 00806 (1.go:109)	MOVQ	$0, (AX)
	0x032d 00813 (1.go:33)	MOVQ	$0, CX
	0x032f 00815 (1.go:110)	MOVQ	CX, "".k+64(SP)
	0x0334 00820 (1.go:110)	CMPQ	CX, $10
	0x0338 00824 (1.go:110)	JGE	$0, 893
	0x033a 00826 (1.go:122)	MOVQ	AX, 16(SP)
	0x033f 00831 (1.go:122)	MOVL	$8, (SP)
	0x0346 00838 (1.go:122)	LEAQ	"".main.func5·f(SB), DX
	0x034d 00845 (1.go:122)	MOVQ	DX, 8(SP)
	0x0352 00850 (1.go:122)	PCDATA	$0, $11
	0x0352 00850 (1.go:122)	CALL	runtime.newproc(SB)
	0x0357 00855 (1.go:123)	MOVQ	$40000, (SP)
	0x035f 00863 (1.go:123)	PCDATA	$0, $11
	0x035f 00863 (1.go:123)	CALL	time.Sleep(SB)
	0x0364 00868 (1.go:110)	MOVQ	"".k+64(SP), AX
	0x0369 00873 (1.go:110)	LEAQ	1(AX), CX
	0x036d 00877 (1.go:122)	MOVQ	"".&g+120(SP), AX
	0x0372 00882 (1.go:110)	MOVQ	CX, "".k+64(SP)
	0x0377 00887 (1.go:110)	CMPQ	CX, $10
	0x037b 00891 (1.go:110)	JLT	$0, 826
	0x037d 00893 (1.go:133)	MOVQ	$100000, 16(SP)
	0x0386 00902 (1.go:133)	MOVL	$8, (SP)
	0x038d 00909 (1.go:133)	LEAQ	"".main.func6·f(SB), CX
	0x0394 00916 (1.go:133)	MOVQ	CX, 8(SP)
	0x0399 00921 (1.go:133)	PCDATA	$0, $11
	0x0399 00921 (1.go:133)	CALL	runtime.newproc(SB)
	0x039e 00926 (1.go:33)	LEAQ	type.int(SB), AX
	0x03a5 00933 (1.go:137)	MOVQ	AX, (SP)
	0x03a9 00937 (1.go:137)	PCDATA	$0, $11
	0x03a9 00937 (1.go:137)	CALL	runtime.newobject(SB)
	0x03ae 00942 (1.go:137)	MOVQ	8(SP), AX
	0x03b3 00947 (1.go:137)	MOVQ	$0, (AX)
	0x03ba 00954 (1.go:148)	MOVQ	"".&g+120(SP), CX
	0x03bf 00959 (1.go:148)	MOVQ	CX, 24(SP)
	0x03c4 00964 (1.go:148)	MOVQ	AX, 16(SP)
	0x03c9 00969 (1.go:148)	MOVQ	""..autotmp_23+80(SP), AX
	0x03ce 00974 (1.go:148)	MOVQ	AX, 32(SP)
	0x03d3 00979 (1.go:148)	MOVL	$24, (SP)
	0x03da 00986 (1.go:148)	LEAQ	"".main.func7·f(SB), AX
	0x03e1 00993 (1.go:148)	MOVQ	AX, 8(SP)
	0x03e6 00998 (1.go:148)	PCDATA	$0, $0
	0x03e6 00998 (1.go:148)	CALL	runtime.newproc(SB)
	0x03eb 01003 (1.go:149)	MOVQ	$1000000000000, AX
	0x03f5 01013 (1.go:149)	MOVQ	AX, (SP)
	0x03f9 01017 (1.go:149)	PCDATA	$0, $0
	0x03f9 01017 (1.go:149)	CALL	time.Sleep(SB)
	0x03fe 01022 (1.go:150)	MOVQ	144(SP), BP
	0x0406 01030 (1.go:150)	ADDQ	$152, SP
	0x040d 01037 (1.go:150)	RET
	0x040e 01038 (1.go:39)	MOVQ	DX, (SP)
	0x0412 01042 (1.go:39)	MOVQ	"".&mm+104(SP), CX
	0x0417 01047 (1.go:39)	MOVQ	CX, 8(SP)
	0x041c 01052 (1.go:39)	PCDATA	$0, $7
	0x041c 01052 (1.go:39)	CALL	runtime.writebarrierptr(SB)
	0x0421 01057 (1.go:79)	MOVQ	""..autotmp_23+80(SP), AX
	0x0426 01062 (1.go:79)	MOVQ	"".&mm+104(SP), CX
	0x042b 01067 (1.go:46)	JMP	545
	0x0430 01072 (1.go:36)	MOVQ	""..autotmp_22+88(SP), DX
	0x0435 01077 (1.go:36)	MOVQ	DX, (SP)
	0x0439 01081 (1.go:36)	MOVQ	AX, 8(SP)
	0x043e 01086 (1.go:36)	PCDATA	$0, $6
	0x043e 01086 (1.go:36)	CALL	runtime.writebarrierptr(SB)
	0x0443 01091 (1.go:79)	MOVQ	"".&deUser2+128(SP), CX
	0x044b 01099 (1.go:79)	MOVQ	"".&deUser+136(SP), DX
	0x0453 01107 (1.go:39)	JMP	480
	0x0458 01112 (1.go:35)	MOVQ	"".&deUser2+128(SP), CX
	0x0460 01120 (1.go:35)	MOVQ	CX, (SP)
	0x0464 01124 (1.go:35)	MOVQ	AX, 8(SP)
	0x0469 01129 (1.go:35)	PCDATA	$0, $5
	0x0469 01129 (1.go:35)	CALL	runtime.writebarrierptr(SB)
	0x046e 01134 (1.go:36)	MOVQ	"".&deUser2+128(SP), CX
	0x0476 01142 (1.go:36)	JMP	451
	0x047b 01147 (1.go:34)	MOVQ	BX, (SP)
	0x047f 01151 (1.go:34)	MOVQ	DX, 8(SP)
	0x0484 01156 (1.go:34)	PCDATA	$0, $3
	0x0484 01156 (1.go:34)	CALL	runtime.writebarrierptr(SB)
	0x0489 01161 (1.go:34)	MOVQ	""..autotmp_22+88(SP), AX
	0x048e 01166 (1.go:34)	MOVQ	AX, (SP)
	0x0492 01170 (1.go:34)	MOVQ	""..autotmp_21+96(SP), CX
	0x0497 01175 (1.go:34)	MOVQ	CX, 8(SP)
	0x049c 01180 (1.go:34)	PCDATA	$0, $4
	0x049c 01180 (1.go:34)	CALL	runtime.writebarrierptr(SB)
	0x04a1 01185 (1.go:36)	MOVQ	"".&deUser+136(SP), AX
	0x04a9 01193 (1.go:36)	MOVQ	""..autotmp_22+88(SP), DI
	0x04ae 01198 (1.go:35)	JMP	345
	0x04b3 01203 (1.go:31)	MOVQ	BX, (SP)
	0x04b7 01207 (1.go:31)	MOVQ	AX, 8(SP)
	0x04bc 01212 (1.go:31)	PCDATA	$0, $1
	0x04bc 01212 (1.go:31)	CALL	runtime.writebarrierptr(SB)
	0x04c1 01217 (1.go:39)	MOVQ	"".&mm+104(SP), DX
	0x04c6 01222 (1.go:33)	JMP	208
	0x04cb 01227 (1.go:30)	LEAQ	type."".MM(SB), AX
	0x04d2 01234 (1.go:30)	MOVQ	AX, (SP)
	0x04d6 01238 (1.go:30)	MOVQ	DI, 8(SP)
	0x04db 01243 (1.go:30)	LEAQ	"".statictmp_10(SB), AX
	0x04e2 01250 (1.go:30)	MOVQ	AX, 16(SP)
	0x04e7 01255 (1.go:30)	PCDATA	$0, $1
	0x04e7 01255 (1.go:30)	CALL	runtime.typedmemmove(SB)
	0x04ec 01260 (1.go:31)	JMP	133
	0x04f1 01265 (1.go:31)	NOP
	0x04f1 01265 (1.go:27)	PCDATA	$0, $-1
	0x04f1 01265 (1.go:27)	CALL	runtime.morestack_noctxt(SB)
	0x04f6 01270 (1.go:27)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 8d 44 24 e8 48 3b 41 10 0f 86 d2 04 00 00 48  H.D$.H;A.......H
	0x0020 81 ec 98 00 00 00 48 89 ac 24 90 00 00 00 48 8d  ......H..$....H.
	0x0030 ac 24 90 00 00 00 48 c7 04 24 f4 01 00 00 e8 00  .$....H..$......
	0x0040 00 00 00 48 8d 05 00 00 00 00 48 89 04 24 e8 00  ...H......H..$..
	0x0050 00 00 00 48 8b 7c 24 08 48 89 7c 24 68 8b 05 00  ...H.|$.H.|$h...
	0x0060 00 00 00 85 c0 0f 85 60 04 00 00 48 8d 35 00 00  .......`...H.5..
	0x0070 00 00 48 89 6c 24 f0 48 8d 6c 24 f0 e8 00 00 00  ..H.l$.H.l$.....
	0x0080 00 48 8b 6d 00 48 8d 05 00 00 00 00 48 89 04 24  .H.m.H......H..$
	0x0090 48 c7 44 24 08 00 00 00 00 48 c7 44 24 10 00 00  H.D$.....H.D$...
	0x00a0 00 00 48 c7 44 24 18 00 00 00 00 e8 00 00 00 00  ..H.D$..........
	0x00b0 48 8b 44 24 20 8b 0d 00 00 00 00 48 8b 54 24 68  H.D$ ......H.T$h
	0x00c0 48 8d 5a 18 85 c9 0f 85 e7 03 00 00 48 89 42 18  H.Z.........H.B.
	0x00d0 48 8d 05 00 00 00 00 48 89 04 24 e8 00 00 00 00  H......H..$.....
	0x00e0 48 8b 44 24 08 48 89 44 24 70 48 c7 00 00 00 00  H.D$.H.D$pH.....
	0x00f0 00 48 8d 0d 00 00 00 00 48 89 0c 24 e8 00 00 00  .H......H..$....
	0x0100 00 48 8b 44 24 08 48 89 84 24 88 00 00 00 48 8b  .H.D$.H..$....H.
	0x0110 0d 00 00 00 00 48 8b 15 00 00 00 00 48 8b 1d 00  .....H......H...
	0x0120 00 00 00 48 8b 35 00 00 00 00 48 89 74 24 60 48  ...H.5....H.t$`H
	0x0130 89 18 48 89 48 10 8b 0d 00 00 00 00 48 8d 58 08  ..H.H.......H.X.
	0x0140 48 8d 78 18 48 89 7c 24 58 85 c9 0f 85 2a 03 00  H.x.H.|$X....*..
	0x0150 00 48 89 50 08 48 89 70 18 48 8d 0d 00 00 00 00  .H.P.H.p.H......
	0x0160 48 89 0c 24 e8 00 00 00 00 48 8b 44 24 08 48 89  H..$.....H.D$.H.
	0x0170 84 24 80 00 00 00 48 8d 0d 00 00 00 00 48 89 0c  .$....H......H..
	0x0180 24 e8 00 00 00 00 48 8b 44 24 08 48 c7 00 00 00  $.....H.D$.H....
	0x0190 00 00 48 c7 40 08 00 00 00 00 48 c7 40 10 00 00  ..H.@.....H.@...
	0x01a0 00 00 48 c7 40 18 00 00 00 00 8b 0d 00 00 00 00  ..H.@...........
	0x01b0 85 c9 0f 85 a0 02 00 00 48 8b 8c 24 80 00 00 00  ........H..$....
	0x01c0 48 89 01 48 8b 01 8b 15 00 00 00 00 85 d2 0f 85  H..H............
	0x01d0 5c 02 00 00 48 8b 94 24 88 00 00 00 48 89 42 18  \...H..$....H.B.
	0x01e0 48 8d 05 00 00 00 00 48 89 04 24 e8 00 00 00 00  H......H..$.....
	0x01f0 48 8b 44 24 08 48 89 44 24 50 48 8d 0d 00 00 00  H.D$.H.D$PH.....
	0x0200 00 48 89 08 84 00 8b 0d 00 00 00 00 48 8d 50 08  .H..........H.P.
	0x0210 85 c9 0f 85 f6 01 00 00 48 8b 4c 24 68 48 89 48  ........H.L$hH.H
	0x0220 08 48 8d 15 00 00 00 00 48 89 14 24 48 c7 44 24  .H......H..$H.D$
	0x0230 08 00 00 00 00 e8 00 00 00 00 48 8b 44 24 10 48  ..........H.D$.H
	0x0240 89 44 24 48 48 8b 8c 24 80 00 00 00 48 89 4c 24  .D$HH..$....H.L$
	0x0250 38 48 c7 44 24 10 a0 86 01 00 48 8b 4c 24 70 48  8H.D$.....H.L$pH
	0x0260 89 4c 24 18 48 8b 54 24 68 48 89 54 24 20 48 8b  .L$.H.T$hH.T$ H.
	0x0270 9c 24 88 00 00 00 48 89 5c 24 28 48 8b 74 24 50  .$....H.\$(H.t$P
	0x0280 48 89 74 24 30 c7 04 24 30 00 00 00 48 8d 3d 00  H.t$0..$0...H.=.
	0x0290 00 00 00 48 89 7c 24 08 e8 00 00 00 00 48 8b 84  ...H.|$......H..
	0x02a0 24 88 00 00 00 48 89 44 24 28 48 8b 44 24 48 48  $....H.D$(H.D$HH
	0x02b0 89 44 24 10 48 8b 44 24 70 48 89 44 24 18 48 8b  .D$.H.D$pH.D$.H.
	0x02c0 4c 24 68 48 89 4c 24 20 c7 04 24 20 00 00 00 48  L$hH.L$ ..$ ...H
	0x02d0 8d 15 00 00 00 00 48 89 54 24 08 e8 00 00 00 00  ......H.T$......
	0x02e0 48 8b 44 24 68 48 89 44 24 18 48 8b 44 24 70 48  H.D$hH.D$.H.D$pH
	0x02f0 89 44 24 10 c7 04 24 10 00 00 00 48 8d 05 00 00  .D$...$....H....
	0x0300 00 00 48 89 44 24 08 e8 00 00 00 00 48 8d 05 00  ..H.D$......H...
	0x0310 00 00 00 48 89 04 24 e8 00 00 00 00 48 8b 44 24  ...H..$.....H.D$
	0x0320 08 48 89 44 24 78 48 c7 00 00 00 00 00 31 c9 48  .H.D$xH......1.H
	0x0330 89 4c 24 40 48 83 f9 0a 7d 43 48 89 44 24 10 c7  .L$@H...}CH.D$..
	0x0340 04 24 08 00 00 00 48 8d 15 00 00 00 00 48 89 54  .$....H......H.T
	0x0350 24 08 e8 00 00 00 00 48 c7 04 24 40 9c 00 00 e8  $......H..$@....
	0x0360 00 00 00 00 48 8b 44 24 40 48 8d 48 01 48 8b 44  ....H.D$@H.H.H.D
	0x0370 24 78 48 89 4c 24 40 48 83 f9 0a 7c bd 48 c7 44  $xH.L$@H...|.H.D
	0x0380 24 10 a0 86 01 00 c7 04 24 08 00 00 00 48 8d 0d  $.......$....H..
	0x0390 00 00 00 00 48 89 4c 24 08 e8 00 00 00 00 48 8d  ....H.L$......H.
	0x03a0 05 00 00 00 00 48 89 04 24 e8 00 00 00 00 48 8b  .....H..$.....H.
	0x03b0 44 24 08 48 c7 00 00 00 00 00 48 8b 4c 24 78 48  D$.H......H.L$xH
	0x03c0 89 4c 24 18 48 89 44 24 10 48 8b 44 24 50 48 89  .L$.H.D$.H.D$PH.
	0x03d0 44 24 20 c7 04 24 18 00 00 00 48 8d 05 00 00 00  D$ ..$....H.....
	0x03e0 00 48 89 44 24 08 e8 00 00 00 00 48 b8 00 10 a5  .H.D$......H....
	0x03f0 d4 e8 00 00 00 48 89 04 24 e8 00 00 00 00 48 8b  .....H..$.....H.
	0x0400 ac 24 90 00 00 00 48 81 c4 98 00 00 00 c3 48 89  .$....H.......H.
	0x0410 14 24 48 8b 4c 24 68 48 89 4c 24 08 e8 00 00 00  .$H.L$hH.L$.....
	0x0420 00 48 8b 44 24 50 48 8b 4c 24 68 e9 f1 fd ff ff  .H.D$PH.L$h.....
	0x0430 48 8b 54 24 58 48 89 14 24 48 89 44 24 08 e8 00  H.T$XH..$H.D$...
	0x0440 00 00 00 48 8b 8c 24 80 00 00 00 48 8b 94 24 88  ...H..$....H..$.
	0x0450 00 00 00 e9 88 fd ff ff 48 8b 8c 24 80 00 00 00  ........H..$....
	0x0460 48 89 0c 24 48 89 44 24 08 e8 00 00 00 00 48 8b  H..$H.D$......H.
	0x0470 8c 24 80 00 00 00 e9 48 fd ff ff 48 89 1c 24 48  .$.....H...H..$H
	0x0480 89 54 24 08 e8 00 00 00 00 48 8b 44 24 58 48 89  .T$......H.D$XH.
	0x0490 04 24 48 8b 4c 24 60 48 89 4c 24 08 e8 00 00 00  .$H.L$`H.L$.....
	0x04a0 00 48 8b 84 24 88 00 00 00 48 8b 7c 24 58 e9 a6  .H..$....H.|$X..
	0x04b0 fc ff ff 48 89 1c 24 48 89 44 24 08 e8 00 00 00  ...H..$H.D$.....
	0x04c0 00 48 8b 54 24 68 e9 05 fc ff ff 48 8d 05 00 00  .H.T$h.....H....
	0x04d0 00 00 48 89 04 24 48 89 7c 24 08 48 8d 05 00 00  ..H..$H.|$.H....
	0x04e0 00 00 48 89 44 24 10 e8 00 00 00 00 e9 94 fb ff  ..H.D$..........
	0x04f0 ff e8 00 00 00 00 e9 05 fb ff ff                 ...........
	rel 12+4 t=16 TLS+0
	rel 63+4 t=8 runtime/debug.SetGCPercent+0
	rel 70+4 t=15 type."".MM+0
	rel 79+4 t=8 runtime.newobject+0
	rel 95+4 t=15 runtime.writeBarrier+0
	rel 110+4 t=15 "".statictmp_10+0
	rel 125+4 t=8 runtime.duffcopy+868
	rel 136+4 t=15 type.map[int]"".User+0
	rel 172+4 t=8 runtime.makemap+0
	rel 183+4 t=15 runtime.writeBarrier+0
	rel 211+4 t=15 type.int+0
	rel 220+4 t=8 runtime.newobject+0
	rel 244+4 t=15 type."".User+0
	rel 253+4 t=8 runtime.newobject+0
	rel 273+4 t=15 "".statictmp_11+16
	rel 280+4 t=15 "".statictmp_11+8
	rel 287+4 t=15 "".statictmp_11+0
	rel 294+4 t=15 "".statictmp_11+24
	rel 312+4 t=15 runtime.writeBarrier+0
	rel 348+4 t=15 type.*"".User+0
	rel 357+4 t=8 runtime.newobject+0
	rel 377+4 t=15 type."".User+0
	rel 386+4 t=8 runtime.newobject+0
	rel 428+4 t=15 runtime.writeBarrier+0
	rel 456+4 t=15 runtime.writeBarrier+0
	rel 483+4 t=15 type.struct { F uintptr; "".mm *"".MM }+0
	rel 492+4 t=8 runtime.newobject+0
	rel 509+4 t=15 "".main.func1+0
	rel 520+4 t=15 runtime.writeBarrier+0
	rel 548+4 t=15 type.chan bool+0
	rel 566+4 t=8 runtime.makechan+0
	rel 655+4 t=15 "".main.func2·f+0
	rel 665+4 t=8 runtime.newproc+0
	rel 722+4 t=15 "".main.func3·f+0
	rel 732+4 t=8 runtime.newproc+0
	rel 766+4 t=15 "".main.func4·f+0
	rel 776+4 t=8 runtime.newproc+0
	rel 783+4 t=15 type.int+0
	rel 792+4 t=8 runtime.newobject+0
	rel 841+4 t=15 "".main.func5·f+0
	rel 851+4 t=8 runtime.newproc+0
	rel 864+4 t=8 time.Sleep+0
	rel 912+4 t=15 "".main.func6·f+0
	rel 922+4 t=8 runtime.newproc+0
	rel 929+4 t=15 type.int+0
	rel 938+4 t=8 runtime.newobject+0
	rel 989+4 t=15 "".main.func7·f+0
	rel 999+4 t=8 runtime.newproc+0
	rel 1018+4 t=8 time.Sleep+0
	rel 1053+4 t=8 runtime.writebarrierptr+0
	rel 1087+4 t=8 runtime.writebarrierptr+0
	rel 1130+4 t=8 runtime.writebarrierptr+0
	rel 1157+4 t=8 runtime.writebarrierptr+0
	rel 1181+4 t=8 runtime.writebarrierptr+0
	rel 1213+4 t=8 runtime.writebarrierptr+0
	rel 1230+4 t=15 type."".MM+0
	rel 1246+4 t=15 "".statictmp_10+0
	rel 1256+4 t=8 runtime.typedmemmove+0
	rel 1266+4 t=8 runtime.morestack_noctxt+0
"".init.1 t=1 size=82 args=0x0 locals=0x20
	0x0000 00000 (1.go:154)	TEXT	"".init.1(SB), $32-0
	0x0000 00000 (1.go:154)	MOVQ	TLS, CX
	0x0009 00009 (1.go:154)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:154)	CMPQ	SP, 16(CX)
	0x0014 00020 (1.go:154)	JLS	75
	0x0016 00022 (1.go:154)	SUBQ	$32, SP
	0x001a 00026 (1.go:154)	MOVQ	BP, 24(SP)
	0x001f 00031 (1.go:154)	LEAQ	24(SP), BP
	0x0024 00036 (1.go:154)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (1.go:154)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (1.go:155)	PCDATA	$0, $0
	0x0024 00036 (1.go:155)	CALL	time.Now(SB)
	0x0029 00041 (1.go:155)	MOVQ	(SP), AX
	0x002d 00045 (1.go:155)	MOVQ	$-62135596800, CX
	0x0037 00055 (1.go:155)	ADDQ	CX, AX
	0x003a 00058 (1.go:155)	MOVQ	AX, "".t1(SB)
	0x0041 00065 (1.go:156)	MOVQ	24(SP), BP
	0x0046 00070 (1.go:156)	ADDQ	$32, SP
	0x004a 00074 (1.go:156)	RET
	0x004b 00075 (1.go:156)	NOP
	0x004b 00075 (1.go:154)	PCDATA	$0, $-1
	0x004b 00075 (1.go:154)	CALL	runtime.morestack_noctxt(SB)
	0x0050 00080 (1.go:154)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 76 35 48 83 ec 20 48 89 6c 24 18 48  H;a.v5H.. H.l$.H
	0x0020 8d 6c 24 18 e8 00 00 00 00 48 8b 04 24 48 b9 00  .l$......H..$H..
	0x0030 09 6e 88 f1 ff ff ff 48 01 c8 48 89 05 00 00 00  .n.....H..H.....
	0x0040 00 48 8b 6c 24 18 48 83 c4 20 c3 e8 00 00 00 00  .H.l$.H.. ......
	0x0050 eb ae                                            ..
	rel 12+4 t=16 TLS+0
	rel 37+4 t=8 time.Now+0
	rel 61+4 t=15 "".t1+0
	rel 76+4 t=8 runtime.morestack_noctxt+0
"".printMem t=1 size=4478 args=0x0 locals=0x1a88
	0x0000 00000 (1.go:158)	TEXT	"".printMem(SB), $6792-0
	0x0000 00000 (1.go:158)	MOVQ	TLS, CX
	0x0009 00009 (1.go:158)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:158)	MOVQ	16(CX), SI
	0x0014 00020 (1.go:158)	CMPQ	SI, $-1314
	0x001b 00027 (1.go:158)	JEQ	4468
	0x0021 00033 (1.go:158)	LEAQ	880(SP), AX
	0x0029 00041 (1.go:158)	SUBQ	SI, AX
	0x002c 00044 (1.go:158)	CMPQ	AX, $7544
	0x0032 00050 (1.go:158)	JLS	4468
	0x0038 00056 (1.go:158)	SUBQ	$6792, SP
	0x003f 00063 (1.go:158)	MOVQ	BP, 6784(SP)
	0x0047 00071 (1.go:158)	LEAQ	6784(SP), BP
	0x004f 00079 (1.go:158)	FUNCDATA	$0, gclocals·5ff117978f2c049ee09a6c769745bf8c(SB)
	0x004f 00079 (1.go:158)	FUNCDATA	$1, gclocals·d5cb556d330213bab77ebaeb5fb03bb3(SB)
	0x004f 00079 (1.go:160)	LEAQ	"".mem+96(SP), DI
	0x0054 00084 (1.go:160)	MOVQ	$722, CX
	0x005b 00091 (1.go:160)	MOVQ	$0, AX
	0x005d 00093 (1.go:160)	REP
	0x005e 00094 (1.go:160)	STOSQ
	0x0060 00096 (1.go:160)	LEAQ	"".mem+96(SP), DX
	0x0065 00101 (1.go:161)	MOVQ	DX, (SP)
	0x0069 00105 (1.go:161)	PCDATA	$0, $0
	0x0069 00105 (1.go:161)	CALL	runtime.ReadMemStats(SB)
	0x006e 00110 (1.go:163)	LEAQ	go.string."Alloc: "(SB), DX
	0x0075 00117 (1.go:163)	MOVQ	DX, ""..autotmp_26+6160(SP)
	0x007d 00125 (1.go:163)	MOVQ	$7, ""..autotmp_26+6168(SP)
	0x0089 00137 (1.go:163)	MOVQ	$0, ""..autotmp_25+6752(SP)
	0x0095 00149 (1.go:163)	MOVQ	$0, ""..autotmp_25+6760(SP)
	0x00a1 00161 (1.go:163)	MOVQ	$0, ""..autotmp_25+6768(SP)
	0x00ad 00173 (1.go:163)	MOVQ	$0, ""..autotmp_25+6776(SP)
	0x00b9 00185 (1.go:163)	LEAQ	type.string(SB), DX
	0x00c0 00192 (1.go:163)	MOVQ	DX, (SP)
	0x00c4 00196 (1.go:163)	LEAQ	""..autotmp_26+6160(SP), BX
	0x00cc 00204 (1.go:163)	MOVQ	BX, 8(SP)
	0x00d1 00209 (1.go:163)	PCDATA	$0, $1
	0x00d1 00209 (1.go:163)	CALL	runtime.convT2E(SB)
	0x00d6 00214 (1.go:163)	MOVQ	24(SP), DX
	0x00db 00219 (1.go:163)	MOVQ	16(SP), BX
	0x00e0 00224 (1.go:163)	MOVQ	BX, ""..autotmp_25+6752(SP)
	0x00e8 00232 (1.go:163)	MOVQ	DX, ""..autotmp_25+6760(SP)
	0x00f0 00240 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x00f7 00247 (1.go:163)	MOVQ	DX, (SP)
	0x00fb 00251 (1.go:163)	LEAQ	"".mem+96(SP), BX
	0x0100 00256 (1.go:163)	MOVQ	BX, 8(SP)
	0x0105 00261 (1.go:163)	PCDATA	$0, $1
	0x0105 00261 (1.go:163)	CALL	runtime.convT2E(SB)
	0x010a 00266 (1.go:163)	MOVQ	16(SP), DX
	0x010f 00271 (1.go:163)	MOVQ	24(SP), BX
	0x0114 00276 (1.go:163)	MOVQ	DX, ""..autotmp_25+6768(SP)
	0x011c 00284 (1.go:163)	MOVQ	BX, ""..autotmp_25+6776(SP)
	0x0124 00292 (1.go:163)	LEAQ	""..autotmp_25+6752(SP), DX
	0x012c 00300 (1.go:163)	MOVQ	DX, (SP)
	0x0130 00304 (1.go:163)	MOVQ	$2, 8(SP)
	0x0139 00313 (1.go:163)	MOVQ	$2, 16(SP)
	0x0142 00322 (1.go:163)	PCDATA	$0, $1
	0x0142 00322 (1.go:163)	CALL	fmt.Println(SB)
	0x0147 00327 (1.go:164)	LEAQ	go.string."EnableGC: "(SB), DX
	0x014e 00334 (1.go:164)	MOVQ	DX, ""..autotmp_28+6144(SP)
	0x0156 00342 (1.go:164)	MOVQ	$10, ""..autotmp_28+6152(SP)
	0x0162 00354 (1.go:164)	MOVQ	$0, ""..autotmp_27+6720(SP)
	0x016e 00366 (1.go:164)	MOVQ	$0, ""..autotmp_27+6728(SP)
	0x017a 00378 (1.go:164)	MOVQ	$0, ""..autotmp_27+6736(SP)
	0x0186 00390 (1.go:164)	MOVQ	$0, ""..autotmp_27+6744(SP)
	0x0192 00402 (1.go:163)	LEAQ	type.string(SB), DX
	0x0199 00409 (1.go:164)	MOVQ	DX, (SP)
	0x019d 00413 (1.go:164)	LEAQ	""..autotmp_28+6144(SP), BX
	0x01a5 00421 (1.go:164)	MOVQ	BX, 8(SP)
	0x01aa 00426 (1.go:164)	PCDATA	$0, $2
	0x01aa 00426 (1.go:164)	CALL	runtime.convT2E(SB)
	0x01af 00431 (1.go:164)	MOVQ	24(SP), DX
	0x01b4 00436 (1.go:164)	MOVQ	16(SP), BX
	0x01b9 00441 (1.go:164)	MOVQ	BX, ""..autotmp_27+6720(SP)
	0x01c1 00449 (1.go:164)	MOVQ	DX, ""..autotmp_27+6728(SP)
	0x01c9 00457 (1.go:164)	LEAQ	type.bool(SB), DX
	0x01d0 00464 (1.go:164)	MOVQ	DX, (SP)
	0x01d4 00468 (1.go:164)	LEAQ	"".mem+4400(SP), DX
	0x01dc 00476 (1.go:164)	MOVQ	DX, 8(SP)
	0x01e1 00481 (1.go:164)	PCDATA	$0, $2
	0x01e1 00481 (1.go:164)	CALL	runtime.convT2E(SB)
	0x01e6 00486 (1.go:164)	MOVQ	16(SP), DX
	0x01eb 00491 (1.go:164)	MOVQ	24(SP), BX
	0x01f0 00496 (1.go:164)	MOVQ	DX, ""..autotmp_27+6736(SP)
	0x01f8 00504 (1.go:164)	MOVQ	BX, ""..autotmp_27+6744(SP)
	0x0200 00512 (1.go:164)	LEAQ	""..autotmp_27+6720(SP), DX
	0x0208 00520 (1.go:164)	MOVQ	DX, (SP)
	0x020c 00524 (1.go:164)	MOVQ	$2, 8(SP)
	0x0215 00533 (1.go:164)	MOVQ	$2, 16(SP)
	0x021e 00542 (1.go:164)	PCDATA	$0, $2
	0x021e 00542 (1.go:164)	CALL	fmt.Println(SB)
	0x0223 00547 (1.go:165)	LEAQ	go.string."Frees: "(SB), DX
	0x022a 00554 (1.go:165)	MOVQ	DX, ""..autotmp_30+6128(SP)
	0x0232 00562 (1.go:165)	MOVQ	$7, ""..autotmp_30+6136(SP)
	0x023e 00574 (1.go:165)	MOVQ	$0, ""..autotmp_29+6688(SP)
	0x024a 00586 (1.go:165)	MOVQ	$0, ""..autotmp_29+6696(SP)
	0x0256 00598 (1.go:165)	MOVQ	$0, ""..autotmp_29+6704(SP)
	0x0262 00610 (1.go:165)	MOVQ	$0, ""..autotmp_29+6712(SP)
	0x026e 00622 (1.go:163)	LEAQ	type.string(SB), DX
	0x0275 00629 (1.go:165)	MOVQ	DX, (SP)
	0x0279 00633 (1.go:165)	LEAQ	""..autotmp_30+6128(SP), BX
	0x0281 00641 (1.go:165)	MOVQ	BX, 8(SP)
	0x0286 00646 (1.go:165)	PCDATA	$0, $3
	0x0286 00646 (1.go:165)	CALL	runtime.convT2E(SB)
	0x028b 00651 (1.go:165)	MOVQ	16(SP), DX
	0x0290 00656 (1.go:165)	MOVQ	24(SP), BX
	0x0295 00661 (1.go:165)	MOVQ	DX, ""..autotmp_29+6688(SP)
	0x029d 00669 (1.go:165)	MOVQ	BX, ""..autotmp_29+6696(SP)
	0x02a5 00677 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x02ac 00684 (1.go:165)	MOVQ	DX, (SP)
	0x02b0 00688 (1.go:165)	LEAQ	"".mem+136(SP), BX
	0x02b8 00696 (1.go:165)	MOVQ	BX, 8(SP)
	0x02bd 00701 (1.go:165)	PCDATA	$0, $3
	0x02bd 00701 (1.go:165)	CALL	runtime.convT2E(SB)
	0x02c2 00706 (1.go:165)	MOVQ	24(SP), DX
	0x02c7 00711 (1.go:165)	MOVQ	16(SP), BX
	0x02cc 00716 (1.go:165)	MOVQ	BX, ""..autotmp_29+6704(SP)
	0x02d4 00724 (1.go:165)	MOVQ	DX, ""..autotmp_29+6712(SP)
	0x02dc 00732 (1.go:165)	LEAQ	""..autotmp_29+6688(SP), DX
	0x02e4 00740 (1.go:165)	MOVQ	DX, (SP)
	0x02e8 00744 (1.go:165)	MOVQ	$2, 8(SP)
	0x02f1 00753 (1.go:165)	MOVQ	$2, 16(SP)
	0x02fa 00762 (1.go:165)	PCDATA	$0, $3
	0x02fa 00762 (1.go:165)	CALL	fmt.Println(SB)
	0x02ff 00767 (1.go:166)	LEAQ	go.string."HeapAlloc: "(SB), DX
	0x0306 00774 (1.go:166)	MOVQ	DX, ""..autotmp_32+6112(SP)
	0x030e 00782 (1.go:166)	MOVQ	$11, ""..autotmp_32+6120(SP)
	0x031a 00794 (1.go:166)	MOVQ	$0, ""..autotmp_31+6656(SP)
	0x0326 00806 (1.go:166)	MOVQ	$0, ""..autotmp_31+6664(SP)
	0x0332 00818 (1.go:166)	MOVQ	$0, ""..autotmp_31+6672(SP)
	0x033e 00830 (1.go:166)	MOVQ	$0, ""..autotmp_31+6680(SP)
	0x034a 00842 (1.go:163)	LEAQ	type.string(SB), DX
	0x0351 00849 (1.go:166)	MOVQ	DX, (SP)
	0x0355 00853 (1.go:166)	LEAQ	""..autotmp_32+6112(SP), BX
	0x035d 00861 (1.go:166)	MOVQ	BX, 8(SP)
	0x0362 00866 (1.go:166)	PCDATA	$0, $4
	0x0362 00866 (1.go:166)	CALL	runtime.convT2E(SB)
	0x0367 00871 (1.go:166)	MOVQ	16(SP), DX
	0x036c 00876 (1.go:166)	MOVQ	24(SP), BX
	0x0371 00881 (1.go:166)	MOVQ	DX, ""..autotmp_31+6656(SP)
	0x0379 00889 (1.go:166)	MOVQ	BX, ""..autotmp_31+6664(SP)
	0x0381 00897 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x0388 00904 (1.go:166)	MOVQ	DX, (SP)
	0x038c 00908 (1.go:166)	LEAQ	"".mem+144(SP), BX
	0x0394 00916 (1.go:166)	MOVQ	BX, 8(SP)
	0x0399 00921 (1.go:166)	PCDATA	$0, $4
	0x0399 00921 (1.go:166)	CALL	runtime.convT2E(SB)
	0x039e 00926 (1.go:166)	MOVQ	16(SP), DX
	0x03a3 00931 (1.go:166)	MOVQ	24(SP), BX
	0x03a8 00936 (1.go:166)	MOVQ	DX, ""..autotmp_31+6672(SP)
	0x03b0 00944 (1.go:166)	MOVQ	BX, ""..autotmp_31+6680(SP)
	0x03b8 00952 (1.go:166)	LEAQ	""..autotmp_31+6656(SP), DX
	0x03c0 00960 (1.go:166)	MOVQ	DX, (SP)
	0x03c4 00964 (1.go:166)	MOVQ	$2, 8(SP)
	0x03cd 00973 (1.go:166)	MOVQ	$2, 16(SP)
	0x03d6 00982 (1.go:166)	PCDATA	$0, $4
	0x03d6 00982 (1.go:166)	CALL	fmt.Println(SB)
	0x03db 00987 (1.go:167)	LEAQ	go.string."HeapIdle: "(SB), DX
	0x03e2 00994 (1.go:167)	MOVQ	DX, ""..autotmp_34+6096(SP)
	0x03ea 01002 (1.go:167)	MOVQ	$10, ""..autotmp_34+6104(SP)
	0x03f6 01014 (1.go:167)	MOVQ	$0, ""..autotmp_33+6624(SP)
	0x0402 01026 (1.go:167)	MOVQ	$0, ""..autotmp_33+6632(SP)
	0x040e 01038 (1.go:167)	MOVQ	$0, ""..autotmp_33+6640(SP)
	0x041a 01050 (1.go:167)	MOVQ	$0, ""..autotmp_33+6648(SP)
	0x0426 01062 (1.go:163)	LEAQ	type.string(SB), DX
	0x042d 01069 (1.go:167)	MOVQ	DX, (SP)
	0x0431 01073 (1.go:167)	LEAQ	""..autotmp_34+6096(SP), BX
	0x0439 01081 (1.go:167)	MOVQ	BX, 8(SP)
	0x043e 01086 (1.go:167)	PCDATA	$0, $5
	0x043e 01086 (1.go:167)	CALL	runtime.convT2E(SB)
	0x0443 01091 (1.go:167)	MOVQ	24(SP), DX
	0x0448 01096 (1.go:167)	MOVQ	16(SP), BX
	0x044d 01101 (1.go:167)	MOVQ	BX, ""..autotmp_33+6624(SP)
	0x0455 01109 (1.go:167)	MOVQ	DX, ""..autotmp_33+6632(SP)
	0x045d 01117 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x0464 01124 (1.go:167)	MOVQ	DX, (SP)
	0x0468 01128 (1.go:167)	LEAQ	"".mem+160(SP), BX
	0x0470 01136 (1.go:167)	MOVQ	BX, 8(SP)
	0x0475 01141 (1.go:167)	PCDATA	$0, $5
	0x0475 01141 (1.go:167)	CALL	runtime.convT2E(SB)
	0x047a 01146 (1.go:167)	MOVQ	16(SP), DX
	0x047f 01151 (1.go:167)	MOVQ	24(SP), BX
	0x0484 01156 (1.go:167)	MOVQ	DX, ""..autotmp_33+6640(SP)
	0x048c 01164 (1.go:167)	MOVQ	BX, ""..autotmp_33+6648(SP)
	0x0494 01172 (1.go:167)	LEAQ	""..autotmp_33+6624(SP), DX
	0x049c 01180 (1.go:167)	MOVQ	DX, (SP)
	0x04a0 01184 (1.go:167)	MOVQ	$2, 8(SP)
	0x04a9 01193 (1.go:167)	MOVQ	$2, 16(SP)
	0x04b2 01202 (1.go:167)	PCDATA	$0, $5
	0x04b2 01202 (1.go:167)	CALL	fmt.Println(SB)
	0x04b7 01207 (1.go:168)	LEAQ	go.string."HeapObjects: "(SB), DX
	0x04be 01214 (1.go:168)	MOVQ	DX, ""..autotmp_36+6080(SP)
	0x04c6 01222 (1.go:168)	MOVQ	$13, ""..autotmp_36+6088(SP)
	0x04d2 01234 (1.go:168)	MOVQ	$0, ""..autotmp_35+6592(SP)
	0x04de 01246 (1.go:168)	MOVQ	$0, ""..autotmp_35+6600(SP)
	0x04ea 01258 (1.go:168)	MOVQ	$0, ""..autotmp_35+6608(SP)
	0x04f6 01270 (1.go:168)	MOVQ	$0, ""..autotmp_35+6616(SP)
	0x0502 01282 (1.go:163)	LEAQ	type.string(SB), DX
	0x0509 01289 (1.go:168)	MOVQ	DX, (SP)
	0x050d 01293 (1.go:168)	LEAQ	""..autotmp_36+6080(SP), BX
	0x0515 01301 (1.go:168)	MOVQ	BX, 8(SP)
	0x051a 01306 (1.go:168)	PCDATA	$0, $6
	0x051a 01306 (1.go:168)	CALL	runtime.convT2E(SB)
	0x051f 01311 (1.go:168)	MOVQ	24(SP), DX
	0x0524 01316 (1.go:168)	MOVQ	16(SP), BX
	0x0529 01321 (1.go:168)	MOVQ	BX, ""..autotmp_35+6592(SP)
	0x0531 01329 (1.go:168)	MOVQ	DX, ""..autotmp_35+6600(SP)
	0x0539 01337 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x0540 01344 (1.go:168)	MOVQ	DX, (SP)
	0x0544 01348 (1.go:168)	LEAQ	"".mem+184(SP), BX
	0x054c 01356 (1.go:168)	MOVQ	BX, 8(SP)
	0x0551 01361 (1.go:168)	PCDATA	$0, $6
	0x0551 01361 (1.go:168)	CALL	runtime.convT2E(SB)
	0x0556 01366 (1.go:168)	MOVQ	24(SP), DX
	0x055b 01371 (1.go:168)	MOVQ	16(SP), BX
	0x0560 01376 (1.go:168)	MOVQ	BX, ""..autotmp_35+6608(SP)
	0x0568 01384 (1.go:168)	MOVQ	DX, ""..autotmp_35+6616(SP)
	0x0570 01392 (1.go:168)	LEAQ	""..autotmp_35+6592(SP), DX
	0x0578 01400 (1.go:168)	MOVQ	DX, (SP)
	0x057c 01404 (1.go:168)	MOVQ	$2, 8(SP)
	0x0585 01413 (1.go:168)	MOVQ	$2, 16(SP)
	0x058e 01422 (1.go:168)	PCDATA	$0, $6
	0x058e 01422 (1.go:168)	CALL	fmt.Println(SB)
	0x0593 01427 (1.go:169)	LEAQ	go.string."HeapReleased: "(SB), DX
	0x059a 01434 (1.go:169)	MOVQ	DX, ""..autotmp_38+6064(SP)
	0x05a2 01442 (1.go:169)	MOVQ	$14, ""..autotmp_38+6072(SP)
	0x05ae 01454 (1.go:169)	MOVQ	$0, ""..autotmp_37+6560(SP)
	0x05ba 01466 (1.go:169)	MOVQ	$0, ""..autotmp_37+6568(SP)
	0x05c6 01478 (1.go:169)	MOVQ	$0, ""..autotmp_37+6576(SP)
	0x05d2 01490 (1.go:169)	MOVQ	$0, ""..autotmp_37+6584(SP)
	0x05de 01502 (1.go:163)	LEAQ	type.string(SB), DX
	0x05e5 01509 (1.go:169)	MOVQ	DX, (SP)
	0x05e9 01513 (1.go:169)	LEAQ	""..autotmp_38+6064(SP), BX
	0x05f1 01521 (1.go:169)	MOVQ	BX, 8(SP)
	0x05f6 01526 (1.go:169)	PCDATA	$0, $7
	0x05f6 01526 (1.go:169)	CALL	runtime.convT2E(SB)
	0x05fb 01531 (1.go:169)	MOVQ	24(SP), DX
	0x0600 01536 (1.go:169)	MOVQ	16(SP), BX
	0x0605 01541 (1.go:169)	MOVQ	BX, ""..autotmp_37+6560(SP)
	0x060d 01549 (1.go:169)	MOVQ	DX, ""..autotmp_37+6568(SP)
	0x0615 01557 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x061c 01564 (1.go:169)	MOVQ	DX, (SP)
	0x0620 01568 (1.go:169)	LEAQ	"".mem+176(SP), BX
	0x0628 01576 (1.go:169)	MOVQ	BX, 8(SP)
	0x062d 01581 (1.go:169)	PCDATA	$0, $7
	0x062d 01581 (1.go:169)	CALL	runtime.convT2E(SB)
	0x0632 01586 (1.go:169)	MOVQ	24(SP), DX
	0x0637 01591 (1.go:169)	MOVQ	16(SP), BX
	0x063c 01596 (1.go:169)	MOVQ	BX, ""..autotmp_37+6576(SP)
	0x0644 01604 (1.go:169)	MOVQ	DX, ""..autotmp_37+6584(SP)
	0x064c 01612 (1.go:169)	LEAQ	""..autotmp_37+6560(SP), DX
	0x0654 01620 (1.go:169)	MOVQ	DX, (SP)
	0x0658 01624 (1.go:169)	MOVQ	$2, 8(SP)
	0x0661 01633 (1.go:169)	MOVQ	$2, 16(SP)
	0x066a 01642 (1.go:169)	PCDATA	$0, $7
	0x066a 01642 (1.go:169)	CALL	fmt.Println(SB)
	0x066f 01647 (1.go:170)	LEAQ	go.string."TotalAlloc: "(SB), DX
	0x0676 01654 (1.go:170)	MOVQ	DX, ""..autotmp_40+6048(SP)
	0x067e 01662 (1.go:170)	MOVQ	$12, ""..autotmp_40+6056(SP)
	0x068a 01674 (1.go:170)	MOVQ	$0, ""..autotmp_39+6528(SP)
	0x0696 01686 (1.go:170)	MOVQ	$0, ""..autotmp_39+6536(SP)
	0x06a2 01698 (1.go:170)	MOVQ	$0, ""..autotmp_39+6544(SP)
	0x06ae 01710 (1.go:170)	MOVQ	$0, ""..autotmp_39+6552(SP)
	0x06ba 01722 (1.go:163)	LEAQ	type.string(SB), DX
	0x06c1 01729 (1.go:170)	MOVQ	DX, (SP)
	0x06c5 01733 (1.go:170)	LEAQ	""..autotmp_40+6048(SP), BX
	0x06cd 01741 (1.go:170)	MOVQ	BX, 8(SP)
	0x06d2 01746 (1.go:170)	PCDATA	$0, $8
	0x06d2 01746 (1.go:170)	CALL	runtime.convT2E(SB)
	0x06d7 01751 (1.go:170)	MOVQ	24(SP), DX
	0x06dc 01756 (1.go:170)	MOVQ	16(SP), BX
	0x06e1 01761 (1.go:170)	MOVQ	BX, ""..autotmp_39+6528(SP)
	0x06e9 01769 (1.go:170)	MOVQ	DX, ""..autotmp_39+6536(SP)
	0x06f1 01777 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x06f8 01784 (1.go:170)	MOVQ	DX, (SP)
	0x06fc 01788 (1.go:170)	LEAQ	"".mem+104(SP), BX
	0x0701 01793 (1.go:170)	MOVQ	BX, 8(SP)
	0x0706 01798 (1.go:170)	PCDATA	$0, $8
	0x0706 01798 (1.go:170)	CALL	runtime.convT2E(SB)
	0x070b 01803 (1.go:170)	MOVQ	24(SP), DX
	0x0710 01808 (1.go:170)	MOVQ	16(SP), BX
	0x0715 01813 (1.go:170)	MOVQ	BX, ""..autotmp_39+6544(SP)
	0x071d 01821 (1.go:170)	MOVQ	DX, ""..autotmp_39+6552(SP)
	0x0725 01829 (1.go:170)	LEAQ	""..autotmp_39+6528(SP), DX
	0x072d 01837 (1.go:170)	MOVQ	DX, (SP)
	0x0731 01841 (1.go:170)	MOVQ	$2, 8(SP)
	0x073a 01850 (1.go:170)	MOVQ	$2, 16(SP)
	0x0743 01859 (1.go:170)	PCDATA	$0, $8
	0x0743 01859 (1.go:170)	CALL	fmt.Println(SB)
	0x0748 01864 (1.go:171)	LEAQ	go.string."StackInuse: "(SB), DX
	0x074f 01871 (1.go:171)	MOVQ	DX, ""..autotmp_42+6032(SP)
	0x0757 01879 (1.go:171)	MOVQ	$12, ""..autotmp_42+6040(SP)
	0x0763 01891 (1.go:171)	MOVQ	$0, ""..autotmp_41+6496(SP)
	0x076f 01903 (1.go:171)	MOVQ	$0, ""..autotmp_41+6504(SP)
	0x077b 01915 (1.go:171)	MOVQ	$0, ""..autotmp_41+6512(SP)
	0x0787 01927 (1.go:171)	MOVQ	$0, ""..autotmp_41+6520(SP)
	0x0793 01939 (1.go:163)	LEAQ	type.string(SB), DX
	0x079a 01946 (1.go:171)	MOVQ	DX, (SP)
	0x079e 01950 (1.go:171)	LEAQ	""..autotmp_42+6032(SP), BX
	0x07a6 01958 (1.go:171)	MOVQ	BX, 8(SP)
	0x07ab 01963 (1.go:171)	PCDATA	$0, $9
	0x07ab 01963 (1.go:171)	CALL	runtime.convT2E(SB)
	0x07b0 01968 (1.go:171)	MOVQ	24(SP), DX
	0x07b5 01973 (1.go:171)	MOVQ	16(SP), BX
	0x07ba 01978 (1.go:171)	MOVQ	BX, ""..autotmp_41+6496(SP)
	0x07c2 01986 (1.go:171)	MOVQ	DX, ""..autotmp_41+6504(SP)
	0x07ca 01994 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x07d1 02001 (1.go:171)	MOVQ	DX, (SP)
	0x07d5 02005 (1.go:171)	LEAQ	"".mem+192(SP), BX
	0x07dd 02013 (1.go:171)	MOVQ	BX, 8(SP)
	0x07e2 02018 (1.go:171)	PCDATA	$0, $9
	0x07e2 02018 (1.go:171)	CALL	runtime.convT2E(SB)
	0x07e7 02023 (1.go:171)	MOVQ	16(SP), DX
	0x07ec 02028 (1.go:171)	MOVQ	24(SP), BX
	0x07f1 02033 (1.go:171)	MOVQ	DX, ""..autotmp_41+6512(SP)
	0x07f9 02041 (1.go:171)	MOVQ	BX, ""..autotmp_41+6520(SP)
	0x0801 02049 (1.go:171)	LEAQ	""..autotmp_41+6496(SP), DX
	0x0809 02057 (1.go:171)	MOVQ	DX, (SP)
	0x080d 02061 (1.go:171)	MOVQ	$2, 8(SP)
	0x0816 02070 (1.go:171)	MOVQ	$2, 16(SP)
	0x081f 02079 (1.go:171)	PCDATA	$0, $9
	0x081f 02079 (1.go:171)	CALL	fmt.Println(SB)
	0x0824 02084 (1.go:172)	LEAQ	go.string."Lookups: "(SB), DX
	0x082b 02091 (1.go:172)	MOVQ	DX, ""..autotmp_44+6016(SP)
	0x0833 02099 (1.go:172)	MOVQ	$9, ""..autotmp_44+6024(SP)
	0x083f 02111 (1.go:172)	MOVQ	$0, ""..autotmp_43+6464(SP)
	0x084b 02123 (1.go:172)	MOVQ	$0, ""..autotmp_43+6472(SP)
	0x0857 02135 (1.go:172)	MOVQ	$0, ""..autotmp_43+6480(SP)
	0x0863 02147 (1.go:172)	MOVQ	$0, ""..autotmp_43+6488(SP)
	0x086f 02159 (1.go:163)	LEAQ	type.string(SB), DX
	0x0876 02166 (1.go:172)	MOVQ	DX, (SP)
	0x087a 02170 (1.go:172)	LEAQ	""..autotmp_44+6016(SP), BX
	0x0882 02178 (1.go:172)	MOVQ	BX, 8(SP)
	0x0887 02183 (1.go:172)	PCDATA	$0, $10
	0x0887 02183 (1.go:172)	CALL	runtime.convT2E(SB)
	0x088c 02188 (1.go:172)	MOVQ	24(SP), DX
	0x0891 02193 (1.go:172)	MOVQ	16(SP), BX
	0x0896 02198 (1.go:172)	MOVQ	BX, ""..autotmp_43+6464(SP)
	0x089e 02206 (1.go:172)	MOVQ	DX, ""..autotmp_43+6472(SP)
	0x08a6 02214 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x08ad 02221 (1.go:172)	MOVQ	DX, (SP)
	0x08b1 02225 (1.go:172)	LEAQ	"".mem+120(SP), BX
	0x08b6 02230 (1.go:172)	MOVQ	BX, 8(SP)
	0x08bb 02235 (1.go:172)	PCDATA	$0, $10
	0x08bb 02235 (1.go:172)	CALL	runtime.convT2E(SB)
	0x08c0 02240 (1.go:172)	MOVQ	16(SP), DX
	0x08c5 02245 (1.go:172)	MOVQ	24(SP), BX
	0x08ca 02250 (1.go:172)	MOVQ	DX, ""..autotmp_43+6480(SP)
	0x08d2 02258 (1.go:172)	MOVQ	BX, ""..autotmp_43+6488(SP)
	0x08da 02266 (1.go:172)	LEAQ	""..autotmp_43+6464(SP), DX
	0x08e2 02274 (1.go:172)	MOVQ	DX, (SP)
	0x08e6 02278 (1.go:172)	MOVQ	$2, 8(SP)
	0x08ef 02287 (1.go:172)	MOVQ	$2, 16(SP)
	0x08f8 02296 (1.go:172)	PCDATA	$0, $10
	0x08f8 02296 (1.go:172)	CALL	fmt.Println(SB)
	0x08fd 02301 (1.go:174)	LEAQ	go.string."LastGC: "(SB), DX
	0x0904 02308 (1.go:174)	MOVQ	DX, ""..autotmp_46+6000(SP)
	0x090c 02316 (1.go:174)	MOVQ	$8, ""..autotmp_46+6008(SP)
	0x0918 02328 (1.go:174)	MOVQ	"".mem+272(SP), AX
	0x0920 02336 (1.go:174)	MOVQ	AX, ""..autotmp_149+88(SP)
	0x0925 02341 (1.go:174)	MOVQ	$1360296554856532783, DX
	0x092f 02351 (1.go:174)	MULQ	DX
	0x0932 02354 (1.go:174)	MOVQ	""..autotmp_149+88(SP), BX
	0x0937 02359 (1.go:174)	ADDQ	DX, BX
	0x093a 02362 (1.go:174)	RCRQ	$1, BX
	0x093d 02365 (1.go:174)	SHRQ	$29, BX
	0x0941 02369 (1.go:174)	MOVQ	BX, ""..autotmp_47+80(SP)
	0x0946 02374 (1.go:174)	MOVQ	$0, ""..autotmp_45+6432(SP)
	0x0952 02386 (1.go:174)	MOVQ	$0, ""..autotmp_45+6440(SP)
	0x095e 02398 (1.go:174)	MOVQ	$0, ""..autotmp_45+6448(SP)
	0x096a 02410 (1.go:174)	MOVQ	$0, ""..autotmp_45+6456(SP)
	0x0976 02422 (1.go:163)	LEAQ	type.string(SB), DX
	0x097d 02429 (1.go:174)	MOVQ	DX, (SP)
	0x0981 02433 (1.go:174)	LEAQ	""..autotmp_46+6000(SP), BX
	0x0989 02441 (1.go:174)	MOVQ	BX, 8(SP)
	0x098e 02446 (1.go:174)	PCDATA	$0, $11
	0x098e 02446 (1.go:174)	CALL	runtime.convT2E(SB)
	0x0993 02451 (1.go:174)	MOVQ	16(SP), DX
	0x0998 02456 (1.go:174)	MOVQ	24(SP), BX
	0x099d 02461 (1.go:174)	MOVQ	DX, ""..autotmp_45+6432(SP)
	0x09a5 02469 (1.go:174)	MOVQ	BX, ""..autotmp_45+6440(SP)
	0x09ad 02477 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x09b4 02484 (1.go:174)	MOVQ	DX, (SP)
	0x09b8 02488 (1.go:174)	LEAQ	""..autotmp_47+80(SP), BX
	0x09bd 02493 (1.go:174)	MOVQ	BX, 8(SP)
	0x09c2 02498 (1.go:174)	PCDATA	$0, $11
	0x09c2 02498 (1.go:174)	CALL	runtime.convT2E(SB)
	0x09c7 02503 (1.go:174)	MOVQ	16(SP), DX
	0x09cc 02508 (1.go:174)	MOVQ	24(SP), BX
	0x09d1 02513 (1.go:174)	MOVQ	DX, ""..autotmp_45+6448(SP)
	0x09d9 02521 (1.go:174)	MOVQ	BX, ""..autotmp_45+6456(SP)
	0x09e1 02529 (1.go:174)	LEAQ	""..autotmp_45+6432(SP), DX
	0x09e9 02537 (1.go:174)	MOVQ	DX, (SP)
	0x09ed 02541 (1.go:174)	MOVQ	$2, 8(SP)
	0x09f6 02550 (1.go:174)	MOVQ	$2, 16(SP)
	0x09ff 02559 (1.go:174)	PCDATA	$0, $11
	0x09ff 02559 (1.go:174)	CALL	fmt.Println(SB)
	0x0a04 02564 (1.go:175)	LEAQ	go.string."NumGC: "(SB), DX
	0x0a0b 02571 (1.go:175)	MOVQ	DX, ""..autotmp_49+5984(SP)
	0x0a13 02579 (1.go:175)	MOVQ	$7, ""..autotmp_49+5992(SP)
	0x0a1f 02591 (1.go:175)	MOVQ	$0, ""..autotmp_48+6400(SP)
	0x0a2b 02603 (1.go:175)	MOVQ	$0, ""..autotmp_48+6408(SP)
	0x0a37 02615 (1.go:175)	MOVQ	$0, ""..autotmp_48+6416(SP)
	0x0a43 02627 (1.go:175)	MOVQ	$0, ""..autotmp_48+6424(SP)
	0x0a4f 02639 (1.go:163)	LEAQ	type.string(SB), DX
	0x0a56 02646 (1.go:175)	MOVQ	DX, (SP)
	0x0a5a 02650 (1.go:175)	LEAQ	""..autotmp_49+5984(SP), BX
	0x0a62 02658 (1.go:175)	MOVQ	BX, 8(SP)
	0x0a67 02663 (1.go:175)	PCDATA	$0, $12
	0x0a67 02663 (1.go:175)	CALL	runtime.convT2E(SB)
	0x0a6c 02668 (1.go:175)	MOVQ	24(SP), DX
	0x0a71 02673 (1.go:175)	MOVQ	16(SP), BX
	0x0a76 02678 (1.go:175)	MOVQ	BX, ""..autotmp_48+6400(SP)
	0x0a7e 02686 (1.go:175)	MOVQ	DX, ""..autotmp_48+6408(SP)
	0x0a86 02694 (1.go:175)	LEAQ	type.uint32(SB), DX
	0x0a8d 02701 (1.go:175)	MOVQ	DX, (SP)
	0x0a91 02705 (1.go:175)	LEAQ	"".mem+4384(SP), DX
	0x0a99 02713 (1.go:175)	MOVQ	DX, 8(SP)
	0x0a9e 02718 (1.go:175)	PCDATA	$0, $12
	0x0a9e 02718 (1.go:175)	CALL	runtime.convT2E(SB)
	0x0aa3 02723 (1.go:175)	MOVQ	24(SP), DX
	0x0aa8 02728 (1.go:175)	MOVQ	16(SP), BX
	0x0aad 02733 (1.go:175)	MOVQ	BX, ""..autotmp_48+6416(SP)
	0x0ab5 02741 (1.go:175)	MOVQ	DX, ""..autotmp_48+6424(SP)
	0x0abd 02749 (1.go:175)	LEAQ	""..autotmp_48+6400(SP), DX
	0x0ac5 02757 (1.go:175)	MOVQ	DX, (SP)
	0x0ac9 02761 (1.go:175)	MOVQ	$2, 8(SP)
	0x0ad2 02770 (1.go:175)	MOVQ	$2, 16(SP)
	0x0adb 02779 (1.go:175)	PCDATA	$0, $12
	0x0adb 02779 (1.go:175)	CALL	fmt.Println(SB)
	0x0ae0 02784 (1.go:176)	LEAQ	go.string."NextGC: "(SB), DX
	0x0ae7 02791 (1.go:176)	MOVQ	DX, ""..autotmp_51+5968(SP)
	0x0aef 02799 (1.go:176)	MOVQ	$8, ""..autotmp_51+5976(SP)
	0x0afb 02811 (1.go:176)	MOVQ	"".mem+264(SP), AX
	0x0b03 02819 (1.go:176)	MOVQ	AX, ""..autotmp_149+88(SP)
	0x0b08 02824 (1.go:174)	MOVQ	$1360296554856532783, DX
	0x0b12 02834 (1.go:176)	MULQ	DX
	0x0b15 02837 (1.go:176)	MOVQ	""..autotmp_149+88(SP), BX
	0x0b1a 02842 (1.go:176)	ADDQ	DX, BX
	0x0b1d 02845 (1.go:176)	RCRQ	$1, BX
	0x0b20 02848 (1.go:176)	SHRQ	$29, BX
	0x0b24 02852 (1.go:176)	MOVQ	BX, ""..autotmp_52+72(SP)
	0x0b29 02857 (1.go:176)	MOVQ	$0, ""..autotmp_50+6368(SP)
	0x0b35 02869 (1.go:176)	MOVQ	$0, ""..autotmp_50+6376(SP)
	0x0b41 02881 (1.go:176)	MOVQ	$0, ""..autotmp_50+6384(SP)
	0x0b4d 02893 (1.go:176)	MOVQ	$0, ""..autotmp_50+6392(SP)
	0x0b59 02905 (1.go:163)	LEAQ	type.string(SB), DX
	0x0b60 02912 (1.go:176)	MOVQ	DX, (SP)
	0x0b64 02916 (1.go:176)	LEAQ	""..autotmp_51+5968(SP), BX
	0x0b6c 02924 (1.go:176)	MOVQ	BX, 8(SP)
	0x0b71 02929 (1.go:176)	PCDATA	$0, $13
	0x0b71 02929 (1.go:176)	CALL	runtime.convT2E(SB)
	0x0b76 02934 (1.go:176)	MOVQ	16(SP), DX
	0x0b7b 02939 (1.go:176)	MOVQ	24(SP), BX
	0x0b80 02944 (1.go:176)	MOVQ	DX, ""..autotmp_50+6368(SP)
	0x0b88 02952 (1.go:176)	MOVQ	BX, ""..autotmp_50+6376(SP)
	0x0b90 02960 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x0b97 02967 (1.go:176)	MOVQ	DX, (SP)
	0x0b9b 02971 (1.go:176)	LEAQ	""..autotmp_52+72(SP), BX
	0x0ba0 02976 (1.go:176)	MOVQ	BX, 8(SP)
	0x0ba5 02981 (1.go:176)	PCDATA	$0, $13
	0x0ba5 02981 (1.go:176)	CALL	runtime.convT2E(SB)
	0x0baa 02986 (1.go:176)	MOVQ	16(SP), DX
	0x0baf 02991 (1.go:176)	MOVQ	24(SP), BX
	0x0bb4 02996 (1.go:176)	MOVQ	DX, ""..autotmp_50+6384(SP)
	0x0bbc 03004 (1.go:176)	MOVQ	BX, ""..autotmp_50+6392(SP)
	0x0bc4 03012 (1.go:176)	LEAQ	""..autotmp_50+6368(SP), DX
	0x0bcc 03020 (1.go:176)	MOVQ	DX, (SP)
	0x0bd0 03024 (1.go:176)	MOVQ	$2, 8(SP)
	0x0bd9 03033 (1.go:176)	MOVQ	$2, 16(SP)
	0x0be2 03042 (1.go:176)	PCDATA	$0, $13
	0x0be2 03042 (1.go:176)	CALL	fmt.Println(SB)
	0x0be7 03047 (1.go:177)	LEAQ	go.string."Mallocs: "(SB), DX
	0x0bee 03054 (1.go:177)	MOVQ	DX, ""..autotmp_54+5952(SP)
	0x0bf6 03062 (1.go:177)	MOVQ	$9, ""..autotmp_54+5960(SP)
	0x0c02 03074 (1.go:177)	MOVQ	$0, ""..autotmp_53+6336(SP)
	0x0c0e 03086 (1.go:177)	MOVQ	$0, ""..autotmp_53+6344(SP)
	0x0c1a 03098 (1.go:177)	MOVQ	$0, ""..autotmp_53+6352(SP)
	0x0c26 03110 (1.go:177)	MOVQ	$0, ""..autotmp_53+6360(SP)
	0x0c32 03122 (1.go:163)	LEAQ	type.string(SB), DX
	0x0c39 03129 (1.go:177)	MOVQ	DX, (SP)
	0x0c3d 03133 (1.go:177)	LEAQ	""..autotmp_54+5952(SP), BX
	0x0c45 03141 (1.go:177)	MOVQ	BX, 8(SP)
	0x0c4a 03146 (1.go:177)	PCDATA	$0, $14
	0x0c4a 03146 (1.go:177)	CALL	runtime.convT2E(SB)
	0x0c4f 03151 (1.go:177)	MOVQ	16(SP), DX
	0x0c54 03156 (1.go:177)	MOVQ	24(SP), BX
	0x0c59 03161 (1.go:177)	MOVQ	DX, ""..autotmp_53+6336(SP)
	0x0c61 03169 (1.go:177)	MOVQ	BX, ""..autotmp_53+6344(SP)
	0x0c69 03177 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x0c70 03184 (1.go:177)	MOVQ	DX, (SP)
	0x0c74 03188 (1.go:177)	LEAQ	"".mem+128(SP), BX
	0x0c7c 03196 (1.go:177)	MOVQ	BX, 8(SP)
	0x0c81 03201 (1.go:177)	PCDATA	$0, $14
	0x0c81 03201 (1.go:177)	CALL	runtime.convT2E(SB)
	0x0c86 03206 (1.go:177)	MOVQ	16(SP), DX
	0x0c8b 03211 (1.go:177)	MOVQ	24(SP), BX
	0x0c90 03216 (1.go:177)	MOVQ	DX, ""..autotmp_53+6352(SP)
	0x0c98 03224 (1.go:177)	MOVQ	BX, ""..autotmp_53+6360(SP)
	0x0ca0 03232 (1.go:177)	LEAQ	""..autotmp_53+6336(SP), DX
	0x0ca8 03240 (1.go:177)	MOVQ	DX, (SP)
	0x0cac 03244 (1.go:177)	MOVQ	$2, 8(SP)
	0x0cb5 03253 (1.go:177)	MOVQ	$2, 16(SP)
	0x0cbe 03262 (1.go:177)	PCDATA	$0, $14
	0x0cbe 03262 (1.go:177)	CALL	fmt.Println(SB)
	0x0cc3 03267 (1.go:178)	LEAQ	go.string."PauseTotalNs: "(SB), DX
	0x0cca 03274 (1.go:178)	MOVQ	DX, ""..autotmp_56+5936(SP)
	0x0cd2 03282 (1.go:178)	MOVQ	$14, ""..autotmp_56+5944(SP)
	0x0cde 03294 (1.go:178)	MOVQ	"".mem+280(SP), DX
	0x0ce6 03302 (1.go:178)	MOVQ	$4835703278458516699, AX
	0x0cf0 03312 (1.go:178)	MULQ	DX
	0x0cf3 03315 (1.go:178)	SHRQ	$18, DX
	0x0cf7 03319 (1.go:178)	MOVQ	DX, ""..autotmp_57+64(SP)
	0x0cfc 03324 (1.go:178)	MOVQ	$0, ""..autotmp_55+6304(SP)
	0x0d08 03336 (1.go:178)	MOVQ	$0, ""..autotmp_55+6312(SP)
	0x0d14 03348 (1.go:178)	MOVQ	$0, ""..autotmp_55+6320(SP)
	0x0d20 03360 (1.go:178)	MOVQ	$0, ""..autotmp_55+6328(SP)
	0x0d2c 03372 (1.go:163)	LEAQ	type.string(SB), DX
	0x0d33 03379 (1.go:178)	MOVQ	DX, (SP)
	0x0d37 03383 (1.go:178)	LEAQ	""..autotmp_56+5936(SP), BX
	0x0d3f 03391 (1.go:178)	MOVQ	BX, 8(SP)
	0x0d44 03396 (1.go:178)	PCDATA	$0, $15
	0x0d44 03396 (1.go:178)	CALL	runtime.convT2E(SB)
	0x0d49 03401 (1.go:178)	MOVQ	16(SP), DX
	0x0d4e 03406 (1.go:178)	MOVQ	24(SP), BX
	0x0d53 03411 (1.go:178)	MOVQ	DX, ""..autotmp_55+6304(SP)
	0x0d5b 03419 (1.go:178)	MOVQ	BX, ""..autotmp_55+6312(SP)
	0x0d63 03427 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x0d6a 03434 (1.go:178)	MOVQ	DX, (SP)
	0x0d6e 03438 (1.go:178)	LEAQ	""..autotmp_57+64(SP), BX
	0x0d73 03443 (1.go:178)	MOVQ	BX, 8(SP)
	0x0d78 03448 (1.go:178)	PCDATA	$0, $15
	0x0d78 03448 (1.go:178)	CALL	runtime.convT2E(SB)
	0x0d7d 03453 (1.go:178)	MOVQ	16(SP), DX
	0x0d82 03458 (1.go:178)	MOVQ	24(SP), BX
	0x0d87 03463 (1.go:178)	MOVQ	DX, ""..autotmp_55+6320(SP)
	0x0d8f 03471 (1.go:178)	MOVQ	BX, ""..autotmp_55+6328(SP)
	0x0d97 03479 (1.go:178)	LEAQ	""..autotmp_55+6304(SP), DX
	0x0d9f 03487 (1.go:178)	MOVQ	DX, (SP)
	0x0da3 03491 (1.go:178)	MOVQ	$2, 8(SP)
	0x0dac 03500 (1.go:178)	MOVQ	$2, 16(SP)
	0x0db5 03509 (1.go:178)	PCDATA	$0, $15
	0x0db5 03509 (1.go:178)	CALL	fmt.Println(SB)
	0x0dba 03514 (1.go:180)	LEAQ	go.string."GCCPUFraction: "(SB), DX
	0x0dc1 03521 (1.go:180)	MOVQ	DX, ""..autotmp_59+5920(SP)
	0x0dc9 03529 (1.go:180)	MOVQ	$15, ""..autotmp_59+5928(SP)
	0x0dd5 03541 (1.go:180)	MOVQ	$0, ""..autotmp_58+6272(SP)
	0x0de1 03553 (1.go:180)	MOVQ	$0, ""..autotmp_58+6280(SP)
	0x0ded 03565 (1.go:180)	MOVQ	$0, ""..autotmp_58+6288(SP)
	0x0df9 03577 (1.go:180)	MOVQ	$0, ""..autotmp_58+6296(SP)
	0x0e05 03589 (1.go:163)	LEAQ	type.string(SB), DX
	0x0e0c 03596 (1.go:180)	MOVQ	DX, (SP)
	0x0e10 03600 (1.go:180)	LEAQ	""..autotmp_59+5920(SP), BX
	0x0e18 03608 (1.go:180)	MOVQ	BX, 8(SP)
	0x0e1d 03613 (1.go:180)	PCDATA	$0, $16
	0x0e1d 03613 (1.go:180)	CALL	runtime.convT2E(SB)
	0x0e22 03618 (1.go:180)	MOVQ	16(SP), DX
	0x0e27 03623 (1.go:180)	MOVQ	24(SP), BX
	0x0e2c 03628 (1.go:180)	MOVQ	DX, ""..autotmp_58+6272(SP)
	0x0e34 03636 (1.go:180)	MOVQ	BX, ""..autotmp_58+6280(SP)
	0x0e3c 03644 (1.go:180)	LEAQ	type.float64(SB), DX
	0x0e43 03651 (1.go:180)	MOVQ	DX, (SP)
	0x0e47 03655 (1.go:180)	LEAQ	"".mem+4392(SP), DX
	0x0e4f 03663 (1.go:180)	MOVQ	DX, 8(SP)
	0x0e54 03668 (1.go:180)	PCDATA	$0, $16
	0x0e54 03668 (1.go:180)	CALL	runtime.convT2E(SB)
	0x0e59 03673 (1.go:180)	MOVQ	16(SP), DX
	0x0e5e 03678 (1.go:180)	MOVQ	24(SP), BX
	0x0e63 03683 (1.go:180)	MOVQ	DX, ""..autotmp_58+6288(SP)
	0x0e6b 03691 (1.go:180)	MOVQ	BX, ""..autotmp_58+6296(SP)
	0x0e73 03699 (1.go:180)	LEAQ	""..autotmp_58+6272(SP), DX
	0x0e7b 03707 (1.go:180)	MOVQ	DX, (SP)
	0x0e7f 03711 (1.go:180)	MOVQ	$2, 8(SP)
	0x0e88 03720 (1.go:180)	MOVQ	$2, 16(SP)
	0x0e91 03729 (1.go:180)	PCDATA	$0, $16
	0x0e91 03729 (1.go:180)	CALL	fmt.Println(SB)
	0x0e96 03734 (1.go:181)	LEAQ	go.string."GCSys: "(SB), DX
	0x0e9d 03741 (1.go:181)	MOVQ	DX, ""..autotmp_61+5904(SP)
	0x0ea5 03749 (1.go:181)	MOVQ	$7, ""..autotmp_61+5912(SP)
	0x0eb1 03761 (1.go:181)	MOVQ	$0, ""..autotmp_60+6240(SP)
	0x0ebd 03773 (1.go:181)	MOVQ	$0, ""..autotmp_60+6248(SP)
	0x0ec9 03785 (1.go:181)	MOVQ	$0, ""..autotmp_60+6256(SP)
	0x0ed5 03797 (1.go:181)	MOVQ	$0, ""..autotmp_60+6264(SP)
	0x0ee1 03809 (1.go:163)	LEAQ	type.string(SB), DX
	0x0ee8 03816 (1.go:181)	MOVQ	DX, (SP)
	0x0eec 03820 (1.go:181)	LEAQ	""..autotmp_61+5904(SP), BX
	0x0ef4 03828 (1.go:181)	MOVQ	BX, 8(SP)
	0x0ef9 03833 (1.go:181)	PCDATA	$0, $17
	0x0ef9 03833 (1.go:181)	CALL	runtime.convT2E(SB)
	0x0efe 03838 (1.go:181)	MOVQ	16(SP), DX
	0x0f03 03843 (1.go:181)	MOVQ	24(SP), BX
	0x0f08 03848 (1.go:181)	MOVQ	DX, ""..autotmp_60+6240(SP)
	0x0f10 03856 (1.go:181)	MOVQ	BX, ""..autotmp_60+6248(SP)
	0x0f18 03864 (1.go:163)	LEAQ	type.uint64(SB), DX
	0x0f1f 03871 (1.go:181)	MOVQ	DX, (SP)
	0x0f23 03875 (1.go:181)	LEAQ	"".mem+248(SP), DX
	0x0f2b 03883 (1.go:181)	MOVQ	DX, 8(SP)
	0x0f30 03888 (1.go:181)	PCDATA	$0, $17
	0x0f30 03888 (1.go:181)	CALL	runtime.convT2E(SB)
	0x0f35 03893 (1.go:181)	MOVQ	16(SP), DX
	0x0f3a 03898 (1.go:181)	MOVQ	24(SP), BX
	0x0f3f 03903 (1.go:181)	MOVQ	DX, ""..autotmp_60+6256(SP)
	0x0f47 03911 (1.go:181)	MOVQ	BX, ""..autotmp_60+6264(SP)
	0x0f4f 03919 (1.go:181)	LEAQ	""..autotmp_60+6240(SP), DX
	0x0f57 03927 (1.go:181)	MOVQ	DX, (SP)
	0x0f5b 03931 (1.go:181)	MOVQ	$2, 8(SP)
	0x0f64 03940 (1.go:181)	MOVQ	$2, 16(SP)
	0x0f6d 03949 (1.go:181)	PCDATA	$0, $17
	0x0f6d 03949 (1.go:181)	CALL	fmt.Println(SB)
	0x0f72 03954 (1.go:183)	LEAQ	go.string."Time: "(SB), DX
	0x0f79 03961 (1.go:183)	MOVQ	DX, ""..autotmp_63+5888(SP)
	0x0f81 03969 (1.go:183)	MOVQ	$6, ""..autotmp_63+5896(SP)
	0x0f8d 03981 (1.go:183)	PCDATA	$0, $18
	0x0f8d 03981 (1.go:183)	CALL	time.Now(SB)
	0x0f92 03986 (1.go:183)	MOVQ	"".t1(SB), DX
	0x0f99 03993 (1.go:183)	MOVQ	(SP), BX
	0x0f9d 03997 (1.go:155)	MOVQ	$-62135596800, SI
	0x0fa7 04007 (1.go:183)	ADDQ	SI, BX
	0x0faa 04010 (1.go:183)	SUBQ	DX, BX
	0x0fad 04013 (1.go:183)	MOVQ	BX, ""..autotmp_64+56(SP)
	0x0fb2 04018 (1.go:183)	MOVQ	$0, ""..autotmp_62+6208(SP)
	0x0fbe 04030 (1.go:183)	MOVQ	$0, ""..autotmp_62+6216(SP)
	0x0fca 04042 (1.go:183)	MOVQ	$0, ""..autotmp_62+6224(SP)
	0x0fd6 04054 (1.go:183)	MOVQ	$0, ""..autotmp_62+6232(SP)
	0x0fe2 04066 (1.go:163)	LEAQ	type.string(SB), DX
	0x0fe9 04073 (1.go:183)	MOVQ	DX, (SP)
	0x0fed 04077 (1.go:183)	LEAQ	""..autotmp_63+5888(SP), BX
	0x0ff5 04085 (1.go:183)	MOVQ	BX, 8(SP)
	0x0ffa 04090 (1.go:183)	PCDATA	$0, $19
	0x0ffa 04090 (1.go:183)	CALL	runtime.convT2E(SB)
	0x0fff 04095 (1.go:183)	MOVQ	16(SP), DX
	0x1004 04100 (1.go:183)	MOVQ	24(SP), BX
	0x1009 04105 (1.go:183)	MOVQ	DX, ""..autotmp_62+6208(SP)
	0x1011 04113 (1.go:183)	MOVQ	BX, ""..autotmp_62+6216(SP)
	0x1019 04121 (1.go:183)	LEAQ	type.int64(SB), DX
	0x1020 04128 (1.go:183)	MOVQ	DX, (SP)
	0x1024 04132 (1.go:183)	LEAQ	""..autotmp_64+56(SP), BX
	0x1029 04137 (1.go:183)	MOVQ	BX, 8(SP)
	0x102e 04142 (1.go:183)	PCDATA	$0, $19
	0x102e 04142 (1.go:183)	CALL	runtime.convT2E(SB)
	0x1033 04147 (1.go:183)	MOVQ	16(SP), DX
	0x1038 04152 (1.go:183)	MOVQ	24(SP), BX
	0x103d 04157 (1.go:183)	MOVQ	DX, ""..autotmp_62+6224(SP)
	0x1045 04165 (1.go:183)	MOVQ	BX, ""..autotmp_62+6232(SP)
	0x104d 04173 (1.go:183)	LEAQ	""..autotmp_62+6208(SP), DX
	0x1055 04181 (1.go:183)	MOVQ	DX, (SP)
	0x1059 04185 (1.go:183)	MOVQ	$2, 8(SP)
	0x1062 04194 (1.go:183)	MOVQ	$2, 16(SP)
	0x106b 04203 (1.go:183)	PCDATA	$0, $19
	0x106b 04203 (1.go:183)	CALL	fmt.Println(SB)
	0x1070 04208 (1.go:184)	LEAQ	go.string."Now: "(SB), DX
	0x1077 04215 (1.go:184)	MOVQ	DX, ""..autotmp_66+5872(SP)
	0x107f 04223 (1.go:184)	MOVQ	$5, ""..autotmp_66+5880(SP)
	0x108b 04235 (1.go:184)	PCDATA	$0, $20
	0x108b 04235 (1.go:184)	CALL	time.Now(SB)
	0x1090 04240 (1.go:184)	MOVQ	(SP), DX
	0x1094 04244 (1.go:155)	MOVQ	$-62135596800, BX
	0x109e 04254 (1.go:184)	ADDQ	BX, DX
	0x10a1 04257 (1.go:184)	MOVQ	DX, ""..autotmp_67+48(SP)
	0x10a6 04262 (1.go:184)	MOVQ	$0, ""..autotmp_65+6176(SP)
	0x10b2 04274 (1.go:184)	MOVQ	$0, ""..autotmp_65+6184(SP)
	0x10be 04286 (1.go:184)	MOVQ	$0, ""..autotmp_65+6192(SP)
	0x10ca 04298 (1.go:184)	MOVQ	$0, ""..autotmp_65+6200(SP)
	0x10d6 04310 (1.go:163)	LEAQ	type.string(SB), DX
	0x10dd 04317 (1.go:184)	MOVQ	DX, (SP)
	0x10e1 04321 (1.go:184)	LEAQ	""..autotmp_66+5872(SP), DX
	0x10e9 04329 (1.go:184)	MOVQ	DX, 8(SP)
	0x10ee 04334 (1.go:184)	PCDATA	$0, $21
	0x10ee 04334 (1.go:184)	CALL	runtime.convT2E(SB)
	0x10f3 04339 (1.go:184)	MOVQ	16(SP), DX
	0x10f8 04344 (1.go:184)	MOVQ	24(SP), BX
	0x10fd 04349 (1.go:184)	MOVQ	DX, ""..autotmp_65+6176(SP)
	0x1105 04357 (1.go:184)	MOVQ	BX, ""..autotmp_65+6184(SP)
	0x110d 04365 (1.go:183)	LEAQ	type.int64(SB), DX
	0x1114 04372 (1.go:184)	MOVQ	DX, (SP)
	0x1118 04376 (1.go:184)	LEAQ	""..autotmp_67+48(SP), DX
	0x111d 04381 (1.go:184)	MOVQ	DX, 8(SP)
	0x1122 04386 (1.go:184)	PCDATA	$0, $21
	0x1122 04386 (1.go:184)	CALL	runtime.convT2E(SB)
	0x1127 04391 (1.go:184)	MOVQ	16(SP), DX
	0x112c 04396 (1.go:184)	MOVQ	24(SP), BX
	0x1131 04401 (1.go:184)	MOVQ	DX, ""..autotmp_65+6192(SP)
	0x1139 04409 (1.go:184)	MOVQ	BX, ""..autotmp_65+6200(SP)
	0x1141 04417 (1.go:184)	LEAQ	""..autotmp_65+6176(SP), DX
	0x1149 04425 (1.go:184)	MOVQ	DX, (SP)
	0x114d 04429 (1.go:184)	MOVQ	$2, 8(SP)
	0x1156 04438 (1.go:184)	MOVQ	$2, 16(SP)
	0x115f 04447 (1.go:184)	PCDATA	$0, $21
	0x115f 04447 (1.go:184)	CALL	fmt.Println(SB)
	0x1164 04452 (1.go:186)	MOVQ	6784(SP), BP
	0x116c 04460 (1.go:186)	ADDQ	$6792, SP
	0x1173 04467 (1.go:186)	RET
	0x1174 04468 (1.go:186)	NOP
	0x1174 04468 (1.go:158)	PCDATA	$0, $-1
	0x1174 04468 (1.go:158)	CALL	runtime.morestack_noctxt(SB)
	0x1179 04473 (1.go:158)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 8b 71 10 48 81 fe de fa ff ff 0f 84 53 11 00  H.q.H........S..
	0x0020 00 48 8d 84 24 70 03 00 00 48 29 f0 48 3d 78 1d  .H..$p...H).H=x.
	0x0030 00 00 0f 86 3c 11 00 00 48 81 ec 88 1a 00 00 48  ....<...H......H
	0x0040 89 ac 24 80 1a 00 00 48 8d ac 24 80 1a 00 00 48  ..$....H..$....H
	0x0050 8d 7c 24 60 48 c7 c1 d2 02 00 00 31 c0 f3 48 ab  .|$`H......1..H.
	0x0060 48 8d 54 24 60 48 89 14 24 e8 00 00 00 00 48 8d  H.T$`H..$.....H.
	0x0070 15 00 00 00 00 48 89 94 24 10 18 00 00 48 c7 84  .....H..$....H..
	0x0080 24 18 18 00 00 07 00 00 00 48 c7 84 24 60 1a 00  $........H..$`..
	0x0090 00 00 00 00 00 48 c7 84 24 68 1a 00 00 00 00 00  .....H..$h......
	0x00a0 00 48 c7 84 24 70 1a 00 00 00 00 00 00 48 c7 84  .H..$p.......H..
	0x00b0 24 78 1a 00 00 00 00 00 00 48 8d 15 00 00 00 00  $x.......H......
	0x00c0 48 89 14 24 48 8d 9c 24 10 18 00 00 48 89 5c 24  H..$H..$....H.\$
	0x00d0 08 e8 00 00 00 00 48 8b 54 24 18 48 8b 5c 24 10  ......H.T$.H.\$.
	0x00e0 48 89 9c 24 60 1a 00 00 48 89 94 24 68 1a 00 00  H..$`...H..$h...
	0x00f0 48 8d 15 00 00 00 00 48 89 14 24 48 8d 5c 24 60  H......H..$H.\$`
	0x0100 48 89 5c 24 08 e8 00 00 00 00 48 8b 54 24 10 48  H.\$......H.T$.H
	0x0110 8b 5c 24 18 48 89 94 24 70 1a 00 00 48 89 9c 24  .\$.H..$p...H..$
	0x0120 78 1a 00 00 48 8d 94 24 60 1a 00 00 48 89 14 24  x...H..$`...H..$
	0x0130 48 c7 44 24 08 02 00 00 00 48 c7 44 24 10 02 00  H.D$.....H.D$...
	0x0140 00 00 e8 00 00 00 00 48 8d 15 00 00 00 00 48 89  .......H......H.
	0x0150 94 24 00 18 00 00 48 c7 84 24 08 18 00 00 0a 00  .$....H..$......
	0x0160 00 00 48 c7 84 24 40 1a 00 00 00 00 00 00 48 c7  ..H..$@.......H.
	0x0170 84 24 48 1a 00 00 00 00 00 00 48 c7 84 24 50 1a  .$H.......H..$P.
	0x0180 00 00 00 00 00 00 48 c7 84 24 58 1a 00 00 00 00  ......H..$X.....
	0x0190 00 00 48 8d 15 00 00 00 00 48 89 14 24 48 8d 9c  ..H......H..$H..
	0x01a0 24 00 18 00 00 48 89 5c 24 08 e8 00 00 00 00 48  $....H.\$......H
	0x01b0 8b 54 24 18 48 8b 5c 24 10 48 89 9c 24 40 1a 00  .T$.H.\$.H..$@..
	0x01c0 00 48 89 94 24 48 1a 00 00 48 8d 15 00 00 00 00  .H..$H...H......
	0x01d0 48 89 14 24 48 8d 94 24 30 11 00 00 48 89 54 24  H..$H..$0...H.T$
	0x01e0 08 e8 00 00 00 00 48 8b 54 24 10 48 8b 5c 24 18  ......H.T$.H.\$.
	0x01f0 48 89 94 24 50 1a 00 00 48 89 9c 24 58 1a 00 00  H..$P...H..$X...
	0x0200 48 8d 94 24 40 1a 00 00 48 89 14 24 48 c7 44 24  H..$@...H..$H.D$
	0x0210 08 02 00 00 00 48 c7 44 24 10 02 00 00 00 e8 00  .....H.D$.......
	0x0220 00 00 00 48 8d 15 00 00 00 00 48 89 94 24 f0 17  ...H......H..$..
	0x0230 00 00 48 c7 84 24 f8 17 00 00 07 00 00 00 48 c7  ..H..$........H.
	0x0240 84 24 20 1a 00 00 00 00 00 00 48 c7 84 24 28 1a  .$ .......H..$(.
	0x0250 00 00 00 00 00 00 48 c7 84 24 30 1a 00 00 00 00  ......H..$0.....
	0x0260 00 00 48 c7 84 24 38 1a 00 00 00 00 00 00 48 8d  ..H..$8.......H.
	0x0270 15 00 00 00 00 48 89 14 24 48 8d 9c 24 f0 17 00  .....H..$H..$...
	0x0280 00 48 89 5c 24 08 e8 00 00 00 00 48 8b 54 24 10  .H.\$......H.T$.
	0x0290 48 8b 5c 24 18 48 89 94 24 20 1a 00 00 48 89 9c  H.\$.H..$ ...H..
	0x02a0 24 28 1a 00 00 48 8d 15 00 00 00 00 48 89 14 24  $(...H......H..$
	0x02b0 48 8d 9c 24 88 00 00 00 48 89 5c 24 08 e8 00 00  H..$....H.\$....
	0x02c0 00 00 48 8b 54 24 18 48 8b 5c 24 10 48 89 9c 24  ..H.T$.H.\$.H..$
	0x02d0 30 1a 00 00 48 89 94 24 38 1a 00 00 48 8d 94 24  0...H..$8...H..$
	0x02e0 20 1a 00 00 48 89 14 24 48 c7 44 24 08 02 00 00   ...H..$H.D$....
	0x02f0 00 48 c7 44 24 10 02 00 00 00 e8 00 00 00 00 48  .H.D$..........H
	0x0300 8d 15 00 00 00 00 48 89 94 24 e0 17 00 00 48 c7  ......H..$....H.
	0x0310 84 24 e8 17 00 00 0b 00 00 00 48 c7 84 24 00 1a  .$........H..$..
	0x0320 00 00 00 00 00 00 48 c7 84 24 08 1a 00 00 00 00  ......H..$......
	0x0330 00 00 48 c7 84 24 10 1a 00 00 00 00 00 00 48 c7  ..H..$........H.
	0x0340 84 24 18 1a 00 00 00 00 00 00 48 8d 15 00 00 00  .$........H.....
	0x0350 00 48 89 14 24 48 8d 9c 24 e0 17 00 00 48 89 5c  .H..$H..$....H.\
	0x0360 24 08 e8 00 00 00 00 48 8b 54 24 10 48 8b 5c 24  $......H.T$.H.\$
	0x0370 18 48 89 94 24 00 1a 00 00 48 89 9c 24 08 1a 00  .H..$....H..$...
	0x0380 00 48 8d 15 00 00 00 00 48 89 14 24 48 8d 9c 24  .H......H..$H..$
	0x0390 90 00 00 00 48 89 5c 24 08 e8 00 00 00 00 48 8b  ....H.\$......H.
	0x03a0 54 24 10 48 8b 5c 24 18 48 89 94 24 10 1a 00 00  T$.H.\$.H..$....
	0x03b0 48 89 9c 24 18 1a 00 00 48 8d 94 24 00 1a 00 00  H..$....H..$....
	0x03c0 48 89 14 24 48 c7 44 24 08 02 00 00 00 48 c7 44  H..$H.D$.....H.D
	0x03d0 24 10 02 00 00 00 e8 00 00 00 00 48 8d 15 00 00  $..........H....
	0x03e0 00 00 48 89 94 24 d0 17 00 00 48 c7 84 24 d8 17  ..H..$....H..$..
	0x03f0 00 00 0a 00 00 00 48 c7 84 24 e0 19 00 00 00 00  ......H..$......
	0x0400 00 00 48 c7 84 24 e8 19 00 00 00 00 00 00 48 c7  ..H..$........H.
	0x0410 84 24 f0 19 00 00 00 00 00 00 48 c7 84 24 f8 19  .$........H..$..
	0x0420 00 00 00 00 00 00 48 8d 15 00 00 00 00 48 89 14  ......H......H..
	0x0430 24 48 8d 9c 24 d0 17 00 00 48 89 5c 24 08 e8 00  $H..$....H.\$...
	0x0440 00 00 00 48 8b 54 24 18 48 8b 5c 24 10 48 89 9c  ...H.T$.H.\$.H..
	0x0450 24 e0 19 00 00 48 89 94 24 e8 19 00 00 48 8d 15  $....H..$....H..
	0x0460 00 00 00 00 48 89 14 24 48 8d 9c 24 a0 00 00 00  ....H..$H..$....
	0x0470 48 89 5c 24 08 e8 00 00 00 00 48 8b 54 24 10 48  H.\$......H.T$.H
	0x0480 8b 5c 24 18 48 89 94 24 f0 19 00 00 48 89 9c 24  .\$.H..$....H..$
	0x0490 f8 19 00 00 48 8d 94 24 e0 19 00 00 48 89 14 24  ....H..$....H..$
	0x04a0 48 c7 44 24 08 02 00 00 00 48 c7 44 24 10 02 00  H.D$.....H.D$...
	0x04b0 00 00 e8 00 00 00 00 48 8d 15 00 00 00 00 48 89  .......H......H.
	0x04c0 94 24 c0 17 00 00 48 c7 84 24 c8 17 00 00 0d 00  .$....H..$......
	0x04d0 00 00 48 c7 84 24 c0 19 00 00 00 00 00 00 48 c7  ..H..$........H.
	0x04e0 84 24 c8 19 00 00 00 00 00 00 48 c7 84 24 d0 19  .$........H..$..
	0x04f0 00 00 00 00 00 00 48 c7 84 24 d8 19 00 00 00 00  ......H..$......
	0x0500 00 00 48 8d 15 00 00 00 00 48 89 14 24 48 8d 9c  ..H......H..$H..
	0x0510 24 c0 17 00 00 48 89 5c 24 08 e8 00 00 00 00 48  $....H.\$......H
	0x0520 8b 54 24 18 48 8b 5c 24 10 48 89 9c 24 c0 19 00  .T$.H.\$.H..$...
	0x0530 00 48 89 94 24 c8 19 00 00 48 8d 15 00 00 00 00  .H..$....H......
	0x0540 48 89 14 24 48 8d 9c 24 b8 00 00 00 48 89 5c 24  H..$H..$....H.\$
	0x0550 08 e8 00 00 00 00 48 8b 54 24 18 48 8b 5c 24 10  ......H.T$.H.\$.
	0x0560 48 89 9c 24 d0 19 00 00 48 89 94 24 d8 19 00 00  H..$....H..$....
	0x0570 48 8d 94 24 c0 19 00 00 48 89 14 24 48 c7 44 24  H..$....H..$H.D$
	0x0580 08 02 00 00 00 48 c7 44 24 10 02 00 00 00 e8 00  .....H.D$.......
	0x0590 00 00 00 48 8d 15 00 00 00 00 48 89 94 24 b0 17  ...H......H..$..
	0x05a0 00 00 48 c7 84 24 b8 17 00 00 0e 00 00 00 48 c7  ..H..$........H.
	0x05b0 84 24 a0 19 00 00 00 00 00 00 48 c7 84 24 a8 19  .$........H..$..
	0x05c0 00 00 00 00 00 00 48 c7 84 24 b0 19 00 00 00 00  ......H..$......
	0x05d0 00 00 48 c7 84 24 b8 19 00 00 00 00 00 00 48 8d  ..H..$........H.
	0x05e0 15 00 00 00 00 48 89 14 24 48 8d 9c 24 b0 17 00  .....H..$H..$...
	0x05f0 00 48 89 5c 24 08 e8 00 00 00 00 48 8b 54 24 18  .H.\$......H.T$.
	0x0600 48 8b 5c 24 10 48 89 9c 24 a0 19 00 00 48 89 94  H.\$.H..$....H..
	0x0610 24 a8 19 00 00 48 8d 15 00 00 00 00 48 89 14 24  $....H......H..$
	0x0620 48 8d 9c 24 b0 00 00 00 48 89 5c 24 08 e8 00 00  H..$....H.\$....
	0x0630 00 00 48 8b 54 24 18 48 8b 5c 24 10 48 89 9c 24  ..H.T$.H.\$.H..$
	0x0640 b0 19 00 00 48 89 94 24 b8 19 00 00 48 8d 94 24  ....H..$....H..$
	0x0650 a0 19 00 00 48 89 14 24 48 c7 44 24 08 02 00 00  ....H..$H.D$....
	0x0660 00 48 c7 44 24 10 02 00 00 00 e8 00 00 00 00 48  .H.D$..........H
	0x0670 8d 15 00 00 00 00 48 89 94 24 a0 17 00 00 48 c7  ......H..$....H.
	0x0680 84 24 a8 17 00 00 0c 00 00 00 48 c7 84 24 80 19  .$........H..$..
	0x0690 00 00 00 00 00 00 48 c7 84 24 88 19 00 00 00 00  ......H..$......
	0x06a0 00 00 48 c7 84 24 90 19 00 00 00 00 00 00 48 c7  ..H..$........H.
	0x06b0 84 24 98 19 00 00 00 00 00 00 48 8d 15 00 00 00  .$........H.....
	0x06c0 00 48 89 14 24 48 8d 9c 24 a0 17 00 00 48 89 5c  .H..$H..$....H.\
	0x06d0 24 08 e8 00 00 00 00 48 8b 54 24 18 48 8b 5c 24  $......H.T$.H.\$
	0x06e0 10 48 89 9c 24 80 19 00 00 48 89 94 24 88 19 00  .H..$....H..$...
	0x06f0 00 48 8d 15 00 00 00 00 48 89 14 24 48 8d 5c 24  .H......H..$H.\$
	0x0700 68 48 89 5c 24 08 e8 00 00 00 00 48 8b 54 24 18  hH.\$......H.T$.
	0x0710 48 8b 5c 24 10 48 89 9c 24 90 19 00 00 48 89 94  H.\$.H..$....H..
	0x0720 24 98 19 00 00 48 8d 94 24 80 19 00 00 48 89 14  $....H..$....H..
	0x0730 24 48 c7 44 24 08 02 00 00 00 48 c7 44 24 10 02  $H.D$.....H.D$..
	0x0740 00 00 00 e8 00 00 00 00 48 8d 15 00 00 00 00 48  ........H......H
	0x0750 89 94 24 90 17 00 00 48 c7 84 24 98 17 00 00 0c  ..$....H..$.....
	0x0760 00 00 00 48 c7 84 24 60 19 00 00 00 00 00 00 48  ...H..$`.......H
	0x0770 c7 84 24 68 19 00 00 00 00 00 00 48 c7 84 24 70  ..$h.......H..$p
	0x0780 19 00 00 00 00 00 00 48 c7 84 24 78 19 00 00 00  .......H..$x....
	0x0790 00 00 00 48 8d 15 00 00 00 00 48 89 14 24 48 8d  ...H......H..$H.
	0x07a0 9c 24 90 17 00 00 48 89 5c 24 08 e8 00 00 00 00  .$....H.\$......
	0x07b0 48 8b 54 24 18 48 8b 5c 24 10 48 89 9c 24 60 19  H.T$.H.\$.H..$`.
	0x07c0 00 00 48 89 94 24 68 19 00 00 48 8d 15 00 00 00  ..H..$h...H.....
	0x07d0 00 48 89 14 24 48 8d 9c 24 c0 00 00 00 48 89 5c  .H..$H..$....H.\
	0x07e0 24 08 e8 00 00 00 00 48 8b 54 24 10 48 8b 5c 24  $......H.T$.H.\$
	0x07f0 18 48 89 94 24 70 19 00 00 48 89 9c 24 78 19 00  .H..$p...H..$x..
	0x0800 00 48 8d 94 24 60 19 00 00 48 89 14 24 48 c7 44  .H..$`...H..$H.D
	0x0810 24 08 02 00 00 00 48 c7 44 24 10 02 00 00 00 e8  $.....H.D$......
	0x0820 00 00 00 00 48 8d 15 00 00 00 00 48 89 94 24 80  ....H......H..$.
	0x0830 17 00 00 48 c7 84 24 88 17 00 00 09 00 00 00 48  ...H..$........H
	0x0840 c7 84 24 40 19 00 00 00 00 00 00 48 c7 84 24 48  ..$@.......H..$H
	0x0850 19 00 00 00 00 00 00 48 c7 84 24 50 19 00 00 00  .......H..$P....
	0x0860 00 00 00 48 c7 84 24 58 19 00 00 00 00 00 00 48  ...H..$X.......H
	0x0870 8d 15 00 00 00 00 48 89 14 24 48 8d 9c 24 80 17  ......H..$H..$..
	0x0880 00 00 48 89 5c 24 08 e8 00 00 00 00 48 8b 54 24  ..H.\$......H.T$
	0x0890 18 48 8b 5c 24 10 48 89 9c 24 40 19 00 00 48 89  .H.\$.H..$@...H.
	0x08a0 94 24 48 19 00 00 48 8d 15 00 00 00 00 48 89 14  .$H...H......H..
	0x08b0 24 48 8d 5c 24 78 48 89 5c 24 08 e8 00 00 00 00  $H.\$xH.\$......
	0x08c0 48 8b 54 24 10 48 8b 5c 24 18 48 89 94 24 50 19  H.T$.H.\$.H..$P.
	0x08d0 00 00 48 89 9c 24 58 19 00 00 48 8d 94 24 40 19  ..H..$X...H..$@.
	0x08e0 00 00 48 89 14 24 48 c7 44 24 08 02 00 00 00 48  ..H..$H.D$.....H
	0x08f0 c7 44 24 10 02 00 00 00 e8 00 00 00 00 48 8d 15  .D$..........H..
	0x0900 00 00 00 00 48 89 94 24 70 17 00 00 48 c7 84 24  ....H..$p...H..$
	0x0910 78 17 00 00 08 00 00 00 48 8b 84 24 10 01 00 00  x.......H..$....
	0x0920 48 89 44 24 58 48 ba 2f 4b 69 6d 82 be e0 12 48  H.D$XH./Kim....H
	0x0930 f7 e2 48 8b 5c 24 58 48 01 d3 48 d1 db 48 c1 eb  ..H.\$XH..H..H..
	0x0940 1d 48 89 5c 24 50 48 c7 84 24 20 19 00 00 00 00  .H.\$PH..$ .....
	0x0950 00 00 48 c7 84 24 28 19 00 00 00 00 00 00 48 c7  ..H..$(.......H.
	0x0960 84 24 30 19 00 00 00 00 00 00 48 c7 84 24 38 19  .$0.......H..$8.
	0x0970 00 00 00 00 00 00 48 8d 15 00 00 00 00 48 89 14  ......H......H..
	0x0980 24 48 8d 9c 24 70 17 00 00 48 89 5c 24 08 e8 00  $H..$p...H.\$...
	0x0990 00 00 00 48 8b 54 24 10 48 8b 5c 24 18 48 89 94  ...H.T$.H.\$.H..
	0x09a0 24 20 19 00 00 48 89 9c 24 28 19 00 00 48 8d 15  $ ...H..$(...H..
	0x09b0 00 00 00 00 48 89 14 24 48 8d 5c 24 50 48 89 5c  ....H..$H.\$PH.\
	0x09c0 24 08 e8 00 00 00 00 48 8b 54 24 10 48 8b 5c 24  $......H.T$.H.\$
	0x09d0 18 48 89 94 24 30 19 00 00 48 89 9c 24 38 19 00  .H..$0...H..$8..
	0x09e0 00 48 8d 94 24 20 19 00 00 48 89 14 24 48 c7 44  .H..$ ...H..$H.D
	0x09f0 24 08 02 00 00 00 48 c7 44 24 10 02 00 00 00 e8  $.....H.D$......
	0x0a00 00 00 00 00 48 8d 15 00 00 00 00 48 89 94 24 60  ....H......H..$`
	0x0a10 17 00 00 48 c7 84 24 68 17 00 00 07 00 00 00 48  ...H..$h.......H
	0x0a20 c7 84 24 00 19 00 00 00 00 00 00 48 c7 84 24 08  ..$........H..$.
	0x0a30 19 00 00 00 00 00 00 48 c7 84 24 10 19 00 00 00  .......H..$.....
	0x0a40 00 00 00 48 c7 84 24 18 19 00 00 00 00 00 00 48  ...H..$........H
	0x0a50 8d 15 00 00 00 00 48 89 14 24 48 8d 9c 24 60 17  ......H..$H..$`.
	0x0a60 00 00 48 89 5c 24 08 e8 00 00 00 00 48 8b 54 24  ..H.\$......H.T$
	0x0a70 18 48 8b 5c 24 10 48 89 9c 24 00 19 00 00 48 89  .H.\$.H..$....H.
	0x0a80 94 24 08 19 00 00 48 8d 15 00 00 00 00 48 89 14  .$....H......H..
	0x0a90 24 48 8d 94 24 20 11 00 00 48 89 54 24 08 e8 00  $H..$ ...H.T$...
	0x0aa0 00 00 00 48 8b 54 24 18 48 8b 5c 24 10 48 89 9c  ...H.T$.H.\$.H..
	0x0ab0 24 10 19 00 00 48 89 94 24 18 19 00 00 48 8d 94  $....H..$....H..
	0x0ac0 24 00 19 00 00 48 89 14 24 48 c7 44 24 08 02 00  $....H..$H.D$...
	0x0ad0 00 00 48 c7 44 24 10 02 00 00 00 e8 00 00 00 00  ..H.D$..........
	0x0ae0 48 8d 15 00 00 00 00 48 89 94 24 50 17 00 00 48  H......H..$P...H
	0x0af0 c7 84 24 58 17 00 00 08 00 00 00 48 8b 84 24 08  ..$X.......H..$.
	0x0b00 01 00 00 48 89 44 24 58 48 ba 2f 4b 69 6d 82 be  ...H.D$XH./Kim..
	0x0b10 e0 12 48 f7 e2 48 8b 5c 24 58 48 01 d3 48 d1 db  ..H..H.\$XH..H..
	0x0b20 48 c1 eb 1d 48 89 5c 24 48 48 c7 84 24 e0 18 00  H...H.\$HH..$...
	0x0b30 00 00 00 00 00 48 c7 84 24 e8 18 00 00 00 00 00  .....H..$.......
	0x0b40 00 48 c7 84 24 f0 18 00 00 00 00 00 00 48 c7 84  .H..$........H..
	0x0b50 24 f8 18 00 00 00 00 00 00 48 8d 15 00 00 00 00  $........H......
	0x0b60 48 89 14 24 48 8d 9c 24 50 17 00 00 48 89 5c 24  H..$H..$P...H.\$
	0x0b70 08 e8 00 00 00 00 48 8b 54 24 10 48 8b 5c 24 18  ......H.T$.H.\$.
	0x0b80 48 89 94 24 e0 18 00 00 48 89 9c 24 e8 18 00 00  H..$....H..$....
	0x0b90 48 8d 15 00 00 00 00 48 89 14 24 48 8d 5c 24 48  H......H..$H.\$H
	0x0ba0 48 89 5c 24 08 e8 00 00 00 00 48 8b 54 24 10 48  H.\$......H.T$.H
	0x0bb0 8b 5c 24 18 48 89 94 24 f0 18 00 00 48 89 9c 24  .\$.H..$....H..$
	0x0bc0 f8 18 00 00 48 8d 94 24 e0 18 00 00 48 89 14 24  ....H..$....H..$
	0x0bd0 48 c7 44 24 08 02 00 00 00 48 c7 44 24 10 02 00  H.D$.....H.D$...
	0x0be0 00 00 e8 00 00 00 00 48 8d 15 00 00 00 00 48 89  .......H......H.
	0x0bf0 94 24 40 17 00 00 48 c7 84 24 48 17 00 00 09 00  .$@...H..$H.....
	0x0c00 00 00 48 c7 84 24 c0 18 00 00 00 00 00 00 48 c7  ..H..$........H.
	0x0c10 84 24 c8 18 00 00 00 00 00 00 48 c7 84 24 d0 18  .$........H..$..
	0x0c20 00 00 00 00 00 00 48 c7 84 24 d8 18 00 00 00 00  ......H..$......
	0x0c30 00 00 48 8d 15 00 00 00 00 48 89 14 24 48 8d 9c  ..H......H..$H..
	0x0c40 24 40 17 00 00 48 89 5c 24 08 e8 00 00 00 00 48  $@...H.\$......H
	0x0c50 8b 54 24 10 48 8b 5c 24 18 48 89 94 24 c0 18 00  .T$.H.\$.H..$...
	0x0c60 00 48 89 9c 24 c8 18 00 00 48 8d 15 00 00 00 00  .H..$....H......
	0x0c70 48 89 14 24 48 8d 9c 24 80 00 00 00 48 89 5c 24  H..$H..$....H.\$
	0x0c80 08 e8 00 00 00 00 48 8b 54 24 10 48 8b 5c 24 18  ......H.T$.H.\$.
	0x0c90 48 89 94 24 d0 18 00 00 48 89 9c 24 d8 18 00 00  H..$....H..$....
	0x0ca0 48 8d 94 24 c0 18 00 00 48 89 14 24 48 c7 44 24  H..$....H..$H.D$
	0x0cb0 08 02 00 00 00 48 c7 44 24 10 02 00 00 00 e8 00  .....H.D$.......
	0x0cc0 00 00 00 48 8d 15 00 00 00 00 48 89 94 24 30 17  ...H......H..$0.
	0x0cd0 00 00 48 c7 84 24 38 17 00 00 0e 00 00 00 48 8b  ..H..$8.......H.
	0x0ce0 94 24 18 01 00 00 48 b8 db 34 b6 d7 82 de 1b 43  .$....H..4.....C
	0x0cf0 48 f7 e2 48 c1 ea 12 48 89 54 24 40 48 c7 84 24  H..H...H.T$@H..$
	0x0d00 a0 18 00 00 00 00 00 00 48 c7 84 24 a8 18 00 00  ........H..$....
	0x0d10 00 00 00 00 48 c7 84 24 b0 18 00 00 00 00 00 00  ....H..$........
	0x0d20 48 c7 84 24 b8 18 00 00 00 00 00 00 48 8d 15 00  H..$........H...
	0x0d30 00 00 00 48 89 14 24 48 8d 9c 24 30 17 00 00 48  ...H..$H..$0...H
	0x0d40 89 5c 24 08 e8 00 00 00 00 48 8b 54 24 10 48 8b  .\$......H.T$.H.
	0x0d50 5c 24 18 48 89 94 24 a0 18 00 00 48 89 9c 24 a8  \$.H..$....H..$.
	0x0d60 18 00 00 48 8d 15 00 00 00 00 48 89 14 24 48 8d  ...H......H..$H.
	0x0d70 5c 24 40 48 89 5c 24 08 e8 00 00 00 00 48 8b 54  \$@H.\$......H.T
	0x0d80 24 10 48 8b 5c 24 18 48 89 94 24 b0 18 00 00 48  $.H.\$.H..$....H
	0x0d90 89 9c 24 b8 18 00 00 48 8d 94 24 a0 18 00 00 48  ..$....H..$....H
	0x0da0 89 14 24 48 c7 44 24 08 02 00 00 00 48 c7 44 24  ..$H.D$.....H.D$
	0x0db0 10 02 00 00 00 e8 00 00 00 00 48 8d 15 00 00 00  ..........H.....
	0x0dc0 00 48 89 94 24 20 17 00 00 48 c7 84 24 28 17 00  .H..$ ...H..$(..
	0x0dd0 00 0f 00 00 00 48 c7 84 24 80 18 00 00 00 00 00  .....H..$.......
	0x0de0 00 48 c7 84 24 88 18 00 00 00 00 00 00 48 c7 84  .H..$........H..
	0x0df0 24 90 18 00 00 00 00 00 00 48 c7 84 24 98 18 00  $........H..$...
	0x0e00 00 00 00 00 00 48 8d 15 00 00 00 00 48 89 14 24  .....H......H..$
	0x0e10 48 8d 9c 24 20 17 00 00 48 89 5c 24 08 e8 00 00  H..$ ...H.\$....
	0x0e20 00 00 48 8b 54 24 10 48 8b 5c 24 18 48 89 94 24  ..H.T$.H.\$.H..$
	0x0e30 80 18 00 00 48 89 9c 24 88 18 00 00 48 8d 15 00  ....H..$....H...
	0x0e40 00 00 00 48 89 14 24 48 8d 94 24 28 11 00 00 48  ...H..$H..$(...H
	0x0e50 89 54 24 08 e8 00 00 00 00 48 8b 54 24 10 48 8b  .T$......H.T$.H.
	0x0e60 5c 24 18 48 89 94 24 90 18 00 00 48 89 9c 24 98  \$.H..$....H..$.
	0x0e70 18 00 00 48 8d 94 24 80 18 00 00 48 89 14 24 48  ...H..$....H..$H
	0x0e80 c7 44 24 08 02 00 00 00 48 c7 44 24 10 02 00 00  .D$.....H.D$....
	0x0e90 00 e8 00 00 00 00 48 8d 15 00 00 00 00 48 89 94  ......H......H..
	0x0ea0 24 10 17 00 00 48 c7 84 24 18 17 00 00 07 00 00  $....H..$.......
	0x0eb0 00 48 c7 84 24 60 18 00 00 00 00 00 00 48 c7 84  .H..$`.......H..
	0x0ec0 24 68 18 00 00 00 00 00 00 48 c7 84 24 70 18 00  $h.......H..$p..
	0x0ed0 00 00 00 00 00 48 c7 84 24 78 18 00 00 00 00 00  .....H..$x......
	0x0ee0 00 48 8d 15 00 00 00 00 48 89 14 24 48 8d 9c 24  .H......H..$H..$
	0x0ef0 10 17 00 00 48 89 5c 24 08 e8 00 00 00 00 48 8b  ....H.\$......H.
	0x0f00 54 24 10 48 8b 5c 24 18 48 89 94 24 60 18 00 00  T$.H.\$.H..$`...
	0x0f10 48 89 9c 24 68 18 00 00 48 8d 15 00 00 00 00 48  H..$h...H......H
	0x0f20 89 14 24 48 8d 94 24 f8 00 00 00 48 89 54 24 08  ..$H..$....H.T$.
	0x0f30 e8 00 00 00 00 48 8b 54 24 10 48 8b 5c 24 18 48  .....H.T$.H.\$.H
	0x0f40 89 94 24 70 18 00 00 48 89 9c 24 78 18 00 00 48  ..$p...H..$x...H
	0x0f50 8d 94 24 60 18 00 00 48 89 14 24 48 c7 44 24 08  ..$`...H..$H.D$.
	0x0f60 02 00 00 00 48 c7 44 24 10 02 00 00 00 e8 00 00  ....H.D$........
	0x0f70 00 00 48 8d 15 00 00 00 00 48 89 94 24 00 17 00  ..H......H..$...
	0x0f80 00 48 c7 84 24 08 17 00 00 06 00 00 00 e8 00 00  .H..$...........
	0x0f90 00 00 48 8b 15 00 00 00 00 48 8b 1c 24 48 be 00  ..H......H..$H..
	0x0fa0 09 6e 88 f1 ff ff ff 48 01 f3 48 29 d3 48 89 5c  .n.....H..H).H.\
	0x0fb0 24 38 48 c7 84 24 40 18 00 00 00 00 00 00 48 c7  $8H..$@.......H.
	0x0fc0 84 24 48 18 00 00 00 00 00 00 48 c7 84 24 50 18  .$H.......H..$P.
	0x0fd0 00 00 00 00 00 00 48 c7 84 24 58 18 00 00 00 00  ......H..$X.....
	0x0fe0 00 00 48 8d 15 00 00 00 00 48 89 14 24 48 8d 9c  ..H......H..$H..
	0x0ff0 24 00 17 00 00 48 89 5c 24 08 e8 00 00 00 00 48  $....H.\$......H
	0x1000 8b 54 24 10 48 8b 5c 24 18 48 89 94 24 40 18 00  .T$.H.\$.H..$@..
	0x1010 00 48 89 9c 24 48 18 00 00 48 8d 15 00 00 00 00  .H..$H...H......
	0x1020 48 89 14 24 48 8d 5c 24 38 48 89 5c 24 08 e8 00  H..$H.\$8H.\$...
	0x1030 00 00 00 48 8b 54 24 10 48 8b 5c 24 18 48 89 94  ...H.T$.H.\$.H..
	0x1040 24 50 18 00 00 48 89 9c 24 58 18 00 00 48 8d 94  $P...H..$X...H..
	0x1050 24 40 18 00 00 48 89 14 24 48 c7 44 24 08 02 00  $@...H..$H.D$...
	0x1060 00 00 48 c7 44 24 10 02 00 00 00 e8 00 00 00 00  ..H.D$..........
	0x1070 48 8d 15 00 00 00 00 48 89 94 24 f0 16 00 00 48  H......H..$....H
	0x1080 c7 84 24 f8 16 00 00 05 00 00 00 e8 00 00 00 00  ..$.............
	0x1090 48 8b 14 24 48 bb 00 09 6e 88 f1 ff ff ff 48 01  H..$H...n.....H.
	0x10a0 da 48 89 54 24 30 48 c7 84 24 20 18 00 00 00 00  .H.T$0H..$ .....
	0x10b0 00 00 48 c7 84 24 28 18 00 00 00 00 00 00 48 c7  ..H..$(.......H.
	0x10c0 84 24 30 18 00 00 00 00 00 00 48 c7 84 24 38 18  .$0.......H..$8.
	0x10d0 00 00 00 00 00 00 48 8d 15 00 00 00 00 48 89 14  ......H......H..
	0x10e0 24 48 8d 94 24 f0 16 00 00 48 89 54 24 08 e8 00  $H..$....H.T$...
	0x10f0 00 00 00 48 8b 54 24 10 48 8b 5c 24 18 48 89 94  ...H.T$.H.\$.H..
	0x1100 24 20 18 00 00 48 89 9c 24 28 18 00 00 48 8d 15  $ ...H..$(...H..
	0x1110 00 00 00 00 48 89 14 24 48 8d 54 24 30 48 89 54  ....H..$H.T$0H.T
	0x1120 24 08 e8 00 00 00 00 48 8b 54 24 10 48 8b 5c 24  $......H.T$.H.\$
	0x1130 18 48 89 94 24 30 18 00 00 48 89 9c 24 38 18 00  .H..$0...H..$8..
	0x1140 00 48 8d 94 24 20 18 00 00 48 89 14 24 48 c7 44  .H..$ ...H..$H.D
	0x1150 24 08 02 00 00 00 48 c7 44 24 10 02 00 00 00 e8  $.....H.D$......
	0x1160 00 00 00 00 48 8b ac 24 80 1a 00 00 48 81 c4 88  ....H..$....H...
	0x1170 1a 00 00 c3 e8 00 00 00 00 e9 82 ee ff ff        ..............
	rel 12+4 t=16 TLS+0
	rel 106+4 t=8 runtime.ReadMemStats+0
	rel 113+4 t=15 go.string."Alloc: "+0
	rel 188+4 t=15 type.string+0
	rel 210+4 t=8 runtime.convT2E+0
	rel 243+4 t=15 type.uint64+0
	rel 262+4 t=8 runtime.convT2E+0
	rel 323+4 t=8 fmt.Println+0
	rel 330+4 t=15 go.string."EnableGC: "+0
	rel 405+4 t=15 type.string+0
	rel 427+4 t=8 runtime.convT2E+0
	rel 460+4 t=15 type.bool+0
	rel 482+4 t=8 runtime.convT2E+0
	rel 543+4 t=8 fmt.Println+0
	rel 550+4 t=15 go.string."Frees: "+0
	rel 625+4 t=15 type.string+0
	rel 647+4 t=8 runtime.convT2E+0
	rel 680+4 t=15 type.uint64+0
	rel 702+4 t=8 runtime.convT2E+0
	rel 763+4 t=8 fmt.Println+0
	rel 770+4 t=15 go.string."HeapAlloc: "+0
	rel 845+4 t=15 type.string+0
	rel 867+4 t=8 runtime.convT2E+0
	rel 900+4 t=15 type.uint64+0
	rel 922+4 t=8 runtime.convT2E+0
	rel 983+4 t=8 fmt.Println+0
	rel 990+4 t=15 go.string."HeapIdle: "+0
	rel 1065+4 t=15 type.string+0
	rel 1087+4 t=8 runtime.convT2E+0
	rel 1120+4 t=15 type.uint64+0
	rel 1142+4 t=8 runtime.convT2E+0
	rel 1203+4 t=8 fmt.Println+0
	rel 1210+4 t=15 go.string."HeapObjects: "+0
	rel 1285+4 t=15 type.string+0
	rel 1307+4 t=8 runtime.convT2E+0
	rel 1340+4 t=15 type.uint64+0
	rel 1362+4 t=8 runtime.convT2E+0
	rel 1423+4 t=8 fmt.Println+0
	rel 1430+4 t=15 go.string."HeapReleased: "+0
	rel 1505+4 t=15 type.string+0
	rel 1527+4 t=8 runtime.convT2E+0
	rel 1560+4 t=15 type.uint64+0
	rel 1582+4 t=8 runtime.convT2E+0
	rel 1643+4 t=8 fmt.Println+0
	rel 1650+4 t=15 go.string."TotalAlloc: "+0
	rel 1725+4 t=15 type.string+0
	rel 1747+4 t=8 runtime.convT2E+0
	rel 1780+4 t=15 type.uint64+0
	rel 1799+4 t=8 runtime.convT2E+0
	rel 1860+4 t=8 fmt.Println+0
	rel 1867+4 t=15 go.string."StackInuse: "+0
	rel 1942+4 t=15 type.string+0
	rel 1964+4 t=8 runtime.convT2E+0
	rel 1997+4 t=15 type.uint64+0
	rel 2019+4 t=8 runtime.convT2E+0
	rel 2080+4 t=8 fmt.Println+0
	rel 2087+4 t=15 go.string."Lookups: "+0
	rel 2162+4 t=15 type.string+0
	rel 2184+4 t=8 runtime.convT2E+0
	rel 2217+4 t=15 type.uint64+0
	rel 2236+4 t=8 runtime.convT2E+0
	rel 2297+4 t=8 fmt.Println+0
	rel 2304+4 t=15 go.string."LastGC: "+0
	rel 2425+4 t=15 type.string+0
	rel 2447+4 t=8 runtime.convT2E+0
	rel 2480+4 t=15 type.uint64+0
	rel 2499+4 t=8 runtime.convT2E+0
	rel 2560+4 t=8 fmt.Println+0
	rel 2567+4 t=15 go.string."NumGC: "+0
	rel 2642+4 t=15 type.string+0
	rel 2664+4 t=8 runtime.convT2E+0
	rel 2697+4 t=15 type.uint32+0
	rel 2719+4 t=8 runtime.convT2E+0
	rel 2780+4 t=8 fmt.Println+0
	rel 2787+4 t=15 go.string."NextGC: "+0
	rel 2908+4 t=15 type.string+0
	rel 2930+4 t=8 runtime.convT2E+0
	rel 2963+4 t=15 type.uint64+0
	rel 2982+4 t=8 runtime.convT2E+0
	rel 3043+4 t=8 fmt.Println+0
	rel 3050+4 t=15 go.string."Mallocs: "+0
	rel 3125+4 t=15 type.string+0
	rel 3147+4 t=8 runtime.convT2E+0
	rel 3180+4 t=15 type.uint64+0
	rel 3202+4 t=8 runtime.convT2E+0
	rel 3263+4 t=8 fmt.Println+0
	rel 3270+4 t=15 go.string."PauseTotalNs: "+0
	rel 3375+4 t=15 type.string+0
	rel 3397+4 t=8 runtime.convT2E+0
	rel 3430+4 t=15 type.uint64+0
	rel 3449+4 t=8 runtime.convT2E+0
	rel 3510+4 t=8 fmt.Println+0
	rel 3517+4 t=15 go.string."GCCPUFraction: "+0
	rel 3592+4 t=15 type.string+0
	rel 3614+4 t=8 runtime.convT2E+0
	rel 3647+4 t=15 type.float64+0
	rel 3669+4 t=8 runtime.convT2E+0
	rel 3730+4 t=8 fmt.Println+0
	rel 3737+4 t=15 go.string."GCSys: "+0
	rel 3812+4 t=15 type.string+0
	rel 3834+4 t=8 runtime.convT2E+0
	rel 3867+4 t=15 type.uint64+0
	rel 3889+4 t=8 runtime.convT2E+0
	rel 3950+4 t=8 fmt.Println+0
	rel 3957+4 t=15 go.string."Time: "+0
	rel 3982+4 t=8 time.Now+0
	rel 3989+4 t=15 "".t1+0
	rel 4069+4 t=15 type.string+0
	rel 4091+4 t=8 runtime.convT2E+0
	rel 4124+4 t=15 type.int64+0
	rel 4143+4 t=8 runtime.convT2E+0
	rel 4204+4 t=8 fmt.Println+0
	rel 4211+4 t=15 go.string."Now: "+0
	rel 4236+4 t=8 time.Now+0
	rel 4313+4 t=15 type.string+0
	rel 4335+4 t=8 runtime.convT2E+0
	rel 4368+4 t=15 type.int64+0
	rel 4387+4 t=8 runtime.convT2E+0
	rel 4448+4 t=8 fmt.Println+0
	rel 4469+4 t=8 runtime.morestack_noctxt+0
"".main.func1 t=1 size=116 args=0x8 locals=0x20
	0x0000 00000 (1.go:39)	TEXT	"".main.func1(SB), $32-8
	0x0000 00000 (1.go:39)	MOVQ	TLS, CX
	0x0009 00009 (1.go:39)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:39)	CMPQ	SP, 16(CX)
	0x0014 00020 (1.go:39)	JLS	109
	0x0016 00022 (1.go:39)	SUBQ	$32, SP
	0x001a 00026 (1.go:39)	MOVQ	BP, 24(SP)
	0x001f 00031 (1.go:39)	LEAQ	24(SP), BP
	0x0024 00036 (1.go:39)	FUNCDATA	$0, gclocals·263043c8f03e3241528dfae4e2812ef4(SB)
	0x0024 00036 (1.go:39)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0024 00036 (1.go:39)	MOVQ	8(DX), AX
	0x0028 00040 (1.go:39)	MOVQ	AX, "".&mm+16(SP)
	0x002d 00045 (1.go:40)	MOVQ	AX, (SP)
	0x0031 00049 (1.go:40)	PCDATA	$0, $1
	0x0031 00049 (1.go:40)	CALL	sync.(*RWMutex).Lock(SB)
	0x0036 00054 (1.go:41)	MOVQ	"".&mm+16(SP), AX
	0x003b 00059 (1.go:41)	MOVQ	24(AX), CX
	0x003f 00063 (1.go:41)	TESTQ	CX, CX
	0x0042 00066 (1.go:41)	JEQ	$0, 105
	0x0044 00068 (1.go:41)	MOVQ	(CX), CX
	0x0047 00071 (1.go:41)	MOVQ	CX, "".l+8(SP)
	0x004c 00076 (1.go:42)	MOVQ	AX, (SP)
	0x0050 00080 (1.go:42)	PCDATA	$0, $0
	0x0050 00080 (1.go:42)	CALL	sync.(*RWMutex).Unlock(SB)
	0x0055 00085 (1.go:43)	MOVQ	"".l+8(SP), AX
	0x005a 00090 (1.go:43)	MOVQ	AX, "".~r0+40(FP)
	0x005f 00095 (1.go:43)	MOVQ	24(SP), BP
	0x0064 00100 (1.go:43)	ADDQ	$32, SP
	0x0068 00104 (1.go:43)	RET
	0x0069 00105 (1.go:39)	MOVQ	$0, CX
	0x006b 00107 (1.go:41)	JMP	71
	0x006d 00109 (1.go:41)	NOP
	0x006d 00109 (1.go:39)	PCDATA	$0, $-1
	0x006d 00109 (1.go:39)	CALL	runtime.morestack(SB)
	0x0072 00114 (1.go:39)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 76 57 48 83 ec 20 48 89 6c 24 18 48  H;a.vWH.. H.l$.H
	0x0020 8d 6c 24 18 48 8b 42 08 48 89 44 24 10 48 89 04  .l$.H.B.H.D$.H..
	0x0030 24 e8 00 00 00 00 48 8b 44 24 10 48 8b 48 18 48  $.....H.D$.H.H.H
	0x0040 85 c9 74 25 48 8b 09 48 89 4c 24 08 48 89 04 24  ..t%H..H.L$.H..$
	0x0050 e8 00 00 00 00 48 8b 44 24 08 48 89 44 24 28 48  .....H.D$.H.D$(H
	0x0060 8b 6c 24 18 48 83 c4 20 c3 31 c9 eb da e8 00 00  .l$.H.. .1......
	0x0070 00 00 eb 8c                                      ....
	rel 12+4 t=16 TLS+0
	rel 50+4 t=8 sync.(*RWMutex).Lock+0
	rel 81+4 t=8 sync.(*RWMutex).Unlock+0
	rel 110+4 t=8 runtime.morestack+0
"".main.func2 t=1 size=1315 args=0x30 locals=0xd8
	0x0000 00000 (1.go:48)	TEXT	"".main.func2(SB), $216-48
	0x0000 00000 (1.go:48)	MOVQ	TLS, CX
	0x0009 00009 (1.go:48)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:48)	LEAQ	-88(SP), AX
	0x0015 00021 (1.go:48)	CMPQ	AX, 16(CX)
	0x0019 00025 (1.go:48)	JLS	1305
	0x001f 00031 (1.go:48)	SUBQ	$216, SP
	0x0026 00038 (1.go:48)	MOVQ	BP, 208(SP)
	0x002e 00046 (1.go:48)	LEAQ	208(SP), BP
	0x0036 00054 (1.go:48)	XORPS	X0, X0
	0x0039 00057 (1.go:48)	MOVUPS	X0, 176(SP)
	0x0041 00065 (1.go:48)	MOVUPS	X0, 192(SP)
	0x0049 00073 (1.go:48)	FUNCDATA	$0, gclocals·faa4501c2717f11e7fce158389399fba(SB)
	0x0049 00073 (1.go:48)	FUNCDATA	$1, gclocals·f92853ccfca3dbb6e8a8fb53149795ad(SB)
	0x0049 00073 (1.go:50)	MOVQ	$0, AX
	0x004b 00075 (1.go:50)	MOVQ	AX, "".i+80(SP)
	0x0050 00080 (1.go:50)	MOVQ	"".N+224(FP), CX
	0x0058 00088 (1.go:50)	MOVQ	CX, BX
	0x005b 00091 (1.go:50)	IMULQ	$100, CX
	0x005f 00095 (1.go:50)	CMPQ	AX, CX
	0x0062 00098 (1.go:50)	JGE	$0, 410
	0x0068 00104 (1.go:51)	MOVQ	"".&index+232(FP), CX
	0x0070 00112 (1.go:51)	MOVQ	(CX), DX
	0x0073 00115 (1.go:51)	INCQ	DX
	0x0076 00118 (1.go:51)	MOVQ	DX, (CX)
	0x0079 00121 (1.go:52)	MOVQ	"".&mm+240(FP), DX
	0x0081 00129 (1.go:52)	MOVQ	DX, (SP)
	0x0085 00133 (1.go:52)	PCDATA	$0, $0
	0x0085 00133 (1.go:52)	CALL	sync.(*RWMutex).Lock(SB)
	0x008a 00138 (1.go:53)	MOVQ	"".&index+232(FP), AX
	0x0092 00146 (1.go:53)	MOVQ	(AX), CX
	0x0095 00149 (1.go:53)	MOVQ	CX, ""..autotmp_151+96(SP)
	0x009a 00154 (1.go:53)	MOVQ	$5000000, (SP)
	0x00a2 00162 (1.go:53)	PCDATA	$0, $0
	0x00a2 00162 (1.go:53)	CALL	math/rand.Intn(SB)
	0x00a7 00167 (1.go:53)	MOVQ	8(SP), AX
	0x00ac 00172 (1.go:53)	MOVQ	AX, (SP)
	0x00b0 00176 (1.go:53)	PCDATA	$0, $0
	0x00b0 00176 (1.go:53)	CALL	strconv.Itoa(SB)
	0x00b5 00181 (1.go:53)	MOVQ	"".&mm+240(FP), AX
	0x00bd 00189 (1.go:53)	MOVQ	24(AX), CX
	0x00c1 00193 (1.go:53)	MOVQ	8(SP), DX
	0x00c6 00198 (1.go:53)	MOVQ	DX, ""..autotmp_172+128(SP)
	0x00ce 00206 (1.go:53)	MOVQ	16(SP), BX
	0x00d3 00211 (1.go:53)	MOVQ	BX, ""..autotmp_173+88(SP)
	0x00d8 00216 (1.go:53)	MOVQ	CX, 8(SP)
	0x00dd 00221 (1.go:53)	LEAQ	type.map[int]"".User(SB), CX
	0x00e4 00228 (1.go:53)	MOVQ	CX, (SP)
	0x00e8 00232 (1.go:53)	LEAQ	""..autotmp_151+96(SP), SI
	0x00ed 00237 (1.go:53)	MOVQ	SI, 16(SP)
	0x00f2 00242 (1.go:53)	PCDATA	$0, $1
	0x00f2 00242 (1.go:53)	CALL	runtime.mapassign(SB)
	0x00f7 00247 (1.go:53)	MOVQ	24(SP), AX
	0x00fc 00252 (1.go:53)	MOVQ	"".i+80(SP), CX
	0x0101 00257 (1.go:53)	MOVQ	CX, (AX)
	0x0104 00260 (1.go:53)	MOVQ	""..autotmp_173+88(SP), DX
	0x0109 00265 (1.go:53)	MOVQ	DX, 16(AX)
	0x010d 00269 (1.go:53)	MOVL	runtime.writeBarrier(SB), DX
	0x0113 00275 (1.go:53)	LEAQ	8(AX), BX
	0x0117 00279 (1.go:53)	LEAQ	24(AX), SI
	0x011b 00283 (1.go:53)	MOVQ	SI, ""..autotmp_174+120(SP)
	0x0120 00288 (1.go:53)	TESTL	DX, DX
	0x0122 00290 (1.go:53)	JNE	$0, 346
	0x0124 00292 (1.go:53)	MOVQ	""..autotmp_172+128(SP), DX
	0x012c 00300 (1.go:53)	MOVQ	DX, 8(AX)
	0x0130 00304 (1.go:53)	MOVQ	"".&deUser+248(FP), DX
	0x0138 00312 (1.go:53)	MOVQ	DX, 24(AX)
	0x013c 00316 (1.go:54)	MOVQ	"".&mm+240(FP), AX
	0x0144 00324 (1.go:54)	MOVQ	AX, (SP)
	0x0148 00328 (1.go:54)	PCDATA	$0, $0
	0x0148 00328 (1.go:54)	CALL	sync.(*RWMutex).Unlock(SB)
	0x014d 00333 (1.go:50)	MOVQ	"".i+80(SP), AX
	0x0152 00338 (1.go:50)	INCQ	AX
	0x0155 00341 (1.go:50)	JMP	75
	0x015a 00346 (1.go:53)	MOVQ	BX, (SP)
	0x015e 00350 (1.go:53)	MOVQ	""..autotmp_172+128(SP), AX
	0x0166 00358 (1.go:53)	MOVQ	AX, 8(SP)
	0x016b 00363 (1.go:53)	PCDATA	$0, $2
	0x016b 00363 (1.go:53)	CALL	runtime.writebarrierptr(SB)
	0x0170 00368 (1.go:53)	MOVQ	""..autotmp_174+120(SP), AX
	0x0175 00373 (1.go:53)	MOVQ	AX, (SP)
	0x0179 00377 (1.go:53)	MOVQ	"".&deUser+248(FP), AX
	0x0181 00385 (1.go:53)	MOVQ	AX, 8(SP)
	0x0186 00390 (1.go:53)	PCDATA	$0, $0
	0x0186 00390 (1.go:53)	CALL	runtime.writebarrierptr(SB)
	0x018b 00395 (1.go:50)	MOVQ	"".i+80(SP), CX
	0x0190 00400 (1.go:53)	MOVQ	"".&deUser+248(FP), DX
	0x0198 00408 (1.go:54)	JMP	316
	0x019a 00410 (1.go:50)	MOVQ	$0, AX
	0x019c 00412 (1.go:56)	MOVQ	AX, "".j+72(SP)
	0x01a1 00417 (1.go:56)	CMPQ	AX, BX
	0x01a4 00420 (1.go:56)	JGE	$0, 985
	0x01aa 00426 (1.go:57)	MOVQ	"".lenMM+256(FP), DX
	0x01b2 00434 (1.go:57)	MOVQ	(DX), CX
	0x01b5 00437 (1.go:57)	PCDATA	$0, $3
	0x01b5 00437 (1.go:57)	CALL	CX
	0x01b7 00439 (1.go:57)	MOVQ	(SP), AX
	0x01bb 00443 (1.go:57)	MOVQ	AX, (SP)
	0x01bf 00447 (1.go:57)	PCDATA	$0, $3
	0x01bf 00447 (1.go:57)	CALL	math/rand.Intn(SB)
	0x01c4 00452 (1.go:57)	MOVQ	8(SP), AX
	0x01c9 00457 (1.go:57)	MOVQ	AX, "".r+64(SP)
	0x01ce 00462 (1.go:58)	MOVQ	"".lenMM+256(FP), DX
	0x01d6 00470 (1.go:58)	MOVQ	(DX), CX
	0x01d9 00473 (1.go:58)	PCDATA	$0, $3
	0x01d9 00473 (1.go:58)	CALL	CX
	0x01db 00475 (1.go:58)	MOVQ	(SP), AX
	0x01df 00479 (1.go:58)	MOVQ	AX, (SP)
	0x01e3 00483 (1.go:58)	PCDATA	$0, $3
	0x01e3 00483 (1.go:58)	CALL	math/rand.Intn(SB)
	0x01e8 00488 (1.go:58)	MOVQ	8(SP), AX
	0x01ed 00493 (1.go:58)	MOVQ	AX, "".r2+56(SP)
	0x01f2 00498 (1.go:59)	MOVQ	"".&mm+240(FP), CX
	0x01fa 00506 (1.go:59)	MOVQ	CX, (SP)
	0x01fe 00510 (1.go:59)	PCDATA	$0, $3
	0x01fe 00510 (1.go:59)	CALL	sync.(*RWMutex).Lock(SB)
	0x0203 00515 (1.go:60)	MOVQ	"".&mm+240(FP), AX
	0x020b 00523 (1.go:60)	MOVQ	24(AX), CX
	0x020f 00527 (1.go:60)	MOVQ	CX, 8(SP)
	0x0214 00532 (1.go:53)	LEAQ	type.map[int]"".User(SB), CX
	0x021b 00539 (1.go:60)	MOVQ	CX, (SP)
	0x021f 00543 (1.go:60)	MOVQ	"".r+64(SP), BX
	0x0224 00548 (1.go:60)	MOVQ	BX, 16(SP)
	0x0229 00553 (1.go:60)	PCDATA	$0, $3
	0x0229 00553 (1.go:60)	CALL	runtime.mapaccess2_fast64(SB)
	0x022e 00558 (1.go:61)	MOVQ	"".&mm+240(FP), AX
	0x0236 00566 (1.go:61)	MOVQ	24(AX), CX
	0x023a 00570 (1.go:61)	MOVQ	CX, 8(SP)
	0x023f 00575 (1.go:53)	LEAQ	type.map[int]"".User(SB), CX
	0x0246 00582 (1.go:61)	MOVQ	CX, (SP)
	0x024a 00586 (1.go:61)	MOVQ	"".r2+56(SP), BX
	0x024f 00591 (1.go:61)	MOVQ	BX, 16(SP)
	0x0254 00596 (1.go:61)	PCDATA	$0, $3
	0x0254 00596 (1.go:61)	CALL	runtime.mapaccess2_fast64(SB)
	0x0259 00601 (1.go:61)	MOVQ	24(SP), AX
	0x025e 00606 (1.go:61)	MOVQ	24(AX), CX
	0x0262 00610 (1.go:61)	MOVQ	8(AX), BX
	0x0266 00614 (1.go:61)	MOVQ	16(AX), SI
	0x026a 00618 (1.go:61)	MOVQ	(AX), AX
	0x026d 00621 (1.go:61)	MOVQ	AX, "".us2+176(SP)
	0x0275 00629 (1.go:61)	MOVQ	BX, "".us2+184(SP)
	0x027d 00637 (1.go:61)	MOVQ	SI, "".us2+192(SP)
	0x0285 00645 (1.go:61)	MOVQ	CX, "".us2+200(SP)
	0x028d 00653 (1.go:62)	MOVQ	"".&mm+240(FP), AX
	0x0295 00661 (1.go:62)	MOVQ	AX, (SP)
	0x0299 00665 (1.go:62)	PCDATA	$0, $3
	0x0299 00665 (1.go:62)	CALL	sync.(*RWMutex).Unlock(SB)
	0x029e 00670 (1.go:68)	LEAQ	type."".User(SB), AX
	0x02a5 00677 (1.go:68)	MOVQ	AX, (SP)
	0x02a9 00681 (1.go:68)	PCDATA	$0, $3
	0x02a9 00681 (1.go:68)	CALL	runtime.newobject(SB)
	0x02ae 00686 (1.go:68)	MOVQ	"".statictmp_168+24(SB), AX
	0x02b5 00693 (1.go:68)	MOVQ	AX, ""..autotmp_175+112(SP)
	0x02ba 00698 (1.go:68)	MOVQ	"".statictmp_168(SB), CX
	0x02c1 00705 (1.go:68)	MOVQ	"".statictmp_168+16(SB), BX
	0x02c8 00712 (1.go:68)	MOVQ	8(SP), SI
	0x02cd 00717 (1.go:68)	MOVQ	SI, "".&u5+136(SP)
	0x02d5 00725 (1.go:68)	MOVQ	"".statictmp_168+8(SB), DI
	0x02dc 00732 (1.go:68)	MOVQ	CX, (SI)
	0x02df 00735 (1.go:68)	MOVQ	BX, 16(SI)
	0x02e3 00739 (1.go:68)	MOVL	runtime.writeBarrier(SB), CX
	0x02e9 00745 (1.go:68)	LEAQ	8(SI), BX
	0x02ed 00749 (1.go:68)	MOVQ	BX, ""..autotmp_176+104(SP)
	0x02f2 00754 (1.go:68)	LEAQ	24(SI), R8
	0x02f6 00758 (1.go:68)	MOVQ	R8, ""..autotmp_174+120(SP)
	0x02fb 00763 (1.go:68)	TESTL	CX, CX
	0x02fd 00765 (1.go:68)	JNE	$0, 1249
	0x0303 00771 (1.go:68)	MOVQ	DI, 8(SI)
	0x0307 00775 (1.go:68)	MOVQ	AX, 24(SI)
	0x030b 00779 (1.go:69)	MOVQ	$5000000, (SP)
	0x0313 00787 (1.go:69)	PCDATA	$0, $5
	0x0313 00787 (1.go:69)	CALL	math/rand.Intn(SB)
	0x0318 00792 (1.go:69)	MOVQ	8(SP), AX
	0x031d 00797 (1.go:69)	MOVQ	AX, (SP)
	0x0321 00801 (1.go:69)	PCDATA	$0, $5
	0x0321 00801 (1.go:69)	CALL	strconv.Itoa(SB)
	0x0326 00806 (1.go:69)	MOVQ	8(SP), AX
	0x032b 00811 (1.go:69)	MOVQ	16(SP), CX
	0x0330 00816 (1.go:69)	MOVQ	$0, (SP)
	0x0338 00824 (1.go:69)	LEAQ	go.string."asdasdasd"(SB), DX
	0x033f 00831 (1.go:69)	MOVQ	DX, 8(SP)
	0x0344 00836 (1.go:69)	MOVQ	$9, 16(SP)
	0x034d 00845 (1.go:69)	MOVQ	AX, 24(SP)
	0x0352 00850 (1.go:69)	MOVQ	CX, 32(SP)
	0x0357 00855 (1.go:69)	PCDATA	$0, $5
	0x0357 00855 (1.go:69)	CALL	runtime.concatstring2(SB)
	0x035c 00860 (1.go:69)	MOVQ	40(SP), AX
	0x0361 00865 (1.go:69)	MOVQ	48(SP), CX
	0x0366 00870 (1.go:69)	MOVQ	"".&u5+136(SP), DX
	0x036e 00878 (1.go:69)	MOVQ	CX, 16(DX)
	0x0372 00882 (1.go:69)	MOVL	runtime.writeBarrier(SB), CX
	0x0378 00888 (1.go:69)	TESTL	CX, CX
	0x037a 00890 (1.go:69)	JNE	$0, 1217
	0x0380 00896 (1.go:69)	MOVQ	AX, 8(DX)
	0x0384 00900 (1.go:70)	MOVQ	"".&deUser2+264(FP), AX
	0x038c 00908 (1.go:70)	MOVQ	(AX), CX
	0x038f 00911 (1.go:70)	TESTB	AL, (CX)
	0x0391 00913 (1.go:70)	MOVL	runtime.writeBarrier(SB), BX
	0x0397 00919 (1.go:70)	LEAQ	24(CX), SI
	0x039b 00923 (1.go:70)	TESTL	BX, BX
	0x039d 00925 (1.go:70)	JNE	$0, 1182
	0x03a3 00931 (1.go:70)	MOVQ	DX, 24(CX)
	0x03a7 00935 (1.go:71)	MOVL	runtime.writeBarrier(SB), CX
	0x03ad 00941 (1.go:71)	TESTL	CX, CX
	0x03af 00943 (1.go:71)	JNE	$0, 1155
	0x03b5 00949 (1.go:71)	MOVQ	DX, (AX)
	0x03b8 00952 (1.go:56)	MOVQ	"".j+72(SP), CX
	0x03bd 00957 (1.go:56)	INCQ	CX
	0x03c0 00960 (1.go:56)	MOVQ	"".N+224(FP), BX
	0x03c8 00968 (1.go:56)	MOVQ	CX, AX
	0x03cb 00971 (1.go:56)	MOVQ	AX, "".j+72(SP)
	0x03d0 00976 (1.go:56)	CMPQ	AX, BX
	0x03d3 00979 (1.go:56)	JLT	$0, 426
	0x03d9 00985 (1.go:73)	MOVQ	$1000000000, (SP)
	0x03e1 00993 (1.go:73)	PCDATA	$0, $7
	0x03e1 00993 (1.go:73)	CALL	time.Sleep(SB)
	0x03e6 00998 (1.go:74)	LEAQ	go.string."++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++="(SB), AX
	0x03ed 01005 (1.go:74)	MOVQ	AX, ""..autotmp_163+144(SP)
	0x03f5 01013 (1.go:74)	MOVQ	$65, ""..autotmp_163+152(SP)
	0x0401 01025 (1.go:74)	MOVQ	$0, ""..autotmp_162+160(SP)
	0x040d 01037 (1.go:74)	MOVQ	$0, ""..autotmp_162+168(SP)
	0x0419 01049 (1.go:74)	LEAQ	type.string(SB), AX
	0x0420 01056 (1.go:74)	MOVQ	AX, (SP)
	0x0424 01060 (1.go:74)	LEAQ	""..autotmp_163+144(SP), AX
	0x042c 01068 (1.go:74)	MOVQ	AX, 8(SP)
	0x0431 01073 (1.go:74)	PCDATA	$0, $8
	0x0431 01073 (1.go:74)	CALL	runtime.convT2E(SB)
	0x0436 01078 (1.go:74)	MOVQ	24(SP), AX
	0x043b 01083 (1.go:74)	MOVQ	16(SP), CX
	0x0440 01088 (1.go:74)	MOVQ	CX, ""..autotmp_162+160(SP)
	0x0448 01096 (1.go:74)	MOVQ	AX, ""..autotmp_162+168(SP)
	0x0450 01104 (1.go:74)	LEAQ	""..autotmp_162+160(SP), AX
	0x0458 01112 (1.go:74)	MOVQ	AX, (SP)
	0x045c 01116 (1.go:74)	MOVQ	$1, 8(SP)
	0x0465 01125 (1.go:74)	MOVQ	$1, 16(SP)
	0x046e 01134 (1.go:74)	PCDATA	$0, $8
	0x046e 01134 (1.go:74)	CALL	fmt.Println(SB)
	0x0473 01139 (1.go:79)	MOVQ	208(SP), BP
	0x047b 01147 (1.go:79)	ADDQ	$216, SP
	0x0482 01154 (1.go:79)	RET
	0x0483 01155 (1.go:71)	MOVQ	AX, (SP)
	0x0487 01159 (1.go:71)	MOVQ	DX, 8(SP)
	0x048c 01164 (1.go:71)	PCDATA	$0, $3
	0x048c 01164 (1.go:71)	CALL	runtime.writebarrierptr(SB)
	0x0491 01169 (1.go:70)	MOVQ	"".&deUser2+264(FP), AX
	0x0499 01177 (1.go:56)	JMP	952
	0x049e 01182 (1.go:70)	MOVQ	SI, (SP)
	0x04a2 01186 (1.go:70)	MOVQ	DX, 8(SP)
	0x04a7 01191 (1.go:70)	PCDATA	$0, $6
	0x04a7 01191 (1.go:70)	CALL	runtime.writebarrierptr(SB)
	0x04ac 01196 (1.go:71)	MOVQ	"".&deUser2+264(FP), AX
	0x04b4 01204 (1.go:71)	MOVQ	"".&u5+136(SP), DX
	0x04bc 01212 (1.go:71)	JMP	935
	0x04c1 01217 (1.go:69)	MOVQ	""..autotmp_176+104(SP), CX
	0x04c6 01222 (1.go:69)	MOVQ	CX, (SP)
	0x04ca 01226 (1.go:69)	MOVQ	AX, 8(SP)
	0x04cf 01231 (1.go:69)	PCDATA	$0, $6
	0x04cf 01231 (1.go:69)	CALL	runtime.writebarrierptr(SB)
	0x04d4 01236 (1.go:70)	MOVQ	"".&u5+136(SP), DX
	0x04dc 01244 (1.go:70)	JMP	900
	0x04e1 01249 (1.go:68)	MOVQ	BX, (SP)
	0x04e5 01253 (1.go:68)	MOVQ	DI, 8(SP)
	0x04ea 01258 (1.go:68)	PCDATA	$0, $4
	0x04ea 01258 (1.go:68)	CALL	runtime.writebarrierptr(SB)
	0x04ef 01263 (1.go:68)	MOVQ	""..autotmp_174+120(SP), AX
	0x04f4 01268 (1.go:68)	MOVQ	AX, (SP)
	0x04f8 01272 (1.go:68)	MOVQ	""..autotmp_175+112(SP), AX
	0x04fd 01277 (1.go:68)	MOVQ	AX, 8(SP)
	0x0502 01282 (1.go:68)	PCDATA	$0, $5
	0x0502 01282 (1.go:68)	CALL	runtime.writebarrierptr(SB)
	0x0507 01287 (1.go:69)	MOVQ	""..autotmp_176+104(SP), BX
	0x050c 01292 (1.go:69)	MOVQ	"".&u5+136(SP), SI
	0x0514 01300 (1.go:69)	JMP	779
	0x0519 01305 (1.go:69)	NOP
	0x0519 01305 (1.go:48)	PCDATA	$0, $-1
	0x0519 01305 (1.go:48)	CALL	runtime.morestack_noctxt(SB)
	0x051e 01310 (1.go:48)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 8d 44 24 a8 48 3b 41 10 0f 86 fa 04 00 00 48  H.D$.H;A.......H
	0x0020 81 ec d8 00 00 00 48 89 ac 24 d0 00 00 00 48 8d  ......H..$....H.
	0x0030 ac 24 d0 00 00 00 0f 57 c0 0f 11 84 24 b0 00 00  .$.....W....$...
	0x0040 00 0f 11 84 24 c0 00 00 00 31 c0 48 89 44 24 50  ....$....1.H.D$P
	0x0050 48 8b 8c 24 e0 00 00 00 48 89 cb 48 6b c9 64 48  H..$....H..Hk.dH
	0x0060 39 c8 0f 8d 32 01 00 00 48 8b 8c 24 e8 00 00 00  9...2...H..$....
	0x0070 48 8b 11 48 ff c2 48 89 11 48 8b 94 24 f0 00 00  H..H..H..H..$...
	0x0080 00 48 89 14 24 e8 00 00 00 00 48 8b 84 24 e8 00  .H..$.....H..$..
	0x0090 00 00 48 8b 08 48 89 4c 24 60 48 c7 04 24 40 4b  ..H..H.L$`H..$@K
	0x00a0 4c 00 e8 00 00 00 00 48 8b 44 24 08 48 89 04 24  L......H.D$.H..$
	0x00b0 e8 00 00 00 00 48 8b 84 24 f0 00 00 00 48 8b 48  .....H..$....H.H
	0x00c0 18 48 8b 54 24 08 48 89 94 24 80 00 00 00 48 8b  .H.T$.H..$....H.
	0x00d0 5c 24 10 48 89 5c 24 58 48 89 4c 24 08 48 8d 0d  \$.H.\$XH.L$.H..
	0x00e0 00 00 00 00 48 89 0c 24 48 8d 74 24 60 48 89 74  ....H..$H.t$`H.t
	0x00f0 24 10 e8 00 00 00 00 48 8b 44 24 18 48 8b 4c 24  $......H.D$.H.L$
	0x0100 50 48 89 08 48 8b 54 24 58 48 89 50 10 8b 15 00  PH..H.T$XH.P....
	0x0110 00 00 00 48 8d 58 08 48 8d 70 18 48 89 74 24 78  ...H.X.H.p.H.t$x
	0x0120 85 d2 75 36 48 8b 94 24 80 00 00 00 48 89 50 08  ..u6H..$....H.P.
	0x0130 48 8b 94 24 f8 00 00 00 48 89 50 18 48 8b 84 24  H..$....H.P.H..$
	0x0140 f0 00 00 00 48 89 04 24 e8 00 00 00 00 48 8b 44  ....H..$.....H.D
	0x0150 24 50 48 ff c0 e9 f1 fe ff ff 48 89 1c 24 48 8b  $PH.......H..$H.
	0x0160 84 24 80 00 00 00 48 89 44 24 08 e8 00 00 00 00  .$....H.D$......
	0x0170 48 8b 44 24 78 48 89 04 24 48 8b 84 24 f8 00 00  H.D$xH..$H..$...
	0x0180 00 48 89 44 24 08 e8 00 00 00 00 48 8b 4c 24 50  .H.D$......H.L$P
	0x0190 48 8b 94 24 f8 00 00 00 eb a2 31 c0 48 89 44 24  H..$......1.H.D$
	0x01a0 48 48 39 d8 0f 8d 2f 02 00 00 48 8b 94 24 00 01  HH9.../...H..$..
	0x01b0 00 00 48 8b 0a ff d1 48 8b 04 24 48 89 04 24 e8  ..H....H..$H..$.
	0x01c0 00 00 00 00 48 8b 44 24 08 48 89 44 24 40 48 8b  ....H.D$.H.D$@H.
	0x01d0 94 24 00 01 00 00 48 8b 0a ff d1 48 8b 04 24 48  .$....H....H..$H
	0x01e0 89 04 24 e8 00 00 00 00 48 8b 44 24 08 48 89 44  ..$.....H.D$.H.D
	0x01f0 24 38 48 8b 8c 24 f0 00 00 00 48 89 0c 24 e8 00  $8H..$....H..$..
	0x0200 00 00 00 48 8b 84 24 f0 00 00 00 48 8b 48 18 48  ...H..$....H.H.H
	0x0210 89 4c 24 08 48 8d 0d 00 00 00 00 48 89 0c 24 48  .L$.H......H..$H
	0x0220 8b 5c 24 40 48 89 5c 24 10 e8 00 00 00 00 48 8b  .\$@H.\$......H.
	0x0230 84 24 f0 00 00 00 48 8b 48 18 48 89 4c 24 08 48  .$....H.H.H.L$.H
	0x0240 8d 0d 00 00 00 00 48 89 0c 24 48 8b 5c 24 38 48  ......H..$H.\$8H
	0x0250 89 5c 24 10 e8 00 00 00 00 48 8b 44 24 18 48 8b  .\$......H.D$.H.
	0x0260 48 18 48 8b 58 08 48 8b 70 10 48 8b 00 48 89 84  H.H.X.H.p.H..H..
	0x0270 24 b0 00 00 00 48 89 9c 24 b8 00 00 00 48 89 b4  $....H..$....H..
	0x0280 24 c0 00 00 00 48 89 8c 24 c8 00 00 00 48 8b 84  $....H..$....H..
	0x0290 24 f0 00 00 00 48 89 04 24 e8 00 00 00 00 48 8d  $....H..$.....H.
	0x02a0 05 00 00 00 00 48 89 04 24 e8 00 00 00 00 48 8b  .....H..$.....H.
	0x02b0 05 00 00 00 00 48 89 44 24 70 48 8b 0d 00 00 00  .....H.D$pH.....
	0x02c0 00 48 8b 1d 00 00 00 00 48 8b 74 24 08 48 89 b4  .H......H.t$.H..
	0x02d0 24 88 00 00 00 48 8b 3d 00 00 00 00 48 89 0e 48  $....H.=....H..H
	0x02e0 89 5e 10 8b 0d 00 00 00 00 48 8d 5e 08 48 89 5c  .^.......H.^.H.\
	0x02f0 24 68 4c 8d 46 18 4c 89 44 24 78 85 c9 0f 85 de  $hL.F.L.D$x.....
	0x0300 01 00 00 48 89 7e 08 48 89 46 18 48 c7 04 24 40  ...H.~.H.F.H..$@
	0x0310 4b 4c 00 e8 00 00 00 00 48 8b 44 24 08 48 89 04  KL......H.D$.H..
	0x0320 24 e8 00 00 00 00 48 8b 44 24 08 48 8b 4c 24 10  $.....H.D$.H.L$.
	0x0330 48 c7 04 24 00 00 00 00 48 8d 15 00 00 00 00 48  H..$....H......H
	0x0340 89 54 24 08 48 c7 44 24 10 09 00 00 00 48 89 44  .T$.H.D$.....H.D
	0x0350 24 18 48 89 4c 24 20 e8 00 00 00 00 48 8b 44 24  $.H.L$ .....H.D$
	0x0360 28 48 8b 4c 24 30 48 8b 94 24 88 00 00 00 48 89  (H.L$0H..$....H.
	0x0370 4a 10 8b 0d 00 00 00 00 85 c9 0f 85 41 01 00 00  J...........A...
	0x0380 48 89 42 08 48 8b 84 24 08 01 00 00 48 8b 08 84  H.B.H..$....H...
	0x0390 01 8b 1d 00 00 00 00 48 8d 71 18 85 db 0f 85 fb  .......H.q......
	0x03a0 00 00 00 48 89 51 18 8b 0d 00 00 00 00 85 c9 0f  ...H.Q..........
	0x03b0 85 ce 00 00 00 48 89 10 48 8b 4c 24 48 48 ff c1  .....H..H.L$HH..
	0x03c0 48 8b 9c 24 e0 00 00 00 48 89 c8 48 89 44 24 48  H..$....H..H.D$H
	0x03d0 48 39 d8 0f 8c d1 fd ff ff 48 c7 04 24 00 ca 9a  H9.......H..$...
	0x03e0 3b e8 00 00 00 00 48 8d 05 00 00 00 00 48 89 84  ;.....H......H..
	0x03f0 24 90 00 00 00 48 c7 84 24 98 00 00 00 41 00 00  $....H..$....A..
	0x0400 00 48 c7 84 24 a0 00 00 00 00 00 00 00 48 c7 84  .H..$........H..
	0x0410 24 a8 00 00 00 00 00 00 00 48 8d 05 00 00 00 00  $........H......
	0x0420 48 89 04 24 48 8d 84 24 90 00 00 00 48 89 44 24  H..$H..$....H.D$
	0x0430 08 e8 00 00 00 00 48 8b 44 24 18 48 8b 4c 24 10  ......H.D$.H.L$.
	0x0440 48 89 8c 24 a0 00 00 00 48 89 84 24 a8 00 00 00  H..$....H..$....
	0x0450 48 8d 84 24 a0 00 00 00 48 89 04 24 48 c7 44 24  H..$....H..$H.D$
	0x0460 08 01 00 00 00 48 c7 44 24 10 01 00 00 00 e8 00  .....H.D$.......
	0x0470 00 00 00 48 8b ac 24 d0 00 00 00 48 81 c4 d8 00  ...H..$....H....
	0x0480 00 00 c3 48 89 04 24 48 89 54 24 08 e8 00 00 00  ...H..$H.T$.....
	0x0490 00 48 8b 84 24 08 01 00 00 e9 1a ff ff ff 48 89  .H..$.........H.
	0x04a0 34 24 48 89 54 24 08 e8 00 00 00 00 48 8b 84 24  4$H.T$......H..$
	0x04b0 08 01 00 00 48 8b 94 24 88 00 00 00 e9 e6 fe ff  ....H..$........
	0x04c0 ff 48 8b 4c 24 68 48 89 0c 24 48 89 44 24 08 e8  .H.L$hH..$H.D$..
	0x04d0 00 00 00 00 48 8b 94 24 88 00 00 00 e9 a3 fe ff  ....H..$........
	0x04e0 ff 48 89 1c 24 48 89 7c 24 08 e8 00 00 00 00 48  .H..$H.|$......H
	0x04f0 8b 44 24 78 48 89 04 24 48 8b 44 24 70 48 89 44  .D$xH..$H.D$pH.D
	0x0500 24 08 e8 00 00 00 00 48 8b 5c 24 68 48 8b b4 24  $......H.\$hH..$
	0x0510 88 00 00 00 e9 f2 fd ff ff e8 00 00 00 00 e9 dd  ................
	0x0520 fa ff ff                                         ...
	rel 12+4 t=16 TLS+0
	rel 134+4 t=8 sync.(*RWMutex).Lock+0
	rel 163+4 t=8 math/rand.Intn+0
	rel 177+4 t=8 strconv.Itoa+0
	rel 224+4 t=15 type.map[int]"".User+0
	rel 243+4 t=8 runtime.mapassign+0
	rel 271+4 t=15 runtime.writeBarrier+0
	rel 329+4 t=8 sync.(*RWMutex).Unlock+0
	rel 364+4 t=8 runtime.writebarrierptr+0
	rel 391+4 t=8 runtime.writebarrierptr+0
	rel 437+0 t=11 +0
	rel 448+4 t=8 math/rand.Intn+0
	rel 473+0 t=11 +0
	rel 484+4 t=8 math/rand.Intn+0
	rel 511+4 t=8 sync.(*RWMutex).Lock+0
	rel 535+4 t=15 type.map[int]"".User+0
	rel 554+4 t=8 runtime.mapaccess2_fast64+0
	rel 578+4 t=15 type.map[int]"".User+0
	rel 597+4 t=8 runtime.mapaccess2_fast64+0
	rel 666+4 t=8 sync.(*RWMutex).Unlock+0
	rel 673+4 t=15 type."".User+0
	rel 682+4 t=8 runtime.newobject+0
	rel 689+4 t=15 "".statictmp_168+24
	rel 701+4 t=15 "".statictmp_168+0
	rel 708+4 t=15 "".statictmp_168+16
	rel 728+4 t=15 "".statictmp_168+8
	rel 741+4 t=15 runtime.writeBarrier+0
	rel 788+4 t=8 math/rand.Intn+0
	rel 802+4 t=8 strconv.Itoa+0
	rel 827+4 t=15 go.string."asdasdasd"+0
	rel 856+4 t=8 runtime.concatstring2+0
	rel 884+4 t=15 runtime.writeBarrier+0
	rel 915+4 t=15 runtime.writeBarrier+0
	rel 937+4 t=15 runtime.writeBarrier+0
	rel 994+4 t=8 time.Sleep+0
	rel 1001+4 t=15 go.string."++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++="+0
	rel 1052+4 t=15 type.string+0
	rel 1074+4 t=8 runtime.convT2E+0
	rel 1135+4 t=8 fmt.Println+0
	rel 1165+4 t=8 runtime.writebarrierptr+0
	rel 1192+4 t=8 runtime.writebarrierptr+0
	rel 1232+4 t=8 runtime.writebarrierptr+0
	rel 1259+4 t=8 runtime.writebarrierptr+0
	rel 1283+4 t=8 runtime.writebarrierptr+0
	rel 1306+4 t=8 runtime.morestack_noctxt+0
"".main.func3 t=1 size=407 args=0x20 locals=0x50
	0x0000 00000 (1.go:81)	TEXT	"".main.func3(SB), $80-32
	0x0000 00000 (1.go:81)	MOVQ	TLS, CX
	0x0009 00009 (1.go:81)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:81)	CMPQ	SP, 16(CX)
	0x0014 00020 (1.go:81)	JLS	397
	0x001a 00026 (1.go:81)	SUBQ	$80, SP
	0x001e 00030 (1.go:81)	MOVQ	BP, 72(SP)
	0x0023 00035 (1.go:81)	LEAQ	72(SP), BP
	0x0028 00040 (1.go:81)	FUNCDATA	$0, gclocals·23224035f3c252fdcd5453a88ab61c35(SB)
	0x0028 00040 (1.go:81)	FUNCDATA	$1, gclocals·78de76ab2bc1c1d963901a10e08a17b9(SB)
	0x0028 00040 (1.go:82)	LEAQ	type.chan bool(SB), AX
	0x002f 00047 (1.go:82)	MOVQ	AX, (SP)
	0x0033 00051 (1.go:82)	MOVQ	"".doned+88(FP), AX
	0x0038 00056 (1.go:82)	MOVQ	AX, 8(SP)
	0x003d 00061 (1.go:82)	MOVQ	$0, 16(SP)
	0x0046 00070 (1.go:82)	PCDATA	$0, $1
	0x0046 00070 (1.go:82)	CALL	runtime.chanrecv1(SB)
	0x004b 00075 (1.go:84)	MOVQ	$0, AX
	0x004d 00077 (1.go:84)	MOVQ	AX, "".i+32(SP)
	0x0052 00082 (1.go:84)	CMPQ	AX, $1000
	0x0058 00088 (1.go:84)	JGE	$0, 324
	0x005e 00094 (1.go:85)	MOVQ	"".&index+96(FP), CX
	0x0063 00099 (1.go:85)	MOVQ	(CX), DX
	0x0066 00102 (1.go:85)	INCQ	DX
	0x0069 00105 (1.go:85)	MOVQ	DX, (CX)
	0x006c 00108 (1.go:86)	MOVQ	"".&mm+104(FP), DX
	0x0071 00113 (1.go:86)	MOVQ	DX, (SP)
	0x0075 00117 (1.go:86)	PCDATA	$0, $1
	0x0075 00117 (1.go:86)	CALL	sync.(*RWMutex).Lock(SB)
	0x007a 00122 (1.go:87)	MOVQ	"".&index+96(FP), AX
	0x007f 00127 (1.go:87)	MOVQ	(AX), CX
	0x0082 00130 (1.go:87)	MOVQ	CX, ""..autotmp_178+48(SP)
	0x0087 00135 (1.go:87)	MOVQ	$5000000, (SP)
	0x008f 00143 (1.go:87)	PCDATA	$0, $1
	0x008f 00143 (1.go:87)	CALL	math/rand.Intn(SB)
	0x0094 00148 (1.go:87)	MOVQ	8(SP), AX
	0x0099 00153 (1.go:87)	MOVQ	AX, (SP)
	0x009d 00157 (1.go:87)	PCDATA	$0, $1
	0x009d 00157 (1.go:87)	CALL	strconv.Itoa(SB)
	0x00a2 00162 (1.go:87)	MOVQ	"".&mm+104(FP), AX
	0x00a7 00167 (1.go:87)	MOVQ	24(AX), CX
	0x00ab 00171 (1.go:87)	MOVQ	8(SP), DX
	0x00b0 00176 (1.go:87)	MOVQ	DX, ""..autotmp_184+64(SP)
	0x00b5 00181 (1.go:87)	MOVQ	16(SP), BX
	0x00ba 00186 (1.go:87)	MOVQ	BX, ""..autotmp_185+40(SP)
	0x00bf 00191 (1.go:87)	MOVQ	CX, 8(SP)
	0x00c4 00196 (1.go:87)	LEAQ	type.map[int]"".User(SB), CX
	0x00cb 00203 (1.go:87)	MOVQ	CX, (SP)
	0x00cf 00207 (1.go:87)	LEAQ	""..autotmp_178+48(SP), SI
	0x00d4 00212 (1.go:87)	MOVQ	SI, 16(SP)
	0x00d9 00217 (1.go:87)	PCDATA	$0, $2
	0x00d9 00217 (1.go:87)	CALL	runtime.mapassign(SB)
	0x00de 00222 (1.go:87)	MOVQ	24(SP), AX
	0x00e3 00227 (1.go:87)	MOVQ	"".i+32(SP), CX
	0x00e8 00232 (1.go:87)	MOVQ	CX, (AX)
	0x00eb 00235 (1.go:87)	MOVQ	""..autotmp_185+40(SP), DX
	0x00f0 00240 (1.go:87)	MOVQ	DX, 16(AX)
	0x00f4 00244 (1.go:87)	MOVL	runtime.writeBarrier(SB), DX
	0x00fa 00250 (1.go:87)	LEAQ	8(AX), BX
	0x00fe 00254 (1.go:87)	LEAQ	24(AX), SI
	0x0102 00258 (1.go:87)	MOVQ	SI, ""..autotmp_186+56(SP)
	0x0107 00263 (1.go:87)	TESTL	DX, DX
	0x0109 00265 (1.go:87)	JNE	$0, 342
	0x010b 00267 (1.go:87)	MOVQ	""..autotmp_184+64(SP), DX
	0x0110 00272 (1.go:87)	MOVQ	DX, 8(AX)
	0x0114 00276 (1.go:87)	MOVQ	"".&deUser+112(FP), DX
	0x0119 00281 (1.go:87)	MOVQ	DX, 24(AX)
	0x011d 00285 (1.go:88)	MOVQ	"".&mm+104(FP), AX
	0x0122 00290 (1.go:88)	MOVQ	AX, (SP)
	0x0126 00294 (1.go:88)	PCDATA	$0, $1
	0x0126 00294 (1.go:88)	CALL	sync.(*RWMutex).Unlock(SB)
	0x012b 00299 (1.go:84)	MOVQ	"".i+32(SP), AX
	0x0130 00304 (1.go:84)	INCQ	AX
	0x0133 00307 (1.go:84)	MOVQ	AX, "".i+32(SP)
	0x0138 00312 (1.go:84)	CMPQ	AX, $1000
	0x013e 00318 (1.go:84)	JLT	$0, 94
	0x0144 00324 (1.go:90)	MOVQ	$50000000, (SP)
	0x014c 00332 (1.go:90)	PCDATA	$0, $1
	0x014c 00332 (1.go:90)	CALL	time.Sleep(SB)
	0x0151 00337 (1.go:84)	JMP	75
	0x0156 00342 (1.go:87)	MOVQ	BX, (SP)
	0x015a 00346 (1.go:87)	MOVQ	""..autotmp_184+64(SP), AX
	0x015f 00351 (1.go:87)	MOVQ	AX, 8(SP)
	0x0164 00356 (1.go:87)	PCDATA	$0, $3
	0x0164 00356 (1.go:87)	CALL	runtime.writebarrierptr(SB)
	0x0169 00361 (1.go:87)	MOVQ	""..autotmp_186+56(SP), AX
	0x016e 00366 (1.go:87)	MOVQ	AX, (SP)
	0x0172 00370 (1.go:87)	MOVQ	"".&deUser+112(FP), AX
	0x0177 00375 (1.go:87)	MOVQ	AX, 8(SP)
	0x017c 00380 (1.go:87)	PCDATA	$0, $1
	0x017c 00380 (1.go:87)	CALL	runtime.writebarrierptr(SB)
	0x0181 00385 (1.go:84)	MOVQ	"".i+32(SP), CX
	0x0186 00390 (1.go:87)	MOVQ	"".&deUser+112(FP), DX
	0x018b 00395 (1.go:88)	JMP	285
	0x018d 00397 (1.go:88)	NOP
	0x018d 00397 (1.go:81)	PCDATA	$0, $-1
	0x018d 00397 (1.go:81)	CALL	runtime.morestack_noctxt(SB)
	0x0192 00402 (1.go:81)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 73 01 00 00 48 83 ec 50 48 89  H;a...s...H..PH.
	0x0020 6c 24 48 48 8d 6c 24 48 48 8d 05 00 00 00 00 48  l$HH.l$HH......H
	0x0030 89 04 24 48 8b 44 24 58 48 89 44 24 08 48 c7 44  ..$H.D$XH.D$.H.D
	0x0040 24 10 00 00 00 00 e8 00 00 00 00 31 c0 48 89 44  $..........1.H.D
	0x0050 24 20 48 3d e8 03 00 00 0f 8d e6 00 00 00 48 8b  $ H=..........H.
	0x0060 4c 24 60 48 8b 11 48 ff c2 48 89 11 48 8b 54 24  L$`H..H..H..H.T$
	0x0070 68 48 89 14 24 e8 00 00 00 00 48 8b 44 24 60 48  hH..$.....H.D$`H
	0x0080 8b 08 48 89 4c 24 30 48 c7 04 24 40 4b 4c 00 e8  ..H.L$0H..$@KL..
	0x0090 00 00 00 00 48 8b 44 24 08 48 89 04 24 e8 00 00  ....H.D$.H..$...
	0x00a0 00 00 48 8b 44 24 68 48 8b 48 18 48 8b 54 24 08  ..H.D$hH.H.H.T$.
	0x00b0 48 89 54 24 40 48 8b 5c 24 10 48 89 5c 24 28 48  H.T$@H.\$.H.\$(H
	0x00c0 89 4c 24 08 48 8d 0d 00 00 00 00 48 89 0c 24 48  .L$.H......H..$H
	0x00d0 8d 74 24 30 48 89 74 24 10 e8 00 00 00 00 48 8b  .t$0H.t$......H.
	0x00e0 44 24 18 48 8b 4c 24 20 48 89 08 48 8b 54 24 28  D$.H.L$ H..H.T$(
	0x00f0 48 89 50 10 8b 15 00 00 00 00 48 8d 58 08 48 8d  H.P.......H.X.H.
	0x0100 70 18 48 89 74 24 38 85 d2 75 4b 48 8b 54 24 40  p.H.t$8..uKH.T$@
	0x0110 48 89 50 08 48 8b 54 24 70 48 89 50 18 48 8b 44  H.P.H.T$pH.P.H.D
	0x0120 24 68 48 89 04 24 e8 00 00 00 00 48 8b 44 24 20  $hH..$.....H.D$ 
	0x0130 48 ff c0 48 89 44 24 20 48 3d e8 03 00 00 0f 8c  H..H.D$ H=......
	0x0140 1a ff ff ff 48 c7 04 24 80 f0 fa 02 e8 00 00 00  ....H..$........
	0x0150 00 e9 f5 fe ff ff 48 89 1c 24 48 8b 44 24 40 48  ......H..$H.D$@H
	0x0160 89 44 24 08 e8 00 00 00 00 48 8b 44 24 38 48 89  .D$......H.D$8H.
	0x0170 04 24 48 8b 44 24 70 48 89 44 24 08 e8 00 00 00  .$H.D$pH.D$.....
	0x0180 00 48 8b 4c 24 20 48 8b 54 24 70 eb 90 e8 00 00  .H.L$ H.T$p.....
	0x0190 00 00 e9 69 fe ff ff                             ...i...
	rel 12+4 t=16 TLS+0
	rel 43+4 t=15 type.chan bool+0
	rel 71+4 t=8 runtime.chanrecv1+0
	rel 118+4 t=8 sync.(*RWMutex).Lock+0
	rel 144+4 t=8 math/rand.Intn+0
	rel 158+4 t=8 strconv.Itoa+0
	rel 199+4 t=15 type.map[int]"".User+0
	rel 218+4 t=8 runtime.mapassign+0
	rel 246+4 t=15 runtime.writeBarrier+0
	rel 295+4 t=8 sync.(*RWMutex).Unlock+0
	rel 333+4 t=8 time.Sleep+0
	rel 357+4 t=8 runtime.writebarrierptr+0
	rel 381+4 t=8 runtime.writebarrierptr+0
	rel 398+4 t=8 runtime.morestack_noctxt+0
"".main.func4 t=1 size=205 args=0x10 locals=0x38
	0x0000 00000 (1.go:94)	TEXT	"".main.func4(SB), $56-16
	0x0000 00000 (1.go:94)	MOVQ	TLS, CX
	0x0009 00009 (1.go:94)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:94)	CMPQ	SP, 16(CX)
	0x0014 00020 (1.go:94)	JLS	195
	0x001a 00026 (1.go:94)	SUBQ	$56, SP
	0x001e 00030 (1.go:94)	MOVQ	BP, 48(SP)
	0x0023 00035 (1.go:94)	LEAQ	48(SP), BP
	0x0028 00040 (1.go:94)	FUNCDATA	$0, gclocals·d2de89f6cf5e99ce0097843c3071e829(SB)
	0x0028 00040 (1.go:94)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (1.go:97)	MOVQ	$0, AX
	0x002a 00042 (1.go:97)	MOVQ	AX, "".i+32(SP)
	0x002f 00047 (1.go:97)	CMPQ	AX, $10
	0x0033 00051 (1.go:97)	JGE	$0, 177
	0x0035 00053 (1.go:98)	MOVQ	"".&index+64(FP), CX
	0x003a 00058 (1.go:98)	MOVQ	(CX), DX
	0x003d 00061 (1.go:98)	MOVQ	DX, (SP)
	0x0041 00065 (1.go:98)	PCDATA	$0, $0
	0x0041 00065 (1.go:98)	CALL	math/rand.Intn(SB)
	0x0046 00070 (1.go:98)	MOVQ	8(SP), AX
	0x004b 00075 (1.go:98)	MOVQ	AX, "".r+24(SP)
	0x0050 00080 (1.go:99)	MOVQ	"".&mm+72(FP), CX
	0x0055 00085 (1.go:99)	MOVQ	CX, (SP)
	0x0059 00089 (1.go:99)	PCDATA	$0, $0
	0x0059 00089 (1.go:99)	CALL	sync.(*RWMutex).Lock(SB)
	0x005e 00094 (1.go:100)	MOVQ	"".r+24(SP), AX
	0x0063 00099 (1.go:100)	MOVQ	AX, ""..autotmp_187+40(SP)
	0x0068 00104 (1.go:100)	MOVQ	"".&mm+72(FP), AX
	0x006d 00109 (1.go:100)	MOVQ	24(AX), CX
	0x0071 00113 (1.go:100)	MOVQ	CX, 8(SP)
	0x0076 00118 (1.go:100)	LEAQ	type.map[int]"".User(SB), CX
	0x007d 00125 (1.go:100)	MOVQ	CX, (SP)
	0x0081 00129 (1.go:100)	LEAQ	""..autotmp_187+40(SP), DX
	0x0086 00134 (1.go:100)	MOVQ	DX, 16(SP)
	0x008b 00139 (1.go:100)	PCDATA	$0, $0
	0x008b 00139 (1.go:100)	CALL	runtime.mapdelete(SB)
	0x0090 00144 (1.go:101)	MOVQ	"".&mm+72(FP), AX
	0x0095 00149 (1.go:101)	MOVQ	AX, (SP)
	0x0099 00153 (1.go:101)	PCDATA	$0, $0
	0x0099 00153 (1.go:101)	CALL	sync.(*RWMutex).Unlock(SB)
	0x009e 00158 (1.go:97)	MOVQ	"".i+32(SP), AX
	0x00a3 00163 (1.go:97)	INCQ	AX
	0x00a6 00166 (1.go:97)	MOVQ	AX, "".i+32(SP)
	0x00ab 00171 (1.go:97)	CMPQ	AX, $10
	0x00af 00175 (1.go:97)	JLT	$0, 53
	0x00b1 00177 (1.go:103)	MOVQ	$50000000, (SP)
	0x00b9 00185 (1.go:103)	PCDATA	$0, $0
	0x00b9 00185 (1.go:103)	CALL	time.Sleep(SB)
	0x00be 00190 (1.go:97)	JMP	40
	0x00c3 00195 (1.go:97)	NOP
	0x00c3 00195 (1.go:94)	PCDATA	$0, $-1
	0x00c3 00195 (1.go:94)	CALL	runtime.morestack_noctxt(SB)
	0x00c8 00200 (1.go:94)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 a9 00 00 00 48 83 ec 38 48 89  H;a.......H..8H.
	0x0020 6c 24 30 48 8d 6c 24 30 31 c0 48 89 44 24 20 48  l$0H.l$01.H.D$ H
	0x0030 83 f8 0a 7d 7c 48 8b 4c 24 40 48 8b 11 48 89 14  ...}|H.L$@H..H..
	0x0040 24 e8 00 00 00 00 48 8b 44 24 08 48 89 44 24 18  $.....H.D$.H.D$.
	0x0050 48 8b 4c 24 48 48 89 0c 24 e8 00 00 00 00 48 8b  H.L$HH..$.....H.
	0x0060 44 24 18 48 89 44 24 28 48 8b 44 24 48 48 8b 48  D$.H.D$(H.D$HH.H
	0x0070 18 48 89 4c 24 08 48 8d 0d 00 00 00 00 48 89 0c  .H.L$.H......H..
	0x0080 24 48 8d 54 24 28 48 89 54 24 10 e8 00 00 00 00  $H.T$(H.T$......
	0x0090 48 8b 44 24 48 48 89 04 24 e8 00 00 00 00 48 8b  H.D$HH..$.....H.
	0x00a0 44 24 20 48 ff c0 48 89 44 24 20 48 83 f8 0a 7c  D$ H..H.D$ H...|
	0x00b0 84 48 c7 04 24 80 f0 fa 02 e8 00 00 00 00 e9 65  .H..$..........e
	0x00c0 ff ff ff e8 00 00 00 00 e9 33 ff ff ff           .........3...
	rel 12+4 t=16 TLS+0
	rel 66+4 t=8 math/rand.Intn+0
	rel 90+4 t=8 sync.(*RWMutex).Lock+0
	rel 121+4 t=15 type.map[int]"".User+0
	rel 140+4 t=8 runtime.mapdelete+0
	rel 154+4 t=8 sync.(*RWMutex).Unlock+0
	rel 186+4 t=8 time.Sleep+0
	rel 196+4 t=8 runtime.morestack_noctxt+0
"".main.func5 t=1 size=400 args=0x8 locals=0x40
	0x0000 00000 (1.go:111)	TEXT	"".main.func5(SB), $64-8
	0x0000 00000 (1.go:111)	MOVQ	TLS, CX
	0x0009 00009 (1.go:111)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:111)	CMPQ	SP, 16(CX)
	0x0014 00020 (1.go:111)	JLS	390
	0x001a 00026 (1.go:111)	SUBQ	$64, SP
	0x001e 00030 (1.go:111)	MOVQ	BP, 56(SP)
	0x0023 00035 (1.go:111)	LEAQ	56(SP), BP
	0x0028 00040 (1.go:111)	FUNCDATA	$0, gclocals·69a9291448fa273f79569cb593f615b2(SB)
	0x0028 00040 (1.go:111)	FUNCDATA	$1, gclocals·55db315ce6b70e62560e945a72b60a76(SB)
	0x0028 00040 (1.go:113)	MOVQ	$0, AX
	0x002a 00042 (1.go:113)	MOVQ	AX, "".i+16(SP)
	0x002f 00047 (1.go:113)	CMPQ	AX, $20000
	0x0035 00053 (1.go:113)	JGE	$0, 237
	0x003b 00059 (1.go:114)	LEAQ	type."".User(SB), CX
	0x0042 00066 (1.go:114)	MOVQ	CX, (SP)
	0x0046 00070 (1.go:114)	PCDATA	$0, $0
	0x0046 00070 (1.go:114)	CALL	runtime.newobject(SB)
	0x004b 00075 (1.go:114)	MOVQ	8(SP), AX
	0x0050 00080 (1.go:114)	MOVQ	AX, "".&u+48(SP)
	0x0055 00085 (1.go:114)	MOVQ	"".statictmp_191+24(SB), CX
	0x005c 00092 (1.go:114)	MOVQ	CX, ""..autotmp_192+40(SP)
	0x0061 00097 (1.go:114)	MOVQ	"".statictmp_191+8(SB), DX
	0x0068 00104 (1.go:114)	MOVQ	"".statictmp_191(SB), BX
	0x006f 00111 (1.go:114)	MOVQ	"".statictmp_191+16(SB), SI
	0x0076 00118 (1.go:114)	MOVQ	BX, (AX)
	0x0079 00121 (1.go:114)	MOVQ	SI, 16(AX)
	0x007d 00125 (1.go:114)	MOVL	runtime.writeBarrier(SB), BX
	0x0083 00131 (1.go:114)	LEAQ	8(AX), SI
	0x0087 00135 (1.go:114)	MOVQ	SI, ""..autotmp_193+32(SP)
	0x008c 00140 (1.go:114)	LEAQ	24(AX), DI
	0x0090 00144 (1.go:114)	MOVQ	DI, ""..autotmp_194+24(SP)
	0x0095 00149 (1.go:114)	TESTL	BX, BX
	0x0097 00151 (1.go:114)	JNE	$0, 337
	0x009d 00157 (1.go:114)	MOVQ	DX, 8(AX)
	0x00a1 00161 (1.go:114)	MOVQ	CX, 24(AX)
	0x00a5 00165 (1.go:115)	MOVQ	$8, 16(AX)
	0x00ad 00173 (1.go:115)	MOVL	runtime.writeBarrier(SB), CX
	0x00b3 00179 (1.go:115)	TESTL	CX, CX
	0x00b5 00181 (1.go:115)	JNE	$0, 299
	0x00b7 00183 (1.go:115)	LEAQ	go.string."lkaskjas"(SB), CX
	0x00be 00190 (1.go:115)	MOVQ	CX, 8(AX)
	0x00c2 00194 (1.go:116)	MOVL	runtime.writeBarrier(SB), DX
	0x00c8 00200 (1.go:116)	TESTL	DX, DX
	0x00ca 00202 (1.go:116)	JNE	$0, 269
	0x00cc 00204 (1.go:116)	MOVQ	AX, "".uuu(SB)
	0x00d3 00211 (1.go:113)	MOVQ	"".i+16(SP), DX
	0x00d8 00216 (1.go:113)	LEAQ	1(DX), AX
	0x00dc 00220 (1.go:113)	MOVQ	AX, "".i+16(SP)
	0x00e1 00225 (1.go:113)	CMPQ	AX, $20000
	0x00e7 00231 (1.go:113)	JLT	$0, 59
	0x00ed 00237 (1.go:119)	MOVQ	"".&g+72(FP), AX
	0x00f2 00242 (1.go:119)	MOVQ	(AX), CX
	0x00f5 00245 (1.go:119)	INCQ	CX
	0x00f8 00248 (1.go:119)	MOVQ	CX, (AX)
	0x00fb 00251 (1.go:120)	MOVQ	$10000, (SP)
	0x0103 00259 (1.go:120)	PCDATA	$0, $0
	0x0103 00259 (1.go:120)	CALL	time.Sleep(SB)
	0x0108 00264 (1.go:113)	JMP	40
	0x010d 00269 (1.go:116)	LEAQ	"".uuu(SB), DX
	0x0114 00276 (1.go:116)	MOVQ	DX, (SP)
	0x0118 00280 (1.go:116)	MOVQ	AX, 8(SP)
	0x011d 00285 (1.go:116)	PCDATA	$0, $0
	0x011d 00285 (1.go:116)	CALL	runtime.writebarrierptr(SB)
	0x0122 00290 (1.go:115)	LEAQ	go.string."lkaskjas"(SB), CX
	0x0129 00297 (1.go:113)	JMP	211
	0x012b 00299 (1.go:115)	MOVQ	SI, (SP)
	0x012f 00303 (1.go:115)	LEAQ	go.string."lkaskjas"(SB), CX
	0x0136 00310 (1.go:115)	MOVQ	CX, 8(SP)
	0x013b 00315 (1.go:115)	PCDATA	$0, $3
	0x013b 00315 (1.go:115)	CALL	runtime.writebarrierptr(SB)
	0x0140 00320 (1.go:116)	MOVQ	"".&u+48(SP), AX
	0x0145 00325 (1.go:115)	LEAQ	go.string."lkaskjas"(SB), CX
	0x014c 00332 (1.go:116)	JMP	194
	0x0151 00337 (1.go:114)	MOVQ	SI, (SP)
	0x0155 00341 (1.go:114)	MOVQ	DX, 8(SP)
	0x015a 00346 (1.go:114)	PCDATA	$0, $1
	0x015a 00346 (1.go:114)	CALL	runtime.writebarrierptr(SB)
	0x015f 00351 (1.go:114)	MOVQ	""..autotmp_194+24(SP), AX
	0x0164 00356 (1.go:114)	MOVQ	AX, (SP)
	0x0168 00360 (1.go:114)	MOVQ	""..autotmp_192+40(SP), AX
	0x016d 00365 (1.go:114)	MOVQ	AX, 8(SP)
	0x0172 00370 (1.go:114)	PCDATA	$0, $2
	0x0172 00370 (1.go:114)	CALL	runtime.writebarrierptr(SB)
	0x0177 00375 (1.go:115)	MOVQ	"".&u+48(SP), AX
	0x017c 00380 (1.go:115)	MOVQ	""..autotmp_193+32(SP), SI
	0x0181 00385 (1.go:115)	JMP	165
	0x0186 00390 (1.go:115)	NOP
	0x0186 00390 (1.go:111)	PCDATA	$0, $-1
	0x0186 00390 (1.go:111)	CALL	runtime.morestack_noctxt(SB)
	0x018b 00395 (1.go:111)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 6c 01 00 00 48 83 ec 40 48 89  H;a...l...H..@H.
	0x0020 6c 24 38 48 8d 6c 24 38 31 c0 48 89 44 24 10 48  l$8H.l$81.H.D$.H
	0x0030 3d 20 4e 00 00 0f 8d b2 00 00 00 48 8d 0d 00 00  = N........H....
	0x0040 00 00 48 89 0c 24 e8 00 00 00 00 48 8b 44 24 08  ..H..$.....H.D$.
	0x0050 48 89 44 24 30 48 8b 0d 00 00 00 00 48 89 4c 24  H.D$0H......H.L$
	0x0060 28 48 8b 15 00 00 00 00 48 8b 1d 00 00 00 00 48  (H......H......H
	0x0070 8b 35 00 00 00 00 48 89 18 48 89 70 10 8b 1d 00  .5....H..H.p....
	0x0080 00 00 00 48 8d 70 08 48 89 74 24 20 48 8d 78 18  ...H.p.H.t$ H.x.
	0x0090 48 89 7c 24 18 85 db 0f 85 b4 00 00 00 48 89 50  H.|$.........H.P
	0x00a0 08 48 89 48 18 48 c7 40 10 08 00 00 00 8b 0d 00  .H.H.H.@........
	0x00b0 00 00 00 85 c9 75 74 48 8d 0d 00 00 00 00 48 89  .....utH......H.
	0x00c0 48 08 8b 15 00 00 00 00 85 d2 75 41 48 89 05 00  H.........uAH...
	0x00d0 00 00 00 48 8b 54 24 10 48 8d 42 01 48 89 44 24  ...H.T$.H.B.H.D$
	0x00e0 10 48 3d 20 4e 00 00 0f 8c 4e ff ff ff 48 8b 44  .H= N....N...H.D
	0x00f0 24 48 48 8b 08 48 ff c1 48 89 08 48 c7 04 24 10  $HH..H..H..H..$.
	0x0100 27 00 00 e8 00 00 00 00 e9 1b ff ff ff 48 8d 15  '............H..
	0x0110 00 00 00 00 48 89 14 24 48 89 44 24 08 e8 00 00  ....H..$H.D$....
	0x0120 00 00 48 8d 0d 00 00 00 00 eb a8 48 89 34 24 48  ..H........H.4$H
	0x0130 8d 0d 00 00 00 00 48 89 4c 24 08 e8 00 00 00 00  ......H.L$......
	0x0140 48 8b 44 24 30 48 8d 0d 00 00 00 00 e9 71 ff ff  H.D$0H.......q..
	0x0150 ff 48 89 34 24 48 89 54 24 08 e8 00 00 00 00 48  .H.4$H.T$......H
	0x0160 8b 44 24 18 48 89 04 24 48 8b 44 24 28 48 89 44  .D$.H..$H.D$(H.D
	0x0170 24 08 e8 00 00 00 00 48 8b 44 24 30 48 8b 74 24  $......H.D$0H.t$
	0x0180 20 e9 1f ff ff ff e8 00 00 00 00 e9 70 fe ff ff   ...........p...
	rel 12+4 t=16 TLS+0
	rel 62+4 t=15 type."".User+0
	rel 71+4 t=8 runtime.newobject+0
	rel 88+4 t=15 "".statictmp_191+24
	rel 100+4 t=15 "".statictmp_191+8
	rel 107+4 t=15 "".statictmp_191+0
	rel 114+4 t=15 "".statictmp_191+16
	rel 127+4 t=15 runtime.writeBarrier+0
	rel 175+4 t=15 runtime.writeBarrier+0
	rel 186+4 t=15 go.string."lkaskjas"+0
	rel 196+4 t=15 runtime.writeBarrier+0
	rel 207+4 t=15 "".uuu+0
	rel 260+4 t=8 time.Sleep+0
	rel 272+4 t=15 "".uuu+0
	rel 286+4 t=8 runtime.writebarrierptr+0
	rel 293+4 t=15 go.string."lkaskjas"+0
	rel 306+4 t=15 go.string."lkaskjas"+0
	rel 316+4 t=8 runtime.writebarrierptr+0
	rel 328+4 t=15 go.string."lkaskjas"+0
	rel 347+4 t=8 runtime.writebarrierptr+0
	rel 371+4 t=8 runtime.writebarrierptr+0
	rel 391+4 t=8 runtime.morestack_noctxt+0
"".main.func6 t=1 size=182 args=0x8 locals=0x20
	0x0000 00000 (1.go:126)	TEXT	"".main.func6(SB), $32-8
	0x0000 00000 (1.go:126)	MOVQ	TLS, CX
	0x0009 00009 (1.go:126)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:126)	CMPQ	SP, 16(CX)
	0x0014 00020 (1.go:126)	JLS	172
	0x001a 00026 (1.go:126)	SUBQ	$32, SP
	0x001e 00030 (1.go:126)	MOVQ	BP, 24(SP)
	0x0023 00035 (1.go:126)	LEAQ	24(SP), BP
	0x0028 00040 (1.go:126)	FUNCDATA	$0, gclocals·2a5305abe05176240e61b8620e19a815(SB)
	0x0028 00040 (1.go:126)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (1.go:128)	MOVQ	$0, AX
	0x002a 00042 (1.go:128)	MOVQ	AX, "".i+16(SP)
	0x002f 00047 (1.go:128)	MOVQ	"".N+40(FP), CX
	0x0034 00052 (1.go:128)	CMPQ	AX, CX
	0x0037 00055 (1.go:128)	JGE	$0, 110
	0x0039 00057 (1.go:129)	MOVQ	$75, "".str+8(SB)
	0x0044 00068 (1.go:129)	MOVL	runtime.writeBarrier(SB), DX
	0x004a 00074 (1.go:129)	TESTL	DX, DX
	0x004c 00076 (1.go:129)	JNE	$0, 125
	0x004e 00078 (1.go:129)	LEAQ	go.string."nkajsdasjdasdjasd asjdbasdhjmabsdasbdasjhdbasdhasbdhjasm asmn asvhcgas cgsa"(SB), DX
	0x0055 00085 (1.go:129)	MOVQ	DX, "".str(SB)
	0x005c 00092 (1.go:128)	INCQ	AX
	0x005f 00095 (1.go:128)	MOVQ	AX, "".i+16(SP)
	0x0064 00100 (1.go:128)	MOVQ	"".N+40(FP), CX
	0x0069 00105 (1.go:128)	CMPQ	AX, CX
	0x006c 00108 (1.go:128)	JLT	$0, 57
	0x006e 00110 (1.go:131)	MOVQ	$10000, (SP)
	0x0076 00118 (1.go:131)	PCDATA	$0, $0
	0x0076 00118 (1.go:131)	CALL	time.Sleep(SB)
	0x007b 00123 (1.go:128)	JMP	40
	0x007d 00125 (1.go:129)	LEAQ	"".str(SB), DX
	0x0084 00132 (1.go:129)	MOVQ	DX, (SP)
	0x0088 00136 (1.go:129)	LEAQ	go.string."nkajsdasjdasdjasd asjdbasdhjmabsdasbdasjhdbasdhasbdhjasm asmn asvhcgas cgsa"(SB), BX
	0x008f 00143 (1.go:129)	MOVQ	BX, 8(SP)
	0x0094 00148 (1.go:129)	PCDATA	$0, $0
	0x0094 00148 (1.go:129)	CALL	runtime.writebarrierptr(SB)
	0x0099 00153 (1.go:128)	MOVQ	"".i+16(SP), AX
	0x009e 00158 (1.go:128)	MOVQ	"".N+40(FP), CX
	0x00a3 00163 (1.go:129)	LEAQ	go.string."nkajsdasjdasdjasd asjdbasdhjmabsdasbdasjhdbasdhasbdhjasm asmn asvhcgas cgsa"(SB), DX
	0x00aa 00170 (1.go:128)	JMP	92
	0x00ac 00172 (1.go:128)	NOP
	0x00ac 00172 (1.go:126)	PCDATA	$0, $-1
	0x00ac 00172 (1.go:126)	CALL	runtime.morestack_noctxt(SB)
	0x00b1 00177 (1.go:126)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 92 00 00 00 48 83 ec 20 48 89  H;a.......H.. H.
	0x0020 6c 24 18 48 8d 6c 24 18 31 c0 48 89 44 24 10 48  l$.H.l$.1.H.D$.H
	0x0030 8b 4c 24 28 48 39 c8 7d 35 48 c7 05 00 00 00 00  .L$(H9.}5H......
	0x0040 4b 00 00 00 8b 15 00 00 00 00 85 d2 75 2f 48 8d  K...........u/H.
	0x0050 15 00 00 00 00 48 89 15 00 00 00 00 48 ff c0 48  .....H......H..H
	0x0060 89 44 24 10 48 8b 4c 24 28 48 39 c8 7c cb 48 c7  .D$.H.L$(H9.|.H.
	0x0070 04 24 10 27 00 00 e8 00 00 00 00 eb ab 48 8d 15  .$.'.........H..
	0x0080 00 00 00 00 48 89 14 24 48 8d 1d 00 00 00 00 48  ....H..$H......H
	0x0090 89 5c 24 08 e8 00 00 00 00 48 8b 44 24 10 48 8b  .\$......H.D$.H.
	0x00a0 4c 24 28 48 8d 15 00 00 00 00 eb b0 e8 00 00 00  L$(H............
	0x00b0 00 e9 4a ff ff ff                                ..J...
	rel 12+4 t=16 TLS+0
	rel 60+4 t=15 "".str+4
	rel 70+4 t=15 runtime.writeBarrier+0
	rel 81+4 t=15 go.string."nkajsdasjdasdjasd asjdbasdhjmabsdasbdasjhdbasdhasbdhjasm asmn asvhcgas cgsa"+0
	rel 88+4 t=15 "".str+0
	rel 119+4 t=8 time.Sleep+0
	rel 128+4 t=15 "".str+0
	rel 139+4 t=15 go.string."nkajsdasjdasdjasd asjdbasdhjmabsdasbdasjhdbasdhasbdhjasm asmn asvhcgas cgsa"+0
	rel 149+4 t=8 runtime.writebarrierptr+0
	rel 166+4 t=15 go.string."nkajsdasjdasdjasd asjdbasdhjmabsdasbdasjhdbasdhasbdhjasm asmn asvhcgas cgsa"+0
	rel 173+4 t=8 runtime.morestack_noctxt+0
"".main.func7 t=1 size=923 args=0x18 locals=0x100
	0x0000 00000 (1.go:138)	TEXT	"".main.func7(SB), $256-24
	0x0000 00000 (1.go:138)	MOVQ	TLS, CX
	0x0009 00009 (1.go:138)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:138)	LEAQ	-128(SP), AX
	0x0015 00021 (1.go:138)	CMPQ	AX, 16(CX)
	0x0019 00025 (1.go:138)	JLS	913
	0x001f 00031 (1.go:138)	SUBQ	$256, SP
	0x0026 00038 (1.go:138)	MOVQ	BP, 248(SP)
	0x002e 00046 (1.go:138)	LEAQ	248(SP), BP
	0x0036 00054 (1.go:138)	FUNCDATA	$0, gclocals·283d2e5c825c49c60b326a4265accb5a(SB)
	0x0036 00054 (1.go:138)	FUNCDATA	$1, gclocals·a7899340dc00631f39a28dacb9997ba7(SB)
	0x0036 00054 (1.go:140)	MOVQ	"".&z+264(FP), AX
	0x003e 00062 (1.go:140)	MOVQ	(AX), CX
	0x0041 00065 (1.go:140)	INCQ	CX
	0x0044 00068 (1.go:140)	MOVQ	CX, (AX)
	0x0047 00071 (1.go:141)	LEAQ	go.string."\n\n"(SB), CX
	0x004e 00078 (1.go:141)	MOVQ	CX, ""..autotmp_198+120(SP)
	0x0053 00083 (1.go:141)	MOVQ	$2, ""..autotmp_198+128(SP)
	0x005f 00095 (1.go:141)	MOVQ	$0, ""..autotmp_197+136(SP)
	0x006b 00107 (1.go:141)	MOVQ	$0, ""..autotmp_197+144(SP)
	0x0077 00119 (1.go:141)	LEAQ	type.string(SB), DX
	0x007e 00126 (1.go:141)	MOVQ	DX, (SP)
	0x0082 00130 (1.go:141)	LEAQ	""..autotmp_198+120(SP), BX
	0x0087 00135 (1.go:141)	MOVQ	BX, 8(SP)
	0x008c 00140 (1.go:141)	PCDATA	$0, $1
	0x008c 00140 (1.go:141)	CALL	runtime.convT2E(SB)
	0x0091 00145 (1.go:141)	MOVQ	24(SP), AX
	0x0096 00150 (1.go:141)	MOVQ	16(SP), CX
	0x009b 00155 (1.go:141)	MOVQ	CX, ""..autotmp_197+136(SP)
	0x00a3 00163 (1.go:141)	MOVQ	AX, ""..autotmp_197+144(SP)
	0x00ab 00171 (1.go:141)	LEAQ	""..autotmp_197+136(SP), AX
	0x00b3 00179 (1.go:141)	MOVQ	AX, (SP)
	0x00b7 00183 (1.go:141)	MOVQ	$1, 8(SP)
	0x00c0 00192 (1.go:141)	MOVQ	$1, 16(SP)
	0x00c9 00201 (1.go:141)	PCDATA	$0, $1
	0x00c9 00201 (1.go:141)	CALL	fmt.Println(SB)
	0x00ce 00206 (1.go:142)	LEAQ	go.string."N: "(SB), AX
	0x00d5 00213 (1.go:142)	MOVQ	AX, ""..autotmp_200+104(SP)
	0x00da 00218 (1.go:142)	MOVQ	$3, ""..autotmp_200+112(SP)
	0x00e3 00227 (1.go:142)	MOVQ	"".&z+264(FP), CX
	0x00eb 00235 (1.go:142)	MOVQ	(CX), DX
	0x00ee 00238 (1.go:142)	MOVQ	DX, ""..autotmp_201+64(SP)
	0x00f3 00243 (1.go:142)	MOVQ	$0, ""..autotmp_199+216(SP)
	0x00ff 00255 (1.go:142)	MOVQ	$0, ""..autotmp_199+224(SP)
	0x010b 00267 (1.go:142)	MOVQ	$0, ""..autotmp_199+232(SP)
	0x0117 00279 (1.go:142)	MOVQ	$0, ""..autotmp_199+240(SP)
	0x0123 00291 (1.go:141)	LEAQ	type.string(SB), DX
	0x012a 00298 (1.go:142)	MOVQ	DX, (SP)
	0x012e 00302 (1.go:142)	LEAQ	""..autotmp_200+104(SP), BX
	0x0133 00307 (1.go:142)	MOVQ	BX, 8(SP)
	0x0138 00312 (1.go:142)	PCDATA	$0, $2
	0x0138 00312 (1.go:142)	CALL	runtime.convT2E(SB)
	0x013d 00317 (1.go:142)	MOVQ	24(SP), AX
	0x0142 00322 (1.go:142)	MOVQ	16(SP), CX
	0x0147 00327 (1.go:142)	MOVQ	CX, ""..autotmp_199+216(SP)
	0x014f 00335 (1.go:142)	MOVQ	AX, ""..autotmp_199+224(SP)
	0x0157 00343 (1.go:142)	LEAQ	type.int(SB), AX
	0x015e 00350 (1.go:142)	MOVQ	AX, (SP)
	0x0162 00354 (1.go:142)	LEAQ	""..autotmp_201+64(SP), CX
	0x0167 00359 (1.go:142)	MOVQ	CX, 8(SP)
	0x016c 00364 (1.go:142)	PCDATA	$0, $2
	0x016c 00364 (1.go:142)	CALL	runtime.convT2E(SB)
	0x0171 00369 (1.go:142)	MOVQ	24(SP), AX
	0x0176 00374 (1.go:142)	MOVQ	16(SP), CX
	0x017b 00379 (1.go:142)	MOVQ	CX, ""..autotmp_199+232(SP)
	0x0183 00387 (1.go:142)	MOVQ	AX, ""..autotmp_199+240(SP)
	0x018b 00395 (1.go:142)	LEAQ	""..autotmp_199+216(SP), AX
	0x0193 00403 (1.go:142)	MOVQ	AX, (SP)
	0x0197 00407 (1.go:142)	MOVQ	$2, 8(SP)
	0x01a0 00416 (1.go:142)	MOVQ	$2, 16(SP)
	0x01a9 00425 (1.go:142)	PCDATA	$0, $2
	0x01a9 00425 (1.go:142)	CALL	fmt.Println(SB)
	0x01ae 00430 (1.go:143)	LEAQ	go.string."g: "(SB), AX
	0x01b5 00437 (1.go:143)	MOVQ	AX, ""..autotmp_203+88(SP)
	0x01ba 00442 (1.go:143)	MOVQ	$3, ""..autotmp_203+96(SP)
	0x01c3 00451 (1.go:143)	MOVQ	"".&g+272(FP), CX
	0x01cb 00459 (1.go:143)	MOVQ	(CX), DX
	0x01ce 00462 (1.go:143)	MOVQ	DX, ""..autotmp_204+56(SP)
	0x01d3 00467 (1.go:143)	MOVQ	$0, ""..autotmp_202+184(SP)
	0x01df 00479 (1.go:143)	MOVQ	$0, ""..autotmp_202+192(SP)
	0x01eb 00491 (1.go:143)	MOVQ	$0, ""..autotmp_202+200(SP)
	0x01f7 00503 (1.go:143)	MOVQ	$0, ""..autotmp_202+208(SP)
	0x0203 00515 (1.go:141)	LEAQ	type.string(SB), DX
	0x020a 00522 (1.go:143)	MOVQ	DX, (SP)
	0x020e 00526 (1.go:143)	LEAQ	""..autotmp_203+88(SP), BX
	0x0213 00531 (1.go:143)	MOVQ	BX, 8(SP)
	0x0218 00536 (1.go:143)	PCDATA	$0, $3
	0x0218 00536 (1.go:143)	CALL	runtime.convT2E(SB)
	0x021d 00541 (1.go:143)	MOVQ	16(SP), AX
	0x0222 00546 (1.go:143)	MOVQ	24(SP), CX
	0x0227 00551 (1.go:143)	MOVQ	AX, ""..autotmp_202+184(SP)
	0x022f 00559 (1.go:143)	MOVQ	CX, ""..autotmp_202+192(SP)
	0x0237 00567 (1.go:142)	LEAQ	type.int(SB), AX
	0x023e 00574 (1.go:143)	MOVQ	AX, (SP)
	0x0242 00578 (1.go:143)	LEAQ	""..autotmp_204+56(SP), CX
	0x0247 00583 (1.go:143)	MOVQ	CX, 8(SP)
	0x024c 00588 (1.go:143)	PCDATA	$0, $3
	0x024c 00588 (1.go:143)	CALL	runtime.convT2E(SB)
	0x0251 00593 (1.go:143)	MOVQ	24(SP), AX
	0x0256 00598 (1.go:143)	MOVQ	16(SP), CX
	0x025b 00603 (1.go:143)	MOVQ	CX, ""..autotmp_202+200(SP)
	0x0263 00611 (1.go:143)	MOVQ	AX, ""..autotmp_202+208(SP)
	0x026b 00619 (1.go:143)	LEAQ	""..autotmp_202+184(SP), AX
	0x0273 00627 (1.go:143)	MOVQ	AX, (SP)
	0x0277 00631 (1.go:143)	MOVQ	$2, 8(SP)
	0x0280 00640 (1.go:143)	MOVQ	$2, 16(SP)
	0x0289 00649 (1.go:143)	PCDATA	$0, $3
	0x0289 00649 (1.go:143)	CALL	fmt.Println(SB)
	0x028e 00654 (1.go:144)	LEAQ	go.string."Map Len: "(SB), AX
	0x0295 00661 (1.go:144)	MOVQ	AX, ""..autotmp_206+72(SP)
	0x029a 00666 (1.go:144)	MOVQ	$9, ""..autotmp_206+80(SP)
	0x02a3 00675 (1.go:144)	MOVQ	"".lenMM+280(FP), DX
	0x02ab 00683 (1.go:144)	MOVQ	(DX), CX
	0x02ae 00686 (1.go:144)	PCDATA	$0, $4
	0x02ae 00686 (1.go:144)	CALL	CX
	0x02b0 00688 (1.go:144)	MOVQ	(SP), AX
	0x02b4 00692 (1.go:144)	MOVQ	AX, ""..autotmp_207+48(SP)
	0x02b9 00697 (1.go:144)	MOVQ	$0, ""..autotmp_205+152(SP)
	0x02c5 00709 (1.go:144)	MOVQ	$0, ""..autotmp_205+160(SP)
	0x02d1 00721 (1.go:144)	MOVQ	$0, ""..autotmp_205+168(SP)
	0x02dd 00733 (1.go:144)	MOVQ	$0, ""..autotmp_205+176(SP)
	0x02e9 00745 (1.go:141)	LEAQ	type.string(SB), AX
	0x02f0 00752 (1.go:144)	MOVQ	AX, (SP)
	0x02f4 00756 (1.go:144)	LEAQ	""..autotmp_206+72(SP), CX
	0x02f9 00761 (1.go:144)	MOVQ	CX, 8(SP)
	0x02fe 00766 (1.go:144)	PCDATA	$0, $5
	0x02fe 00766 (1.go:144)	CALL	runtime.convT2E(SB)
	0x0303 00771 (1.go:144)	MOVQ	16(SP), AX
	0x0308 00776 (1.go:144)	MOVQ	24(SP), CX
	0x030d 00781 (1.go:144)	MOVQ	AX, ""..autotmp_205+152(SP)
	0x0315 00789 (1.go:144)	MOVQ	CX, ""..autotmp_205+160(SP)
	0x031d 00797 (1.go:142)	LEAQ	type.int(SB), AX
	0x0324 00804 (1.go:144)	MOVQ	AX, (SP)
	0x0328 00808 (1.go:144)	LEAQ	""..autotmp_207+48(SP), CX
	0x032d 00813 (1.go:144)	MOVQ	CX, 8(SP)
	0x0332 00818 (1.go:144)	PCDATA	$0, $5
	0x0332 00818 (1.go:144)	CALL	runtime.convT2E(SB)
	0x0337 00823 (1.go:144)	MOVQ	16(SP), AX
	0x033c 00828 (1.go:144)	MOVQ	24(SP), CX
	0x0341 00833 (1.go:144)	MOVQ	AX, ""..autotmp_205+168(SP)
	0x0349 00841 (1.go:144)	MOVQ	CX, ""..autotmp_205+176(SP)
	0x0351 00849 (1.go:144)	LEAQ	""..autotmp_205+152(SP), AX
	0x0359 00857 (1.go:144)	MOVQ	AX, (SP)
	0x035d 00861 (1.go:144)	MOVQ	$2, 8(SP)
	0x0366 00870 (1.go:144)	MOVQ	$2, 16(SP)
	0x036f 00879 (1.go:144)	PCDATA	$0, $5
	0x036f 00879 (1.go:144)	CALL	fmt.Println(SB)
	0x0374 00884 (1.go:145)	PCDATA	$0, $0
	0x0374 00884 (1.go:145)	CALL	"".printMem(SB)
	0x0379 00889 (1.go:146)	MOVQ	$5000000000, AX
	0x0383 00899 (1.go:146)	MOVQ	AX, (SP)
	0x0387 00903 (1.go:146)	PCDATA	$0, $0
	0x0387 00903 (1.go:146)	CALL	time.Sleep(SB)
	0x038c 00908 (1.go:140)	JMP	54
	0x0391 00913 (1.go:140)	NOP
	0x0391 00913 (1.go:138)	PCDATA	$0, $-1
	0x0391 00913 (1.go:138)	CALL	runtime.morestack_noctxt(SB)
	0x0396 00918 (1.go:138)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 8d 44 24 80 48 3b 41 10 0f 86 72 03 00 00 48  H.D$.H;A...r...H
	0x0020 81 ec 00 01 00 00 48 89 ac 24 f8 00 00 00 48 8d  ......H..$....H.
	0x0030 ac 24 f8 00 00 00 48 8b 84 24 08 01 00 00 48 8b  .$....H..$....H.
	0x0040 08 48 ff c1 48 89 08 48 8d 0d 00 00 00 00 48 89  .H..H..H......H.
	0x0050 4c 24 78 48 c7 84 24 80 00 00 00 02 00 00 00 48  L$xH..$........H
	0x0060 c7 84 24 88 00 00 00 00 00 00 00 48 c7 84 24 90  ..$........H..$.
	0x0070 00 00 00 00 00 00 00 48 8d 15 00 00 00 00 48 89  .......H......H.
	0x0080 14 24 48 8d 5c 24 78 48 89 5c 24 08 e8 00 00 00  .$H.\$xH.\$.....
	0x0090 00 48 8b 44 24 18 48 8b 4c 24 10 48 89 8c 24 88  .H.D$.H.L$.H..$.
	0x00a0 00 00 00 48 89 84 24 90 00 00 00 48 8d 84 24 88  ...H..$....H..$.
	0x00b0 00 00 00 48 89 04 24 48 c7 44 24 08 01 00 00 00  ...H..$H.D$.....
	0x00c0 48 c7 44 24 10 01 00 00 00 e8 00 00 00 00 48 8d  H.D$..........H.
	0x00d0 05 00 00 00 00 48 89 44 24 68 48 c7 44 24 70 03  .....H.D$hH.D$p.
	0x00e0 00 00 00 48 8b 8c 24 08 01 00 00 48 8b 11 48 89  ...H..$....H..H.
	0x00f0 54 24 40 48 c7 84 24 d8 00 00 00 00 00 00 00 48  T$@H..$........H
	0x0100 c7 84 24 e0 00 00 00 00 00 00 00 48 c7 84 24 e8  ..$........H..$.
	0x0110 00 00 00 00 00 00 00 48 c7 84 24 f0 00 00 00 00  .......H..$.....
	0x0120 00 00 00 48 8d 15 00 00 00 00 48 89 14 24 48 8d  ...H......H..$H.
	0x0130 5c 24 68 48 89 5c 24 08 e8 00 00 00 00 48 8b 44  \$hH.\$......H.D
	0x0140 24 18 48 8b 4c 24 10 48 89 8c 24 d8 00 00 00 48  $.H.L$.H..$....H
	0x0150 89 84 24 e0 00 00 00 48 8d 05 00 00 00 00 48 89  ..$....H......H.
	0x0160 04 24 48 8d 4c 24 40 48 89 4c 24 08 e8 00 00 00  .$H.L$@H.L$.....
	0x0170 00 48 8b 44 24 18 48 8b 4c 24 10 48 89 8c 24 e8  .H.D$.H.L$.H..$.
	0x0180 00 00 00 48 89 84 24 f0 00 00 00 48 8d 84 24 d8  ...H..$....H..$.
	0x0190 00 00 00 48 89 04 24 48 c7 44 24 08 02 00 00 00  ...H..$H.D$.....
	0x01a0 48 c7 44 24 10 02 00 00 00 e8 00 00 00 00 48 8d  H.D$..........H.
	0x01b0 05 00 00 00 00 48 89 44 24 58 48 c7 44 24 60 03  .....H.D$XH.D$`.
	0x01c0 00 00 00 48 8b 8c 24 10 01 00 00 48 8b 11 48 89  ...H..$....H..H.
	0x01d0 54 24 38 48 c7 84 24 b8 00 00 00 00 00 00 00 48  T$8H..$........H
	0x01e0 c7 84 24 c0 00 00 00 00 00 00 00 48 c7 84 24 c8  ..$........H..$.
	0x01f0 00 00 00 00 00 00 00 48 c7 84 24 d0 00 00 00 00  .......H..$.....
	0x0200 00 00 00 48 8d 15 00 00 00 00 48 89 14 24 48 8d  ...H......H..$H.
	0x0210 5c 24 58 48 89 5c 24 08 e8 00 00 00 00 48 8b 44  \$XH.\$......H.D
	0x0220 24 10 48 8b 4c 24 18 48 89 84 24 b8 00 00 00 48  $.H.L$.H..$....H
	0x0230 89 8c 24 c0 00 00 00 48 8d 05 00 00 00 00 48 89  ..$....H......H.
	0x0240 04 24 48 8d 4c 24 38 48 89 4c 24 08 e8 00 00 00  .$H.L$8H.L$.....
	0x0250 00 48 8b 44 24 18 48 8b 4c 24 10 48 89 8c 24 c8  .H.D$.H.L$.H..$.
	0x0260 00 00 00 48 89 84 24 d0 00 00 00 48 8d 84 24 b8  ...H..$....H..$.
	0x0270 00 00 00 48 89 04 24 48 c7 44 24 08 02 00 00 00  ...H..$H.D$.....
	0x0280 48 c7 44 24 10 02 00 00 00 e8 00 00 00 00 48 8d  H.D$..........H.
	0x0290 05 00 00 00 00 48 89 44 24 48 48 c7 44 24 50 09  .....H.D$HH.D$P.
	0x02a0 00 00 00 48 8b 94 24 18 01 00 00 48 8b 0a ff d1  ...H..$....H....
	0x02b0 48 8b 04 24 48 89 44 24 30 48 c7 84 24 98 00 00  H..$H.D$0H..$...
	0x02c0 00 00 00 00 00 48 c7 84 24 a0 00 00 00 00 00 00  .....H..$.......
	0x02d0 00 48 c7 84 24 a8 00 00 00 00 00 00 00 48 c7 84  .H..$........H..
	0x02e0 24 b0 00 00 00 00 00 00 00 48 8d 05 00 00 00 00  $........H......
	0x02f0 48 89 04 24 48 8d 4c 24 48 48 89 4c 24 08 e8 00  H..$H.L$HH.L$...
	0x0300 00 00 00 48 8b 44 24 10 48 8b 4c 24 18 48 89 84  ...H.D$.H.L$.H..
	0x0310 24 98 00 00 00 48 89 8c 24 a0 00 00 00 48 8d 05  $....H..$....H..
	0x0320 00 00 00 00 48 89 04 24 48 8d 4c 24 30 48 89 4c  ....H..$H.L$0H.L
	0x0330 24 08 e8 00 00 00 00 48 8b 44 24 10 48 8b 4c 24  $......H.D$.H.L$
	0x0340 18 48 89 84 24 a8 00 00 00 48 89 8c 24 b0 00 00  .H..$....H..$...
	0x0350 00 48 8d 84 24 98 00 00 00 48 89 04 24 48 c7 44  .H..$....H..$H.D
	0x0360 24 08 02 00 00 00 48 c7 44 24 10 02 00 00 00 e8  $.....H.D$......
	0x0370 00 00 00 00 e8 00 00 00 00 48 b8 00 f2 05 2a 01  .........H....*.
	0x0380 00 00 00 48 89 04 24 e8 00 00 00 00 e9 a5 fc ff  ...H..$.........
	0x0390 ff e8 00 00 00 00 e9 65 fc ff ff                 .......e...
	rel 12+4 t=16 TLS+0
	rel 74+4 t=15 go.string."\n\n"+0
	rel 122+4 t=15 type.string+0
	rel 141+4 t=8 runtime.convT2E+0
	rel 202+4 t=8 fmt.Println+0
	rel 209+4 t=15 go.string."N: "+0
	rel 294+4 t=15 type.string+0
	rel 313+4 t=8 runtime.convT2E+0
	rel 346+4 t=15 type.int+0
	rel 365+4 t=8 runtime.convT2E+0
	rel 426+4 t=8 fmt.Println+0
	rel 433+4 t=15 go.string."g: "+0
	rel 518+4 t=15 type.string+0
	rel 537+4 t=8 runtime.convT2E+0
	rel 570+4 t=15 type.int+0
	rel 589+4 t=8 runtime.convT2E+0
	rel 650+4 t=8 fmt.Println+0
	rel 657+4 t=15 go.string."Map Len: "+0
	rel 686+0 t=11 +0
	rel 748+4 t=15 type.string+0
	rel 767+4 t=8 runtime.convT2E+0
	rel 800+4 t=15 type.int+0
	rel 819+4 t=8 runtime.convT2E+0
	rel 880+4 t=8 fmt.Println+0
	rel 885+4 t=8 "".printMem+0
	rel 904+4 t=8 time.Sleep+0
	rel 914+4 t=8 runtime.morestack_noctxt+0
"".init t=1 size=136 args=0x0 locals=0x8
	0x0000 00000 (1.go:187)	TEXT	"".init(SB), $8-0
	0x0000 00000 (1.go:187)	MOVQ	TLS, CX
	0x0009 00009 (1.go:187)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:187)	CMPQ	SP, 16(CX)
	0x0014 00020 (1.go:187)	JLS	126
	0x0016 00022 (1.go:187)	SUBQ	$8, SP
	0x001a 00026 (1.go:187)	MOVQ	BP, (SP)
	0x001e 00030 (1.go:187)	LEAQ	(SP), BP
	0x0022 00034 (1.go:187)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0022 00034 (1.go:187)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0022 00034 (1.go:187)	MOVBLZX	"".initdone·(SB), AX
	0x0029 00041 (1.go:187)	CMPB	AL, $1
	0x002b 00043 (1.go:187)	JLS	$0, 54
	0x002d 00045 (1.go:187)	MOVQ	(SP), BP
	0x0031 00049 (1.go:187)	ADDQ	$8, SP
	0x0035 00053 (1.go:187)	RET
	0x0036 00054 (1.go:187)	JNE	$0, 63
	0x0038 00056 (1.go:187)	PCDATA	$0, $0
	0x0038 00056 (1.go:187)	CALL	runtime.throwinit(SB)
	0x003d 00061 (1.go:187)	UNDEF
	0x003f 00063 (1.go:187)	MOVB	$1, "".initdone·(SB)
	0x0046 00070 (1.go:187)	PCDATA	$0, $0
	0x0046 00070 (1.go:187)	CALL	fmt.init(SB)
	0x004b 00075 (1.go:187)	PCDATA	$0, $0
	0x004b 00075 (1.go:187)	CALL	math/rand.init(SB)
	0x0050 00080 (1.go:187)	PCDATA	$0, $0
	0x0050 00080 (1.go:187)	CALL	runtime.init(SB)
	0x0055 00085 (1.go:187)	PCDATA	$0, $0
	0x0055 00085 (1.go:187)	CALL	strconv.init(SB)
	0x005a 00090 (1.go:187)	PCDATA	$0, $0
	0x005a 00090 (1.go:187)	CALL	sync.init(SB)
	0x005f 00095 (1.go:187)	PCDATA	$0, $0
	0x005f 00095 (1.go:187)	CALL	time.init(SB)
	0x0064 00100 (1.go:187)	PCDATA	$0, $0
	0x0064 00100 (1.go:187)	CALL	runtime/debug.init(SB)
	0x0069 00105 (1.go:187)	PCDATA	$0, $0
	0x0069 00105 (1.go:187)	CALL	"".init.1(SB)
	0x006e 00110 (1.go:187)	MOVB	$2, "".initdone·(SB)
	0x0075 00117 (1.go:187)	MOVQ	(SP), BP
	0x0079 00121 (1.go:187)	ADDQ	$8, SP
	0x007d 00125 (1.go:187)	RET
	0x007e 00126 (1.go:187)	NOP
	0x007e 00126 (1.go:187)	PCDATA	$0, $-1
	0x007e 00126 (1.go:187)	CALL	runtime.morestack_noctxt(SB)
	0x0083 00131 (1.go:187)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 76 68 48 83 ec 08 48 89 2c 24 48 8d  H;a.vhH...H.,$H.
	0x0020 2c 24 0f b6 05 00 00 00 00 3c 01 76 09 48 8b 2c  ,$.......<.v.H.,
	0x0030 24 48 83 c4 08 c3 75 07 e8 00 00 00 00 0f 0b c6  $H....u.........
	0x0040 05 00 00 00 00 01 e8 00 00 00 00 e8 00 00 00 00  ................
	0x0050 e8 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 e8  ................
	0x0060 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 c6 05  ................
	0x0070 00 00 00 00 02 48 8b 2c 24 48 83 c4 08 c3 e8 00  .....H.,$H......
	0x0080 00 00 00 e9 78 ff ff ff                          ....x...
	rel 12+4 t=16 TLS+0
	rel 37+4 t=15 "".initdone·+0
	rel 57+4 t=8 runtime.throwinit+0
	rel 65+4 t=15 "".initdone·+-1
	rel 71+4 t=8 fmt.init+0
	rel 76+4 t=8 math/rand.init+0
	rel 81+4 t=8 runtime.init+0
	rel 86+4 t=8 strconv.init+0
	rel 91+4 t=8 sync.init+0
	rel 96+4 t=8 time.init+0
	rel 101+4 t=8 runtime/debug.init+0
	rel 106+4 t=8 "".init.1+0
	rel 112+4 t=15 "".initdone·+-1
	rel 127+4 t=8 runtime.morestack_noctxt+0
type..hash."".User t=1 dupok size=168 args=0x18 locals=0x28
	0x0000 00000 (1.go:1)	TEXT	type..hash."".User(SB), $40-24
	0x0000 00000 (1.go:1)	MOVQ	TLS, CX
	0x0009 00009 (1.go:1)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:1)	CMPQ	SP, 16(CX)
	0x0014 00020 (1.go:1)	JLS	158
	0x001a 00026 (1.go:1)	SUBQ	$40, SP
	0x001e 00030 (1.go:1)	MOVQ	BP, 32(SP)
	0x0023 00035 (1.go:1)	LEAQ	32(SP), BP
	0x0028 00040 (1.go:1)	FUNCDATA	$0, gclocals·e6397a44f8e1b6e77d0f200b4fba5269(SB)
	0x0028 00040 (1.go:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0028 00040 (1.go:1)	MOVQ	"".p+48(FP), AX
	0x002d 00045 (1.go:1)	MOVQ	AX, (SP)
	0x0031 00049 (1.go:1)	MOVQ	"".h+56(FP), CX
	0x0036 00054 (1.go:1)	MOVQ	CX, 8(SP)
	0x003b 00059 (1.go:1)	MOVQ	$8, 16(SP)
	0x0044 00068 (1.go:1)	PCDATA	$0, $0
	0x0044 00068 (1.go:1)	CALL	runtime.memhash(SB)
	0x0049 00073 (1.go:1)	MOVQ	24(SP), AX
	0x004e 00078 (1.go:1)	MOVQ	"".p+48(FP), CX
	0x0053 00083 (1.go:1)	LEAQ	8(CX), DX
	0x0057 00087 (1.go:1)	MOVQ	DX, (SP)
	0x005b 00091 (1.go:1)	MOVQ	AX, 8(SP)
	0x0060 00096 (1.go:1)	PCDATA	$0, $0
	0x0060 00096 (1.go:1)	CALL	runtime.strhash(SB)
	0x0065 00101 (1.go:1)	MOVQ	16(SP), AX
	0x006a 00106 (1.go:1)	MOVQ	"".p+48(FP), CX
	0x006f 00111 (1.go:1)	ADDQ	$24, CX
	0x0073 00115 (1.go:1)	MOVQ	CX, (SP)
	0x0077 00119 (1.go:1)	MOVQ	AX, 8(SP)
	0x007c 00124 (1.go:1)	MOVQ	$8, 16(SP)
	0x0085 00133 (1.go:1)	PCDATA	$0, $1
	0x0085 00133 (1.go:1)	CALL	runtime.memhash(SB)
	0x008a 00138 (1.go:1)	MOVQ	24(SP), AX
	0x008f 00143 (1.go:1)	MOVQ	AX, "".~r2+64(FP)
	0x0094 00148 (1.go:1)	MOVQ	32(SP), BP
	0x0099 00153 (1.go:1)	ADDQ	$40, SP
	0x009d 00157 (1.go:1)	RET
	0x009e 00158 (1.go:1)	NOP
	0x009e 00158 (1.go:1)	PCDATA	$0, $-1
	0x009e 00158 (1.go:1)	CALL	runtime.morestack_noctxt(SB)
	0x00a3 00163 (1.go:1)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 84 00 00 00 48 83 ec 28 48 89  H;a.......H..(H.
	0x0020 6c 24 20 48 8d 6c 24 20 48 8b 44 24 30 48 89 04  l$ H.l$ H.D$0H..
	0x0030 24 48 8b 4c 24 38 48 89 4c 24 08 48 c7 44 24 10  $H.L$8H.L$.H.D$.
	0x0040 08 00 00 00 e8 00 00 00 00 48 8b 44 24 18 48 8b  .........H.D$.H.
	0x0050 4c 24 30 48 8d 51 08 48 89 14 24 48 89 44 24 08  L$0H.Q.H..$H.D$.
	0x0060 e8 00 00 00 00 48 8b 44 24 10 48 8b 4c 24 30 48  .....H.D$.H.L$0H
	0x0070 83 c1 18 48 89 0c 24 48 89 44 24 08 48 c7 44 24  ...H..$H.D$.H.D$
	0x0080 10 08 00 00 00 e8 00 00 00 00 48 8b 44 24 18 48  ..........H.D$.H
	0x0090 89 44 24 40 48 8b 6c 24 20 48 83 c4 28 c3 e8 00  .D$@H.l$ H..(...
	0x00a0 00 00 00 e9 58 ff ff ff                          ....X...
	rel 12+4 t=16 TLS+0
	rel 69+4 t=8 runtime.memhash+0
	rel 97+4 t=8 runtime.strhash+0
	rel 134+4 t=8 runtime.memhash+0
	rel 159+4 t=8 runtime.morestack_noctxt+0
type..eq."".User t=1 dupok size=175 args=0x18 locals=0x30
	0x0000 00000 (1.go:1)	TEXT	type..eq."".User(SB), $48-24
	0x0000 00000 (1.go:1)	MOVQ	TLS, CX
	0x0009 00009 (1.go:1)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:1)	CMPQ	SP, 16(CX)
	0x0014 00020 (1.go:1)	JLS	165
	0x001a 00026 (1.go:1)	SUBQ	$48, SP
	0x001e 00030 (1.go:1)	MOVQ	BP, 40(SP)
	0x0023 00035 (1.go:1)	LEAQ	40(SP), BP
	0x0028 00040 (1.go:1)	FUNCDATA	$0, gclocals·8f9cec06d1ae35cc9900c511c5e4bdab(SB)
	0x0028 00040 (1.go:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (1.go:1)	MOVQ	"".p+56(FP), AX
	0x002d 00045 (1.go:1)	MOVQ	(AX), CX
	0x0030 00048 (1.go:1)	MOVQ	"".q+64(FP), DX
	0x0035 00053 (1.go:1)	MOVQ	(DX), BX
	0x0038 00056 (1.go:1)	CMPQ	CX, BX
	0x003b 00059 (1.go:1)	JNE	161
	0x003d 00061 (1.go:1)	MOVQ	16(AX), CX
	0x0041 00065 (1.go:1)	MOVQ	16(DX), BX
	0x0045 00069 (1.go:1)	MOVQ	8(DX), SI
	0x0049 00073 (1.go:1)	MOVQ	8(AX), DI
	0x004d 00077 (1.go:1)	CMPQ	CX, BX
	0x0050 00080 (1.go:1)	JEQ	$0, 120
	0x0052 00082 (1.go:1)	MOVL	$0, CX
	0x0054 00084 (1.go:1)	TESTB	CL, CL
	0x0056 00086 (1.go:1)	JEQ	116
	0x0058 00088 (1.go:1)	MOVQ	24(AX), AX
	0x005c 00092 (1.go:1)	MOVQ	24(DX), CX
	0x0060 00096 (1.go:1)	CMPQ	AX, CX
	0x0063 00099 (1.go:1)	SETEQ	AL
	0x0066 00102 (1.go:1)	MOVB	AL, "".~r2+72(FP)
	0x006a 00106 (1.go:1)	MOVQ	40(SP), BP
	0x006f 00111 (1.go:1)	ADDQ	$48, SP
	0x0073 00115 (1.go:1)	RET
	0x0074 00116 (1.go:1)	MOVL	$0, AX
	0x0076 00118 (1.go:1)	JMP	102
	0x0078 00120 (1.go:1)	MOVQ	DI, (SP)
	0x007c 00124 (1.go:1)	MOVQ	CX, 8(SP)
	0x0081 00129 (1.go:1)	MOVQ	SI, 16(SP)
	0x0086 00134 (1.go:1)	MOVQ	BX, 24(SP)
	0x008b 00139 (1.go:1)	PCDATA	$0, $0
	0x008b 00139 (1.go:1)	CALL	runtime.eqstring(SB)
	0x0090 00144 (1.go:1)	MOVBLZX	32(SP), CX
	0x0095 00149 (1.go:1)	MOVQ	"".p+56(FP), AX
	0x009a 00154 (1.go:1)	MOVQ	"".q+64(FP), DX
	0x009f 00159 (1.go:1)	JMP	84
	0x00a1 00161 (1.go:1)	MOVL	$0, CX
	0x00a3 00163 (1.go:1)	JMP	84
	0x00a5 00165 (1.go:1)	NOP
	0x00a5 00165 (1.go:1)	PCDATA	$0, $-1
	0x00a5 00165 (1.go:1)	CALL	runtime.morestack_noctxt(SB)
	0x00aa 00170 (1.go:1)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 8b 00 00 00 48 83 ec 30 48 89  H;a.......H..0H.
	0x0020 6c 24 28 48 8d 6c 24 28 48 8b 44 24 38 48 8b 08  l$(H.l$(H.D$8H..
	0x0030 48 8b 54 24 40 48 8b 1a 48 39 d9 75 64 48 8b 48  H.T$@H..H9.udH.H
	0x0040 10 48 8b 5a 10 48 8b 72 08 48 8b 78 08 48 39 d9  .H.Z.H.r.H.x.H9.
	0x0050 74 26 31 c9 84 c9 74 1c 48 8b 40 18 48 8b 4a 18  t&1...t.H.@.H.J.
	0x0060 48 39 c8 0f 94 c0 88 44 24 48 48 8b 6c 24 28 48  H9.....D$HH.l$(H
	0x0070 83 c4 30 c3 31 c0 eb ee 48 89 3c 24 48 89 4c 24  ..0.1...H.<$H.L$
	0x0080 08 48 89 74 24 10 48 89 5c 24 18 e8 00 00 00 00  .H.t$.H.\$......
	0x0090 0f b6 4c 24 20 48 8b 44 24 38 48 8b 54 24 40 eb  ..L$ H.D$8H.T$@.
	0x00a0 b3 31 c9 eb af e8 00 00 00 00 e9 51 ff ff ff     .1.........Q...
	rel 12+4 t=16 TLS+0
	rel 140+4 t=8 runtime.eqstring+0
	rel 166+4 t=8 runtime.morestack_noctxt+0
"".(*MM).Lock t=1 dupok size=17 args=0x8 locals=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*MM).Lock(SB), $0-8
	0x0000 00000 (<autogenerated>:1)	FUNCDATA	$0, gclocals·a36216b97439c93dafebe03e7f0808b5(SB)
	0x0000 00000 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (<autogenerated>:1)	MOVQ	""..this+8(FP), AX
	0x0005 00005 (<autogenerated>:1)	TESTB	AL, (AX)
	0x0007 00007 (<autogenerated>:1)	MOVQ	AX, ""..this+8(FP)
	0x000c 00012 (<autogenerated>:1)	JMP	sync.(*RWMutex).Lock(SB)
	0x0000 48 8b 44 24 08 84 00 48 89 44 24 08 e9 00 00 00  H.D$...H.D$.....
	0x0010 00                                               .
	rel 13+4 t=15 sync.(*RWMutex).Lock+0
"".(*MM).RLock t=1 dupok size=17 args=0x8 locals=0x0
	0x0000 00000 (<autogenerated>:2)	TEXT	"".(*MM).RLock(SB), $0-8
	0x0000 00000 (<autogenerated>:2)	FUNCDATA	$0, gclocals·a36216b97439c93dafebe03e7f0808b5(SB)
	0x0000 00000 (<autogenerated>:2)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (<autogenerated>:2)	MOVQ	""..this+8(FP), AX
	0x0005 00005 (<autogenerated>:2)	TESTB	AL, (AX)
	0x0007 00007 (<autogenerated>:2)	MOVQ	AX, ""..this+8(FP)
	0x000c 00012 (<autogenerated>:2)	JMP	sync.(*RWMutex).RLock(SB)
	0x0000 48 8b 44 24 08 84 00 48 89 44 24 08 e9 00 00 00  H.D$...H.D$.....
	0x0010 00                                               .
	rel 13+4 t=15 sync.(*RWMutex).RLock+0
"".(*MM).RLocker t=1 dupok size=35 args=0x18 locals=0x0
	0x0000 00000 (<autogenerated>:3)	TEXT	"".(*MM).RLocker(SB), $0-24
	0x0000 00000 (<autogenerated>:3)	FUNCDATA	$0, gclocals·d4dc2f11db048877dbc0f60a22b4adb3(SB)
	0x0000 00000 (<autogenerated>:3)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (<autogenerated>:3)	MOVQ	""..this+8(FP), AX
	0x0005 00005 (<autogenerated>:3)	TESTB	AL, (AX)
	0x0007 00007 (<autogenerated>:3)	MOVQ	AX, ""..this+8(FP)
	0x000c 00012 (<autogenerated>:3)	MOVQ	$0, "".~r1+16(FP)
	0x0015 00021 (<autogenerated>:3)	MOVQ	$0, "".~r1+24(FP)
	0x001e 00030 (<autogenerated>:3)	JMP	sync.(*RWMutex).RLocker(SB)
	0x0000 48 8b 44 24 08 84 00 48 89 44 24 08 48 c7 44 24  H.D$...H.D$.H.D$
	0x0010 10 00 00 00 00 48 c7 44 24 18 00 00 00 00 e9 00  .....H.D$.......
	0x0020 00 00 00                                         ...
	rel 31+4 t=15 sync.(*RWMutex).RLocker+0
"".(*MM).RUnlock t=1 dupok size=17 args=0x8 locals=0x0
	0x0000 00000 (<autogenerated>:4)	TEXT	"".(*MM).RUnlock(SB), $0-8
	0x0000 00000 (<autogenerated>:4)	FUNCDATA	$0, gclocals·a36216b97439c93dafebe03e7f0808b5(SB)
	0x0000 00000 (<autogenerated>:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (<autogenerated>:4)	MOVQ	""..this+8(FP), AX
	0x0005 00005 (<autogenerated>:4)	TESTB	AL, (AX)
	0x0007 00007 (<autogenerated>:4)	MOVQ	AX, ""..this+8(FP)
	0x000c 00012 (<autogenerated>:4)	JMP	sync.(*RWMutex).RUnlock(SB)
	0x0000 48 8b 44 24 08 84 00 48 89 44 24 08 e9 00 00 00  H.D$...H.D$.....
	0x0010 00                                               .
	rel 13+4 t=15 sync.(*RWMutex).RUnlock+0
"".(*MM).Unlock t=1 dupok size=17 args=0x8 locals=0x0
	0x0000 00000 (<autogenerated>:5)	TEXT	"".(*MM).Unlock(SB), $0-8
	0x0000 00000 (<autogenerated>:5)	FUNCDATA	$0, gclocals·a36216b97439c93dafebe03e7f0808b5(SB)
	0x0000 00000 (<autogenerated>:5)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (<autogenerated>:5)	MOVQ	""..this+8(FP), AX
	0x0005 00005 (<autogenerated>:5)	TESTB	AL, (AX)
	0x0007 00007 (<autogenerated>:5)	MOVQ	AX, ""..this+8(FP)
	0x000c 00012 (<autogenerated>:5)	JMP	sync.(*RWMutex).Unlock(SB)
	0x0000 48 8b 44 24 08 84 00 48 89 44 24 08 e9 00 00 00  H.D$...H.D$.....
	0x0010 00                                               .
	rel 13+4 t=15 sync.(*RWMutex).Unlock+0
type..hash.[2]interface {} t=1 dupok size=126 args=0x18 locals=0x28
	0x0000 00000 (1.go:1)	TEXT	type..hash.[2]interface {}(SB), $40-24
	0x0000 00000 (1.go:1)	MOVQ	TLS, CX
	0x0009 00009 (1.go:1)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:1)	CMPQ	SP, 16(CX)
	0x0014 00020 (1.go:1)	JLS	119
	0x0016 00022 (1.go:1)	SUBQ	$40, SP
	0x001a 00026 (1.go:1)	MOVQ	BP, 32(SP)
	0x001f 00031 (1.go:1)	LEAQ	32(SP), BP
	0x0024 00036 (1.go:1)	FUNCDATA	$0, gclocals·d4dc2f11db048877dbc0f60a22b4adb3(SB)
	0x0024 00036 (1.go:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0024 00036 (1.go:1)	MOVQ	$0, AX
	0x0026 00038 (1.go:1)	MOVQ	"".h+56(FP), CX
	0x002b 00043 (1.go:1)	MOVQ	AX, "".i+24(SP)
	0x0030 00048 (1.go:1)	CMPQ	AX, $2
	0x0034 00052 (1.go:1)	JGE	$0, 104
	0x0036 00054 (1.go:1)	SHLQ	$4, AX
	0x003a 00058 (1.go:1)	MOVQ	"".p+48(FP), BX
	0x003f 00063 (1.go:1)	ADDQ	BX, AX
	0x0042 00066 (1.go:1)	MOVQ	AX, (SP)
	0x0046 00070 (1.go:1)	MOVQ	CX, 8(SP)
	0x004b 00075 (1.go:1)	PCDATA	$0, $0
	0x004b 00075 (1.go:1)	CALL	runtime.nilinterhash(SB)
	0x0050 00080 (1.go:1)	MOVQ	16(SP), CX
	0x0055 00085 (1.go:1)	MOVQ	"".i+24(SP), AX
	0x005a 00090 (1.go:1)	INCQ	AX
	0x005d 00093 (1.go:1)	MOVQ	AX, "".i+24(SP)
	0x0062 00098 (1.go:1)	CMPQ	AX, $2
	0x0066 00102 (1.go:1)	JLT	$0, 54
	0x0068 00104 (1.go:1)	MOVQ	CX, "".~r2+64(FP)
	0x006d 00109 (1.go:1)	MOVQ	32(SP), BP
	0x0072 00114 (1.go:1)	ADDQ	$40, SP
	0x0076 00118 (1.go:1)	RET
	0x0077 00119 (1.go:1)	NOP
	0x0077 00119 (1.go:1)	PCDATA	$0, $-1
	0x0077 00119 (1.go:1)	CALL	runtime.morestack_noctxt(SB)
	0x007c 00124 (1.go:1)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 76 61 48 83 ec 28 48 89 6c 24 20 48  H;a.vaH..(H.l$ H
	0x0020 8d 6c 24 20 31 c0 48 8b 4c 24 38 48 89 44 24 18  .l$ 1.H.L$8H.D$.
	0x0030 48 83 f8 02 7d 32 48 c1 e0 04 48 8b 5c 24 30 48  H...}2H...H.\$0H
	0x0040 01 d8 48 89 04 24 48 89 4c 24 08 e8 00 00 00 00  ..H..$H.L$......
	0x0050 48 8b 4c 24 10 48 8b 44 24 18 48 ff c0 48 89 44  H.L$.H.D$.H..H.D
	0x0060 24 18 48 83 f8 02 7c ce 48 89 4c 24 40 48 8b 6c  $.H...|.H.L$@H.l
	0x0070 24 20 48 83 c4 28 c3 e8 00 00 00 00 eb 82        $ H..(........
	rel 12+4 t=16 TLS+0
	rel 76+4 t=8 runtime.nilinterhash+0
	rel 120+4 t=8 runtime.morestack_noctxt+0
type..eq.[2]interface {} t=1 dupok size=183 args=0x18 locals=0x38
	0x0000 00000 (1.go:1)	TEXT	type..eq.[2]interface {}(SB), $56-24
	0x0000 00000 (1.go:1)	MOVQ	TLS, CX
	0x0009 00009 (1.go:1)	MOVQ	(CX)(TLS*2), CX
	0x0010 00016 (1.go:1)	CMPQ	SP, 16(CX)
	0x0014 00020 (1.go:1)	JLS	173
	0x001a 00026 (1.go:1)	SUBQ	$56, SP
	0x001e 00030 (1.go:1)	MOVQ	BP, 48(SP)
	0x0023 00035 (1.go:1)	LEAQ	48(SP), BP
	0x0028 00040 (1.go:1)	FUNCDATA	$0, gclocals·8f9cec06d1ae35cc9900c511c5e4bdab(SB)
	0x0028 00040 (1.go:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0028 00040 (1.go:1)	MOVQ	$0, AX
	0x002a 00042 (1.go:1)	MOVQ	AX, "".i+40(SP)
	0x002f 00047 (1.go:1)	CMPQ	AX, $2
	0x0033 00051 (1.go:1)	JGE	$0, 143
	0x0035 00053 (1.go:1)	SHLQ	$4, AX
	0x0039 00057 (1.go:1)	MOVQ	"".p+64(FP), DX
	0x003e 00062 (1.go:1)	MOVQ	8(DX)(AX*1), BX
	0x0043 00067 (1.go:1)	MOVQ	(DX)(AX*1), SI
	0x0047 00071 (1.go:1)	MOVQ	"".q+72(FP), DI
	0x004c 00076 (1.go:1)	MOVQ	8(DI)(AX*1), R8
	0x0051 00081 (1.go:1)	MOVQ	(DI)(AX*1), AX
	0x0055 00085 (1.go:1)	CMPQ	SI, AX
	0x0058 00088 (1.go:1)	JNE	$0, 158
	0x005a 00090 (1.go:1)	MOVQ	SI, (SP)
	0x005e 00094 (1.go:1)	MOVQ	BX, 8(SP)
	0x0063 00099 (1.go:1)	MOVQ	AX, 16(SP)
	0x0068 00104 (1.go:1)	MOVQ	R8, 24(SP)
	0x006d 00109 (1.go:1)	PCDATA	$0, $0
	0x006d 00109 (1.go:1)	CALL	runtime.efaceeq(SB)
	0x0072 00114 (1.go:1)	MOVBLZX	32(SP), AX
	0x0077 00119 (1.go:1)	TESTB	AL, AL
	0x0079 00121 (1.go:1)	JEQ	$0, 158
	0x007b 00123 (1.go:1)	MOVQ	"".i+40(SP), CX
	0x0080 00128 (1.go:1)	LEAQ	1(CX), AX
	0x0084 00132 (1.go:1)	MOVQ	AX, "".i+40(SP)
	0x0089 00137 (1.go:1)	CMPQ	AX, $2
	0x008d 00141 (1.go:1)	JLT	$0, 53
	0x008f 00143 (1.go:1)	MOVB	$1, "".~r2+80(FP)
	0x0094 00148 (1.go:1)	MOVQ	48(SP), BP
	0x0099 00153 (1.go:1)	ADDQ	$56, SP
	0x009d 00157 (1.go:1)	RET
	0x009e 00158 (1.go:1)	MOVB	$0, "".~r2+80(FP)
	0x00a3 00163 (1.go:1)	MOVQ	48(SP), BP
	0x00a8 00168 (1.go:1)	ADDQ	$56, SP
	0x00ac 00172 (1.go:1)	RET
	0x00ad 00173 (1.go:1)	NOP
	0x00ad 00173 (1.go:1)	PCDATA	$0, $-1
	0x00ad 00173 (1.go:1)	CALL	runtime.morestack_noctxt(SB)
	0x00b2 00178 (1.go:1)	JMP	0
	0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
	0x0010 48 3b 61 10 0f 86 93 00 00 00 48 83 ec 38 48 89  H;a.......H..8H.
	0x0020 6c 24 30 48 8d 6c 24 30 31 c0 48 89 44 24 28 48  l$0H.l$01.H.D$(H
	0x0030 83 f8 02 7d 5a 48 c1 e0 04 48 8b 54 24 40 48 8b  ...}ZH...H.T$@H.
	0x0040 5c 02 08 48 8b 34 02 48 8b 7c 24 48 4c 8b 44 07  \..H.4.H.|$HL.D.
	0x0050 08 48 8b 04 07 48 39 c6 75 44 48 89 34 24 48 89  .H...H9.uDH.4$H.
	0x0060 5c 24 08 48 89 44 24 10 4c 89 44 24 18 e8 00 00  \$.H.D$.L.D$....
	0x0070 00 00 0f b6 44 24 20 84 c0 74 23 48 8b 4c 24 28  ....D$ ..t#H.L$(
	0x0080 48 8d 41 01 48 89 44 24 28 48 83 f8 02 7c a6 c6  H.A.H.D$(H...|..
	0x0090 44 24 50 01 48 8b 6c 24 30 48 83 c4 38 c3 c6 44  D$P.H.l$0H..8..D
	0x00a0 24 50 00 48 8b 6c 24 30 48 83 c4 38 c3 e8 00 00  $P.H.l$0H..8....
	0x00b0 00 00 e9 49 ff ff ff                             ...I...
	rel 12+4 t=16 TLS+0
	rel 110+4 t=8 runtime.efaceeq+0
	rel 174+4 t=8 runtime.morestack_noctxt+0
gclocals·6dbd0f82d210f8c1253a4f13f77840be t=8 dupok size=32
	0x0000 0c 00 00 00 09 00 00 00 00 00 10 00 30 00 3c 01  ............0.<.
	0x0010 34 01 b4 01 b0 01 b2 01 33 01 32 00 02 00 42 00  4.......3.2...B.
gclocals·c89758d07e85e5e9784036c1fc126388 t=8 dupok size=8
	0x0000 0c 00 00 00 00 00 00 00                          ........
go.info."".main t=45 size=166
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 00 00 00 00 00 00 00 00 01 04 6b 00 05 9c 11  ...........k....
	0x0020 a0 7f 22 00 00 00 00 00 00 00 00 04 64 6f 6e 65  ..".........done
	0x0030 64 00 05 9c 11 a8 7f 22 00 00 00 00 00 00 00 00  d......"........
	0x0040 04 26 6d 6d 00 04 9c 11 48 22 00 00 00 00 00 00  .&mm....H"......
	0x0050 00 00 04 26 69 6e 64 65 78 00 04 9c 11 50 22 00  ...&index....P".
	0x0060 00 00 00 00 00 00 00 04 26 67 00 04 9c 11 58 22  ........&g....X"
	0x0070 00 00 00 00 00 00 00 00 04 26 64 65 55 73 65 72  .........&deUser
	0x0080 32 00 04 9c 11 60 22 00 00 00 00 00 00 00 00 04  2....`".........
	0x0090 26 64 65 55 73 65 72 00 04 9c 11 68 22 00 00 00  &deUser....h"...
	0x00a0 00 00 00 00 00 00                                ......
	rel 9+8 t=1 "".main+0
	rel 17+8 t=1 "".main+1275
	rel 35+8 t=28 go.info.int+0
	rel 56+8 t=28 go.info.chan bool+0
	rel 74+8 t=28 go.info.*"".MM+0
	rel 95+8 t=28 go.info.*int+0
	rel 112+8 t=28 go.info.*int+0
	rel 135+8 t=28 go.info.**"".User+0
	rel 157+8 t=28 go.info.*"".User+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb t=8 dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
go.info."".init.1 t=45 size=29
	0x0000 02 22 22 2e 69 6e 69 74 2e 31 00 00 00 00 00 00  ."".init.1......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 01 00           .............
	rel 11+8 t=1 "".init.1+0
	rel 19+8 t=1 "".init.1+82
go.string."Alloc: " t=8 dupok size=7
	0x0000 41 6c 6c 6f 63 3a 20                             Alloc: 
go.string."EnableGC: " t=8 dupok size=10
	0x0000 45 6e 61 62 6c 65 47 43 3a 20                    EnableGC: 
go.string."Frees: " t=8 dupok size=7
	0x0000 46 72 65 65 73 3a 20                             Frees: 
go.string."HeapAlloc: " t=8 dupok size=11
	0x0000 48 65 61 70 41 6c 6c 6f 63 3a 20                 HeapAlloc: 
go.string."HeapIdle: " t=8 dupok size=10
	0x0000 48 65 61 70 49 64 6c 65 3a 20                    HeapIdle: 
go.string."HeapObjects: " t=8 dupok size=13
	0x0000 48 65 61 70 4f 62 6a 65 63 74 73 3a 20           HeapObjects: 
go.string."HeapReleased: " t=8 dupok size=14
	0x0000 48 65 61 70 52 65 6c 65 61 73 65 64 3a 20        HeapReleased: 
go.string."TotalAlloc: " t=8 dupok size=12
	0x0000 54 6f 74 61 6c 41 6c 6c 6f 63 3a 20              TotalAlloc: 
go.string."StackInuse: " t=8 dupok size=12
	0x0000 53 74 61 63 6b 49 6e 75 73 65 3a 20              StackInuse: 
go.string."Lookups: " t=8 dupok size=9
	0x0000 4c 6f 6f 6b 75 70 73 3a 20                       Lookups: 
go.string."LastGC: " t=8 dupok size=8
	0x0000 4c 61 73 74 47 43 3a 20                          LastGC: 
go.string."NumGC: " t=8 dupok size=7
	0x0000 4e 75 6d 47 43 3a 20                             NumGC: 
go.string."NextGC: " t=8 dupok size=8
	0x0000 4e 65 78 74 47 43 3a 20                          NextGC: 
go.string."Mallocs: " t=8 dupok size=9
	0x0000 4d 61 6c 6c 6f 63 73 3a 20                       Mallocs: 
go.string."PauseTotalNs: " t=8 dupok size=14
	0x0000 50 61 75 73 65 54 6f 74 61 6c 4e 73 3a 20        PauseTotalNs: 
go.string."GCCPUFraction: " t=8 dupok size=15
	0x0000 47 43 43 50 55 46 72 61 63 74 69 6f 6e 3a 20     GCCPUFraction: 
go.string."GCSys: " t=8 dupok size=7
	0x0000 47 43 53 79 73 3a 20                             GCSys: 
go.string."Time: " t=8 dupok size=6
	0x0000 54 69 6d 65 3a 20                                Time: 
go.string."Now: " t=8 dupok size=5
	0x0000 4e 6f 77 3a 20                                   Now: 
gclocals·d5cb556d330213bab77ebaeb5fb03bb3 t=8 dupok size=338
	0x0000 16 00 00 00 72 00 00 00 00 00 00 00 00 00 00 00  ....r...........
	0x0010 00 00 00 00 00 00 00 00 00 00 00 10 00 00 00 00  ................
	0x0020 00 00 00 00 c0 03 00 00 00 00 04 00 00 00 00 00  ................
	0x0030 00 00 00 3c 00 00 00 00 00 01 00 00 00 00 00 00  ...<............
	0x0040 00 c0 03 00 00 00 00 40 00 00 00 00 00 00 00 00  .......@........
	0x0050 3c 00 00 00 00 00 10 00 00 00 00 00 00 00 c0 03  <...............
	0x0060 00 00 00 00 00 04 00 00 00 00 00 00 00 3c 00 00  .............<..
	0x0070 00 00 00 00 01 00 00 00 00 00 00 c0 03 00 00 00  ................
	0x0080 00 00 40 00 00 00 00 00 00 00 3c 00 00 00 00 00  ..@.......<.....
	0x0090 00 10 00 00 00 00 00 00 c0 03 00 00 00 00 00 00  ................
	0x00a0 04 00 00 00 00 00 00 3c 00 00 00 00 00 00 00 01  .......<........
	0x00b0 00 00 00 00 00 c0 03 00 00 00 00 00 00 40 00 00  .............@..
	0x00c0 00 00 00 00 3c 00 00 00 00 00 00 00 10 00 00 00  ....<...........
	0x00d0 00 00 c0 03 00 00 00 00 00 00 00 04 00 00 00 00  ................
	0x00e0 00 3c 00 00 00 00 00 00 00 00 01 00 00 00 00 c0  .<..............
	0x00f0 03 00 00 00 00 00 00 00 40 00 00 00 00 00 3c 00  ........@.....<.
	0x0100 00 00 00 00 00 00 00 10 00 00 00 00 c0 03 00 00  ................
	0x0110 00 00 00 00 00 00 04 00 00 00 00 00 00 00 00 00  ................
	0x0120 00 00 00 00 00 04 00 00 00 00 3c 00 00 00 00 00  ..........<.....
	0x0130 00 00 00 00 01 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0140 00 00 00 01 00 00 00 c0 03 00 00 00 00 00 00 00  ................
	0x0150 00 00                                            ..
gclocals·5ff117978f2c049ee09a6c769745bf8c t=8 dupok size=8
	0x0000 16 00 00 00 00 00 00 00                          ........
go.info."".printMem t=45 size=50
	0x0000 02 22 22 2e 70 72 69 6e 74 4d 65 6d 00 00 00 00  ."".printMem....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 01 04 6d  ...............m
	0x0020 65 6d 00 05 9c 11 d0 4b 22 00 00 00 00 00 00 00  em.....K".......
	0x0030 00 00                                            ..
	rel 13+8 t=1 "".printMem+0
	rel 21+8 t=1 "".printMem+4478
	rel 41+8 t=28 go.info.runtime.MemStats+0
gclocals·9fb7f0986f647f17cb53dda1484e0f7a t=8 dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
gclocals·263043c8f03e3241528dfae4e2812ef4 t=8 dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 00                    ..........
go.info."".main.func1 t=45 size=82
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 31 00 00  ."".main.func1..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 04 6c 00 04 9c 11 60 22 00 00 00 00 00 00 00 00  .l....`"........
	0x0030 04 26 6d 6d 00 04 9c 11 68 22 00 00 00 00 00 00  .&mm....h"......
	0x0040 00 00 05 7e 72 30 00 01 9c 00 00 00 00 00 00 00  ...~r0..........
	0x0050 00 00                                            ..
	rel 15+8 t=1 "".main.func1+0
	rel 23+8 t=1 "".main.func1+116
	rel 40+8 t=28 go.info.int+0
	rel 58+8 t=28 go.info.*"".MM+0
	rel 73+8 t=28 go.info.int+0
go.string."asdasdasd" t=8 dupok size=9
	0x0000 61 73 64 61 73 64 61 73 64                       asdasdasd
go.string."++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++=" t=8 dupok size=65
	0x0000 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b  ++++++++++++++++
	0x0010 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b  ++++++++++++++++
	0x0020 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b  ++++++++++++++++
	0x0030 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b 2b  ++++++++++++++++
	0x0040 3d                                               =
gclocals·f92853ccfca3dbb6e8a8fb53149795ad t=8 dupok size=26
	0x0000 09 00 00 00 0d 00 00 00 00 00 08 00 04 00 00 14  ................
	0x0010 17 14 11 14 10 14 00 14 a0 15                    ..........
gclocals·faa4501c2717f11e7fce158389399fba t=8 dupok size=17
	0x0000 09 00 00 00 06 00 00 00 3e 3e 3e 34 34 34 34 00  ........>>>4444.
	0x0010 00                                               .
go.info."".main.func2 t=45 size=256
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 32 00 00  ."".main.func2..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 04 72 32 00 05 9c 11 d8 7e 22 00 00 00 00 00 00  .r2.....~"......
	0x0030 00 00 04 72 00 05 9c 11 e0 7e 22 00 00 00 00 00  ...r.....~".....
	0x0040 00 00 00 04 6a 00 05 9c 11 e8 7e 22 00 00 00 00  ....j.....~"....
	0x0050 00 00 00 00 04 69 00 05 9c 11 f0 7e 22 00 00 00  .....i.....~"...
	0x0060 00 00 00 00 00 04 26 75 35 00 05 9c 11 a8 7f 22  ......&u5......"
	0x0070 00 00 00 00 00 00 00 00 04 75 73 32 00 04 9c 11  .........us2....
	0x0080 50 22 00 00 00 00 00 00 00 00 05 4e 00 01 9c 00  P".........N....
	0x0090 00 00 00 00 00 00 00 05 26 69 6e 64 65 78 00 04  ........&index..
	0x00a0 9c 11 08 22 00 00 00 00 00 00 00 00 05 26 6d 6d  ...".........&mm
	0x00b0 00 04 9c 11 10 22 00 00 00 00 00 00 00 00 05 26  .....".........&
	0x00c0 64 65 55 73 65 72 00 04 9c 11 18 22 00 00 00 00  deUser....."....
	0x00d0 00 00 00 00 05 6c 65 6e 4d 4d 00 04 9c 11 20 22  .....lenMM.... "
	0x00e0 00 00 00 00 00 00 00 00 05 26 64 65 55 73 65 72  .........&deUser
	0x00f0 32 00 04 9c 11 28 22 00 00 00 00 00 00 00 00 00  2....(".........
	rel 15+8 t=1 "".main.func2+0
	rel 23+8 t=1 "".main.func2+1315
	rel 42+8 t=28 go.info.int+0
	rel 59+8 t=28 go.info.int+0
	rel 76+8 t=28 go.info.int+0
	rel 93+8 t=28 go.info.int+0
	rel 112+8 t=28 go.info.*"".User+0
	rel 130+8 t=28 go.info."".User+0
	rel 143+8 t=28 go.info.int+0
	rel 164+8 t=28 go.info.*int+0
	rel 182+8 t=28 go.info.*"".MM+0
	rel 204+8 t=28 go.info.*"".User+0
	rel 224+8 t=28 go.info.func() int+0
	rel 247+8 t=28 go.info.**"".User+0
gclocals·78de76ab2bc1c1d963901a10e08a17b9 t=8 dupok size=12
	0x0000 04 00 00 00 02 00 00 00 00 00 02 01              ............
gclocals·23224035f3c252fdcd5453a88ab61c35 t=8 dupok size=12
	0x0000 04 00 00 00 04 00 00 00 0f 0e 0e 0e              ............
go.info."".main.func3 t=45 size=127
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 33 00 00  ."".main.func3..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 04 69 00 04 9c 11 48 22 00 00 00 00 00 00 00 00  .i....H"........
	0x0030 05 64 6f 6e 65 64 00 01 9c 00 00 00 00 00 00 00  .doned..........
	0x0040 00 05 26 69 6e 64 65 78 00 04 9c 11 08 22 00 00  ..&index....."..
	0x0050 00 00 00 00 00 00 05 26 6d 6d 00 04 9c 11 10 22  .......&mm....."
	0x0060 00 00 00 00 00 00 00 00 05 26 64 65 55 73 65 72  .........&deUser
	0x0070 00 04 9c 11 18 22 00 00 00 00 00 00 00 00 00     .....".........
	rel 15+8 t=1 "".main.func3+0
	rel 23+8 t=1 "".main.func3+407
	rel 40+8 t=28 go.info.int+0
	rel 57+8 t=28 go.info.chan bool+0
	rel 78+8 t=28 go.info.*int+0
	rel 96+8 t=28 go.info.*"".MM+0
	rel 118+8 t=28 go.info.*"".User+0
gclocals·d2de89f6cf5e99ce0097843c3071e829 t=8 dupok size=9
	0x0000 01 00 00 00 02 00 00 00 03                       .........
go.info."".main.func4 t=45 size=101
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 34 00 00  ."".main.func4..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 04 72 00 04 9c 11 58 22 00 00 00 00 00 00 00 00  .r....X"........
	0x0030 04 69 00 04 9c 11 60 22 00 00 00 00 00 00 00 00  .i....`"........
	0x0040 05 26 69 6e 64 65 78 00 01 9c 00 00 00 00 00 00  .&index.........
	0x0050 00 00 05 26 6d 6d 00 04 9c 11 08 22 00 00 00 00  ...&mm....."....
	0x0060 00 00 00 00 00                                   .....
	rel 15+8 t=1 "".main.func4+0
	rel 23+8 t=1 "".main.func4+205
	rel 40+8 t=28 go.info.int+0
	rel 56+8 t=28 go.info.int+0
	rel 74+8 t=28 go.info.*int+0
	rel 92+8 t=28 go.info.*"".MM+0
go.string."lkaskjas" t=8 dupok size=8
	0x0000 6c 6b 61 73 6b 6a 61 73                          lkaskjas
gclocals·55db315ce6b70e62560e945a72b60a76 t=8 dupok size=12
	0x0000 04 00 00 00 04 00 00 00 00 0f 0a 08              ............
gclocals·69a9291448fa273f79569cb593f615b2 t=8 dupok size=12
	0x0000 04 00 00 00 01 00 00 00 01 01 01 01              ............
go.info."".main.func5 t=45 size=80
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 35 00 00  ."".main.func5..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 04 69 00 04 9c 11 48 22 00 00 00 00 00 00 00 00  .i....H"........
	0x0030 04 26 75 00 04 9c 11 68 22 00 00 00 00 00 00 00  .&u....h".......
	0x0040 00 05 26 67 00 01 9c 00 00 00 00 00 00 00 00 00  ..&g............
	rel 15+8 t=1 "".main.func5+0
	rel 23+8 t=1 "".main.func5+400
	rel 40+8 t=28 go.info.int+0
	rel 57+8 t=28 go.info.*"".User+0
	rel 71+8 t=28 go.info.*int+0
go.string."nkajsdasjdasdjasd asjdbasdhjmabsdasbdasjhdbasdhasbdhjasm asmn asvhcgas cgsa" t=8 dupok size=75
	0x0000 6e 6b 61 6a 73 64 61 73 6a 64 61 73 64 6a 61 73  nkajsdasjdasdjas
	0x0010 64 20 61 73 6a 64 62 61 73 64 68 6a 6d 61 62 73  d asjdbasdhjmabs
	0x0020 64 61 73 62 64 61 73 6a 68 64 62 61 73 64 68 61  dasbdasjhdbasdha
	0x0030 73 62 64 68 6a 61 73 6d 20 61 73 6d 6e 20 61 73  sbdhjasm asmn as
	0x0040 76 68 63 67 61 73 20 63 67 73 61                 vhcgas cgsa
gclocals·2a5305abe05176240e61b8620e19a815 t=8 dupok size=9
	0x0000 01 00 00 00 01 00 00 00 00                       .........
go.info."".main.func6 t=45 size=62
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 36 00 00  ."".main.func6..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 04 69 00 04 9c 11 68 22 00 00 00 00 00 00 00 00  .i....h"........
	0x0030 05 4e 00 01 9c 00 00 00 00 00 00 00 00 00        .N............
	rel 15+8 t=1 "".main.func6+0
	rel 23+8 t=1 "".main.func6+182
	rel 40+8 t=28 go.info.int+0
	rel 53+8 t=28 go.info.int+0
go.string."\n\n" t=8 dupok size=2
	0x0000 0a 0a                                            ..
go.string."N: " t=8 dupok size=3
	0x0000 4e 3a 20                                         N: 
go.string."g: " t=8 dupok size=3
	0x0000 67 3a 20                                         g: 
go.string."Map Len: " t=8 dupok size=9
	0x0000 4d 61 70 20 4c 65 6e 3a 20                       Map Len: 
gclocals·a7899340dc00631f39a28dacb9997ba7 t=8 dupok size=26
	0x0000 06 00 00 00 16 00 00 00 00 00 00 40 03 00 10 00  ...........@....
	0x0010 3c 04 c0 03 01 00 00 01 3c 00                    <.......<.
gclocals·283d2e5c825c49c60b326a4265accb5a t=8 dupok size=14
	0x0000 06 00 00 00 03 00 00 00 07 07 07 07 07 07        ..............
go.info."".main.func7 t=45 size=84
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 37 00 00  ."".main.func7..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 05 26 7a 00 01 9c 00 00 00 00 00 00 00 00 05 26  .&z............&
	0x0030 67 00 04 9c 11 08 22 00 00 00 00 00 00 00 00 05  g.....".........
	0x0040 6c 65 6e 4d 4d 00 04 9c 11 10 22 00 00 00 00 00  lenMM.....".....
	0x0050 00 00 00 00                                      ....
	rel 15+8 t=1 "".main.func7+0
	rel 23+8 t=1 "".main.func7+923
	rel 38+8 t=28 go.info.*int+0
	rel 55+8 t=28 go.info.*int+0
	rel 75+8 t=28 go.info.func() int+0
go.info."".init t=45 size=27
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 00 00 00 00 00 00 00 00 01 00                 ...........
	rel 9+8 t=1 "".init+0
	rel 17+8 t=1 "".init+136
"".uuu t=31 size=8
"".str t=31 size=16
"".t1 t=32 size=8
"".statictmp_10 t=8 size=32
"".statictmp_11 t=8 size=32
"".statictmp_168 t=8 size=32
"".statictmp_191 t=8 size=32
"".initdone· t=32 size=1
"".main.func2·f t=8 dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".main.func2+0
"".main.func3·f t=8 dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".main.func3+0
"".main.func4·f t=8 dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".main.func4+0
"".main.func5·f t=8 dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".main.func5+0
"".main.func6·f t=8 dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".main.func6+0
"".main.func7·f t=8 dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 "".main.func7+0
runtime.gcbits.01 t=8 dupok size=1
	0x0000 01                                               .
type..namedata.***main.User. t=8 dupok size=15
	0x0000 00 00 0c 2a 2a 2a 6d 61 69 6e 2e 55 73 65 72     ...***main.User
type.***"".User t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 43 df de 79 00 08 08 36 00 00 00 00 00 00 00 00  C..y...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.***main.User.+0
	rel 48+8 t=1 type.**"".User+0
type..namedata.**main.User. t=8 dupok size=14
	0x0000 00 00 0b 2a 2a 6d 61 69 6e 2e 55 73 65 72        ...**main.User
type.**"".User t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 e7 1f 07 2e 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.**main.User.+0
	rel 44+4 t=6 type.***"".User+0
	rel 48+8 t=1 type.*"".User+0
type..namedata.*main.User. t=8 dupok size=13
	0x0000 01 00 0a 2a 6d 61 69 6e 2e 55 73 65 72           ...*main.User
type.*"".User t=8 size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 02 fa 53 6d 00 08 08 36 00 00 00 00 00 00 00 00  ..Sm...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.User.+0
	rel 44+4 t=6 type.**"".User+0
	rel 48+8 t=1 type."".User+0
gclocals·69c1753bd5f81501d95132d08af04464 t=8 dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·e6397a44f8e1b6e77d0f200b4fba5269 t=8 dupok size=10
	0x0000 02 00 00 00 03 00 00 00 01 00                    ..........
go.info.type..hash."".User t=45 dupok size=85
	0x0000 02 74 79 70 65 2e 2e 68 61 73 68 2e 22 22 2e 55  .type..hash."".U
	0x0010 73 65 72 00 00 00 00 00 00 00 00 00 00 00 00 00  ser.............
	0x0020 00 00 00 00 01 05 70 00 01 9c 00 00 00 00 00 00  ......p.........
	0x0030 00 00 05 68 00 04 9c 11 08 22 00 00 00 00 00 00  ...h....."......
	0x0040 00 00 05 7e 72 32 00 04 9c 11 10 22 00 00 00 00  ...~r2....."....
	0x0050 00 00 00 00 00                                   .....
	rel 20+8 t=1 type..hash."".User+0
	rel 28+8 t=1 type..hash."".User+168
	rel 42+8 t=28 go.info.*"".User+0
	rel 58+8 t=28 go.info.uintptr+0
	rel 76+8 t=28 go.info.uintptr+0
gclocals·8f9cec06d1ae35cc9900c511c5e4bdab t=8 dupok size=9
	0x0000 01 00 00 00 03 00 00 00 03                       .........
go.info.type..eq."".User t=45 dupok size=83
	0x0000 02 74 79 70 65 2e 2e 65 71 2e 22 22 2e 55 73 65  .type..eq."".Use
	0x0010 72 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  r...............
	0x0020 00 00 01 05 70 00 01 9c 00 00 00 00 00 00 00 00  ....p...........
	0x0030 05 71 00 04 9c 11 08 22 00 00 00 00 00 00 00 00  .q....."........
	0x0040 05 7e 72 32 00 04 9c 11 10 22 00 00 00 00 00 00  .~r2....."......
	0x0050 00 00 00                                         ...
	rel 18+8 t=1 type..eq."".User+0
	rel 26+8 t=1 type..eq."".User+175
	rel 40+8 t=28 go.info.*"".User+0
	rel 56+8 t=28 go.info.*"".User+0
	rel 74+8 t=28 go.info.bool+0
type..hashfunc."".User t=8 dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..hash."".User+0
type..eqfunc."".User t=8 dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq."".User+0
type..alg."".User t=8 dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc."".User+0
	rel 8+8 t=1 type..eqfunc."".User+0
runtime.gcbits.0a t=8 dupok size=1
	0x0000 0a                                               .
type..namedata.id. t=8 dupok size=5
	0x0000 00 00 02 69 64                                   ...id
type..namedata.name. t=8 dupok size=7
	0x0000 00 00 04 6e 61 6d 65                             ...name
type..namedata.u. t=8 dupok size=4
	0x0000 00 00 01 75                                      ...u
type."".User t=8 size=168
	0x0000 20 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00   ....... .......
	0x0010 64 15 5a 77 07 08 08 19 00 00 00 00 00 00 00 00  d.Zw............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 03 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 58 00 00 00 00 00 00 00  ........X.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 18 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..alg."".User+0
	rel 32+8 t=1 runtime.gcbits.0a+0
	rel 40+4 t=5 type..namedata.*main.User.+0
	rel 44+4 t=5 type.*"".User+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".User+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.id.+0
	rel 104+8 t=1 type.int+0
	rel 120+8 t=1 type..namedata.name.+0
	rel 128+8 t=1 type.string+0
	rel 144+8 t=1 type..namedata.u.+0
	rel 152+8 t=1 type.*"".User+0
type..namedata.*[]uint8. t=8 dupok size=11
	0x0000 00 00 08 2a 5b 5d 75 69 6e 74 38                 ...*[]uint8
type.*[]uint8 t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a5 8e d0 69 00 08 08 36 00 00 00 00 00 00 00 00  ...i...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8.+0
	rel 48+8 t=1 type.[]uint8+0
type.[]uint8 t=8 dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 df 7e 2e 38 02 08 08 17 00 00 00 00 00 00 00 00  .~.8............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]uint8.+0
	rel 44+4 t=6 type.*[]uint8+0
	rel 48+8 t=1 type.uint8+0
type..namedata.*[8]uint8. t=8 dupok size=12
	0x0000 00 00 09 2a 5b 38 5d 75 69 6e 74 38              ...*[8]uint8
type.*[8]uint8 t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 a9 89 a5 7a 00 08 08 36 00 00 00 00 00 00 00 00  ...z...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[8]uint8.+0
	rel 48+8 t=1 type.[8]uint8+0
runtime.gcbits. t=8 dupok size=0
type.[8]uint8 t=8 dupok size=72
	0x0000 08 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 3e f9 30 b4 02 01 01 91 00 00 00 00 00 00 00 00  >.0.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 08 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[8]uint8.+0
	rel 44+4 t=6 type.*[8]uint8+0
	rel 48+8 t=1 type.uint8+0
	rel 56+8 t=1 type.[]uint8+0
type..namedata.*[]int. t=8 dupok size=9
	0x0000 00 00 06 2a 5b 5d 69 6e 74                       ...*[]int
type.*[]int t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1b 31 52 88 00 08 08 36 00 00 00 00 00 00 00 00  .1R....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int.+0
	rel 48+8 t=1 type.[]int+0
type.[]int t=8 dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8e 66 f9 1b 02 08 08 17 00 00 00 00 00 00 00 00  .f..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]int.+0
	rel 44+4 t=6 type.*[]int+0
	rel 48+8 t=1 type.int+0
type..namedata.*[8]int. t=8 dupok size=10
	0x0000 00 00 07 2a 5b 38 5d 69 6e 74                    ...*[8]int
type.*[8]int t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 3f a8 3b 00 08 08 36 00 00 00 00 00 00 00 00  .?.;...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[8]int.+0
	rel 48+8 t=1 type.noalg.[8]int+0
type.noalg.[8]int t=8 dupok size=72
	0x0000 40 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  @...............
	0x0010 96 99 d5 05 02 08 08 91 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 08 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*[8]int.+0
	rel 44+4 t=6 type.*[8]int+0
	rel 48+8 t=1 type.int+0
	rel 56+8 t=1 type.[]int+0
type..namedata.*[]main.User. t=8 dupok size=15
	0x0000 00 00 0c 2a 5b 5d 6d 61 69 6e 2e 55 73 65 72     ...*[]main.User
type.*[]"".User t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 8a bd a3 4d 00 08 08 36 00 00 00 00 00 00 00 00  ...M...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]main.User.+0
	rel 48+8 t=1 type.[]"".User+0
type.[]"".User t=8 dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 42 93 15 3a 02 08 08 17 00 00 00 00 00 00 00 00  B..:............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]main.User.+0
	rel 44+4 t=6 type.*[]"".User+0
	rel 48+8 t=1 type."".User+0
type..namedata.*[8]main.User. t=8 dupok size=16
	0x0000 00 00 0d 2a 5b 38 5d 6d 61 69 6e 2e 55 73 65 72  ...*[8]main.User
type.*[8]"".User t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 5b 1a f9 ff 00 08 08 36 00 00 00 00 00 00 00 00  [......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[8]main.User.+0
	rel 48+8 t=1 type.noalg.[8]"".User+0
runtime.gcbits.aaaaaaaa t=8 dupok size=4
	0x0000 aa aa aa aa                                      ....
type.noalg.[8]"".User t=8 dupok size=72
	0x0000 00 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00  ................
	0x0010 33 d0 88 9e 02 08 08 11 00 00 00 00 00 00 00 00  3...............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 08 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.aaaaaaaa+0
	rel 40+4 t=5 type..namedata.*[8]main.User.+0
	rel 44+4 t=6 type.*[8]"".User+0
	rel 48+8 t=1 type."".User+0
	rel 56+8 t=1 type.[]"".User+0
type..namedata.*map.bucket[int]main.User. t=8 dupok size=28
	0x0000 00 00 19 2a 6d 61 70 2e 62 75 63 6b 65 74 5b 69  ...*map.bucket[i
	0x0010 6e 74 5d 6d 61 69 6e 2e 55 73 65 72              nt]main.User
type.*map.bucket[int]"".User t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 08 c7 a2 7e 00 08 08 36 00 00 00 00 00 00 00 00  ...~...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map.bucket[int]main.User.+0
	rel 48+8 t=1 type.noalg.map.bucket[int]"".User+0
runtime.gcbits.005455555503 t=8 dupok size=6
	0x0000 00 54 55 55 55 03                                .TUUU.
type..importpath.. t=8 dupok size=3
	0x0000 00 00 00                                         ...
type..namedata.topbits. t=8 dupok size=10
	0x0000 00 00 07 74 6f 70 62 69 74 73                    ...topbits
type..namedata.keys. t=8 dupok size=7
	0x0000 00 00 04 6b 65 79 73                             ...keys
type..namedata.values. t=8 dupok size=9
	0x0000 00 00 06 76 61 6c 75 65 73                       ...values
type..namedata.overflow. t=8 dupok size=11
	0x0000 00 00 08 6f 76 65 72 66 6c 6f 77                 ...overflow
type.noalg.map.bucket[int]"".User t=8 dupok size=176
	0x0000 50 01 00 00 00 00 00 00 50 01 00 00 00 00 00 00  P.......P.......
	0x0010 52 fa e0 e7 02 08 08 19 00 00 00 00 00 00 00 00  R...............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 04 00 00 00 00 00 00 00 04 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 48 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  H...............
	0x00a0 00 00 00 00 00 00 00 00 48 01 00 00 00 00 00 00  ........H.......
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.005455555503+0
	rel 40+4 t=5 type..namedata.*map.bucket[int]main.User.+0
	rel 44+4 t=6 type.*map.bucket[int]"".User+0
	rel 48+8 t=1 type..importpath..+0
	rel 56+8 t=1 type.noalg.map.bucket[int]"".User+80
	rel 80+8 t=1 type..namedata.topbits.+0
	rel 88+8 t=1 type.[8]uint8+0
	rel 104+8 t=1 type..namedata.keys.+0
	rel 112+8 t=1 type.noalg.[8]int+0
	rel 128+8 t=1 type..namedata.values.+0
	rel 136+8 t=1 type.noalg.[8]"".User+0
	rel 152+8 t=1 type..namedata.overflow.+0
	rel 160+8 t=1 type.*map.bucket[int]"".User+0
type..namedata.*map.hdr[int]main.User. t=8 dupok size=25
	0x0000 00 00 16 2a 6d 61 70 2e 68 64 72 5b 69 6e 74 5d  ...*map.hdr[int]
	0x0010 6d 61 69 6e 2e 55 73 65 72                       main.User
type.*map.hdr[int]"".User t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 3d 82 69 8b 00 08 08 36 00 00 00 00 00 00 00 00  =.i....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map.hdr[int]main.User.+0
	rel 48+8 t=1 type.noalg.map.hdr[int]"".User+0
runtime.gcbits.2c t=8 dupok size=1
	0x0000 2c                                               ,
type..namedata.count. t=8 dupok size=8
	0x0000 00 00 05 63 6f 75 6e 74                          ...count
type..namedata.flags. t=8 dupok size=8
	0x0000 00 00 05 66 6c 61 67 73                          ...flags
type..namedata.B. t=8 dupok size=4
	0x0000 01 00 01 42                                      ...B
type..namedata.noverflow. t=8 dupok size=12
	0x0000 00 00 09 6e 6f 76 65 72 66 6c 6f 77              ...noverflow
type..namedata.hash0. t=8 dupok size=8
	0x0000 00 00 05 68 61 73 68 30                          ...hash0
type..namedata.buckets. t=8 dupok size=10
	0x0000 00 00 07 62 75 63 6b 65 74 73                    ...buckets
type..namedata.oldbuckets. t=8 dupok size=13
	0x0000 00 00 0a 6f 6c 64 62 75 63 6b 65 74 73           ...oldbuckets
type..namedata.nevacuate. t=8 dupok size=12
	0x0000 00 00 09 6e 65 76 61 63 75 61 74 65              ...nevacuate
type.noalg.map.hdr[int]"".User t=8 dupok size=296
	0x0000 30 00 00 00 00 00 00 00 30 00 00 00 00 00 00 00  0.......0.......
	0x0010 ae 35 2e f8 02 08 08 19 00 00 00 00 00 00 00 00  .5..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 09 00 00 00 00 00 00 00 09 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 09 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00a0 00 00 00 00 00 00 00 00 0a 00 00 00 00 00 00 00  ................
	0x00b0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00c0 0c 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00d0 00 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x00e0 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x00f0 18 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0100 00 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00  ........ .......
	0x0110 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0120 28 00 00 00 00 00 00 00                          (.......
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.2c+0
	rel 40+4 t=5 type..namedata.*map.hdr[int]main.User.+0
	rel 44+4 t=6 type.*map.hdr[int]"".User+0
	rel 48+8 t=1 type..importpath..+0
	rel 56+8 t=1 type.noalg.map.hdr[int]"".User+80
	rel 80+8 t=1 type..namedata.count.+0
	rel 88+8 t=1 type.int+0
	rel 104+8 t=1 type..namedata.flags.+0
	rel 112+8 t=1 type.uint8+0
	rel 128+8 t=1 type..namedata.B.+0
	rel 136+8 t=1 type.uint8+0
	rel 152+8 t=1 type..namedata.noverflow.+0
	rel 160+8 t=1 type.uint16+0
	rel 176+8 t=1 type..namedata.hash0.+0
	rel 184+8 t=1 type.uint32+0
	rel 200+8 t=1 type..namedata.buckets.+0
	rel 208+8 t=1 type.*map.bucket[int]"".User+0
	rel 224+8 t=1 type..namedata.oldbuckets.+0
	rel 232+8 t=1 type.*map.bucket[int]"".User+0
	rel 248+8 t=1 type..namedata.nevacuate.+0
	rel 256+8 t=1 type.uintptr+0
	rel 272+8 t=1 type..namedata.overflow.+0
	rel 280+8 t=1 type.unsafe.Pointer+0
type..namedata.**map[int]main.User. t=8 dupok size=22
	0x0000 00 00 13 2a 2a 6d 61 70 5b 69 6e 74 5d 6d 61 69  ...**map[int]mai
	0x0010 6e 2e 55 73 65 72                                n.User
type.**map[int]"".User t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 02 6b ec d3 00 08 08 36 00 00 00 00 00 00 00 00  .k.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.**map[int]main.User.+0
	rel 48+8 t=1 type.*map[int]"".User+0
type..namedata.*map[int]main.User. t=8 dupok size=21
	0x0000 00 00 12 2a 6d 61 70 5b 69 6e 74 5d 6d 61 69 6e  ...*map[int]main
	0x0010 2e 55 73 65 72                                   .User
type.*map[int]"".User t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 c2 f2 ec bb 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map[int]main.User.+0
	rel 44+4 t=6 type.**map[int]"".User+0
	rel 48+8 t=1 type.map[int]"".User+0
type.map[int]"".User t=8 dupok size=88
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 d2 b1 b0 53 02 08 08 35 00 00 00 00 00 00 00 00  ...S...5........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 08 00 20 00 50 01 01 00                          .. .P...
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*map[int]main.User.+0
	rel 44+4 t=6 type.*map[int]"".User+0
	rel 48+8 t=1 type.int+0
	rel 56+8 t=1 type."".User+0
	rel 64+8 t=1 type.noalg.map.bucket[int]"".User+0
	rel 72+8 t=1 type.noalg.map.hdr[int]"".User+0
type..namedata.***main.MM. t=8 dupok size=13
	0x0000 00 00 0a 2a 2a 2a 6d 61 69 6e 2e 4d 4d           ...***main.MM
type.***"".MM t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 6a 39 7c 4b 00 08 08 36 00 00 00 00 00 00 00 00  j9|K...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.***main.MM.+0
	rel 48+8 t=1 type.**"".MM+0
type..namedata.**main.MM. t=8 dupok size=12
	0x0000 00 00 09 2a 2a 6d 61 69 6e 2e 4d 4d              ...**main.MM
type.**"".MM t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4d b0 a7 77 00 08 08 36 00 00 00 00 00 00 00 00  M..w...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.**main.MM.+0
	rel 44+4 t=6 type.***"".MM+0
	rel 48+8 t=1 type.*"".MM+0
gclocals·a36216b97439c93dafebe03e7f0808b5 t=8 dupok size=9
	0x0000 01 00 00 00 01 00 00 00 01                       .........
go.info."".(*MM).Lock t=45 dupok size=50
	0x0000 02 22 22 2e 28 2a 4d 4d 29 2e 4c 6f 63 6b 00 00  ."".(*MM).Lock..
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 01  ................
	0x0020 05 2e 74 68 69 73 00 01 9c 00 00 00 00 00 00 00  ..this..........
	0x0030 00 00                                            ..
	rel 15+8 t=1 "".(*MM).Lock+0
	rel 23+8 t=1 "".(*MM).Lock+17
	rel 41+8 t=28 go.info.*"".MM+0
go.info."".(*MM).RLock t=45 dupok size=51
	0x0000 02 22 22 2e 28 2a 4d 4d 29 2e 52 4c 6f 63 6b 00  ."".(*MM).RLock.
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 01 05 2e 74 68 69 73 00 01 9c 00 00 00 00 00 00  ...this.........
	0x0030 00 00 00                                         ...
	rel 16+8 t=1 "".(*MM).RLock+0
	rel 24+8 t=1 "".(*MM).RLock+17
	rel 42+8 t=28 go.info.*"".MM+0
gclocals·d4dc2f11db048877dbc0f60a22b4adb3 t=8 dupok size=9
	0x0000 01 00 00 00 03 00 00 00 01                       .........
go.info."".(*MM).RLocker t=45 dupok size=71
	0x0000 02 22 22 2e 28 2a 4d 4d 29 2e 52 4c 6f 63 6b 65  ."".(*MM).RLocke
	0x0010 72 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  r...............
	0x0020 00 00 01 05 2e 74 68 69 73 00 01 9c 00 00 00 00  .....this.......
	0x0030 00 00 00 00 05 7e 72 31 00 04 9c 11 08 22 00 00  .....~r1....."..
	0x0040 00 00 00 00 00 00 00                             .......
	rel 18+8 t=1 "".(*MM).RLocker+0
	rel 26+8 t=1 "".(*MM).RLocker+35
	rel 44+8 t=28 go.info.*"".MM+0
	rel 62+8 t=28 go.info.sync.Locker+0
go.info."".(*MM).RUnlock t=45 dupok size=53
	0x0000 02 22 22 2e 28 2a 4d 4d 29 2e 52 55 6e 6c 6f 63  ."".(*MM).RUnloc
	0x0010 6b 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  k...............
	0x0020 00 00 01 05 2e 74 68 69 73 00 01 9c 00 00 00 00  .....this.......
	0x0030 00 00 00 00 00                                   .....
	rel 18+8 t=1 "".(*MM).RUnlock+0
	rel 26+8 t=1 "".(*MM).RUnlock+17
	rel 44+8 t=28 go.info.*"".MM+0
go.info."".(*MM).Unlock t=45 dupok size=52
	0x0000 02 22 22 2e 28 2a 4d 4d 29 2e 55 6e 6c 6f 63 6b  ."".(*MM).Unlock
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 01 05 2e 74 68 69 73 00 01 9c 00 00 00 00 00  ....this........
	0x0030 00 00 00 00                                      ....
	rel 17+8 t=1 "".(*MM).Unlock+0
	rel 25+8 t=1 "".(*MM).Unlock+17
	rel 43+8 t=28 go.info.*"".MM+0
type..namedata.*main.MM. t=8 dupok size=11
	0x0000 01 00 08 2a 6d 61 69 6e 2e 4d 4d                 ...*main.MM
type..namedata.*func(*main.MM). t=8 dupok size=18
	0x0000 00 00 0f 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 4d  ...*func(*main.M
	0x0010 4d 29                                            M)
type.*func(*"".MM) t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 e5 24 fd 77 00 08 08 36 00 00 00 00 00 00 00 00  .$.w...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.MM).+0
	rel 48+8 t=1 type.func(*"".MM)+0
type.func(*"".MM) t=8 dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1c 1a ef 91 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.MM).+0
	rel 44+4 t=6 type.*func(*"".MM)+0
	rel 56+8 t=1 type.*"".MM+0
type..namedata.*func(*main.MM) sync.Locker. t=8 dupok size=30
	0x0000 00 00 1b 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 4d  ...*func(*main.M
	0x0010 4d 29 20 73 79 6e 63 2e 4c 6f 63 6b 65 72        M) sync.Locker
type.*func(*"".MM) sync.Locker t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 b7 e7 27 1e 00 08 08 36 00 00 00 00 00 00 00 00  ..'....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.MM) sync.Locker.+0
	rel 48+8 t=1 type.func(*"".MM) sync.Locker+0
type.func(*"".MM) sync.Locker t=8 dupok size=72
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 03 02 b4 21 02 08 08 33 00 00 00 00 00 00 00 00  ...!...3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 01 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.MM) sync.Locker.+0
	rel 44+4 t=6 type.*func(*"".MM) sync.Locker+0
	rel 56+8 t=1 type.*"".MM+0
	rel 64+8 t=1 type.sync.Locker+0
type..namedata.Lock. t=8 dupok size=7
	0x0000 01 00 04 4c 6f 63 6b                             ...Lock
type..namedata.*func(). t=8 dupok size=10
	0x0000 00 00 07 2a 66 75 6e 63 28 29                    ...*func()
type.*func() t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 9b 90 75 1b 00 08 08 36 00 00 00 00 00 00 00 00  ..u....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func().+0
	rel 48+8 t=1 type.func()+0
type.func() t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f6 bc 82 f6 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00                                      ....
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func().+0
	rel 44+4 t=6 type.*func()+0
type..namedata.RLock. t=8 dupok size=8
	0x0000 01 00 05 52 4c 6f 63 6b                          ...RLock
type..namedata.RLocker. t=8 dupok size=10
	0x0000 01 00 07 52 4c 6f 63 6b 65 72                    ...RLocker
type..namedata.*func() sync.Locker. t=8 dupok size=22
	0x0000 00 00 13 2a 66 75 6e 63 28 29 20 73 79 6e 63 2e  ...*func() sync.
	0x0010 4c 6f 63 6b 65 72                                Locker
type.*func() sync.Locker t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 57 74 29 ee 00 08 08 36 00 00 00 00 00 00 00 00  Wt)....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() sync.Locker.+0
	rel 48+8 t=1 type.func() sync.Locker+0
type.func() sync.Locker t=8 dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 c8 f6 04 19 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() sync.Locker.+0
	rel 44+4 t=6 type.*func() sync.Locker+0
	rel 56+8 t=1 type.sync.Locker+0
type..namedata.RUnlock. t=8 dupok size=10
	0x0000 01 00 07 52 55 6e 6c 6f 63 6b                    ...RUnlock
type..namedata.Unlock. t=8 dupok size=9
	0x0000 01 00 06 55 6e 6c 6f 63 6b                       ...Unlock
type.*"".MM t=8 size=152
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 1a 8c ea 7b 01 08 08 36 00 00 00 00 00 00 00 00  ...{...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 05 00 00 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0090 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.MM.+0
	rel 44+4 t=6 type.**"".MM+0
	rel 48+8 t=1 type."".MM+0
	rel 56+4 t=5 type..importpath."".+0
	rel 72+4 t=5 type..namedata.Lock.+0
	rel 76+4 t=24 type.func()+0
	rel 80+4 t=24 "".(*MM).Lock+0
	rel 84+4 t=24 "".(*MM).Lock+0
	rel 88+4 t=5 type..namedata.RLock.+0
	rel 92+4 t=24 type.func()+0
	rel 96+4 t=24 "".(*MM).RLock+0
	rel 100+4 t=24 "".(*MM).RLock+0
	rel 104+4 t=5 type..namedata.RLocker.+0
	rel 108+4 t=24 type.func() sync.Locker+0
	rel 112+4 t=24 "".(*MM).RLocker+0
	rel 116+4 t=24 "".(*MM).RLocker+0
	rel 120+4 t=5 type..namedata.RUnlock.+0
	rel 124+4 t=24 type.func()+0
	rel 128+4 t=24 "".(*MM).RUnlock+0
	rel 132+4 t=24 "".(*MM).RUnlock+0
	rel 136+4 t=5 type..namedata.Unlock.+0
	rel 140+4 t=24 type.func()+0
	rel 144+4 t=24 "".(*MM).Unlock+0
	rel 148+4 t=24 "".(*MM).Unlock+0
runtime.gcbits.08 t=8 dupok size=1
	0x0000 08                                               .
type..namedata.-noname-exported. t=8 dupok size=3
	0x0000 01 00 00                                         ...
type..namedata.m. t=8 dupok size=4
	0x0000 00 00 01 6d                                      ...m
type."".MM t=8 size=144
	0x0000 20 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00   ....... .......
	0x0010 36 87 28 60 07 08 08 19 00 00 00 00 00 00 00 00  6.(`............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 40 00 00 00 00 00 00 00  ........@.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00 18 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.08+0
	rel 40+4 t=5 type..namedata.*main.MM.+0
	rel 44+4 t=5 type.*"".MM+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".MM+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.-noname-exported.+0
	rel 104+8 t=1 type.sync.RWMutex+0
	rel 120+8 t=1 type..namedata.m.+0
	rel 128+8 t=1 type.map[int]"".User+0
type..namedata.*struct { F uintptr; mm *main.MM }. t=8 dupok size=37
	0x0000 00 00 22 2a 73 74 72 75 63 74 20 7b 20 46 20 75  .."*struct { F u
	0x0010 69 6e 74 70 74 72 3b 20 6d 6d 20 2a 6d 61 69 6e  intptr; mm *main
	0x0020 2e 4d 4d 20 7d                                   .MM }
type.*struct { F uintptr; "".mm *"".MM } t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 57 92 6a c6 00 08 08 36 00 00 00 00 00 00 00 00  W.j....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; mm *main.MM }.+0
	rel 48+8 t=1 type.struct { F uintptr; "".mm *"".MM }+0
runtime.gcbits.02 t=8 dupok size=1
	0x0000 02                                               .
type..namedata..F. t=8 dupok size=5
	0x0000 00 00 02 2e 46                                   ....F
type..namedata.mm. t=8 dupok size=5
	0x0000 00 00 02 6d 6d                                   ...mm
type.struct { F uintptr; "".mm *"".MM } t=8 dupok size=128
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 16 b1 57 d6 02 08 08 19 00 00 00 00 00 00 00 00  ..W.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+96
	rel 32+8 t=1 runtime.gcbits.02+0
	rel 40+4 t=5 type..namedata.*struct { F uintptr; mm *main.MM }.+0
	rel 44+4 t=6 type.*struct { F uintptr; "".mm *"".MM }+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type.struct { F uintptr; "".mm *"".MM }+80
	rel 80+8 t=1 type..namedata..F.+0
	rel 88+8 t=1 type.uintptr+0
	rel 104+8 t=1 type..namedata.mm.+0
	rel 112+8 t=1 type.*"".MM+0
type..namedata.*chan bool. t=8 dupok size=13
	0x0000 00 00 0a 2a 63 68 61 6e 20 62 6f 6f 6c           ...*chan bool
type.*chan bool t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 13 c4 b2 cb 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*chan bool.+0
	rel 48+8 t=1 type.chan bool+0
type.chan bool t=8 dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 b8 48 df dd 02 08 08 32 00 00 00 00 00 00 00 00  .H.....2........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 03 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*chan bool.+0
	rel 44+4 t=6 type.*chan bool+0
	rel 48+8 t=1 type.bool+0
type..namedata.*interface {}. t=8 dupok size=16
	0x0000 00 00 0d 2a 69 6e 74 65 72 66 61 63 65 20 7b 7d  ...*interface {}
type.*interface {} t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 4f 0f 96 9d 00 08 08 36 00 00 00 00 00 00 00 00  O......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*interface {}.+0
	rel 48+8 t=1 type.interface {}+0
runtime.gcbits.03 t=8 dupok size=1
	0x0000 03                                               .
type.interface {} t=8 dupok size=80
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 e7 57 a0 18 02 08 08 14 00 00 00 00 00 00 00 00  .W..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*interface {}.+0
	rel 44+4 t=6 type.*interface {}+0
	rel 56+8 t=1 type.interface {}+80
type..namedata.*[]interface {}. t=8 dupok size=18
	0x0000 00 00 0f 2a 5b 5d 69 6e 74 65 72 66 61 63 65 20  ...*[]interface 
	0x0010 7b 7d                                            {}
type.*[]interface {} t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 f3 04 9a e7 00 08 08 36 00 00 00 00 00 00 00 00  .......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}.+0
	rel 48+8 t=1 type.[]interface {}+0
type.[]interface {} t=8 dupok size=56
	0x0000 18 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 70 93 ea 2f 02 08 08 17 00 00 00 00 00 00 00 00  p../............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[]interface {}.+0
	rel 44+4 t=6 type.*[]interface {}+0
	rel 48+8 t=1 type.interface {}+0
go.info.type..hash.[2]interface {} t=45 dupok size=109
	0x0000 02 74 79 70 65 2e 2e 68 61 73 68 2e 5b 32 5d 69  .type..hash.[2]i
	0x0010 6e 74 65 72 66 61 63 65 20 7b 7d 00 00 00 00 00  nterface {}.....
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 01 04 69 00  ..............i.
	0x0030 04 9c 11 68 22 00 00 00 00 00 00 00 00 05 70 00  ...h".........p.
	0x0040 01 9c 00 00 00 00 00 00 00 00 05 68 00 04 9c 11  ...........h....
	0x0050 08 22 00 00 00 00 00 00 00 00 05 7e 72 32 00 04  .".........~r2..
	0x0060 9c 11 10 22 00 00 00 00 00 00 00 00 00           ...".........
	rel 28+8 t=1 type..hash.[2]interface {}+0
	rel 36+8 t=1 type..hash.[2]interface {}+126
	rel 53+8 t=28 go.info.int+0
	rel 66+8 t=28 go.info.*[2]interface {}+0
	rel 82+8 t=28 go.info.uintptr+0
	rel 100+8 t=28 go.info.uintptr+0
go.info.type..eq.[2]interface {} t=45 dupok size=107
	0x0000 02 74 79 70 65 2e 2e 65 71 2e 5b 32 5d 69 6e 74  .type..eq.[2]int
	0x0010 65 72 66 61 63 65 20 7b 7d 00 00 00 00 00 00 00  erface {}.......
	0x0020 00 00 00 00 00 00 00 00 00 00 01 04 69 00 04 9c  ............i...
	0x0030 11 68 22 00 00 00 00 00 00 00 00 05 70 00 01 9c  .h".........p...
	0x0040 00 00 00 00 00 00 00 00 05 71 00 04 9c 11 08 22  .........q....."
	0x0050 00 00 00 00 00 00 00 00 05 7e 72 32 00 04 9c 11  .........~r2....
	0x0060 10 22 00 00 00 00 00 00 00 00 00                 .".........
	rel 26+8 t=1 type..eq.[2]interface {}+0
	rel 34+8 t=1 type..eq.[2]interface {}+183
	rel 51+8 t=28 go.info.int+0
	rel 64+8 t=28 go.info.*[2]interface {}+0
	rel 80+8 t=28 go.info.*[2]interface {}+0
	rel 98+8 t=28 go.info.bool+0
type..hashfunc.[2]interface {} t=8 dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..hash.[2]interface {}+0
type..eqfunc.[2]interface {} t=8 dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 type..eq.[2]interface {}+0
type..alg.[2]interface {} t=8 dupok size=16
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 0+8 t=1 type..hashfunc.[2]interface {}+0
	rel 8+8 t=1 type..eqfunc.[2]interface {}+0
type..namedata.**[2]interface {}. t=8 dupok size=20
	0x0000 00 00 11 2a 2a 5b 32 5d 69 6e 74 65 72 66 61 63  ...**[2]interfac
	0x0010 65 20 7b 7d                                      e {}
type.**[2]interface {} t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 17 95 86 56 00 08 08 36 00 00 00 00 00 00 00 00  ...V...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.**[2]interface {}.+0
	rel 48+8 t=1 type.*[2]interface {}+0
type..namedata.*[2]interface {}. t=8 dupok size=19
	0x0000 00 00 10 2a 5b 32 5d 69 6e 74 65 72 66 61 63 65  ...*[2]interface
	0x0010 20 7b 7d                                          {}
type.*[2]interface {} t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 be 73 2d 71 00 08 08 36 00 00 00 00 00 00 00 00  .s-q...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[2]interface {}.+0
	rel 44+4 t=6 type.**[2]interface {}+0
	rel 48+8 t=1 type.[2]interface {}+0
runtime.gcbits.0f t=8 dupok size=1
	0x0000 0f                                               .
type.[2]interface {} t=8 dupok size=72
	0x0000 20 00 00 00 00 00 00 00 20 00 00 00 00 00 00 00   ....... .......
	0x0010 2c 59 a4 f1 02 08 08 11 00 00 00 00 00 00 00 00  ,Y..............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 02 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 type..alg.[2]interface {}+0
	rel 32+8 t=1 runtime.gcbits.0f+0
	rel 40+4 t=5 type..namedata.*[2]interface {}.+0
	rel 44+4 t=6 type.*[2]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..namedata.*func() int. t=8 dupok size=14
	0x0000 00 00 0b 2a 66 75 6e 63 28 29 20 69 6e 74        ...*func() int
type.*func() int t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 51 b2 ed c6 00 08 08 36 00 00 00 00 00 00 00 00  Q......6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() int.+0
	rel 48+8 t=1 type.func() int+0
type.func() int t=8 dupok size=64
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 e5 86 39 e0 02 08 08 33 00 00 00 00 00 00 00 00  ..9....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 24+8 t=1 runtime.algarray+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func() int.+0
	rel 44+4 t=6 type.*func() int+0
	rel 56+8 t=1 type.int+0
type..namedata.*[1]interface {}. t=8 dupok size=19
	0x0000 00 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65  ...*[1]interface
	0x0010 20 7b 7d                                          {}
type.*[1]interface {} t=8 dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 bf 03 a8 35 00 08 08 36 00 00 00 00 00 00 00 00  ...5...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+80
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*[1]interface {}.+0
	rel 48+8 t=1 type.[1]interface {}+0
type.[1]interface {} t=8 dupok size=72
	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
	0x0010 50 91 5b fa 02 08 08 11 00 00 00 00 00 00 00 00  P.[.............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.algarray+144
	rel 32+8 t=1 runtime.gcbits.03+0
	rel 40+4 t=5 type..namedata.*[1]interface {}.+0
	rel 44+4 t=6 type.*[1]interface {}+0
	rel 48+8 t=1 type.interface {}+0
	rel 56+8 t=1 type.[]interface {}+0
type..importpath.fmt. t=8 dupok size=6
	0x0000 00 00 03 66 6d 74                                ...fmt
type..importpath.math/rand. t=8 dupok size=12
	0x0000 00 00 09 6d 61 74 68 2f 72 61 6e 64              ...math/rand
type..importpath.runtime. t=8 dupok size=10
	0x0000 00 00 07 72 75 6e 74 69 6d 65                    ...runtime
type..importpath.runtime/debug. t=8 dupok size=16
	0x0000 00 00 0d 72 75 6e 74 69 6d 65 2f 64 65 62 75 67  ...runtime/debug
type..importpath.strconv. t=8 dupok size=10
	0x0000 00 00 07 73 74 72 63 6f 6e 76                    ...strconv
type..importpath.sync. t=8 dupok size=7
	0x0000 00 00 04 73 79 6e 63                             ...sync
type..importpath.time. t=8 dupok size=7
	0x0000 00 00 04 74 69 6d 65                             ...time
