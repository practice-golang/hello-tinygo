# Build and Flash
all:
	tinygo build -target pico -o ./hello.uf2
	tinygo flash -target pico ./hello.uf2

build:
	tinygo build -target pico -o ./hello.uf2

flash:
	tinygo flash -target=pico

clean:
	rm -f ./hello.uf2

