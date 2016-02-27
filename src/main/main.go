package main

import (
  "net/http"
  "fmt"
  "os"
  "text/template"
  "bufio"
  "strings"
  "viewmodels"
)

func form(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    r.ParseForm()
    fmt.Println("first:", r.Form["item1"])
    fmt.Println("second", r.Form["item2"])
}

func main() {
	http.HandleFunc("/postForm", form)
		templates := populateTemplates()

		http.HandleFunc("/", 
		func(w http.ResponseWriter, req *http.Request) {
			requestedFile := req.URL.Path[1:]
			template :=
				templates.Lookup(requestedFile + ".html")
			fmt.Println("test: ", req.URL.Path)	
	
			var context interface{} = nil
			fmt.Println(len(req.URL.Path))

//			if len(req.URL.Path) == 1 {
//				fmt.Println("in if statement")
//				context = viewmodels.GetHome()
//				template = templates.Lookup("index.html")
//				}
			switch requestedFile {
			case "index":
				context = viewmodels.GetHome()
			case "page1":
				context = viewmodels.GetPage1()
			case "":
				context = viewmodels.GetHome()
//			case "page2":
//				context = viewmodels.GetPage2()	
			}
			if template != nil {
				template.Execute(w, context)
			} else {
				w.WriteHeader(404)
			}
		})
		
		http.HandleFunc("/html/", serveResource)
		http.HandleFunc("/img/", serveResource)
		http.HandleFunc("/css/", serveResource)
		http.HandleFunc("/js/", serveResource)
		http.HandleFunc("/fonts/", serveResource)

		http.ListenAndServe(":8000", nil)
	//http.ListenAndServe(":8000", http.FileServer(http.Dir("public")));
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else if strings.HasSuffix(path, ".js") {
        contentType = "application/javascript"
    } else if strings.HasSuffix(path, ".html") {
        contentType = "text/html"
    } else {
		contentType = "text/plain"
	}
	
	f, err := os.Open(path)
	
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}


func populateTemplates() *template.Template {
	result := template.New("templates")
	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()
	
	templatePathsRaw, _ := templateFolder.Readdir(-1) //Negative number means we want all the content in one pass
	
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
				basePath + "/" + pathInfo.Name())
		}
	}
	
	result.ParseFiles(*templatePaths...)
	
	return result
	}

