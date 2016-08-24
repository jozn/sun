package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/fetchbot"
    "io/ioutil"
    "io"
    "os"
    "strings"
    "golang.org/x/net/html"
    "bytes"
    "math/rand"
)

var queue *fetchbot.Queue

func main() {
	f := fetchbot.New(fetchbot.HandlerFunc(handler))
    f.DisablePoliteness = true
	queue = f.Start()
	//queue.SendStringGet("http://golang.org", "http://golang.org/doc")
	queue.SendStringGet("http://www.shemalepornsex.com/")
	//queue.Close()
    queue.Block()
}

func handler(ctx *fetchbot.Context, res *http.Response, err error) {
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Printf("[%d] %s %s\n", res.StatusCode, ctx.Cmd.Method(), ctx.Cmd.URL())

    b,_:=ioutil.ReadAll(res.Body)
    newBuf := bytes.NewBuffer(b)
    links(newBuf, ctx.Cmd.URL().String())

    bs := string(b)
    name:= ctx.Cmd.URL().String()
    name= strings.Replace(name,":","_",-1)
    name= strings.Replace(name,"/","_",-1)
    name= strings.Replace(name,"?","_",-1)
    name= strings.Replace(name,"&","",-1)

    if strings.HasSuffix(name,".jpg") || strings.HasSuffix(name,".jpeg"){
        f ,err := os.Create("./grabs4/"+name+".jpg")
        if err == nil {
            f.Write(b)
            //f.WriteString( strings.Join(strings.Split(bs," ") , "\n") )
            f.Close()
        }else {
            fmt.Println(err)
        }
    _ = bs
    }
    /*f ,err := os.Create("./grabs4/"+name+".hrml")
    if err == nil {
        f.Write(b)
        //f.WriteString( strings.Join(strings.Split(bs," ") , "\n") )
        f.Close()
    }else {
        fmt.Println(err)
    }
    _ = bs*/

    //fmt.Println(bs[:50])
}

func links(r io.Reader, siteUrl string)  {
    //~~~~~~~~~~~~~~~~~~~~~~~~~~~~//
    // Parse HTML for Anchor Tags //
    //~~~~~~~~~~~~~~~~~~~~~~~~~~~~//

    z := html.NewTokenizer(r)

    for {
        tt := z.Next()

        switch {
        case tt == html.ErrorToken:
            // End of the document, we're done
            return



        case tt == html.StartTagToken:
            t := z.Token()
            tag := t.Data

            isAnchor := tag == "a"
            if isAnchor {
                //fmt.Println("We found a link!")
                for _, a := range t.Attr {
                    if a.Key == "href" {
                        //path := fmt.Println("Found href:", a.Val)
                        path :=  a.Val
                        url := path
                        if ! strings.HasPrefix(path,"http") {
                            url = siteUrl + path
                        }

                        if rand.Intn(100) == 2 {
                            queue.SendStringGet(url)
                        }

                    }
                }

            }

            //imgs
            if tag == "img" {
                for _, a := range t.Attr {
                    if a.Key == "src" {
                        //path := fmt.Println("Found href:", a.Val)
                        path :=  a.Val
                        url := path
                        if ! strings.HasPrefix(path,"http") {
                            url = siteUrl + path
                        }

                        if strings.HasSuffix(url,".jpg") || strings.HasSuffix(url,".jpeg"){
                            queue.SendStringGet(url)
                        }

                    }
                }
            }

        }
    }

}