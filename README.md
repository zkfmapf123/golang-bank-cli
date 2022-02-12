## golang-cli

### init path

```
    1.
    go env
    기존 PATH를 현재폴더로 명시적으로 지정한다.
    
    export GOPATH = 현재 파일 위치
    export PATH = "$GOPATH/bin:${PATH}"

    2.
    mkdir bin pkg src
    go mod init (module name)
```

### use makefile 

```
    touch Makefile
```

### execute

```
    make prod-rund

    파일안에 main.exec 파일을 실행
```