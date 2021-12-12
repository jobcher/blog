# mysql 笔记（1）

# mysql 学习笔记（1）
本文章不涉及到关于mysql开放上的问题，主要记录关于mysql出现的问题，以及如何去维护mysql数据的日常。  
  
## mysql各类信息的收集
1. 收集变量信息
```sql
show global variables;
```
2. 收集进程信息
```sql
show PROCESSLIST;
```
3. 收集错误日志
```sql
show global variables like 'log_error';
```
4. 收集慢日志信息
```sql
show global variables like 'slow_querry_log_file';
```
5. 收集锁信息，高峰时期运行三次，每次间隔10s
```sql
SELECT  locked_table,
        locked_index,
        locked_type,
        blocking_pid,
        T2.USER  blocking_user,
        T2.HOST blocking_host,
        blocking_lock_mode,
        blocking_trx_rows_modified,
        waiting_pid,
        T3.USER waiting_user,
        T3.HOST waiting_host,
        waiting_lock_mode,
        waiting_trx_row_modified,
        wait_age_secs,
        waiting_query
FROM sys.x$innodb_lock_waits T1
LEFT JOIN INFROMATION_SCHEMA.processlist T2 ON T1.blocking_pid=T2.ID
LEFT JOIN INFROMATION_SCHEMA.processlist T3 ON T3.ID=T1.waiting_pid;

```
6. 收集mysql状态信息
```sql
show global status;
show engine innodb status;
show engine innodb mutex;
```

## mysql 基础语法

1. 连接数据库
```sql
mysql -u <用户名> -p
```
  
2. 创建数据库
```sql
CREATE DATABASE <数据库名称>;
```
  
3. 删除数据库
```sql
drop database <数据库名称>;
```
  
4. 选择数据库
```sql
use <数据库名称>;
```
  
