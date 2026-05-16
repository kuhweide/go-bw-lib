package bw

type Login struct {
	Uris     []uri  `json:"uris"`
	Username string `json:"username"`
	Password string `json:"password"`
	Totp     string `json:"totp"`
}

func (login *Login) AddUri(uri uri) {
	login.Uris = append(login.Uris, uri)
}
