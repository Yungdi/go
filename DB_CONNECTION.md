# 환경 구성
```bash
$ wget https://golang.org/dl/go1.15.4.linux-amd64.tar.gz
$ sudo tar -xvf go1.15.4.linux-amd64.tar.gz
$ sudo mv go /usr/local
$ export GOROOT=/usr/local/go
$ export GOPATH=$HOME/go
$ export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
$ go get -u github.com/go-sql-driver/mysql
```
