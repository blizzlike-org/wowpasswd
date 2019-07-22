all: wowpasswd

wowpasswd:
	go get golang.org/x/crypto/ssh/terminal
	go build

clean:
	rm -rf wowpasswd
