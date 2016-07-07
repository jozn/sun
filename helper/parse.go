package helper

import "strconv"

func StrToInt(str string, defualt int) int{
    r64,err :=strconv.ParseInt(str,10,64)
    if err!= nil{
        return defualt
    }
    return int(r64)
}

func IntToStr(num int) string  {
    return strconv.Itoa(num)
}
