proto:
	protoc -I mouselive mouselive/mouselive.proto --go_out=plugins=grpc:mouselive