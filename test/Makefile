#
# Makefile for cinema_test
#
PROG=cinema_test
usage:
	@echo "make [edit|build|run|clean]"

edit e:
	vi test.go

build b:
	go build -o $(PROG) test.go

run r:
	./$(PROG)

clean:
	rm -f $(PROG) *.mp4 *.mov

check:
	ffprobe test_input.mp4
	ffprobe test_output.mov
