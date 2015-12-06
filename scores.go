package emotions

import (
	"bytes"
	"fmt"
)

// Scores - emotion values
type Scores struct {
	Anger     float32
	Contempt  float32
	Disgust   float32
	Fear      float32
	Happiness float32
	Neutral   float32
	Sadness   float32
	Surprise  float32
}

// Get scores as String
func (scores *Scores) String() string {
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
