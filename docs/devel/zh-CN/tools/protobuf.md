# protobuf安装与环境配置

## 安装 protobuf
```shell
# 在任意位置下载 protobuf 源文件包
cd /tmp/
git clone --depth=1 https://github.com/protocolbuffers/protobuf
# 解压之后进入
cd protobuf
sudo ./autogen.sh
sudo ./configure
sudo make
sudo make install
# 查看 protoc 版本，成功输出版本号，说明安装成功
protoc --version 
```
## 安装 protoc-gen-go
```shell
go install github.com/golang/protobuf/protoc-gen-go
```

## 可能存在的问题

### autoreconf: not found

#### 【Q】具体问题
>   ./autogen.sh: 4: autoreconf: not found

#### 【A】解决方案

```shell
sudo apt-get install autoconf automake libtool
```


### error: C++ preprocessor

#### 【Q】具体问题
> checking how to run the C++ preprocessor... /lib/cpp
configure: error: in `/home/wind/Wind/apps/common/protobuf':
configure: error: C++ preprocessor "/lib/cpp" fails sanity check
See `config.log' for more details
>

#### 【A】解决方案

```shell
sudo apt-get install g++
```

### no configuration information is in third_party/googletest

#### 【Q】具体问题
> configure: WARNING: no configuration information is in third_party/googletest

#### 【A】解决方案

执行如下命令，可以检测并补全`third_party`目录下所有文件，包括但不限于`googletest`，是最简单的方式：
```shell

git submodule update --init --recursive
```

也可以直接通过如下完成： 
- 从 https://github.com/google/googletest/releases 获取源文件
- 如 https://github.com/google/googletest/archive/refs/tags/release-1.11.0.tar.gz
- 解压并重命名为`googletest`（即去掉文件夹中的版本号）,并放到 `third_party/`即可。

##  libprotoc.so.30: cannot open shared object file: No such file or directory

#### 【Q】具体问题

> protoc: error while loading shared libraries: libprotoc.so.30: cannot open shared object file: No such file or directory

#### 【A】解决方案

方案一： 最简单的方式是在`/etc/profile` 或` .bashrc` 等相关文件中加入：` export LD_LIBRARY_PATH=/usr/local/lib`,如：
```shell
# vim /etc/profile
# 文件最下面添加：
export LD_LIBRARY_PATH=/usr/local/lib 
# source /etc/profile 
```

方案二： 但使用过程中发现仅当次有效，注销系统后又失效了，故采用了如下方式：

1. 通过 `sudo vim /etc/ld.so.conf.d/libprotobuf.conf ` 新建【/etc/ld.so.conf.d/libprotobuf.conf】 ，内容如下：
    ```shell
    /usr/local/lib
    ```
2. 执行命令 `sudo ldconfig`
3. `/etc/profile` 添加路径
   ```shell
    # vim /etc/profile
    export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH
    # source /etc/profile
    ```