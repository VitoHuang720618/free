

#### entities/ 底下放http request response struct
#### svc/ 底下放 mysql redis etcd grpc 實體
#### logic/ 邏輯業務
#### handler/router.go。  gin group / gin func 邏輯與入口
#### etc/file.yaml 需要與 internal/config/filename.go 中struct 配合