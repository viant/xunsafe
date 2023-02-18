package xunsafe

import "os"

var debugEnabled = false

func init() {
	if os.Getenv("XUNSAFE_DEBUG") == "true" {
		debugEnabled = true
	}
}
