# Template

為了解決[gin-gonic/gin](https://github.com/gin-gonic/gin)的LoadHTMLGlob方法在巢狀資料夾中無法讀取及轉換模板上的問題

##  安裝

```

$ go get -u github.com/RickyJian/template

```

## 使用方法

```go

import(
	"github.com/RickyJian/template"
	"github.com/gin-gonic/gin"
)

func main (){
	router := gin.Default()
	// 參數1：檔案位置、參數2：副檔名
	t := template.TemplateWalk("assets/template/",".html")
	router.SetHTMLTemplate(t)
}

```