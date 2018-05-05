package config

import "testing"

const test_info = "正在测试 %s 方法..."

// NOTE: 由于代码路径原因，要通过本测试，您有可能需要手动将 PATH_FILE_CONFIG 常量修改为 ../config.toml 方能通过。
func TestLoadConfig(t *testing.T) {
	t.Logf(test_info, "LoadConfig")
	LoadConfig()

	// NOTE: 下面的测试仅用作调试输出用途
	t.Logf("最大工作线程数: %d", ServerCfg.GoMaxProcs)
	t.Logf("监听地址: %s", ServerCfg.ListenAddr)

	t.Logf("测试完成。")
}