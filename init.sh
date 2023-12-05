if ! command -v go &> /dev/null; then
    # Go is not installed, proceed with installation
    echo "Go is not installed. Installing..."
    
    # 下载Go语言的安装包
    wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz
    
    # 解压安装包到/usr/local目录
    sudo tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
    
    # 添加Go语言的可执行文件路径到环境变量
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    
    # 刷新当前终端的环境变量配置
    source ~/.bashrc
    
    echo "Go has been installed."
else
    # Go is already installed
    echo "Go is already installed."
fi

# 验证安装是否成功
go version