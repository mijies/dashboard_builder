WINFLAGS	= GOOS=windows GOARCH=amd64
PROGRAM		= dashboard_builder.exe

win: $(PROGRAM)

$(PROGRAM):
	$(WINFLAGS) go build -o $(PROGRAM) main.go

dep:
	-go get golang.org/x/sys/windows > /dev/null
	-go get gopkg.in/Knetic/govaluate.v3
	-go get github.com/akavel/rsrc
	-go get github.com/lxn/win
	-go get github.com/lxn/walk
	-go get github.com/360EntSecGroup-Skylar/excelize
	

.PHONY: clean
clean:
	-@rm *.exe
	@echo cleaned up
