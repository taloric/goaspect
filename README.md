# goaspect
An implement for golang-version AspectF

record execute process and act as an AOP

functions list :

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
| TrackPanic |  | track function panic |

------
Retry / Delay / RunAsync : Describe how to run a function

With / Until / While : Describe a function with condition

Log / Watch / TrackPanic : Track function execute process

------
call a function with aspect like:
```go
RunnerStart()
.Retry()
.With()
.Watch()
.Log().Log()
.Execute(func(){
})
```

-----
#### todolist
1. aspect function extension (API Design) 
2. reflect and object type validation on Complete method
3. maybe another implement of interface