package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

/**
加盐：先将用户输入的密码进行一次MD5（或其它哈希算法）加密；
     将得到的 MD5 值前后加上一些只有管理员自己知道的随机串，再进行一次MD5加密。
     这个随机串中可以包括某些固定的串，也可以包括用户名（用来保证每个用户加密使用的密钥都不一样）

随着并行计算能力的提升，只要时间与资源允许，没有破译不了的密码，因为攻击者有足够的资源建立这么多的rainbow table
所以方案是:故意增加密码计算所需耗费的资源和时间，使得任何人都不可获得足够的资源建立所需的rainbow table
推荐scrypt方案：dk := scrypt.Key([]byte("some password"), []byte(salt), 16384, 8, 1, 32)
*/
func main() {
	// 假设用户名abc，密码123456
	h := md5.New()
	io.WriteString(h, "123456")

	//pw_md5等于e10adc3949ba59abbe56e057f20f883e
	pwMd5 := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("md5密码：", pwMd5)
	//指定两个 salt： salt1 = @#$%   salt2 = ^&*()
	salt1 := "@#$%"
	salt2 := "^&*()"

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, "abc")
	io.WriteString(h, salt2)
	io.WriteString(h, pwMd5)

	last := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println("最终密码：", last)
}
