build:
	go build

dev-docker:
	go build -o kirb.tmp


dev-stage1:
	go build && ./kirb && rm -f ./kirb