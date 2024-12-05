# generate-goclient-for-sonarqube

## 获取api.json
```bash
curl -u username:password "http://sonarqube:9000/api/webservices/list?include_internals=true" -o ./assets/api.json
```
## 修改Makefile
修改参数

## 生成代码
``` bash
make gen
```

## 更新代码
``` bash
make update
```