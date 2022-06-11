package photoTools

import "fmt"

type RandomPhotoToolInterface interface {
	GenerateImageFromUrlEndpoint(url string)
	GetExtraMessage()
	RandomPhotoTool
}

type RandomPhotoTool struct {
}

func (rpt RandomPhotoTool) printSnap() {
	fmt.Print("snap")
}

//needs to have a display Image function?
