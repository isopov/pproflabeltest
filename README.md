Unfortunately I have failed to reproduce this failure without bazel, but here is minimalistic test with bazel. I use bazelisk to run bazel and simply took latest bazel and rules_go versions.

To observe failing test run it with:
```
bazel test --runs_per_test=100 foo_test
```
I get 
```
INFO: Found 1 test target...
Target //:foo_test up-to-date:
  bazel-bin/foo_test_/foo_test
INFO: Elapsed time: 0.194s, Critical Path: 0.05s
INFO: 101 processes: 1 internal, 100 processwrapper-sandbox.
INFO: Build completed, 1 test FAILED, 101 total actions
//:foo_test                                                              FAILED in 18 out of 100 in 0.0s
```
this way with test logs
```
Executing tests from //:foo_test
-----------------------------------------------------------------------------
--- FAIL: TestFoo (0.00s)
    foo_test.go:32: no pprof label whatdoyousay in goroutine profile goroutine profile: total 4
        1 @ 0x415fa1 0x477361
        #	0x415fa0	runtime.runfinq+0x0	GOROOT/src/runtime/mfinal.go:176
        
        1 @ 0x42f9d1 0x46e9fd 0x537c71 0x537aa5 0x5348cb 0x570e0b 0x56bdb4 0x477361
        #	0x537c70	runtime/pprof.writeRuntimeProfile+0xb0	GOROOT/src/runtime/pprof/pprof.go:793
        #	0x537aa4	runtime/pprof.writeGoroutine+0x44	GOROOT/src/runtime/pprof/pprof.go:752
        #	0x5348ca	runtime/pprof.(*Profile).WriteTo+0x14a	GOROOT/src/runtime/pprof/pprof.go:374
        #	0x570e0a	foo_test.TestFoo+0xca			foo_test.go:26
        #	0x56bdb3	testing.tRunner+0xf3			GOROOT/src/testing/testing.go:1690
        
        1 @ 0x46f5ae 0x406c7c 0x406852 0x56cc4b 0x56eeb7 0x56bdb4 0x56ed9d 0x56d7ca 0x571754 0x439ecb 0x477361
        #	0x56cc4a	testing.(*T).Run+0x3aa		GOROOT/src/testing/testing.go:1751
        #	0x56eeb6	testing.runTests.func1+0x36	GOROOT/src/testing/testing.go:2168
        #	0x56bdb3	testing.tRunner+0xf3		GOROOT/src/testing/testing.go:1690
        #	0x56ed9c	testing.runTests+0x43c		GOROOT/src/testing/testing.go:2166
        #	0x56d7c9	testing.(*M).Run+0x649		GOROOT/src/testing/testing.go:2034
        #	0x571753	main.main+0x353			bazel-out/k8-fastbuild/bin/foo_test_/testmain.go:136
        #	0x439eca	runtime.main+0x28a		GOROOT/src/runtime/proc.go:272
        
        1 @ 0x5297e1 0x477361
        #	0x5297e0	os/signal.loop+0x0	GOROOT/src/os/signal/signal_unix.go:21
        
FAIL
```