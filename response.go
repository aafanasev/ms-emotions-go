package emotions

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
