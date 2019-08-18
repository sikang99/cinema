#
# Makefile for cinema_test
#
PROG=cinema_test
usage:
	@echo "usage: make [edit|build|run|clean]"

edit e:
	vi test.go

build b:
	go build -o $(PROG) test.go

run r:
	./$(PROG)

play p:
	ffplay *.mov

clean:
	rm -f $(PROG) *.mp4 *.mov

check:
	ffprobe test_input.mp4
	ffprobe test_output.mov
