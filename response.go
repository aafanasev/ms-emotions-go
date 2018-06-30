package emotions

import (
	"bytes"
	"fmt"
)

// Face - Detected face
type Face struct {
	FaceId         string
	FaceRectangle  FaceRectangle
	FaceAttributes FaceAttributes
}

// FaceRectangle - coordinates of face
type FaceRectangle struct {
	Left   int
	Top    int
	Width  int
	Height int
}

type FaceAttributes struct {
	Smile      float32
	HeadPose   HeadPose
	Gender     string
	Age        int
	FacialHair FacialHair
	Glasses    string
	Emotion    Emotion
	Blur       Blur
	Exposure   Exposure
	Noise      Noise
	Makeup     Makeup
	Accessory  []Accessory
	Occlusion  Occlusion
	Hair       Hair
}

type HeadPose struct {
	Pitch float32
	Roll  float32
	Yaw   float32
}

type FacialHair struct {
	Moustache float32
	Beard     float32
	Sideburns float32
}

type Emotion struct {
	Anger     float32
	Contempt  float32
	Disgust   float32
	Fear      float32
	Happiness float32
	Neutral   float32
	Sadness   float32
	Surprise  float32
}

type Blur struct {
	BlurLevel string
	Value     float32
}

type Exposure struct {
	ExposureLevel string
	Value         float32
}

type Noise struct {
	NoiseLevel string
	Value      float32
}

type Makeup struct {
	EyeMakeup bool
	LipMakeup bool
}

type Accessory struct {
}

type Occlusion struct {
	ForeheadOccluded bool
	EyeOccluded      bool
	MouthOccluded    bool
}

type Hair struct {
	Bald      float32
	Invisible bool
	HairColor []HairColor
}

type HairColor struct {
	Color      string
	Confidence float32
}

// Get scores as String
func (scores *Emotion) String() string {
	var buffer bytes.Buffer

	if anger := round(scores.Anger * 100); anger > 0 {
		buffer.WriteString(fmt.Sprintf("Anger: %d%%\n", anger))
	}
	if contempt := round(scores.Contempt * 100); contempt > 0 {
		buffer.WriteString(fmt.Sprintf("Contempt: %d%%\n", contempt))
	}
	if disgust := round(scores.Disgust * 100); disgust > 0 {
		buffer.WriteString(fmt.Sprintf("Disgust: %d%%\n", disgust))
	}
	if fear := round(scores.Fear * 100); fear > 0 {
		buffer.WriteString(fmt.Sprintf("Fear: %d%%\n", fear))
	}
	if happiness := round(scores.Happiness * 100); happiness > 0 {
		buffer.WriteString(fmt.Sprintf("Happiness: %d%%\n", happiness))
	}
	if neutral := round(scores.Neutral * 100); neutral > 0 {
		buffer.WriteString(fmt.Sprintf("Neutral: %d%%\n", neutral))
	}
	if sadness := round(scores.Sadness * 100); sadness > 0 {
		buffer.WriteString(fmt.Sprintf("Sadness: %d%%\n", sadness))
	}
	if surprise := round(scores.Surprise * 100); surprise > 0 {
		buffer.WriteString(fmt.Sprintf("Surprise: %d%%\n", surprise))
	}

	return buffer.String()
}
