package photoTools

type RandomPhotoToolInterface interface {
	GenerateImageFromUrlEndpoint(url string)
	GetExtraMessage()
}

type RandomPhotoTool struct {
}
