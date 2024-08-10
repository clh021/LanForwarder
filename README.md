# LanForwarder (小小接线员)
一个用 Go 语言编写的简单但强大的端口转发工具，旨在帮助用户轻松地将本地计算机上的指定端口流量转发到局域网内的任何设备上。

## 功能特点
端口转发：能够将本地端口上的流量无缝转发到局域网内的目标 IP 和端口。
简单易用：通过命令行参数即可快速配置和启动。
高性能：利用 Go 语言的并发特性，能够高效处理大量并发连接。
跨平台：支持 Windows、Linux 和 macOS 等操作系统。
安全性：虽然本工具不包含加密功能，但它为网络调试和测试提供了便利的基础。

## 使用方法
```bash
make build # output: ./lan-forwarder
./lan-forwarder
# curl -vvv http://<runIP>:<listen_port>/<target_ip>/<target_port>
# curl -vvv http://<runIP>:<listen_port>/<target_host_tag>
# curl -vvv http://<runIP>:<proxy_port>
# curl -vvv http://<runDomain>:<proxy_port>
# curl -vvv http://<customDomainAsProxy>
```

## 注意事项
在生产环境中使用本工具时，请确保遵守网络安全政策。
本工具仅适用于合法合规的网络测试和调试目的。
