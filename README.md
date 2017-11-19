# ping
Ping a server.  Maybe don't use this in Prod yet.  K thanks!

```go
import github.com/guitarbum722/ping

func main() {
    results, err := ping.Ping("www.google.com")
    if err != nil {
		log.Fatalln("err pinging host : ", err.Error())
	}
    fmt.Println(summary)
}
```