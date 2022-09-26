package GetTreePermission

import (
	"user_center/app/Model"
)

type Resp struct {
	ID       uint       `comment:"权限id" json:"id"`
	Name     string     `comment:"名称" json:"name"`
	ParentID uint       `comment:"父级id" json:"parent_id"`
	Children *RoleTrees `comment:"子集" json:"children"`
}

// RoleTrees 二叉树列表
type RoleTrees []*Resp

func Item(permission []Model.Permission) RoleTrees {
	var list RoleTrees
	for _, v := range permission {
		var info *Resp
		info = new(Resp)
		info.ID = v.ID
		info.Name = v.Name
		info.ParentID = v.ParentID
		list = append(list, info)
	}
	list = toTree(list)
	return list
}

// ToTree 转换为树形结构
func toTree(data RoleTrees) RoleTrees {
	// 定义 HashMap 的变量，并初始化
	TreeData := make(map[uint]*Resp)
	// 先重组数据：以数据的ID作为外层的key编号，以便下面进行子树的数据组合
	for _, item := range data {
		TreeData[item.ID] = item
	}
	// 定义 RoleTrees 结构体
	var TreeDataList RoleTrees
	// 开始生成树形
	for _, item := range TreeData {
		// 如果没有根节点树，则为根节点
		if item.ParentID == 0 {
			// 追加到 TreeDataList 结构体中
			TreeDataList = append(TreeDataList, item)
			// 跳过该次循环
			continue
		}
		// 通过 上面的 TreeData HashMap的组合，进行判断是否存在根节点
		// 如果存在根节点，则对应该节点进行处理
		if pItem, ok := TreeData[item.ParentID]; ok {
			// 判断当次循环是否存在子节点，如果没有则作为子节点进行组合
			if pItem.Children == nil {
				// 写入子节点
				children := RoleTrees{item}
				// 插入到 当次结构体的子节点字段中，以指针的方式
				pItem.Children = &children
				// 跳过当前循环
				continue
			}
			// 以指针地址的形式进行追加到结构体中
			*pItem.Children = append(*pItem.Children, item)
		}
	}
	return TreeDataList
}
