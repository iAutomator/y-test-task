package visits

import (
	"html/template"
	"net/http"
	"sync"
)

type counterViewModel struct {
	Visits uint
}

type visitsCounter struct {
	nVisits uint
	m       sync.RWMutex
}

func (c *visitsCounter) visitsCount() uint {
	c.m.RLock()
	v := c.nVisits
	c.m.RUnlock()
	return v
}
func (c *visitsCounter) newVisit() {
	c.m.Lock()
	c.nVisits++
	c.m.Unlock()

}

type Controller struct {
	view *template.Template
	visitsCounter
}

func NewController(t *template.Template) *Controller {
	return &Controller{view: t}
}
func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.newVisit()
	c.view.Execute(w, counterViewModel{c.visitsCount()})
}
