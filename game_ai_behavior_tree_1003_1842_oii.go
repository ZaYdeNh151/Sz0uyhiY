// 代码生成时间: 2025-10-03 18:42:38
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/labstack/echo"
)

// BehaviorTreeNode 定义了行为树中节点的接口
type BehaviorTreeNode interface {
    Tick() (bool, error)
}

// BaseNode 行为树节点的基础结构体
type BaseNode struct {
    children []BehaviorTreeNode
}

// AddChild 添加子节点
func (node *BaseNode) AddChild(child BehaviorTreeNode) {
    node.children = append(node.children, child)
}

// Tick 遍历所有子节点
func (node *BaseNode) Tick() (bool, error) {
    for _, child := range node.children {
        if success, err := child.Tick(); err != nil {
            return false, err
        } else if success {
            return true, nil
        }
    }
    return false, nil
}

// SequenceNode 序列节点，所有子节点必须成功执行
type SequenceNode struct {
    BaseNode
}

// Tick 实现 SequenceNode 的 Tick 方法
func (node *SequenceNode) Tick() (bool, error) {
    for _, child := range node.children {
        if success, err := child.Tick(); err != nil {
            return false, err
        } else if !success {
            return false, nil
        }
    }
    return true, nil
}

// SelectorNode 选择节点，任意一个子节点成功即成功
type SelectorNode struct {
    BaseNode
}

// Tick 实现 SelectorNode 的 Tick 方法
func (node *SelectorNode) Tick() (bool, error) {
    for _, child := range node.children {
        if success, err := child.Tick(); err != nil {
            return false, err
        } else if success {
            return true, nil
        }
    }
    return false, nil
}

// ActionNode 动作节点，执行具体动作
type ActionNode struct {
    name string
    action func() (bool, error)
}

// NewActionNode 创建一个新的 ActionNode
func NewActionNode(name string, action func() (bool, error)) *ActionNode {
    return &ActionNode{name: name, action: action}
}

// Tick 实现 ActionNode 的 Tick 方法
func (node *ActionNode) Tick() (bool, error) {
    return node.action()
}

// GameAIBehaviorTreeHandler 行为树的 HTTP 处理函数
func GameAIBehaviorTreeHandler(c echo.Context) error {
    behaviorTree := createBehaviorTree()
    success, err := behaviorTree.Tick()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error occurred: %s
", err.Error()))
    }
    if success {
        return c.JSON(http.StatusOK, "Behavior tree executed successfully
")
    } else {
        return c.JSON(http.StatusOK, "Behavior tree execution failed
")
    }
}

// createBehaviorTree 创建行为树
func createBehaviorTree() BehaviorTreeNode {
    root := &SelectorNode{}

    root.AddChild(&SequenceNode{
        BaseNode: BaseNode{
            children: []BehaviorTreeNode{
                NewActionNode("Action1", func() (bool, error) {
                    fmt.Println("Executing Action 1")
                    return true, nil
                }),
            },
        },
    })

    root.AddChild(&SequenceNode{
        BaseNode: BaseNode{
            children: []BehaviorTreeNode{
                NewActionNode("Action2", func() (bool, error) {
                    fmt.Println("Executing Action 2")
                    return true, nil
                }),
            },
        },
    })

    return root
}

func main() {
    e := echo.New()
    e.GET("/behavior-tree", GameAIBehaviorTreeHandler)
    log.Fatal(e.Start(":8080"))
}