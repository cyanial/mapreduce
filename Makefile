
.PHONY: wc-app, clean, mrsequential

words=$(wildcard ./main/words/pg*.txt)

echo:
	@echo $(words)

mrsequential: wc-app
	@go run main/mrsequential.go wc.so $(words)

wc-app:
	@go build -buildmode=plugin ./mrapp/wc.go


clean:
	rm -f wc.so
	rm -f mr-out-*