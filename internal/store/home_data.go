package store

import (
	"net/http"

	"github.com/jl54/jonasluethi-web/internal/models"
)

const (
	DefaultLayoutPath = "web/template/layout.html"
	DefaultTitle      = "jonasluethi.com"
	ProfilePicPath    = "/static/images/profile-pic.jpg"
	SwissDevIcon      = "/static/images/swiss-dev.png"
	FhnwIcon          = "/static/images/fhnw.png"
)

func GetPage(r *http.Request) *models.Page {
	title := DefaultTitle

	if r.URL.Path == "/" {
		title += " > Home"
	}

	return &models.Page{
		Layout: DefaultLayoutPath,
		PageMeta: models.PageMeta{
			Title:       title,
			Description: "Jonas Lüthi - Full-Stack Developer",
		},
		Navigation: getNavigation(r),
		PageData: []models.ContentBlock{
			getAboutBlock(),
			getExperienceBlock(),
		},
	}
}

func getNavigation(r *http.Request) models.Navigation {
	currentPath := r.URL.Path

	return models.Navigation{
		Items: []models.NavigationItem{
			{Name: "Home", Href: "/", Active: currentPath == "/"},
			{Name: "Projects", Href: "/projects", Active: currentPath == "projects"},
		},
	}
}

func newAboutSection(title string, items ...string) models.AboutBlockSection {
	content := make([]models.AboutBlockListItem, len(items))

	for i, item := range items {
		content[i] = models.AboutBlockListItem{Text: item}
	}

	return models.AboutBlockSection{Title: title, Content: content}
}

func getAboutBlock() models.AboutBlock {
	return models.AboutBlock{
		AboutPic:   ProfilePicPath,
		AboutTitle: "About me",
		AboutSections: []models.AboutBlockSection{
			newAboutSection(
				"Summary",
				"Passionate full-stack developer from Switzerland",
				"Always learning new technologies and improving my homelab",
				"Currently building automation and infrastructure tools for minecraft servers",
			),
			newAboutSection(
				"Currently Building",
				"Chunk Servers - CLI-Tool to deploy minecraft servers automatically",
				"Homelab - Self hosted services, monitoring and automation",
				"Portfolio site - It's this website",
			),
			newAboutSection(
				"Tech Stack",
				"Languages - PHP, JavaScript, Python, Go",
				"Frameworks - Symfony (Pimcore), Laravel, Angular, Vue",
				"Tools - Docker, Docker Compose, Kubernetes, Proxmox",
			),
			newAboutSection(
				"Get in touch",
				"to@jonasluethi.dev",
			),
		},
	}
}

func getExperienceBlock() models.ExperienceBlock {
	return models.ExperienceBlock{
		ExperienceTitle: "Experience",
		Experiences: []models.Experience{
			{
				Icon:             SwissDevIcon,
				Year:             "Today",
				Company:          "Swiss-Development GmbH",
				Position:         "Full-stack developer",
				EmploymentPeriod: "Sep 2021 - today",
				Description:      "I specialize in developing websites for small and medium-sized businesses using modern PHP frameworks such as Pimcore, Symfony, and Laravel. I also manage our automated deployment pipelines with GitHub Actions and maintain our server infrastructure to ensure security, stability, and reliability.",
			},
			{
				Icon:             SwissDevIcon,
				Year:             "2019",
				Company:          "Swiss-Development GmbH",
				Position:         "Frontend developer",
				EmploymentPeriod: "Sep 2019 - Aug 2021",
				Description:      "I started my career in a frontend developer role, implementing designs with HTML, CSS, and JavaScript. This position also gave me the opportunity to build websites with modern frameworks like Angular and Vue.js. After gaining experience on the frontend, I expanded my skill set to backend development, allowing me to contribute to both sides of the development process.",
			},
			{
				Icon:             FhnwIcon,
				Year:             "2014",
				Company:          "Fachhochschule Nordwestschweiz",
				Position:         "Apprenticeship application developer",
				EmploymentPeriod: "Aug 2014 - Sep 2018",
			},
		},
	}
}
