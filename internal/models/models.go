package models

type Page struct {
	Layout string
	PageMeta PageMeta  
	Navigation Navigation	
	PageData []ContentBlock
}

type PageMeta struct {
	Title string
	Description string
}

type Navigation struct {
	Items []NavigationItem
}

type NavigationItem struct {
	Name   string
	Href   string
	Active bool
}

type ContentBlock interface {
	IsContentBlock() bool
	GetContentBlockType() string 
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

func (AboutBlock) IsContentBlock() bool {
	return true
}

func (AboutBlock) GetContentBlockType() string {
	return "about-block"
}

func (ExperienceBlock) IsContentBlock() bool {
	return true
}

func (ExperienceBlock) GetContentBlockType() string {
	return "experience-block"
}
