



.PHONY: wc-app, clean

wc-app:
	@go build -buildmode=plugin ./mrapp/wc.go


clean:
	rm -f wc.so