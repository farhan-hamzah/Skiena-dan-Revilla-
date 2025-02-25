package main
import "fmt"

func main(){
    var masukan int
    fmt.Scan(&masukan)

    for masukan != 1{
        if masukan%2 == 0{
            masukan = masukan/2
            fmt.Println(masukan)
        }else{
            masukan = masukan*3 + 1
            fmt.Println(masukan)
        }
    }
}