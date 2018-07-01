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

// FaceAttributes
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

// HeadPose
type HeadPose struct {
	Pitch float32
	Roll  float32
	Yaw   float32
}

// FacialHair
type FacialHair struct {
	Moustache float32
	Beard     float32
	Sideburns float32
}

// Emotion
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

// Blur
type Blur struct {
	BlurLevel string
	Value     float32
}

// Exposure
type Exposure struct {
	ExposureLevel string
	Value         float32
}

// Noise
type Noise struct {
	NoiseLevel string
	Value      float32
}

// Makeup
type Makeup struct {
	EyeMakeup bool
	LipMakeup bool
}

// Accessory
type Accessory struct {
}

// Occlusion
type Occlusion struct {
	ForeheadOccluded bool
	EyeOccluded      bool
	MouthOccluded    bool
}

// Hair
type Hair struct {
	Bald      float32
	Invisible bool
	HairColor []HairColor
}

// HairColor
type HairColor struct {
	Color      string
	Confidence float32
}

// ErrorAPI - Emotions API error response
type ErrorAPI struct {
	StatusCode int
	Message    string
}
