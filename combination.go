package main

import "fmt"

type Base struct {
  Name string
}

func (b *Base) SetName(name string) {
  b.Name = name
}

func (b *Base) GetName() string {
  return b.Name
}

// 组合，实现集成
type Child struct {
  base Base    // 这里保存的是Base类型
}

// 重写GetName方法
func (c *Child) GetName() string {
  c.base.SetName("modify...")
  return c.base.GetName()
}

// 实现集成，但需要外部提供一个base的实例
type Child2 struct {
  base *Base  // 这里是指针
}

func (c *Child2) GetName() string {
  c.base.SetName("canuse?")
  return c.base.GetName()
}

func main() {
  c:=new(Child)
  c.base.SetName("world")
  fmt.Println(c.GetName())

  c2:= new(Child2)
  c2.base = new(Base) // 因为Child2里面的Base是指针类型，所以必须提供一个Base的实例
  fmt.Println(c2.GetName())
}
