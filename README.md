# gin MVC 框架
严格意义来说 只有 控制器，模型 2层, 目前流行前后端分离，Go 专做 后端开发

控制器调度模型负责数据组装，
模型负责按业务逻辑 从数据库查询数据

目前演示了 分页查询功能

## 数据库 
采用 gorm 

## 目录结构

### controller 目录
控制器
以 数据表 items, member_items 为例
```bash
controllers/表单词1/[表单词2]/action.go
controllers/items/action.go 
controllers/member/items/action.go 

```

### models 目录

模型
以 数据表 items, member_items 为例
```bash
models/表前缀/表名.go
models/items/items.go
models/member/member_items.go
```


MIT | Copyright © 2015-2022 Kis | 微信: vvveric