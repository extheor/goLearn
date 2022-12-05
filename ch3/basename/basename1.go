/*
不能使用任何库，全部手工硬编码实现。
basename(s)将看起来像是系统路径的前缀删除，同时将看似文件类型的后缀名部分删除。
*/
package basename

func Basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
