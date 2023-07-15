export PATH=$PATH:$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin
protoc -I proto --go_out=internal ./proto/*.proto