```shell
# Go1.16及以前
go get -u github.com/gin-gonic/gin 
```

```shell
#  Go1.17
go install github.com/gin-gonic/gin@latest
go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest
```

```shell
# go1.7提示
go get: installing executables with 'go get' in module mode is deprecated.
        To adjust and download dependencies of the current module, use 'go get -d'.
        To install using requirements of the current module, use 'go install'.
        To install ignoring the current module, use 'go install' with a version,
        like 'go install example.com/cmd@latest'.
        For more information, see https://golang.org/doc/go-get-install-deprecation
        or run 'go help get' or 'go help install'.

```

git tag -a 0.1.0  -m "Test Edition v0.1.0"

选择
- ? What is the URL of your repository? https://gitee.com/windcoder/qingyublog
- ? What is your favorite style? github
- ? Choose the format of your favorite commit message <type>(<scope>): <subject>
- ? What is your favorite template style? standard
- ? Do you include Merge Commit in CHANGELOG? No
- ? Do you include Revert Commit in CHANGELOG? Yes
- ? In which directory do you output configuration files and templates? .chglog
 

https://www.topgoer.com/gin%E6%A1%86%E6%9E%B6/%E7%AE%80%E4%BB%8B.html

https://gin-gonic.com/