package validation

type User struct {
	Name string `form:"name" json:"name" xml:"name" bind:"required"`
}