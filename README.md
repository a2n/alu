# alu
*The toolset for go programming.*

## Usage
```
import (
	"fmt"
	"alu"
)

func main() {
	// It prints out "The caller is: main.main"
	fmt.Printf("The caller is: %s\n", alu.Caller())
}
```

## Inventory
### Log
-   `Caller`
-   `ToDateString`
-   `ToDateTimeString`
-   `NewLogger`
