To observe failing test run it with:
```
for run in {1..100}; do go test; done
```
I get something like
```
PASS
ok  	github.com/pproflabeltest	0.001s
...
PASS
ok  	github.com/pproflabeltest	0.001s
--- FAIL: TestFoo (0.00s)
    foo_test.go:33: no pprof label whatdoyousay in goroutine profile goroutine profile: total 3
        1 @ 0x41a761 0x478a01
        #	0x41a760	runtime.runfinq+0x0	/usr/local/go/src/runtime/mfinal.go:179
        
        1 @ 0x434131 0x470a7d 0x50c931 0x50c765 0x50958b 0x522da8 0x4dba74 0x478a01
        #	0x50c930	runtime/pprof.writeRuntimeProfile+0xb0		/usr/local/go/src/runtime/pprof/pprof.go:796
        #	0x50c764	runtime/pprof.writeGoroutine+0x44		/usr/local/go/src/runtime/pprof/pprof.go:755
        #	0x50958a	runtime/pprof.(*Profile).WriteTo+0x14a		/usr/local/go/src/runtime/pprof/pprof.go:377
        #	0x522da7	github.com/isopov/pproflabeltest.TestFoo+0x107	/home/isopov/tmp/pproflabeltest/foo_test.go:27
        #	0x4dba73	testing.tRunner+0xf3				/usr/local/go/src/testing/testing.go:1792
        
        1 @ 0x4716ce 0x40d225 0x40cdd2 0x4dc991 0x4dec97 0x4dba74 0x4deb74 0x4dd52a 0x52311b 0x43e5eb 0x478a01
        #	0x4dc990	testing.(*T).Run+0x430		/usr/local/go/src/testing/testing.go:1859
        #	0x4dec96	testing.runTests.func1+0x36	/usr/local/go/src/testing/testing.go:2279
        #	0x4dba73	testing.tRunner+0xf3		/usr/local/go/src/testing/testing.go:1792
        #	0x4deb73	testing.runTests+0x4b3		/usr/local/go/src/testing/testing.go:2277
        #	0x4dd529	testing.(*M).Run+0x649		/usr/local/go/src/testing/testing.go:2142
        #	0x52311a	main.main+0x9a			_testmain.go:45
        #	0x43e5ea	runtime.main+0x28a		/usr/local/go/src/runtime/proc.go:283
        
FAIL
exit status 1
FAIL	github.com/pproflabeltest	0.002s
PASS
ok  	github.com/pproflabeltest	0.001s
PASS
...
PASS
ok  	github.com/pproflabeltest	0.001s
```