CC := gcc

FLAGS := -lm

SRCS := $(wildcard */*.cpp)
BINS := $(SRCS:%.cpp=%)

.PHONY = all clean

all: ${BINS}

*/%: */%.o
	${CC} ${FLAGS} $< -o $@

*/%.o: */%.cpp */%.hpp
	${CC} -c $<

clean:
	@echo "Cleaning up..."
	@rm -rvf *.o ${BINS}
