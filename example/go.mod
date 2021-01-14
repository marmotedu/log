module github.com/marmotedu/log/example

go 1.14

replace github.com/marmotedu/log => /home/colin/workspace/golang/src/github.com/marmotedu/log/

require (
	github.com/marmotedu/log v0.0.0-00010101000000-000000000000
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.0 // indirect
	go.uber.org/zap v1.15.0
	k8s.io/klog v1.0.0 // indirect
)
