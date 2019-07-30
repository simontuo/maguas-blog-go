package validation

type User struct {
	Name string `form:"name" json:"name" xml:"name" bind:"required"`
}

type Article struct {
	Title string `form:"title" json:"title" xml:"title" bind:"title"`
}

type Comment struct {
	Content string `form:"content" json:"content" xml:"content" bind:"content"`
}

type Tga struct {
	Name string `form:"name" json:"name" xml:"name" bind:"name"`
}
