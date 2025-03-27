# 基于现有镜像
FROM mirrors.tencent.com/todacc/trpc-python-compile_tlinux3.1:0.1.3

# 安装 Python 3.8+ 的依赖
RUN yum install -y gcc openssl-devel bzip2-devel libffi-devel zlib-devel wget make

# 下载并安装 Python 3.11+
RUN wget https://www.python.org/ftp/python/3.11.6/Python-3.11.6.tgz && \
    tar xzf Python-3.11.6.tgz && \
    cd Python-3.11.6 && \
    ./configure --enable-optimizations && \
    make altinstall && \
    cd .. && \
    rm -rf Python-3.11.6 Python-3.11.6.tgz

# 设置 Python 3.11 为默认 Python 版本
RUN alternatives --install /usr/bin/python3 python3 /usr/local/bin/python3.11 1 && \
    alternatives --set python3 /usr/local/bin/python3.11

# 验证 Python 版本
RUN python3 --version

# 安装 pip
RUN curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py && \
    python3 get-pip.py && \
    rm get-pip.py

# 验证 pip 版本
RUN pip3 --version
