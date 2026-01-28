package store

import (
	"net/http"

	"github.com/jl54/jonasluethi-web/internal/models"
)

const (
	DefaultLayoutPath = "web/template/layout.html"
	DefaultTitle = "jonasluethi.dev"
	ProfilePicPath = "/static/images/profile-pic.jpg"
	SwissDevIcon = "/static/images/swiss-dev.png"
	FhnwIcon = "/static/images/fhnw.png"
)

func GetPage(r *http.Request) *models.Page {
	title := DefaultTitle

	if r.URL.Path == "/" {
		title += " > Home"
	}

	return &models.Page{
		Layout: DefaultLayoutPath,
		PageMeta: models.PageMeta{
			Title: title,
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
		AboutPic: ProfilePicPath,
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
				Description: "I am responsible for developing websites for small to medium-sized businesses using modern PHP frameworks such as Pimcore, Symfony, and Laravel. In addition, I set up automated deployment pipelines with GitHub Actions to deploy these websites to virtual private servers (VPS). I also maintain and updat these servers to ensure security, stability, and reliable operation.",
			},
			{
				Icon:             SwissDevIcon,
				Year:             "2019",
				Company:          "Swiss-Development GmbH",
				Position:         "Frontend developer",
				EmploymentPeriod: "Sep 2019 - Aug 2021",
				Description: "My first job after my apprenticeship. I was responsible for implementing given designs with HTML, CSS and Javascript. There were also some opportunities to create some websites using Angular and Vuejs. After a while of only working on the frontend, I've got the chance to work on the backends of some websites.",
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
