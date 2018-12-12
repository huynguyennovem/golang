### Import trong Golang
Khi vọc source code Kubernetes, bạn có thể thấy những dòng import như sau:
```go
import (
	"fmt"
	
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/master/ports"
	"k8s.io/kubernetes/test/e2e/framework"
	imageutils "k8s.io/kubernetes/test/utils/image"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "k8s.io/kubernetes/pkg/api/testapi"
)
```
Những "ký hiệu" như **imageutils**, **.** , **_** sẽ khiến ta bỡ ngỡ khi lần đầu tiếp xúc với Go, vậy chúng là gì và chức năng như nào thì ta cùng tìm hiểu nhé.

**1. Simple import**

Một trong những chương trình đầu tiên khi ta học một ngôn ngữ là print ra 1 chuỗi như sau:
```go
import "fmt"

func main() {
    fmt.Println("Hello world!")
}
```
Khi import fmt package, mọi structs và functions trong package đó sẽ được sử dụng với *fmt.* như fmt.Println ở trên.

**2. Multiple package import**

Để import nhiều package, thay vì gõ nhiều lần import từng package, ta wrap các package lại với *import()*. 
Ví dụ:
```go
import (
    "fmt"
    "bytes"
)
```

**3. Import alias**

Alias cung cấp cho dev một các viết ngắn gọn và dễ nhớ khi dùng package.
Hãy xem xét một đoạn import sau:
```go
import(
	"math/rand"
    "crypto/rand"
)
```
Việc call rand dễ gây ra conflict giữa 2 packages, vì thế trong case này dùng alias là rất hợp lý
```go
import(
	"math/rand"
    crand "crypto/rand"
)
```
**4. Dot import**

Với những dev đã từng dùng Python thì có thể hình dung dot import như sau:

Go's import "os" == Python's import os

Go's import . "os" == Python's from os import *

Format thường dùng là PackageName "." exported_identifier. 
Khi sử dụng dot import thì ta có thể call các exported identifier mà không cần gọi package name

*Note*: exported_identifier: variable/function/... của package mà được exported (public)*

*Nhược điểm*: Khó đọc code cho dev vì không biết được func, struct được call từ đâu. Đa số ý kiến từ community là hạn chế dùng thằng này

Ref: https://golang.org/ref/spec#Import_declarations

**5. Blank import**
Hạn chế việc optimize tự động của Go đối với những unused import. Tức là khi dùng blank import thì bạn vẫn có thể run chương trình khi chưa dùng package đã import.

Ví dụ:
```go
package main

import (
	_ "fmt"
)
func main() {
}

```

**6. Coding convention trong import**

- Import nên được nhóm theo group, cách nhau bởi 1 blank line
- Các package standard luôn đặt trên top, dưới là các third party packages
- Ví dụ:
```go
package main

import (
	"fmt"
	"hash/adler32"
	"os"

	"appengine/foo"
	"appengine/user"

        "github.com/foo/bar"
	"rsc.io/goversion/version"
)
```

- Refer:
	- https://github.com/golang/go/wiki/CodeReviewComments#import-dot
	- https://scene-si.org/2018/01/25/go-tips-and-tricks-almost-everything-about-imports/
