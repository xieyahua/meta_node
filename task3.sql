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




----------------------------------------------进阶gorm------------------------------------------------------------------------------------------------------------------------
-- 题目1：模型定义
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Posts    []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"not null"`
	Content  string    `gorm:"type:text"`
	UserID   uint      `gorm:"not null"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text"`
	PostID  uint   `gorm:"not null"`
	UserID  uint   `gorm:"not null"`
}

func main() {
	dsn := "root:lucaxie123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移创建表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		panic("failed to migrate database")
	}
}


-- 题目2：关联查询-------------------
package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:50;unique;not null"`
	Email    string `gorm:"size:100;unique;not null"`
	Password string `gorm:"size:255;not null"`
	Posts    []Post // 一对多关系：用户拥有多篇文章
}

type Post struct {
	gorm.Model
	Title    string    `gorm:"size:100;not null"`
	Content  string    `gorm:"type:text;not null"`
	UserID   uint      // 外键
	User     User      // 反向引用
	Comments []Comment // 一对多关系：文章拥有多条评论
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	UserID  uint   // 评论作者ID
	PostID  uint   // 外键
	Post    Post   // 反向引用
}

// 查询用户所有文章及评论
func GetUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Preload("Comments").
		Where("user_id = ?", userID).
		Find(&posts).Error
	return posts, err
}

// 查询评论最多的文章
func GetMostCommentedPost(db *gorm.DB) (Post, error) {
	var post Post
	err := db.Model(&Post{}).
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		First(&post).Error
	return post, err
}

func main() {
	dsn := "root:lucaxie123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}

	// 示例1：查询用户1的所有文章及评论
	posts, err := GetUserPostsWithComments(db, 1)
	if err != nil {
		fmt.Println("查询失败:", err)
	} else {
		fmt.Printf("用户1的文章数量: %d\n", len(posts))
		for _, post := range posts {
			fmt.Printf("文章《%s》有%d条评论\n", post.Title, len(post.Comments))
		}
	}

	// 示例2：查询评论最多的文章
	mostCommented, err := GetMostCommentedPost(db)
	if err != nil {
		fmt.Println("查询失败:", err)
	} else {
		fmt.Printf("评论最多的文章是《%s》\n", mostCommented.Title)
	}
}



-- 题目3：钩子函数---------------------------------

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Posts    []Post // 一对多关系：用户拥有多篇文章
}

type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint      // 外键，关联用户
	User     User      // 属于用户
	Comments []Comment // 一对多关系：文章拥有多条评论
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint // 外键，关联文章
	Post    Post // 属于文章
	UserID  uint // 外键，关联评论者
	User    User // 属于评论者
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).Where("id = ?", p.UserID).
		UpdateColumn("post_count", gorm.Expr("post_count + 1")).Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	if err := tx.Model(&Comment{}).
		Where("post_id = ?", c.PostID).
		Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		return tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Update("comment_status", "无评论").Error
	}
	return nil
}

