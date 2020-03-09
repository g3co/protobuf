proto:
	protoc -I . mouselive.proto --go_out=plugins=grpc:.

#test