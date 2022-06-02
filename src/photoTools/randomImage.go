package photoTools

type RandomPhotoToolInterface interface {
	GenerateImage()
	GetExtraMessage()
}

type RandomPhotoTool struct {
}
