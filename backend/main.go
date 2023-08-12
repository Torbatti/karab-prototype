package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type Job struct {
	Id           string
	CompanyTitle string
	JobTitle     string
	City         string
	Time         string
}

func main() {
	fmt.Println("Server Started...")

	css := http.FileServer(http.Dir("./public/css"))
	http.Handle("/css/", http.StripPrefix("/css", css))

	svg := http.FileServer(http.Dir("./public/svg"))
	http.Handle("/svg/", http.StripPrefix("/svg", svg))

	fonts := http.FileServer(http.Dir("./public/fonts"))
	http.Handle("/fonts/", http.StripPrefix("/fonts", fonts))

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./frontend/index.html"))
		tmpl.Execute(w, nil)
	}
	http.HandleFunc("/", h1)

	jobs := map[string][]Job{
		"Jobs": {
			{Id: "1", CompanyTitle: "شرکت سهامی عام اول", JobTitle: "منشی", City: "تهران", Time: "امروز"},
			{Id: "2", CompanyTitle: "شرکت سهامی عام اول", JobTitle: "کارگر", City: "تهران", Time: "امروز"},
			{Id: "3", CompanyTitle: "شرکت سهامی عام اول", JobTitle: "کارشناس برنامه نویسی", City: "تهران", Time: "دیروز"},
			{Id: "4", CompanyTitle: "شرکت سهامی عام اول", JobTitle: "منشی", City: "تهران", Time: "دیروز"},

			{Id: "5", CompanyTitle: "شرکت سهامی عام سوم", JobTitle: "مدرس زبان", City: "تهران", Time: "امروز"},
			{Id: "6", CompanyTitle: "شرکت سهامی عام سوم", JobTitle: "حسابداری", City: "تهران", Time: "دیروز"},

			{Id: "7", CompanyTitle: "شرکت سهامی عام چهاردهم", JobTitle: "طراح لباس", City: "تهران", Time: "امروز"},
		},
	}
	random_string := []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Sem fringilla ut morbi tincidunt augue interdum velit. Lacus sed turpis tincidunt id aliquet risus feugiat. Quis auctor elit sed vulputate mi sit. Lorem dolor sed viverra ipsum nunc aliquet bibendum. Ipsum a arcu cursus vitae. Mauris pellentesque pulvinar pellentesque habitant morbi tristique senectus. Cursus mattis molestie a iaculis at erat pellentesque adipiscing commodo. Vestibulum mattis ullamcorper velit sed ullamcorper. Ut ornare lectus sit amet est placerat in egestas. Sapien et ligula ullamcorper malesuada proin libero nunc consequat interdum. Bibendum at varius vel pharetra vel turpis. Egestas pretium aenean pharetra magna. Pharetra massa massa ultricies mi quis hendrerit. Vestibulum sed arcu non odio euismod lacinia. Egestas congue quisque egestas diam in arcu. Sit amet nisl purus in mollis. Maecenas volutpat blandit aliquam etiam erat velit scelerisque. Dui nunc mattis enim ut tellus elementum sagittis. Mi in nulla posuere sollicitudin aliquam.",
		"Nisi scelerisque eu ultrices vitae auctor eu augue. Purus in massa tempor nec feugiat nisl pretium fusce id. Erat imperdiet sed euismod nisi porta. Malesuada bibendum arcu vitae elementum. Ultricies mi quis hendrerit dolor magna eget est. Purus non enim praesent elementum facilisis leo vel fringilla est. Vestibulum lectus mauris ultrices eros. At elementum eu facilisis sed. Pellentesque diam volutpat commodo sed egestas. Tortor vitae purus faucibus ornare suspendisse sed nisi lacus sed. Rhoncus dolor purus non enim. Vestibulum mattis ullamcorper velit sed ullamcorper morbi. Mattis nunc sed blandit libero volutpat sed cras. Egestas egestas fringilla phasellus faucibus scelerisque. Felis donec et odio pellentesque diam volutpat. Pretium aenean pharetra magna ac placerat vestibulum.",
		"Sagittis purus sit amet volutpat consequat mauris nunc congue. Lacinia at quis risus sed vulputate. Arcu odio ut sem nulla. Varius morbi enim nunc faucibus a pellentesque sit amet porttitor. Risus nullam eget felis eget nunc. Lacus sed viverra tellus in hac. Cum sociis natoque penatibus et magnis dis parturient montes nascetur. Risus in hendrerit gravida rutrum quisque non. Nunc id cursus metus aliquam eleifend mi in. Non curabitur gravida arcu ac tortor dignissim convallis aenean et. Eget egestas purus viverra accumsan. A iaculis at erat pellentesque. Netus et malesuada fames ac turpis egestas sed tempus urna. Nec sagittis aliquam malesuada bibendum. Magna sit amet purus gravida quis blandit turpis cursus in.",
		"Enim praesent elementum facilisis leo. Donec adipiscing tristique risus nec feugiat in fermentum. Maecenas volutpat blandit aliquam etiam erat. Purus gravida quis blandit turpis cursus in. At in tellus integer feugiat scelerisque. Quam id leo in vitae. Convallis posuere morbi leo urna. Justo laoreet sit amet cursus. Mi sit amet mauris commodo. Quam quisque id diam vel quam. Felis eget velit aliquet sagittis id consectetur purus ut faucibus. Adipiscing elit ut aliquam purus sit amet luctus venenatis. Urna et pharetra pharetra massa massa ultricies mi. Nullam eget felis eget nunc. Nec dui nunc mattis enim ut tellus elementum sagittis vitae. Dis parturient montes nascetur ridiculus. Justo donec enim diam vulputate. Dictum non consectetur a erat nam at lectus urna. Nunc pulvinar sapien et ligula ullamcorper.",
		"Purus sit amet volutpat consequat. A pellentesque sit amet porttitor eget dolor morbi. Ac turpis egestas sed tempus urna et. Iaculis urna id volutpat lacus laoreet non curabitur gravida. Lorem donec massa sapien faucibus et molestie ac. Proin fermentum leo vel orci porta non pulvinar neque. Nisi quis eleifend quam adipiscing vitae proin sagittis nisl. Quisque id diam vel quam elementum pulvinar. Nisi scelerisque eu ultrices vitae auctor eu augue. Suspendisse faucibus interdum posuere lorem ipsum dolor. Congue eu consequat ac felis donec. Cursus risus at ultrices mi tempus imperdiet. Laoreet sit amet cursus sit amet. Hac habitasse platea dictumst quisque. Justo donec enim diam vulputate ut pharetra sit amet. Nullam ac tortor vitae purus faucibus.",
		"At consectetur lorem donec massa sapien. Diam donec adipiscing tristique risus nec feugiat in. Velit dignissim sodales ut eu. Morbi tincidunt ornare massa eget. Porttitor leo a diam sollicitudin tempor id eu. Vivamus arcu felis bibendum ut tristique et egestas quis. Senectus et netus et malesuada fames ac turpis egestas integer. Felis imperdiet proin fermentum leo vel orci porta non pulvinar. Lobortis scelerisque fermentum dui faucibus in ornare quam viverra. Commodo viverra maecenas accumsan lacus vel facilisis volutpat est velit. Malesuada fames ac turpis egestas integer eget aliquet nibh praesent. Purus faucibus ornare suspendisse sed. Fermentum et sollicitudin ac orci phasellus egestas tellus rutrum. Tellus integer feugiat scelerisque varius morbi enim. Commodo sed egestas egestas fringilla phasellus faucibus scelerisque eleifend donec. Ut pharetra sit amet aliquam id diam maecenas ultricies. Congue eu consequat ac felis donec et. Nunc scelerisque viverra mauris in aliquam.",
		"Phasellus egestas tellus rutrum tellus pellentesque eu tincidunt tortor. Ornare suspendisse sed nisi lacus sed. Neque aliquam vestibulum morbi blandit cursus risus. Egestas pretium aenean pharetra magna ac placerat vestibulum lectus. Risus commodo viverra maecenas accumsan lacus vel. Curabitur gravida arcu ac tortor dignissim convallis aenean. Nunc aliquet bibendum enim facilisis gravida neque convallis. Sed egestas egestas fringilla phasellus faucibus scelerisque eleifend. Erat nam at lectus urna duis convallis convallis tellus. Id cursus metus aliquam eleifend. Iaculis nunc sed augue lacus. Quis blandit turpis cursus in. Amet consectetur adipiscing elit pellentesque habitant morbi tristique. Congue eu consequat ac felis donec et odio pellentesque. Eu consequat ac felis donec et odio pellentesque. Sapien eget mi proin sed libero enim. Vitae semper quis lectus nulla at volutpat diam ut.",
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(350 * time.Microsecond)

		// fmt.Println(r)

		tmpl := template.Must(template.ParseFiles("./frontend/search.html"))
		tmpl.Execute(w, jobs)
	}
	http.HandleFunc("/search.html", h2)

	h3 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(250 * time.Microsecond)

		rand_num := rand.Intn(7)
		fmt.Printf("jobs[\"Jobs\"]: %v\n", jobs["Jobs"][rand_num])

		// htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		htmlStr := fmt.Sprintf(`
		<div class=''>
		<h1>
			%s - %s - %s - %s
		</h1>
		<h1>%s</h1>
		</div>
		`, jobs["Jobs"][rand_num].CompanyTitle,
			jobs["Jobs"][rand_num].JobTitle,
			jobs["Jobs"][rand_num].City,
			jobs["Jobs"][rand_num].Time,
			random_string[rand_num],
		)
		tmpl, _ := template.New("t").Parse(htmlStr)

		tmpl.Execute(w, tmpl)
	}
	http.HandleFunc("/jobs", h3)

	h4 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(250 * time.Millisecond)

		for i := 0; i < 7; i++ {
			rand_num := rand.Intn(7)

			tmpl := template.Must(template.ParseFiles("./frontend/search.html"))
			tmpl.ExecuteTemplate(w, "job-list-element", Job{Id: jobs["Jobs"][rand_num].Id, CompanyTitle: jobs["Jobs"][rand_num].CompanyTitle, JobTitle: jobs["Jobs"][rand_num].JobTitle, City: jobs["Jobs"][rand_num].City, Time: jobs["Jobs"][rand_num].Time})
		}
	}
	http.HandleFunc("/next-page", h4)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
