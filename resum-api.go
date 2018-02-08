package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		switch api_request := r.URL.Path; api_request {
		case "/api/joey":
			JoeyResponse(w)
		case "/api/skills":
			SkillsResponse(w)
		case "/api/contact":
			ContactResponse(w)
		case "/api/jobs":
			JobsResponse(w)
		default:
			DefaultResponse(w)
		}
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}

func DefaultResponse(w http.ResponseWriter) {
	fmt.Fprintf(w, "Invalid API Call Check github.com/joeygrasso/resum-api for documentation. \n")
}

func JoeyResponse(w http.ResponseWriter) {
	fmt.Fprintf(w, "You know my name! \n")
}

func SkillsResponse(w http.ResponseWriter) {
	skills := []string{"Linux", "System Administration", "Puppet", "Python", "Bash", "Go", "Nginx", "Git", "Virtualization", "KVM", "Docker", "APIs", "System Monitoring"}
	skills_formatted, _ := json.Marshal(skills)
	fmt.Fprintf(w, string(skills_formatted))
}

func ContactResponse(w http.ResponseWriter) {
	contact_info := map[string]string{
		"email":    "me@joeygrasso.com",
		"twitter":  "@grassojoey",
		"linkedin": "linkedin.com/in/grassojoey",
		"github":   "github.com/joeygrasso",
		"web":      "JoeyGrasso.com",
	}

	contact_info_formatted, err := json.MarshalIndent(contact_info, "", " ")
	if err != nil {
		log.Fatalf("JSON Marshaling failed: %s", err)
	}
	fmt.Fprintf(w, string(contact_info_formatted))
}

func JobsResponse(w http.ResponseWriter) {
	type Job struct {
		Title, StartDate, EndDate, Location, Employer, Description string
		Accomplishments                                            []string
	}

	var jobs = []Job{
		{Title: "Web Application Programmer", StartDate: "March 2012", EndDate: "January2014", Location: "Savannah, GA", Employer: "Armstrong State University", Description: "Support campus constituents and other IT teams with strong emphasis on application development", Accomplishments: []string{"Developed a data visualization application for displaying project management and tracking dashboards.", "Developed a service status dashboard that displays current status of a system service and sends out a tweet when updated.", "Developed a searchable employee & department directory that is updated by parsing a Google Spreadsheet and is updated via cron."}},
		{Title: "IT Technician", StartDate: "August 2014", EndDate: "May 2015", Location: "Savannah, GA", Employer: "Armstrong State University", Description: "Provide research and hardware/software support for the Dept. of Computer Science and Information Technology", Accomplishments: []string{"Imaged and maintained a Windows Server lab of 24 machines", "Created, documented, and maintained images for 3 labs consisting of over 100 PCs and 50 Macs"}},
		{Title: "Systems Specialist", StartDate: "January 2014", EndDate: "May 2015", Location: "Savannah, GA", Employer: "Vaden Automotive Group", Description: "Assist in the development, maintenance, and support the Revlink Auto application including infrastructure, documentation, and end user support.", Accomplishments: []string{"Developed an automated deployment script using Rackspace's API that created server instances, user accounts, and infrastructure necessary for application deployment", "Developed a unified login portal that allowed users to log into software across multiple domains through one interface.", "Conducted research and analysis of configuration management tools to automate deployment of web based applications."}},
		{Title: "Systems Engineer", StartDate: "June 2015", EndDate: "Present", Location: "Atlanta, GA", Employer: "MailChimp", Description: "work closely with developers, marketing, and data scientists to build and support the infrastructure for large-scale projects and products that reach millions of users around the world.", Accomplishments: []string{"Still in progress."}},
	}

	jobs_formatted, err := json.MarshalIndent(jobs, "", " ")
	if err != nil {
		log.Fatalf("JSON Marshaling failed: %s", err)
	}
	fmt.Fprintf(w, string(jobs_formatted))

}
