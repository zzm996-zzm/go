package main

import (
  "context"
  "fmt"
  "github.com/pkg/errors"
  "golang.org/x/sync/errgroup"
  "log"
  "net/http"
  "os"
  "os/signal"
  "syscall"
  "time"
)

func main(){
  //我设置了整个程序的声明周期是1分钟
  rootCtx,cancel:=context.WithTimeout(context.Background(),time.Second * 60)
  defer cancel()
  g, ctx := errgroup.WithContext(rootCtx)

  mux := http.NewServeMux()
  //模拟不断在消费数据
  mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {

      w.Write([]byte("pong"))

  })


  // 模拟该接口要报错
  Out := make(chan struct{})
  mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
    Out <- struct{}{}
  })



  server := http.Server{
    Handler: mux,
    Addr:    ":8080",
  }


  //goroutine id:1
  //如果1号goroutine退出那么其他的子goroutine都退出 然后main函数退出
  g.Go(func() error{
    return server.ListenAndServe()
  })

  g.Go(func() error{
    select{
      case <-ctx.Done():
        fmt.Println("context：exit")
      case <-Out:
        fmt.Println("Out: exit")
    }
    //手动shutdown如果比较卡住了可以timeout自动关闭
    timeCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    log.Println("shutting down server...")
    return server.Shutdown(timeCtx)
  })


  g.Go(func() error {
    quit := make(chan os.Signal)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    select {
    case <-ctx.Done():
      fmt.Println("ctx <- DONE")
      return ctx.Err()
    case sig := <-quit:
      return errors.Errorf("get os signal: %v", sig)
    }
  })

  fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}