package voicevox

/*
各種構造体
json -> struct -> json と扱いが変わるタイミングで 1.0 -> 1 のように値が丸められることがあるが動作には問題ないはず
値がnullになるようなやつはポインタで扱うことで対応している
doc: https://voicevox.github.io/voicevox_engine/api/
*/

type Mora struct {
	Text            string   `json:"text"`
	Consonant       *string  `json:"consonant"`
	ConsonantLength *float64 `json:"consonant_length"`
	Vowel           string   `json:"vowel"`
	VowelLength     float64  `json:"vowel_length"`
	Pitch           float64  `json:"pitch"`
}

type AccentPhrases struct {
	Moras           []Mora `json:"moras"`
	Accent          int    `json:"accent"`
	PauseMora       *Mora  `json:"pause_mora"`
	IsInterrogative *bool  `json:"is_interrogative"`
}

type Query struct {
	AccentPhrases      []AccentPhrases `json:"accent_phrases"`
	SpeedScale         float64         `json:"speedScale"`
	PitchScale         float64         `json:"pitchScale"`
	IntonationScale    float64         `json:"intonationScale"`
	VolumeScale        float64         `json:"volumeScale"`
	PrePhonemeLength   float64         `json:"prePhonemeLength"`
	PostPhonemeLength  float64         `json:"postPhonemeLength"`
	OutputSamplingRate int             `json:"outputSamplingRate"`
	OutputStereo       bool            `json:"outputStereo"`
	Kana               *string         `json:"kana"`
}

type Preset struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	SpeakerUuid       string  `json:"speaker_uuid"`
	StyleId           int     `json:"style_id"`
	SpeedScale        float64 `json:"speedScale"`
	PitchScale        float64 `json:"pitchScale"`
	IntonationScale   float64 `json:"intonationScale"`
	VolumeScale       float64 `json:"volumeScale"`
	PrePhonemeLength  float64 `json:"prePhonemeLength"`
	PostPhonemeLength float64 `json:"postPhonemeLength"`
}

type Style struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Speaker struct {
	Name        string  `json:"name"`
	SpeakerUuid string  `json:"speaker_uuid"`
	Version     *string `json:"version"`
	Styles      []Style `json:"styles"`
}

type StyleInfo struct {
	Id          int      `json:"id"`
	Icon        string   `json:"icon"`
	VoiceSample []string `json:"voice_samples"`
}

type SpeakerInfo struct {
	Policy     string      `json:"policy"`
	Portrait   string      `json:"portrait"`
	StyleInfos []StyleInfo `json:"style_infos"`
}

type DictWord struct {
	Surface               string `json:"surface"`
	Priority              int    `json:"priority"`
	ContextId             *int   `json:"context_id"`
	PartOfSpeech          string `json:"part_of_speech"`
	PartOfSpeechDetail1   string `json:"part_of_speech_detail_1"`
	PartOfSpeechDetail2   string `json:"part_of_speech_detail_2"`
	PartOfSpeechDetail3   string `json:"part_of_speech_detail_3"`
	InflectionalType      string `json:"inflectional_type"`
	InflectionalForm      string `json:"inflectional_form"`
	Stem                  string `json:"stem"`
	Yomi                  string `json:"yomi"`
	Pronunciation         string `json:"pronunciation"`
	AccentType            int    `json:"accent_type"`
	MoraCount             *int   `json:"mora_count"`
	AccentAssociativeRule string `json:"accent_associative_rule"`
}

type Dict map[string]DictWord
