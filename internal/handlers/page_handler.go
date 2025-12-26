package handlers

import (
	"net/http"

	"github.com/jl54/jonasluethi-web/internal/responders"
)

type PageHandler struct {}

type NavigationItem struct {
	Name   string
	Href   string
	Active bool
}

type Navigation struct {
	Items []NavigationItem
}

type AboutBlockListItem struct {
	Text string
}

type AboutBlockSection struct {
	Title   string
	Content []AboutBlockListItem
}

type AboutBlock struct {
	AboutPic      string
	AboutTitle    string
	AboutSections []AboutBlockSection
}

type Experience struct {
	Icon             string
	Year             string
	Company          string
	Position         string
	EmploymentPeriod string
	Description      string
}

type ExperienceBlock struct {
	ExperienceTitle string
	Experiences     []Experience
}

type PageData struct {
	Title           string
	Navigation      Navigation
	AboutBlock      AboutBlock
	ExperienceBlock ExperienceBlock
}

func getNavigation(r *http.Request) Navigation {
	navigation := Navigation{}

	homeLink := NavigationItem{
		Name:   "Home",
		Href:   "/",
		Active: r.URL.Path == "/",
	}

	projectsLink := NavigationItem{
		Name:   "Projects",
		Href:   "/projects",
		Active: r.URL.Path == "/projects",
	}

	navigation.Items = []NavigationItem{homeLink, projectsLink}

	return navigation
}

func getAboutBlock() AboutBlock {
	aboutBlock := AboutBlock{
		AboutPic:   "https://media.licdn.com/dms/image/v2/C4D03AQEPAelAt8RBAA/profile-displayphoto-shrink_200_200/profile-displayphoto-shrink_200_200/0/1539166851386?e=1767830400&v=beta&t=hoFzbIeIGH5spofpLPodFFydIzH12Supt7pFoF1nR9M",
		AboutTitle: "jl54 / README.md",
	}

	aboutMeSection := AboutBlockSection{
		Title: "About me",
		Content: []AboutBlockListItem{
			{Text: "Hi, I'm Jonas"},
			{Text: "Passionate full-stack developer from Switzerland"},
			{Text: "Always learning new technologies and improving my homelab"},
			{Text: "Currently building automation and infrastructure tools for minecraft servers"},
		},
	}

	currentProjectSection := AboutBlockSection{
		Title: "Current Projects",
		Content: []AboutBlockListItem{
			{Text: "Chunk Servers - CLI-Tool to deploy minecraft servers automatically"},
			{Text: "Homelab - Self hosted services, monitoring and automation"},
			{Text: "Portfolio site - It's this website"},
		},
	}

	techStackSection := AboutBlockSection{
		Title: "Tech Stack",
		Content: []AboutBlockListItem{
			{Text: "Languages - PHP, JavaScript, Python, Go"},
			{Text: "Frameworks - Symfony (Pimcore), Laravel, Angular, Vue"},
			{Text: "Tools - Docker, Docker Compose, Kubernetes, Proxmox"},
		},
	}

	getInTouchSection := AboutBlockSection{
		Title: "Get in touch",
		Content: []AboutBlockListItem{
			{Text: "to@jonasluethi.dev"},
		},
	}

	aboutBlock.AboutSections = []AboutBlockSection{
		aboutMeSection,
		currentProjectSection,
		techStackSection,
		getInTouchSection,
	}

	return aboutBlock
}

func getExperienceBlock() ExperienceBlock {
	experienceBlock := ExperienceBlock{
		ExperienceTitle: "Experience",
		Experiences: []Experience{
			{
				Icon:             "/static/images/swiss-dev.png",
				Year:             "Today",
				Company:          "Swiss-Development GmbH",
				Position:         "Full-stack developer",
				EmploymentPeriod: "Sep 2021 - today",
			},
			{
				Icon:             "/static/images/swiss-dev.png",
				Year:             "2019",
				Company:          "Swiss-Development GmbH",
				Position:         "Frontend developer",
				EmploymentPeriod: "Sep 2018 - Aug 2021",
			},
			{
				Icon:             "/static/images/fhnw.png",
				Year:             "2014",
				Company:          "Fachhochschule Nordwestschweiz",
				Position:         "Apprenticeship application developer",
				EmploymentPeriod: "Aug 2014 - Sep 2018",
			},
		},
	}

	return experienceBlock
}

func getPageData(r *http.Request) *PageData {
	pageData := &PageData{}

	if r.URL.Path == "/" {
		pageData.Title = "jonasuethi.dev > Home"
	} else {
		pageData.Title = "jonasluethi.dev"
	}

	pageData.Navigation = getNavigation(r)
	pageData.AboutBlock = getAboutBlock()
	pageData.ExperienceBlock = getExperienceBlock()

	return pageData
}

func (handler PageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	responders.ResponWithHtml(w, "web/template/layout.html", "web/template/landing.html", getPageData(r))	
}
