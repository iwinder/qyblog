# go-gitlint使用指南

检查历史提交的 Commit Message 是否符合 Angular 规范，可以将该工具添加在 CI 流程中，确保 Commit Message 都是符合规范的。

## 安装 

```shell
go install github.com/marmotedu/go-gitlint/cmd/go-gitlint@latest
```

## 使用
## githook: commit-msg配置
```shell
# Create hook
echo 'gitlint --msg-file=$1' > .git/hooks/commit-msg
# Make it executable
chmod +x .git/hooks/commit-msg 
``` 

### .gitlint配置

```shell
--subject-regex=^((Merge branch.*of.*)|((revert: )?(feat|fix|perf|style|refactor|test|ci|docs|chore)(\(.+\))?: [^A-Z].*[^.]$))
--subject-maxlen=72
--body-regex=^([^\r\n]{0,72}(\r?\n|$))*$
```

### 检测

```shell
go-gitlint 
```
