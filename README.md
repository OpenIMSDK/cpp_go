# cpp_go
```
go build -o libcallback.so  -buildmode=c-shared a.go b.go
```

```
gcc -v  m.cpp  -o m ./libcallback.so 
```

```
./m
```
