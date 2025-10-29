package maillist

import (
	"fmt"
	"net/http"

	"github.com/titpetric/platform"
)

type MailList struct{}

func NewMailList() *MailList {
	return &MailList{}
}

func (*MailList) Name() string {
	return "maillist"
}

func (*MailList) Close() {
}

func (m *MailList) Mount(r platform.Router) {
	r.Route("/maillist", func(r platform.Router) {
		r.Get("/", m.Index)
		r.Post("/create", m.Create)
	})
}

func (m *MailList) Index(w http.ResponseWriter, r *http.Request) {
	perms := NewPermissions(r)

	fmt.Fprintf(w, "Permissions: %#v", perms)
}

func (m *MailList) Create(w http.ResponseWriter, r *http.Request) {
}
