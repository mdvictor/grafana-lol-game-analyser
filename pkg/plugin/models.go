package plugin

import "time"

type LolClient struct {
	ConnectionSettings
}

type ConnectionSettings struct {
	URL          string
	ApiToken     string
	SummonerName string
	SummonerId   string
	UID          string
	Platform     string
	Region       string
	PUUID        string
}

type queryModel struct {
	MatchId            string `json:"matchID"`
	Player             string `json:"player"`
	TimelineData       string `json:"timelineData"`
	RefID              string `json:"refId"`
	NormalizeTimerange bool   `json:"normalizeTimerange"`
	WithStreaming      bool   `json:"withStreaming"`
}

type SettingsJsonData struct {
	Platform     string `json:"platform"`
	SummonerName string `json:"summonerName"`
}

type SummonerData struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	SummonerLevel int    `json:"summonerLevel"`
}

type RiotError struct {
	Status struct {
		Message    string `json:"message"`
		StatusCode string `json:"status_code"`
	} `json:"status"`
}

type MatchJson struct {
	Metadata struct {
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		Participants []MatchInfo `json:"participants"`
	}
}

type MatchInfo struct {
	ChampionName       string `json:"championName"`
	PUUID              string `json:"puuid"`
	SummonerName       string `json:"summonerName"`
	TeamId             int    `json:"teamId"`
	IndividualPosition string `json:"individualPosition"`
	Win                bool   `json:"win"`
	Kills              int    `json:"kills"`
	Deaths             int    `json:"deaths"`
	Assists            int    `json:"assists"`
	MatchId            string `json:"matchId"`
}

type Participant struct {
	ParticipantId int    `json:"participantId"`
	PUUID         string `json:"puuid"`
}

type ParticipantFrame struct {
	CurrentGold         int     `json:"currentGold"`
	GoldPerSecond       float64 `json:"goldPerSecond"`
	MinionsKilled       int     `json:"minionsKilled"`
	JungleMinionsKilled int     `json:"jungleMinionsKilled"`
	Level               int     `json:"level"`
	TotalGold           int     `json:"totalGold"`
	ChampionStats       struct {
		AbilityHaste      int `json:"abilityHaste"`
		AbilityPower      int `json:"abilityPower"`
		Armor             int `json:"armor"`
		ArmorPen          int `json:"armorPen"`
		AttackDamage      int `json:"attackDamage"`
		AttackSpeed       int `json:"attackSpeed"`
		CcReduction       int `json:"ccReduction"`
		CooldownReduction int `json:"cooldownReduction"`
		HealthMax         int `json:"healthMax"`
		HealthRegen       int `json:"healthRegen"`
		Lifesteal         int `json:"lifesteal"`
		MagicPen          int `json:"magicPen"`
		MagicResist       int `json:"magicResist"`
		MovementSpeed     int `json:"movementSpeed"`
		Omnivamp          int `json:"omnivamp"`
		PhysicalVamp      int `json:"physicalVamp"`
		PowerMax          int `json:"powerMax"`
		SpellVamp         int `json:"spellVamp"`
	} `json:"championStats"`
	DamageStats struct {
		MagicDamageDone               int `json:"magicDamageDone"`
		MagicDamageDoneToChampions    int `json:"magicDamageDoneToChampions"`
		MagicDamageTaken              int `json:"magicDamageTaken"`
		PhysicalDamageDone            int `json:"physicalDamageDone"`
		PhysicalDamageDoneToChampions int `json:"physicalDamageDoneToChampions"`
		PhysicalDamageTaken           int `json:"physicalDamageTaken"`
		TotalDamageDone               int `json:"totalDamageDone"`
		TotalDamageDoneToChampions    int `json:"totalDamageDoneToChampions"`
		TotalDamageTaken              int `json:"totalDamageTaken"`
		TrueDamageDone                int `json:"trueDamageDone"`
		TrueDamageDoneToChampions     int `json:"trueDamageDoneToChampions"`
		TrueDamageTaken               int `json:"trueDamageTaken"`
	} `json:"damageStats"`
}

type Frame struct {
	Events []struct {
		RealTimestamp int64  `json:"realTimestamp"`
		Type          string `json:"type"`
	}
	ParticipantFrames struct {
		ParticipantOneFrame   ParticipantFrame `json:"1"`
		ParticipantTwoFrame   ParticipantFrame `json:"2"`
		ParticipantThreeFrame ParticipantFrame `json:"3"`
		ParticipantFourFrame  ParticipantFrame `json:"4"`
		ParticipantFiveFrame  ParticipantFrame `json:"5"`
		ParticipantSixFrame   ParticipantFrame `json:"6"`
		ParticipantSevenFrame ParticipantFrame `json:"7"`
		ParticipantEightFrame ParticipantFrame `json:"8"`
		ParticipantNineFrame  ParticipantFrame `json:"9"`
		ParticipantTenFrame   ParticipantFrame `json:"10"`
	}
	Timestamp int64 `json:"timestamp"`
}

type MatchTimeline struct {
	Info struct {
		Frames       []Frame       `json:"frames"`
		Participants []Participant `json:"participants"`
	}
}

type TimeLineDataFrameValues struct {
	Timestamp time.Time `json:"timestamp"`
	Value     float64   `json:"value"`
}

type MatchParticipant struct {
	PUUID              string `json:"puuid"`
	SummonerName       string `json:"summonerName"`
	ChampionName       string `json:"championName"`
	IndividualPosition string `json:"individualPosition"`
	TeamId             int    `json:"teamId"`
}
