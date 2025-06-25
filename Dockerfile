# Base images 基础镜像
FROM mirrors.tencent.com/devcloud/codev-tlinux4:0.0.1

#MAINTAINER 维护者信息
LABEL MAINTAINER="v_wensqi"

ENV PYTHON_VERSION=3.12.2 \
    PYTHON_PIP_VERSION=24.0 \
    PYTHONUNBUFFERED=1 \
    DEBIAN_FRONTEND=noninteractive

#ENV 设置环境变量
ENV PATH /usr/local/nginx/sbin:$PATH

#RUN 执行安装基础工具
RUN yum install -y wget && yum clean all
RUN yum install -y \
    python3-pip \
    cmake \
    gcc \
    gcc-c++ \
    make \
    zlib-devel \
    bzip2-devel \
    openssl-devel \
    sqlite-devel \
    readline-devel \
    libffi-devel \
    xz-devel \
    tree \
    cpio \
    && yum clean all

# 下载并安装 Python 3.12
RUN cd /tmp \
    && wget https://www.python.org/ftp/python/${PYTHON_VERSION}/Python-${PYTHON_VERSION}.tgz \
    && tar xzf Python-${PYTHON_VERSION}.tgz \
    && cd Python-${PYTHON_VERSION} \
    && ./configure \
        --enable-shared \
        --prefix=/usr/local/python3.12 \
        --disable-optimizations \
        --with-ensurepip=install \
    && make -j$(nproc) \
    && make install \
    && ldconfig \
    && cd .. \
    && rm -rf Python-${PYTHON_VERSION} Python-${PYTHON_VERSION}.tgz

# 设置 Python 路径
ENV PATH=/usr/local/python3.12/bin:$PATH \
    LD_LIBRARY_PATH=/usr/local/python3.12/lib:$LD_LIBRARY_PATH


# 安装 pip（分步执行）
RUN /usr/local/python3.12/bin/python3.12 -m ensurepip --upgrade
RUN /usr/local/python3.12/bin/pip3.12 install --upgrade pip==${PYTHON_PIP_VERSION}
RUN /usr/local/python3.12/bin/pip3.12 install --upgrade setuptools wheel


# 创建软链接
RUN ln -sf /usr/local/python3.12/bin/python3.12 /usr/local/bin/python3 \
    && ln -sf /usr/local/python3.12/bin/python3.12 /usr/local/bin/python \
    && ln -sf /usr/local/python3.12/bin/pip3.12 /usr/local/bin/pip3 \
    && ln -sf /usr/local/python3.12/bin/pip3.12 /usr/local/bin/pip

# 加入tRPC环境支持
RUN pip3 install concurrent-log-asyncio>=0.10.1 --index-url https://mirrors.cloud.tencent.com/pypi/simple/
RUN pip3 install automaxprocs>=2.0.0 --index-url https://mirrors.cloud.tencent.com/pypi/simple/
RUN pip3 install trpc_pb>=0.3.0 --index-url https://mirrors.cloud.tencent.com/pypi/simple/
RUN wget https://mirrors.tencent.com/repository/pypi/tencent_pypi/packages/trpc/0.9.0a1/trpc-0.9.0a1-cp312-cp312-linux_x86_64.whl
RUN rm trpc-0.9.0a1-cp312-cp312-linux_x86_64.whl

#WORKDIR 相当于cd
WORKDIR /usr/local/app
# 验证安装
RUN /usr/local/python3.12/bin/python3.12 --version \
    && /usr/local/python3.12/bin/pip3.12 --version
