package main

import (
	"path/filepath"
	"os/exec"
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
	"fmt"
)

type CONFIG struct{
	Url      string `json:"url"`
	Asset   string `json:"asset"`
}

func main() {
	fmt.Println("欢迎输入热更生成器");
	config,err :=getConfig(getCurrentDirectory()+"/config.json");
    if(err!=nil){
		fmt.Println(getCurrentDirectory()+"/config.json is not exist");
		return ;
	}

	fmt.Println("请选择热更新模式:1 表示制作base包; 2表示制作remote包");
	reader := bufio.NewReader(os.Stdin)
	data1, _, _ := reader.ReadLine()
	_type := string(data1)
	var model = ""
	var asset = ""

	if(_type=="2"){
		fmt.Println("12312312");
	}

    switch _type {
	case "1":
		model = "制作base包";
		asset = "assets/";
	   break;	
    case "2":
	   model = "制作remote包";
	   asset = "remote-assets/";
	  break; 
	default:
		fmt.Println("热更新模式不正确,请重新来过");
		return; 	
   }

	fmt.Println("请选择热更新模式"+model);

	fmt.Print("请输入当前的热更版本号:");

    data, _, _ := reader.ReadLine()
    version := string(data)
	fmt.Println("当前的热更版本为"+version);
	var nodeShell = "cd "+getCurrentDirectory()+" && node version_generator.js -v "+version+" -u "+config.Url+" -s "+config.Asset+" -d "+asset;

	if(_type=="2"){
		nodeShell = "cd "+getCurrentDirectory()+" && rm -R remote-assets && node version_generator.js -v "+version+" -u "+config.Url+" -s "+config.Asset+" -d "+asset+" && cp -R "+config.Asset+"/src remote-assets/src && cp -R "+config.Asset+"/res remote-assets/res && rm remote-assets.tar.gz && tar -cvf remote-assets.tar.gz remote-assets";
	}

	fmt.Println(nodeShell);
	strData := execCommand(nodeShell);
	fmt.Println("Execute finished:" + strData)
	fmt.Print("创建成功");
}

func ReadAll(filePth string) ([]byte,error) {
	f, err := os.Open(filePth)
	if err != nil {
	   return nil, err;
	}
	ret,err := ioutil.ReadAll(f);
	return ret,nil;
   }

func getConfig(filePth string)(CONFIG,error){
	var con CONFIG
	  str,err := ReadAll(filePth);
	  json.Unmarshal(str, &con)
      return con,err;  
   }

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		
	}
	return dir;
}

func execCommand(strCommand string)(string){
    cmd := exec.Command("/bin/bash", "-c", strCommand)
 
 
    stdout, _ := cmd.StdoutPipe()
    if err := cmd.Start(); err != nil{
        fmt.Println("Execute failed when Start:" + err.Error())
        return ""
    }
 
    out_bytes, _ := ioutil.ReadAll(stdout)
    stdout.Close()
 
    if err := cmd.Wait(); err != nil {
        fmt.Println("Execute failed when Wait:" + err.Error())
        return ""
    }
    return string(out_bytes)
}