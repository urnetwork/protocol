package protocol

//go:generate protoc -I=. --go_out=. --go_opt=paths=source_relative audit.proto extender.proto frame.proto ip.proto test.proto transfer.proto
