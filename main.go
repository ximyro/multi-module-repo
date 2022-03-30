package main

import(
 "github.com/ximyro/multi-module-repo/b"
 "github.com/ximyro/multi-module-repo/a"
)


func main(){
b := b.NewB(a.NewA())
b.CallA()
}
