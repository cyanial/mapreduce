
.PHONY: wc-app, clean, mrsequential

words=$(wildcard ./main/words/pg*.txt)

echo:
	@echo $(words)

mrsequential: wc-app
	@go run -race main/mrsequential.go wc.so $(words)

mrcoordinator:
	@go run -race

mrworker: wc-app
	@go run 



# apps
wc-app:
	@go build -buildmode=plugin ./mrapp/wc.go

clean:
	rm -f wc.so
	rm -f mr-out-*