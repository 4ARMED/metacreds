EXECUTABLE := $(EXECUTABLE)

clean:
	@rm -f ${EXECUTABLE}

install: clean
	go install -a -tags netgo -ldflags "-w -extldflags '-static'"

docker:
	docker build -f build/Dockerfile . -t ${EXECUTABLE} --build-arg EXECUTABLE=$(EXECUTABLE)
