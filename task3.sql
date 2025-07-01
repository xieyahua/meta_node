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
SELECT @balance = balance FROM accounts WHERE id = A; -- 账户A的ID为1

IF @balance < 100
BEGIN
    ROLLBACK TRANSACTION;
    PRINT '转账失败：账户余额不足';
    RETURN;
END

-- 2. 从账户A扣除100元
UPDATE accounts SET balance = balance - 100 WHERE id = A;

-- 3. 向账户B增加100元
UPDATE accounts SET balance = balance + 100 WHERE id = B; -- 账户B的ID为2

-- 4. 记录交易信息
INSERT INTO transactions (from_account_id, to_account_id, amount)
VALUES (A, B, 100);

COMMIT TRANSACTION;
PRINT '转账成功';



----------------------------------------------------------------------------------------------------------------------------------------------------------------------
-- 题目1：使用SQL扩展库进行查询
package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 定义Employee结构体映射表字段
type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
}

func main() {
	// 假设已建立数据库连接
	db, err := sqlx.Connect("mysql", "root:lucaxie123456@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 查询技术部员工
	var techEmployees []Employee
	err = db.Select(&techEmployees,
		"SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatal(err)
	}

	// 打印结果

	fmt.Println("技术部员工：")
	for _, emp := range techEmployees {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %d\n",
			emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// 查询工资最高员工
	var topSalaryEmployee Employee
	err = db.Get(&topSalaryEmployee,
		"SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n工资最高员工：")
	fmt.Printf("%+v\n", topSalaryEmployee)
}


-- 题目2：实现类型安全映射

package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	db, err := sqlx.Connect("mysql", "root:lucaxie123456@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var expensiveBooks []Book
	err = db.Select(&expensiveBooks,
		"SELECT id, title, author, price FROM books WHERE price > ?", 50.0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("价格大于50元的书籍：")
	for _, book := range expensiveBooks {
		fmt.Printf("ID: %d, 书名: %s, 作者: %s, 价格: %.2f\n",
			book.ID, book.Title, book.Author, book.Price)
	}
}

