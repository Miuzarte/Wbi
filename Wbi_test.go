package Wbi

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	mid := 1850091
	url := fmt.Sprintf("https://api.bilibili.com/x/space/wbi/acc/info?mid=%v", mid)
	urlWithWbi, err := Sign(url)
	fmt.Println(urlWithWbi)
	fmt.Println(err)
}

func TestSign(t *testing.T) {
	fmt.Println(Sign("https://api.bilibili.com/x/space/wbi/acc/info?mid=1850091"))
}

func TestGetWbiKeysCached(t *testing.T) {
	fmt.Println(GetWbiKeysCached())
}

func TestGetWbiKeys(t *testing.T) {
	fmt.Println(GetWbiKeys())
}
