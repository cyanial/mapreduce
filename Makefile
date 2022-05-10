
.PHONY: wc-app, clean, mrsequential

words=$(wildcard ./main/words/pg*.txt)

echo:
	@echo $(words)

mrsequential: wc-app
	@go run -race main/mrsequential.go wc.so $(words)

mrcoordinator:
	@go run -race main/mrcoordinator.go $(words)

# mrworker: wc-app
# 	@go run -race main/mrworker.go wc.so


# apps
wc-app: 
	@go build -buildmode=plugin ./mrapp/wc.go

clean:
	rm -f wc.so
	rm -f mr-out-*