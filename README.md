# json2test
Reverts the result of `go tool test2json`.

## Reason
Imagine you want an _Automated Go Tests_ report in two different formats.
- To produce [HTML](https://en.wikipedia.org/wiki/HTML) you can use [`go-test-report`](https://github.com/vakenbolt/go-test-report)
- To obtain [`junit.xml`](https://www.google.com/search?q=junit.xml) you may use [`go-junit-report`](https://github.com/jstemmer/go-junit-report)

But!

The first one wants json from `go test -json` in its input, while the second one wants a clean output from `go test` command.

## Steps towards solution
### Very naive attempt
Run tests twice.
1. `go test -v -json | go-test-report`
1. `go test -v | go-junit-report`

### A little dissatisfaction leads to thinking.
The above solution is far far away from a touch of elegance. Almost like Lord Farquaad.

When we look at this pseudo json we notice that it contains the full output of the stdout of the command `go test`.
```json
{"Time":"1989-06-04T12:00:00.000000000+01:00","Action":"run","Package":"inteligo.com.pl/srv-template-go","Test":"TestAllCases"}
{"Time":"1989-06-04T12:00:00.000000000+01:00","Action":"output","Package":"inteligo.com.pl/srv-template-go","Test":"TestAllCases","Output":"=== RUN   TestAllCases\n"}
{"Time":"1989-06-04T12:00:00.000000000+01:00","Action":"run","Package":"inteligo.com.pl/srv-template-go","Test":"TestAllCases/La_configuration."}
{"Time":"1989-06-04T12:00:00.000000000+01:00","Action":"output","Package":"inteligo.com.pl/srv-template-go","Test":"TestAllCases/La_configuration.","Output":"=== RUN   TestAllCases/La_configuration.\n"}
⋮
{"Time":"1989-06-04T12:00:00.000000000+01:00","Action":"output","Package":"inteligo.com.pl/srv-template-go","Test":"TestAllCases/La_configuration.","Output":"    --- PASS: TestAllCases/La_configuration. (0.30s)\n"}
⋮
{"Time":"1989-06-04T12:00:00.000000000+01:00","Action":"output","Package":"inteligo.com.pl/srv-template-go","Output":"PASS\n"}
{"Time":"1989-06-04T12:00:00.000000000+01:00","Action":"pass","Package":"inteligo.com.pl/srv-template-go","Elapsed":0.416}
```
Notice the lines that have "Action": "output" `!

So maybe I could get it out? This idea leads us to the next step.

### Less naive attempt
Run tests once and use it's pseudo json output as [single source of truth](https://en.wikipedia.org/wiki/Single_source_of_truth)
```sh
go test -v -json > ssot
cat ssot | go-test-report -o report.html
cat ssot | json2test | go-junit-report > junit.xml
```

## Conclusion
It is obvious that the approach taken is patching.

## Two directions to check.
1. Make `go-junit-report` ingest pseudo json from the command `go test -json`.
1. Get interested in [gopogh](https://github.com/medyagh/gopogh)

## Last word.
However, if my tool is useful to you in any way, use it and _if you can_ let me know that it was useful to someone more than just me.
