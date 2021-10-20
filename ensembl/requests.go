package main

import "github.com/asmcos/requests"

func main (){

        resp,err := requests.Get("http://www.zhanluejia.net.cn")
        if err != nil{
          return
        }
        println(resp.Text())
}
