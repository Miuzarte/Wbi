# Wbi
Bilibili wbi sign / 哔哩哔哩wbi鉴权签名

Standard libraries only / 纯原生库实现

### From

[github.com/SocialSisterYi/bilibili-API-collect/.../wbi.md#Golang](https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/misc/sign/wbi.md#Golang)

### Usage

```go
import (
	"fmt"
	"github.com/Miuzarte/Wbi"
)

func main() {
	mid := 1850091
	url := fmt.Sprintf("https://api.bilibili.com/x/space/wbi/acc/info?mid=%v", mid)
	urlWithWbi, err := Wbi.Sign(url)
	fmt.Println(urlWithWbi)
	fmt.Println(err)
}

func getKeys1() (imgKey, subKey string) {
	imgKey, subKey = Wbi.GetWbiKeysCached()
	return
}

func getKeys2() (imgKey, subKey string) {
	imgKey, subKey = Wbi.GetWbiKeys()
	return
}
```