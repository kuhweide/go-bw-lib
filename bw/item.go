package bw

import "time"

type ItemType int

const (
	ItemTypeLogin ItemType = iota + 1
	ItemTypeSecureNote
	ItemTypeCard
	ItemTypeIdentity
)

type item struct {
	Id              string      `json:"id,omitempty"`
	PasswordHistory []string    `json:"passwordHistory"`
	RevisionDate    *time.Time  `json:"revisionDate"`
	CreationDate    *time.Time  `json:"creationDate"`
	DeletedDate     *time.Time  `json:"deletedDate"`
	ArchivedDate    *time.Time  `json:"archivedDate"`
	OrganizationId  *string     `json:"organizationId"`
	CollectionIds   []string    `json:"collectionIds"`
	FolderId        *string     `json:"folderId"`
	Type            ItemType    `json:"type"`
	Name            string      `json:"name"`
	Notes           string      `json:"notes"`
	Favorite        bool        `json:"favorite"`
	Fields          []field     `json:"fields"`
	Login           *Login      `json:"login"`
	SecureNote      *SecureNote `json:"secureNote"`
	Card            *Card       `json:"card"`
	Identity        *Identity   `json:"identity"`
	SshKey          *string     `json:"sshKey"`
	Reprompt        int         `json:"reprompt"`
}

func NewLoginItem(name string) *item {
	item := &item{}

	item.Name = name
	item.Login = &Login{}
	item.Type = ItemTypeLogin

	return item
}

func NewSecureNoteItem(name string) *item {
	item := &item{}

	item.Name = name
	item.SecureNote = &SecureNote{}
	item.Type = ItemTypeSecureNote

	return item
}

func NewCardItem(name string) *item {
	item := &item{}

	item.Name = name
	item.Card = &Card{}
	item.Type = ItemTypeCard

	return item
}

func NewCardIdentity(name string) *item {
	item := &item{}

	item.Name = name
	item.Identity = &Identity{}
	item.Type = ItemTypeIdentity

	return item
}

func (item *item) AddField(field field) {
	item.Fields = append(item.Fields, field)
}
