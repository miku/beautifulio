SHELL := /bin/bash

SIZE = 1073741824
B = 1

# 1048576       |   1 MB
# 10485760      |  10 MB
# 104857600     | 100 MB
# 1073741824    |   1 GB
# 10737418240   |  10 GB
# 107374182400  | 100 GB
# 1099511627776 |   1 TB

noopfilter: main.go
	go build -o $@ $<

.PHONY: run
run: noopfilter
	cat /dev/zero | head -c $(SIZE) | ./noopfilter > /dev/null

.PHONY: runb
runb: noopfilter
	cat /dev/zero | head -c $(SIZE) | ./noopfilter -b $(B) > /dev/null

.PHONY: clean
clean:
	rm -f noopfilter

