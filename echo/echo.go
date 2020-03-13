package echo

import "fmt"

func Hello(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hello, %s\n", name), nil
	case "zh":
		return fmt.Sprintf("你好, %s\n", name), nil
	default:
		return "", fmt.Errorf("unknow lang")
	}
}
