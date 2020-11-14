package main

import (
    "fmt"
    "io/ioutil"
    "os/exec"
)

func main(){
    cmd := exec.Command("/bin/bash", "-c", " kubectl  top node ;echo -----------------------;clock;sleep 1;  kubectl top pod -n websocket-master;kubectl top pod -n websocket-admin;kubectl top pod -n websocket-assist;kubectl top pod -n websocket-magic;clock;netstat -an | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a, S[a]}'|grep --color . ;sleep 1; ")

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

    fmt.Println(" pod占用cpu 和内存如下图所示:" + string(out_bytes))
}

