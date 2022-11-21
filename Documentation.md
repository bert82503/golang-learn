

Documentation-文档
======
> Practice-『实践出真知识』
> https://golang.google.cn/doc/

The Go programming language is an open source project to make programmers more productive.
Go编程语言是一个开源项目，旨在提高程序员的工作效率。

Go is expressive, concise, clean, and efficient. 
Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, 
while its novel type system enables flexible and modular program construction. 
Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. 
It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.
Go富有表现力、简洁、干净且高效。
它的并发机制使编写程序变得容易，以充分利用多核和联网机器，而其新颖的类型系统支持灵活和模块化的程序构建。
Go可以快速编译为机器代码，同时还具有垃圾收集的便利性和运行时反射的强大功能。
它是一种快速、静态类型、编译的语言，感觉就像一种动态类型、解释的语言。


# Getting started 新手入门
## Installing Go
Instructions for downloading and installing Go.

## Tutorial: Getting started
A brief Hello, World tutorial to get started. Learn a bit about Go code, tools, packages, and modules.
一个简短的"你好，世界"教程开始。了解一些关于 Go 代码、工具、包和模块的信息。

## Tutorial: Create a module
A tutorial of short topics introducing functions, error handling, arrays, maps, unit testing, and compiling.
介绍函数、错误处理、数组、映射、单元测试和编译的简短主题教程。

## A Tour of Go 指南
An interactive introduction to Go in three sections. 
The first section covers basic syntax and data structures; the second discusses methods and interfaces; 
and the third introduces Go's concurrency primitives. 
Each section concludes with a few exercises, so you can practice what you've learned. 
You can take the tour online or install it locally with:
分三个部分对Go进行交互式介绍。
第一部分介绍基本语法和数据结构，第二部分讨论方法和接口，第三部分介绍Go的并发原语。
每个部分都以一些练习结束，这样您就可以练习所学的内容。
您可以通过以下方式在线或在本地安装教程：
```
$ go install golang.org/x/website/tour@latest
```
This will place the tour binary in your GOPATH's bin directory.
这将把旅行二进制文件放在GOPATH的bin目录中。

