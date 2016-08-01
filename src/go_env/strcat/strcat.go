package strcat

import (
	"encoding/json"
	"expvar"
	"net/http"
)

type strcatArg struct {
	Strings []string `json:"strings"`
}

type strcatReply struct {
	String string `json:"string"`
}

type strcat struct {
	count *expvar.Int
}

func (s *strcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var arg strcatArg
	if err := json.NewDecoder(r.Body).Decode(&arg); err != nil {
		http.Error(w, "invalid arg:"+err.Error(), http.StatusBadRequest)
		return
	}
	reply := strcatReply{
		String: joinStrings(arg.Strings),
	}
	if err := json.NewEncoder(w).Encode(reply); err != nil {
		http.Error(w, "reply error:"+err.Error(), http.StatusInternalServerError)
		return
	}
	s.count.Add(1)
}

func joinStrings(strings []string) string {
	ret := ""
	for _, s := range strings {
		ret += s
	}
	return ret
}
