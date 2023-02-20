
## 使用
GORM 使用 database/sql 维护连接池（文档：https://pkg.go.dev/database/sql#DB）

## 连接
当我们运行sql.Open()方法时，将打开一个连接池，但并不进行实际的连接。实际的连接定义在以下方法中：

```
// conn returns a newly-opened or cached *driverConn.
func (db *DB) conn(ctx context.Context, strategy connReuseStrategy) (*driverConn, error)
```

### 参数
```
//根据负载测试结果和实际吞吐量级别进行优化。
//设置空闲连接池中链接的最大数量
sqlDb.SetMaxIdleConns(25)
//设置打开数据库链接的最大数量
sqlDb.SetMaxOpenConns(1000)
//设置链接可复用的最大时间
sqlDb.SetConnMaxLifetime(5 * time.Minute)
```

## 释放连接
能过go.sql中的putConnDBLocked()来释放连接


