package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/golang/glog"
	"github.com/magicsong/generate-go-for-sonarqube/pkg/api"
	"github.com/magicsong/generate-go-for-sonarqube/pkg/generate"
)

var (
	h           bool
	JsonPath    string
	OutputPath  string
	PackageName string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&JsonPath, "f", "", "specify location of api file(only support local file)")
	flag.StringVar(&OutputPath, "o", ".", "specify the destination dir, default is current workspace")
	flag.StringVar(&PackageName, "n", "sonarqube", "specify the name of generated package,default is \"sonarqube\"")
	flag.Usage = usage
}

func validate() error {
	reg := regexp.MustCompile(`[a-z]+`)
	if !reg.MatchString(PackageName) {
		return errors.New("Ilegal package name:" + PackageName)
	}
	if JsonPath == "" {
		return errors.New("Must specify the json location,please add -f [filepath]")
	}
	_, err := os.Stat(JsonPath) //os.Stat获取文件信息
	if err != nil {
		glog.Errorln(err)
		return errors.New("No such api file")
	}
	return nil
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	err := validate()
	if err != nil {
		glog.Error(err)
		os.Exit(1)
	}
	file, err := os.Open(JsonPath)
	if err != nil {
		glog.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	myapi := new(api.API)
	err = decoder.Decode(myapi)
	if err != nil {
		glog.Error(err)
		glog.Errorln("cannot decode api file")
		os.Exit(1)
	}
	err = generate.Build(PackageName, OutputPath, myapi)
	if err != nil {
		glog.Fatal(err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, ` generate-go-for-sonarqube version: 0.0.1
Usage: main.go [-h] -f jsonpath [-o outputpath] 

Options:
`)
	flag.PrintDefaults()
}