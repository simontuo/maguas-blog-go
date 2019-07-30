package validation

// 用户
type User struct {
	Name string `form:"name" json:"name" xml:"name" bind:"required"`
}

// 文章
type Article struct {
	Title string `form:"title" json:"title" xml:"title" bind:"title"`
}

// 评论
type Comment struct {
	Content string `form:"content" json:"content" xml:"content" bind:"content"`
}

// 标签
type Tga struct {
	Name string `form:"name" json:"name" xml:"name" bind:"name"`
}
