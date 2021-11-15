# GOBCO - Golang Branch Coverage

Branch coverage measurement tool for golang.


## Install and Usage
```sh
git clone https://github.com/ShriprasadM/gobco.git
cd gobco
go get...
cd sample
go test -v -cover -toolexec 'gobco-tool'  -coverpkg ./...
```



## Sample output (With Current Error)

### Tests and modules for POC

```sh
(base) shriprasad@L881 sample % tree
.
├── bar.go
├── foo.go
├── foo_test.go
└── shri_test
    ├── module1
    │   └── module1.go
    ├── module2
    │   └── module2.go
    ├── module3
    │   └── module3.go
    └── modules.go

4 directories, 7 files
```
We can see in following sample output, nested packages are instrumented, but, i think when go compiler tries to load them, somehow itis not able to load it.
Any guess, how to resolve this ?

If it is resolved, I think it would be able to get branch coverage on nested packages

```sh
(base) shriprasad@L881 sample % go1.16.10 test -v -cover -toolexec 'gobco-tool'  -coverpkg ./... 
# cover github.com/junhwi/gobco/sample/shri_test/module1
file = /tmp/gobco/sample/shri_test/module1/module1.go
# github.com/junhwi/gobco/sample/shri_test/module1
/var/folders/_q/40r7ttwx7gvcgy_nv46hz0x40000gp/T/go-build4283635409/b055/module1.cover.go:3:8: can't find import: "github.com/junhwi/gobco"
gobco: exit status 2
# cover github.com/junhwi/gobco/sample/shri_test/module2
file = /tmp/gobco/sample/shri_test/module2/module2.go
# github.com/junhwi/gobco/sample/shri_test/module2
/var/folders/_q/40r7ttwx7gvcgy_nv46hz0x40000gp/T/go-build4283635409/b056/module2.cover.go:3:8: can't find import: "github.com/junhwi/gobco"
gobco: exit status 2
# cover github.com/junhwi/gobco/sample/shri_test/module3
file = /tmp/gobco/sample/shri_test/module3/module3.go
# github.com/junhwi/gobco/sample/shri_test/module3
/var/folders/_q/40r7ttwx7gvcgy_nv46hz0x40000gp/T/go-build4283635409/b057/module3.cover.go:3:8: can't find import: "github.com/junhwi/gobco"
gobco: exit status 2
FAIL	github.com/junhwi/gobco/sample [build failed]
```

Generate an HTML report from profile:
```
$ gobco -html=c.out -o gobco.html
```
