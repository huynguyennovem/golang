#### Description
godown (Go Downloader)

#### Usage
Command:
```go
$ go build godown.go
$ ./godown --urlOption="https://www.shopify.com/sitemap.xml"
```
Result:
```go

Requesting https://www.shopify.com/sitemap.xml ...
  Download completed.
  --------------------
  {sitemap.xml /home/huynq/GoDownDir/sitemap.xml 2018-12-06 15:00:53.612344971 +0700 +07 -rw-r--r-- 330677}

```

Use this everywhere:
```xpath
sudo cp godown /usr/local/bin
```