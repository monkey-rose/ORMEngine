# ORMEngine
### 实现功能
- 1.自动插入数据（利用反射技术，根据输入数据的类型自动生成插入 sql 语句，实现函数 Insert(o interface{})）
- 2.查询结果自动映射（利用反射技术，根据输入数据的类型自动生成查询 sql 语句，并将结果集合根据数据类型自动映射到对象，并加入结果表，Find(o interface{})）
### 数据库
MySql数据库，数据库表：
```
CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `departname` VARCHAR(64) NULL DEFAULT NULL,
    `createat` VARCHAR(64) NULL DEFAULT NULL,
    PRIMARY KEY (`uid`)
);
```
### 测试样例

