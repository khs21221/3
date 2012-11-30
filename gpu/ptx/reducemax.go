package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["reducemax"] = REDUCEMAX }

const REDUCEMAX = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00000ada_00000000-9_reducemax.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/reducemax.cu"
	.file	3 "/usr/local/cuda-5.0/nvvm/ci_include.h"
	.file	4 "/usr/local/cuda/bin/../include/sm_11_atomic_functions.h"
// __cuda_local_var_33851_32_non_const_sdata has been demoted

.visible .entry reducemax(
	.param .u64 reducemax_param_0,
	.param .u64 reducemax_param_1,
	.param .u32 reducemax_param_2
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<40>;
	.reg .f32 	%f<31>;
	.reg .s64 	%rd<13>;
	// demoted variable
	.shared .align 4 .b8 __cuda_local_var_33851_32_non_const_sdata[2048];

	ld.param.u64 	%rd4, [reducemax_param_0];
	ld.param.u64 	%rd5, [reducemax_param_1];
	ld.param.u32 	%r9, [reducemax_param_2];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	.loc 2 8 1
	mov.u32 	%r39, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r38, %r39, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r39, %r11;
	.loc 2 8 1
	setp.ge.s32 	%p1, %r38, %r9;
	mov.f32 	%f29, 0fFF7FFFFF;
	mov.f32 	%f30, %f29;
	@%p1 bra 	BB0_2;

BB0_1:
	.loc 2 8 1
	mul.wide.s32 	%rd6, %r38, 4;
	add.s64 	%rd7, %rd2, %rd6;
	ld.global.f32 	%f6, [%rd7];
	.loc 3 435 5
	max.f32 	%f30, %f30, %f6;
	.loc 2 8 1
	add.s32 	%r38, %r38, %r4;
	.loc 2 8 1
	setp.lt.s32 	%p2, %r38, %r9;
	mov.f32 	%f29, %f30;
	@%p2 bra 	BB0_1;

BB0_2:
	.loc 2 8 1
	mul.wide.s32 	%rd8, %r2, 4;
	mov.u64 	%rd9, __cuda_local_var_33851_32_non_const_sdata;
	add.s64 	%rd3, %rd9, %rd8;
	st.shared.f32 	[%rd3], %f29;
	bar.sync 	0;
	.loc 2 8 1
	setp.lt.u32 	%p3, %r39, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	.loc 2 8 1
	mov.u32 	%r7, %r39;
	shr.u32 	%r39, %r7, 1;
	.loc 2 8 1
	setp.ge.u32 	%p4, %r2, %r39;
	@%p4 bra 	BB0_5;

	.loc 2 8 1
	ld.shared.f32 	%f7, [%rd3];
	add.s32 	%r15, %r39, %r2;
	mul.wide.u32 	%rd10, %r15, 4;
	add.s64 	%rd12, %rd9, %rd10;
	ld.shared.f32 	%f8, [%rd12];
	.loc 3 435 5
	max.f32 	%f9, %f7, %f8;
	.loc 2 8 1
	st.shared.f32 	[%rd3], %f9;

BB0_5:
	.loc 2 8 1
	bar.sync 	0;
	.loc 2 8 1
	setp.gt.u32 	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	.loc 2 8 1
	setp.gt.s32 	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	.loc 2 8 1
	ld.volatile.shared.f32 	%f10, [%rd3];
	ld.volatile.shared.f32 	%f11, [%rd3+128];
	.loc 3 435 5
	max.f32 	%f12, %f10, %f11;
	.loc 2 8 1
	st.volatile.shared.f32 	[%rd3], %f12;
	ld.volatile.shared.f32 	%f13, [%rd3+64];
	ld.volatile.shared.f32 	%f14, [%rd3];
	.loc 3 435 5
	max.f32 	%f15, %f14, %f13;
	.loc 2 8 1
	st.volatile.shared.f32 	[%rd3], %f15;
	ld.volatile.shared.f32 	%f16, [%rd3+32];
	ld.volatile.shared.f32 	%f17, [%rd3];
	.loc 3 435 5
	max.f32 	%f18, %f17, %f16;
	.loc 2 8 1
	st.volatile.shared.f32 	[%rd3], %f18;
	ld.volatile.shared.f32 	%f19, [%rd3+16];
	ld.volatile.shared.f32 	%f20, [%rd3];
	.loc 3 435 5
	max.f32 	%f21, %f20, %f19;
	.loc 2 8 1
	st.volatile.shared.f32 	[%rd3], %f21;
	ld.volatile.shared.f32 	%f22, [%rd3+8];
	ld.volatile.shared.f32 	%f23, [%rd3];
	.loc 3 435 5
	max.f32 	%f24, %f23, %f22;
	.loc 2 8 1
	st.volatile.shared.f32 	[%rd3], %f24;
	ld.volatile.shared.f32 	%f25, [%rd3+4];
	ld.volatile.shared.f32 	%f26, [%rd3];
	.loc 3 435 5
	max.f32 	%f27, %f26, %f25;
	.loc 2 8 1
	st.volatile.shared.f32 	[%rd3], %f27;

BB0_8:
	.loc 2 8 1
	setp.ne.s32 	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	.loc 2 8 1
	ld.shared.u32 	%r36, [__cuda_local_var_33851_32_non_const_sdata];
	.loc 3 1881 5
	atom.global.max.s32 	%r37, [%rd1], %r36;

BB0_10:
	.loc 2 9 2
	ret;
}


`
