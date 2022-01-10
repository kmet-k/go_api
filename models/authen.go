package models

type AuthenMember struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Authen struct {
	ID           string `json:"id,omitempty" bson:"id,omitempty"`
	Firstname    string `json:"firstname" bson:"firstname"`
	Lastname     string `json:"lastname" bson:"lastname"`
	Username     string `json:"username" bson:"username"`
	AccessToken  string `json:"accessToken" bson:"accessToken"`
	RefreshToken string `json:"refreshToken" bson:"refreshToken"`
}

type RespondAuthen struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Data    Authen `json:"data"`
}

type Message struct {
	Th string `json:"th"`
	En string `json:"en"`
}
