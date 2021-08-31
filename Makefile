# Build and Flash
build:
	tinygo build -target pico -o ./hello.uf2

flash:
	tinygo flash -target=pico

clean:
	rm -f ./hello.uf2

