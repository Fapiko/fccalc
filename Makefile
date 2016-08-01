install:
	go install github.com/Fapiko/fccalc
	go build -o $$GOPATH/bin/fccalc-autocompletion github.com/Fapiko/fccalc/autocompletion
	sudo GOPATH=$$GOPATH $$GOPATH/bin/fccalc-autocompletion
