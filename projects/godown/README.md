#### Description
godown (Go Downloader)

#### Features:
- Download file to directory
- Auto create directory when user does not specify destination path
- Go technique:
    - goroutines
    - flag
    - colly
    - defer
    - file

#### Usage
Command:
```xpath
$ go build godown.go
$ ./godown --urlOption="https://www.shopify.com/sitemap.xml" --desPath="/home/huynq"
```
Result:
```xpath
Requesting https://www.shopify.com/sitemap.xml ...
  Download completed.
  --------------------
  {sitemap.xml /home/huynq/GoDownDir/sitemap.xml 2018-12-06 15:00:53.612344971 +0700 +07 -rw-r--r-- 330677}

```

Use this everywhere:
```xpath
$ sudo cp godown /usr/local/bin
$ godown --url="https://www.shopify.com/sitemap.xml" --desPath="/home/huynq/abc"
```