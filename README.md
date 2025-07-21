# tunnel-proxy

轻量化的隧道代理，使用 OpenResty 做请求转发

## 启动 OpenResty

```bash
docker-compose up
```

启动成功之后，可以通过访问 `http://localhost:8080` 来测试，应该会有类似以下输出内容

```markdown
Hello, World from OpenResty + Docker Compose!
```

## 隧道代理

配置文件详见 `conf/runnel_proxy_redis.conf`，可以根据实际需要进行修改。 需要修改其中的 redis 配置。

其中 redis zset 数据示例类似

```bash
zadd tunnel_proxy_pool 0 37.27.253.44:8099 0 46.218.15.74:80 0 185.49.31.207:8081
```

### 隧道代理示例请求

```bash
go run main.go
```