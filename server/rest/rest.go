package rest

// REST APIハンドラー群
type RestHandlers interface{}

type restHandlersImpl struct{}

// 新しいREST APIハンドラー群を生成
func NewRestHandlers() RestHandlers {
	return &restHandlersImpl{}
}
