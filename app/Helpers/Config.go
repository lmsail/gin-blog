package Helpers

import (
	config "github.com/Unknwon/goconfig"
	"path/filepath"
	"strings"
)

// 获取配置文件信息 / 获取整块数据时需手动转换成 map 类型
// @params key 			关键字 - 对应根目录 .env 文件中 [key] 部分
// @params defaultValue 默认值 - 未匹配到数据时则返回默认值
func Config(key string, defaultValue interface{}) interface{} {
	split := strings.Split(key, ".")
	cfg, err := config.LoadConfigFile(filepath.Join(GetGoRunPath(), ".env"))
	if err != nil {
		return defaultValue
	}
	section, err := cfg.GetSection(split[0])
	if err != nil {
		return defaultValue
	}
	if index := strings.Index(key, "."); index != -1 {
		if value := section[split[1]]; len(value) > 0 {
			return value
		}
		return defaultValue
	}
	return section
}

// 获取单个配置
func ConfigOne(key, defaultValue string) string {
	cfg, err := config.LoadConfigFile(filepath.Join(GetGoRunPath(), ".env"))
	if err != nil {
		return defaultValue
	}
	if index := strings.Index(key, "."); index < 0 {
		return defaultValue
	}
	split := strings.Split(key, ".")
	section, err := cfg.GetSection(split[0])
	if err != nil {
		return defaultValue
	}
	return section[split[1]]
}

// 获取多个配置
func ConfigMultiple(key string) (list map[string]string) {
	cfg, err := config.LoadConfigFile(filepath.Join(GetGoRunPath(), ".env"))
	if err != nil {
		return list
	}
	section, err := cfg.GetSection(key)
	if err != nil {
		return list
	}
	return section
}
