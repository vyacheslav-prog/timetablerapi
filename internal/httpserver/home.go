package httpserver

import "net/http"

func handleHome(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	writeResponse(w, []byte(`
{
	"actions": [
		{
			"http": "GET /performer-boards/{boardId}",
			"title": "view a performer board"
		},
		{
			"http": "POST /performers",
			"title": "add performer"
		},
		{
			"http": "POST /tasks",
			"title": "add task"
		}
	]
}
	`))
}
