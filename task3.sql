-- 题目1：基本CRUD操作--------

-- 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"
insert into students
select
  "张三" as name,
  "20" as age,
  "三年级" as grade;

-- 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息
select * from students where age > 18;
-- 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"
update students set grade="四年级" where name = "张三";
-- 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录
delete from students where age < 15;


-- 题目2：事务语句--------------
-- 编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务

BEGIN TRANSACTION;

-- 1. 检查账户A余额是否足够
DECLARE @balance DECIMAL(18,2);
SELECT @balance = balance FROM accounts WHERE id = 1; -- 账户A的ID为1

IF @balance < 100
BEGIN
    ROLLBACK TRANSACTION;
    PRINT '转账失败：账户余额不足';
    RETURN;
END

-- 2. 从账户A扣除100元
UPDATE accounts SET balance = balance - 100 WHERE id = 1;

-- 3. 向账户B增加100元
UPDATE accounts SET balance = balance + 100 WHERE id = 2; -- 账户B的ID为2

-- 4. 记录交易信息
INSERT INTO transactions (from_account_id, to_account_id, amount)
VALUES (1, 2, 100);

COMMIT TRANSACTION;
PRINT '转账成功';
