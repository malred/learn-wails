package include

import (
	"context"
	"fmt"
)

// App struct
type Calculator struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewCalculator() *Calculator {
	return &Calculator{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (c *Calculator) startup(ctx context.Context) {
	c.ctx = ctx
}

func (c *Calculator) SetContext(ctx context.Context) {
	c.ctx = ctx
}

// 让前端调用的方法
// 计算a+b
func (c *Calculator) Add(a int, b int) string {
	return fmt.Sprintf("a + b = %v", (a + b))
}
