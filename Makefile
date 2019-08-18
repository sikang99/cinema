#
# Makefile for mixderdl
#
.PHONY: usage edit build play clean git

usage:
	@echo "usage: make [edit|clean]"

edit e:
	vi cinema.go

build b:
	go build

clean:
	cd test; make clean
#----------------------------------------------------------------------------------
git g:
	@echo "> make (git:g) [update|store]"
git-update gu:
	make clean
	git add .
	git commit -a -m "forked and modify test code"
	git push
git-store gs:
	git config credential.helper store
#----------------------------------------------------------------------------------

