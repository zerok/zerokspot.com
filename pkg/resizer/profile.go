package resizer

import "fmt"

const FormatJPEG = "jpg"

type Profile struct {
	Format string
	Width  int
}

func (p *Profile) Filename() string {
	return fmt.Sprintf("%d.%s", p.Width, p.Format)
}

type Profiles struct {
	p map[string]Profile
}

func NewProfiles() *Profiles {
	return &Profiles{
		p: make(map[string]Profile),
	}
}

func (p *Profiles) Add(name string, prof Profile) {
	p.p[name] = prof
}

func (p *Profiles) Get(name string) (Profile, bool) {
	prof, ok := p.p[name]
	if ok {
		return prof, true
	}
	return Profile{}, false
}

func (p *Profiles) Names() []string {
	names := make([]string, 0, len(p.p))
	for n := range p.p {
		names = append(names, n)
	}
	return names
}
