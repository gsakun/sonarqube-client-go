# generate-goclient-for-sonarqube

## 获取api.json
```bash
curl -u username:password "http://sonarqube:9000/api/webservices/list?include_internals=true" -o ./assets/api.json
```
## 修改Makefile