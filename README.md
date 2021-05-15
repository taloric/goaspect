# goaspect
An implement for golang-version AspectF

record execute process and Act as an AOP

Functions should have :

| function name | params | behavior |
| :--- | :---: | --- |
| Retry |  | function retry [] times |
| Delay |  | function run after delay [] seconds |
| RunAsync |  | function run async by go |
| With |  | function execute with condition |
| Until |  | function will not execute until condition returns true  |
| While |  | function will keep run while condition returns true |
| Log |  | log something |
| Watch |  | record execution time |
| TrackError |  | track function error |
| RecoverPanic |  | recover function panic(how?) |

-----
call a function with aspect like:
```go
RunnerStart()
.Retry()
.With()
.Watch()
.Log().Log()
.Do()
```