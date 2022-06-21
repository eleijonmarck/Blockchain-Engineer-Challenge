package blog

const (
	ModuleName = "blog"
	StoreKey   = ModuleName

	PostKey    = "post"
	CommentKey = "post_slug"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
