# Game Programming in Golang Code

This repository contains the sample source code for [*Game Programming in C++* by Sanjay Madhav](https://github.com/gameprogcpp/code) trying to implement by Golang.

Trying to reproduce the original implementation as much as possible,
but since Golang doesn't have Inheritance, using `interface` and `embedding` instead.

In addition, processing load and efficiency are not taken into consideration.

This repository is for studying purposes only and is unrelated with the original authors.

# Requirements
* Go 1.22 or later
* SDL2([go-sdl2](https://github.com/veandco/go-sdl2))

# How to run
* Edit `import` in `main.go` to choose the chapter you want to run.
* Download assets from [original repository](https://github.com/gameprogcpp/code) and put them in the `Assets` directory.
* Run the following command.

```bash
$ go run main.go
```
