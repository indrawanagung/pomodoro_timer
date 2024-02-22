# Pomodoro Timer

## Description


The Pomodoro Technique was developed in the late 1980s by then-university student Francesco Cirillo. Cirillo was struggling to focus on his studies and complete assignments. Feeling overwhelmed, he asked himself to commit to just 10 minutes of focused study time. Encouraged by the challenge, he found a tomato (pomodoro in Italian) shaped kitchen timer, and the Pomodoro technique was born.

### Tools Used:

In this project, I use some tools listed below.

- UI Progress Bar : https://github.com/gosuri/uiprogress
- Sending Desktop Notification : https://github.com/gen2brain/beeep
- Displaying Dialogs and Input Boxes : https://github.com/gen2brain/dlgs

### How To Run :
#### Before you start
- Install Go https://go.dev/doc/install

#### Install All Dependency
To install all dependency in this project, run command `go get .`
```bash
### Install all dependency
$ go get .
```
#### Usage
Usage of pomodoro:
```bash
go run main.go -work=[work duration in minute] -rest=[rest duration in minute]
```
Work duration default is 30 minutes and rest duration default  is 5 minutes.

#### Start Pomodoro
```bash
go run main.go -work=30 -rest=5  
```
