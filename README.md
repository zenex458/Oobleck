<h1 align="center">Oobleck</h1>
<p align="center">Your own personal pentester</p>

## About
Oobleck is a system hardener/auto-pentesting tool which provides diagnostics and suggestions to increase the security and anonymity of your system.

At its core, it is a series of scripts that test some isolated portion of your system and report on any potential vulnerabilities it finds. 

## Use
At the moment, Oobleck must be compiled from source.

### Compiling from source and running the program
1. Install Golang from your package manager or the [Golang website](https://go.dev/)
2. Install git from your package manager or the [git-scm website](https://git-scm.com/)
3. Clone the repo with `git clone https://github.com/NCFAGlobal/Oobleck`
4. Run `make run` from the root of the project

Alternatively, if you do not wish to compile all the built-in tests, you can simply run `make build/Oobleck` instead of step 4.
