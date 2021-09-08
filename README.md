# cpp_go
```
go build -o libcallback.so  -buildmode=c-shared a.go b.go
```
![1](https://user-images.githubusercontent.com/72860476/132514036-43bcec97-5949-4847-b18a-7d5ba44df25f.png)


```
gcc -v  m.cpp  -o m ./libcallback.so 
```
![2](https://user-images.githubusercontent.com/72860476/132514055-1613ff07-0976-41dd-913f-c3bad2a0eb1f.png)


```
./m
```
![3](https://user-images.githubusercontent.com/72860476/132514074-0534a57c-1ffa-43e4-920e-4cb689a4717d.png)

