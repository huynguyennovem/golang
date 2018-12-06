package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"os/user"
	"strings"
)

const BUG_GOD_DIR_SUBFIX string = "/BugGodDir"
const SUCCESS_CODE int = 200

type File struct {
	name       string
	path       string
	date       string
	permission string
	capacity   int64
}

func getFileName(url string) string {
	arrUrl := strings.Split(url, "/")
	if len(arrUrl) <= 0 {
		return ""
	}
	fileName := arrUrl[len(arrUrl)-1]
	return fileName
}

func readFileInfo(fileAbsPath string) File {
	var f File
	fi, err := os.Stat(fileAbsPath)
	if err != nil {
		fmt.Println(err)
		return f
	}
	f.capacity = fi.Size()
	f.name = fi.Name()
	f.path = fileAbsPath
	f.date = fi.ModTime().String()
	f.permission = fi.Mode().String()
	return f
}

func goroutinesCollyDownload(url string, downloadChan chan File, pathCache string) {

	var f File

	c := colly.NewCollector(
		colly.CacheDir(pathCache))

	//rp, err := proxy.RoundRobinProxySwitcher(proxies)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//c.SetProxyFunc(rp)

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Requesting " + request.URL.String() + " ...")
	})

	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode != SUCCESS_CODE {
			log.Fatal("Fail to get.")
			downloadChan <- f
		}
		fileName := getFileName(url)
		fileDir := pathCache + "/" + fileName
		err := response.Save(fileDir)
		if err != nil {
			log.Fatal("Fail to save.")
		}
		f = readFileInfo(fileDir)
	})
	c.Visit(url)
	downloadChan <- f
}

func downloadCompleted(file File) {
	fmt.Println("Download completed.\n--------------------")
	fmt.Println(file)
}

func getDesPath() string {
	usrDir, err := user.Current()
	if err != nil {
		return ""
	}
	bugGoPath := usrDir.HomeDir + BUG_GOD_DIR_SUBFIX
	if _, err := os.Stat(bugGoPath); os.IsNotExist(err) {
		err := os.Mkdir(usrDir.HomeDir+BUG_GOD_DIR_SUBFIX, 0700)
		if err != nil {
			log.Fatal("Fail to create folder.")
		}
	}
	return bugGoPath
}

func main() {

	// define options
	urlOption := flag.String("urlOption", "", "Desinate urlOption (Required)")
	desPathOption := flag.String("desPath", "", "Destinate path (Optional)")

	// parse command flags from os.flag(1)
	flag.Parse()

	// validate flag
	if flag.Parsed() == false {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *urlOption == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// build a download channel
	downloadChannel := make(chan File)

	// passing flag param to command
	var desPath string
	if *desPathOption == "" {
		desPath = getDesPath()
	} else {
		desPath = *desPathOption
	}

	go goroutinesCollyDownload(*urlOption, downloadChannel, desPath)

	file := <-downloadChannel
	// TODO: handle this when file is nil

	// job after download task has finished
	defer downloadCompleted(file)
}
