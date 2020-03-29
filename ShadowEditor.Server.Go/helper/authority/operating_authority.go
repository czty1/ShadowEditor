package authority

// All return all operating authorities
var All = map[string]string{
	"Administrator":    Administrator,
	"Login":            Login,
	"ListAnimation":    ListAnimation,
	"AddAnimation":     AddAnimation,
	"EditAnimation":    EditAnimation,
	"DeleteAnimation":  DeleteAnimation,
	"ListAudio":        ListAudio,
	"AddAudio":         AddAudio,
	"EditAudio":        EditAudio,
	"DeleteAudio":      DeleteAudio,
	"ListCategory":     ListCategory,
	"SaveCategory":     SaveCategory,
	"DeleteCategory":   DeleteCategory,
	"ListCharacter":    ListCharacter,
	"EditCharacter":    EditCharacter,
	"SaveCharacter":    SaveCharacter,
	"DeleteCharacter":  DeleteCharacter,
	"ListMap":          ListMap,
	"AddMap":           AddMap,
	"EditMap":          EditMap,
	"DeleteMap":        DeleteMap,
	"ListMaterial":     ListMaterial,
	"EditMaterial":     EditMaterial,
	"SaveMaterial":     SaveMaterial,
	"DeleteMaterial":   DeleteMaterial,
	"ListMesh":         ListMesh,
	"AddMesh":          AddMesh,
	"EditMesh":         EditMesh,
	"DeleteMesh":       DeleteMesh,
	"ListParticle":     ListParticle,
	"EditParticle":     EditParticle,
	"SaveParticle":     SaveParticle,
	"DeleteParticle":   DeleteParticle,
	"ListPrefab":       ListPrefab,
	"EditPrefab":       EditPrefab,
	"SavePrefab":       SavePrefab,
	"DeletePrefab":     DeletePrefab,
	"EditScene":        EditScene,
	"SaveScene":        SaveScene,
	"PublishScene":     PublishScene,
	"DeleteScene":      DeleteScene,
	"ListScreenshot":   ListScreenshot,
	"AddScreenshot":    AddScreenshot,
	"EditScreenshot":   EditScreenshot,
	"DeleteScreenshot": DeleteScreenshot,
	"ListVideo":        ListVideo,
	"AddVideo":         AddVideo,
	"EditVideo":        EditVideo,
	"DeleteVideo":      DeleteVideo,
}

const (
	// Administrator 管理员权限
	Administrator string = "Administrator"
	// Login 只有登录用户才能具有的权限
	Login string = "Login"
	// ListAnimation 获取动画列表（不具备该权限无法查看动画列表）
	ListAnimation string = "List Animation"
	// AddAnimation 添加动画
	AddAnimation string = "Add Animation"
	// EditAnimation 编辑动画
	EditAnimation string = "Edit Animation"
	// DeleteAnimation 删除动画
	DeleteAnimation string = "Delete Animation"
	// ListAudio 获取音频列表（不具备该权限无法查看音频列表）
	ListAudio string = "List Audio"
	// AddAudio 添加音频
	AddAudio string = "Add Audio"
	// EditAudio 编辑音频
	EditAudio string = "Edit Audio"
	// DeleteAudio 删除音频
	DeleteAudio string = "Delete Audio"
	// ListCategory 获取分类列表（不具有该权限无法进入分类窗口）
	ListCategory string = "List Category"
	// SaveCategory 保存分类
	SaveCategory string = "Save Category"
	// DeleteCategory 删除分类
	DeleteCategory string = "Delete Category"
	// ListCharacter 获取人物列表（不具备该权限无法进入人物列表）
	ListCharacter string = "List Character"
	// EditCharacter 编辑人物
	EditCharacter string = "Edit Character"
	// SaveCharacter 保存人物
	SaveCharacter string = "Save Character"
	// DeleteCharacter 删除人物
	DeleteCharacter string = "Delete Character"
	// ListMap 获取贴图列表（不具有该权限无法查看贴图列表）
	ListMap string = "List Map"
	// AddMap 添加贴图
	AddMap string = "Add Map"
	// EditMap 编辑贴图
	EditMap string = "Edit Map"
	// DeleteMap 删除贴图
	DeleteMap string = "Delete Map"
	// ListMaterial 获取材质列表（不具有该权限无法查看材质列表）
	ListMaterial string = "List Material"
	// EditMaterial 编辑材质
	EditMaterial string = "Edit Material"
	// SaveMaterial 保存材质
	SaveMaterial string = "Save Material"
	// DeleteMaterial 删除材质
	DeleteMaterial string = "Delete Material"
	// ListMesh 获取模型列表（不具有该权限无法查看模型列表）
	ListMesh string = "List Mesh"
	// AddMesh 添加模型
	AddMesh string = "Add Mesh"
	// EditMesh 编辑模型
	EditMesh string = "Edit Mesh"
	// DeleteMesh 删除模型
	DeleteMesh string = "Delete Mesh"
	// ListParticle 获取粒子列表（不具有该权限无法查看粒子列表）
	ListParticle string = "List Particle"
	// EditParticle 编辑粒子
	EditParticle string = "Edit Particle"
	// SaveParticle 保存粒子
	SaveParticle string = "Save Particle"
	// DeleteParticle 删除粒子
	DeleteParticle string = "Delete Particle"
	// ListPrefab 获取预设体列表（不具有该权限无法查看预设体列表）
	ListPrefab string = "List Prefab"
	// EditPrefab 编辑预设体
	EditPrefab string = "Edit Prefab"
	// SavePrefab 保存预设体
	SavePrefab string = "SavePrefab"
	// DeletePrefab 删除预设体
	DeletePrefab string = "Delete Prefab"
	// EditScene 编辑场景
	EditScene string = "Edit Scene"
	// SaveScene 保存场景
	SaveScene string = "Save Scene"
	// PublishScene 发布场景
	PublishScene string = "Publish Scene"
	// DeleteScene 删除场景
	DeleteScene string = "Delete Scene"
	// ListScreenshot 获取截图列表（不具有该权限无法查看截图列表）
	ListScreenshot string = "List Screenshot"
	// AddScreenshot 添加截图
	AddScreenshot string = "Add Screenshot"
	// EditScreenshot 编辑截图
	EditScreenshot string = "Edit Screenshot"
	// DeleteScreenshot 删除截图
	DeleteScreenshot string = "Delete Screenshot"
	// ListVideo 获取视频列表（不具有该权限无法查看视频列表）
	ListVideo string = "List Video"
	// AddVideo 添加视频
	AddVideo string = "Add Video"
	// EditVideo 编辑视频
	EditVideo string = "Edit Video"
	// DeleteVideo 删除视频
	DeleteVideo string = "Delete Video"
)
