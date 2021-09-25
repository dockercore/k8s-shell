package main

import (
    _ "embed"
    "fmt"
    "io/ioutil"
    "os"
    "os/exec"
)
//go:embed config-ceshi
var s string

//go:embed kubectl
var kube []byte
func main(){
    ioutil.WriteFile("config-t", []byte(s), 0644)
    defer os.Remove("config-t")
    defer os.Remove("kubectl")
    
    ioutil.WriteFile("kubectl", kube, 0777)
    //cp kubectl /usr/local/bin/
    cmdText := "./kubectl  top no --kubeconfig=./config-t | awk 'NR>1 {print \"conduct  \",echo \"nodeIP地址=>\",$1,echo \"   CPU已经使用=> \"$2, echo \"   CPU使用率占比=>\",$3,echo \"  内存已经使用=>\",$4  echo \"   内存使用率占比=> \"$5}';"
   // fmt.Println(cmdText)
   
    fmt.Println("k8s集群 -监控情况如下")
    cmd := exec.Command("/bin/bash", "-c", cmdText)

    stdin, _ := cmd.StdinPipe()
    stdout, _ := cmd.StdoutPipe()

    if err := cmd.Start(); err != nil{
        fmt.Println("Execute failed when Start:" + err.Error())
        return
    }

    stdin.Write([]byte("go text for grep\n"))
    stdin.Write([]byte("go test text for grep\n"))
    stdin.Close()

    out_bytes, _ := ioutil.ReadAll(stdout)
    stdout.Close()

    if err := cmd.Wait(); err != nil {
        fmt.Println("Execute failed when Wait:" + err.Error())
        return
    }

    fmt.Println(" pod占用cpu 和内存如下图所示----\n" + string(out_bytes))
}


