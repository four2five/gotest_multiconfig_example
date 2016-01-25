# gotest_multiconfig_example
An example showing an issue that surfaces when using ginkgo in conjunction with go test.

Running "go test" in the simple_stringutil/ directory succeeds, as expected.
Running "go test -v" yields this error:
joseph:simple_stringutil$ go test -v
=== RUN   TestReverse
flag provided but not defined: -test.v
Usage of /var/folders/24/vx7zzdy56wzf64sz7kw57hyh0000gp/T/go-build006013306/github.com/four2five/simple_stringutil/_test/simple_stringutil.test:
  -encoding
    	Change value of Encoding. (default Binary)
  -interesting
    	Change value of Interesting.

Generated environment variables:
   STRINGMETADATA_ENCODING
   STRINGMETADATA_INTERESTING

exit status 2
FAIL	github.com/four2five/simple_stringutil	0.006s
joseph:simple_stringutil$
