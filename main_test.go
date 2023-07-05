package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHookHandler(t *testing.T) {
	tt := map[string]struct {
		contentType string
		body        string
	}{
		"empty": {
			contentType: "application/x-www-form-urlencoded",
			body:        "",
		},
		"json": {
			body:        `{"test": "test"}`,
			contentType: "application/json",
		},
		"text": {
			body:        "this is a plain text message",
			contentType: "plain/text",
		},
		"form": {
			contentType: "multipart/form-data; boundary=------------------------e9b6fd531fc73c30",
			body: `--------------------------e9b6fd531fc73c30
Content-Disposition: form-data; name="payload"
Content-Type: application/json

{"event":"media.pause","user":false,"owner":true,"Account":{"id":19267583,"thumb":"https://plex.tv/users/055a1ca820bd5eab/avatar?c=1648049779","title":"melissar26"},"Server":{"title":"storm","uuid":"544b62f0b4f85d5d8f2c91696763d13578f5264a"},"Player":{"local":true,"publicAddress":"68.193.140.209","title":"Melissas-Air-2.localdomain","uuid":"lc67jl41c87g9tqmftb3n8le"},"Metadata":{"librarySectionType":"show","ratingKey":"146734","key":"/library/metadata/146734","parentRatingKey":"146719","grandparentRatingKey":"146560","guid":"plex://episode/5de8dbdf18b6d6001db44001","parentGuid":"plex://season/602e65269b7e9c002d714336","grandparentGuid":"plex://show/5d9c0855ffd9ef001e991de0","type":"episode","title":"The Fresh Princ-ipal","titleSort":"Fresh Princ-ipal","grandparentKey":"/library/metadata/146560","parentKey":"/library/metadata/146719","librarySectionTitle":"TV","librarySectionID":2,"librarySectionKey":"/library/sections/2","grandparentTitle":"Bob's Burgers","parentTitle":"Season 9","contentRating":"TV-PG","summary":"Things get out of hand when Louise wins a contest to be Principal for a Day. Meanwhile, Teddy attempts to help Bob when he suddenly finds himself unable to flip burgers","index":15,"parentIndex":9,"audienceRating":8.5,"viewOffset":622000,"lastViewedAt":1687821924,"year":2019,"thumb":"/library/metadata/146734/thumb/1684719871","art":"/library/metadata/146560/art/1687820065","parentThumb":"/library/metadata/146719/thumb/1667536815","grandparentThumb":"/library/metadata/146560/thumb/1687820065","grandparentArt":"/library/metadata/146560/art/1687820065","grandparentTheme":"/library/metadata/146560/theme/1687820065","duration":1320000,"originallyAvailableAt":"2019-03-03","addedAt":1667511134,"updatedAt":1684719871,"audienceRatingImage":"themoviedb://image.rating","Guid":[{"id":"imdb://tt9413736"},{"id":"tmdb://1707381"},{"id":"tvdb://7047159"}],"Rating":[{"image":"themoviedb://image.rating","value":8.5,"type":"audience"}],"Director":[{"id":512241,"filter":"director=512241","tag":"Tyree Dillihay","tagKey":"5e164efdef1040003f263741"}],"Writer":[{"id":512271,"filter":"writer=512271","tag":"Greg Thompson","tagKey":"6323590c7ced2ab196a729a1"}],"Role":[{"id":352424,"filter":"actor=352424","tag":"Joe Lo Truglio","tagKey":"5d77682a4de0ee001fcc979e","role":"Don (voice)","thumb":"https://metadata-static.plex.tv/2/people/20ee2d145dbae3b1bb680e58b56bcd5a.jpg"},{"id":319798,"filter":"actor=319798","tag":"Sarah Silverman","tagKey":"5d776826999c64001ec2c60b","role":"Ms. Schnur (voice)","thumb":"https://metadata-static.plex.tv/9/people/9706b7a4ef2194dd245292bda89fd1c2.jpg"},{"id":512112,"filter":"actor=512112","tag":"Melissa Bardin Galsky","tagKey":"5d9c07fdd4f2a9001f7fce1f","role":"Miss Jacobson (voice)"},{"id":344123,"filter":"actor=344123","tag":"David Herman","tagKey":"5d77682a961905001eb91e4a","role":"Mr. Frond (voice) / Mr. Branca (voice)","thumb":"https://metadata-static.plex.tv/people/5d77682a961905001eb91e4a.jpg"},{"id":512111,"filter":"actor=512111","tag":"Jay Johnston","tagKey":"5d7768378a7581001f12d990","role":"Jimmy Pesto (voice)","thumb":"https://metadata-static.plex.tv/people/5d7768378a7581001f12d990.jpg"},{"id":355183,"filter":"actor=355183","tag":"Bobby Tisdale","tagKey":"5d77682a54c0f0001f3022b6","role":"Zeke (voice)","thumb":"https://metadata-static.plex.tv/6/people/687d2722fa896dca276051d126621f17.jpg"},{"id":345492,"filter":"actor=345492","tag":"Jenny Slate","tagKey":"5d776896431c830024c100cd","role":"Tammy (voice)","thumb":"https://metadata-static.plex.tv/8/people/8a2386041a3d151bb8474dfd0072d85c.jpg"},{"id":322950,"filter":"actor=322950","tag":"Brian Huskey","tagKey":"5d7768327228e5001f1de168","role":"Regular Sized Rudy (voice)","thumb":"https://metadata-static.plex.tv/3/people/398a1c3fe26aa6cfba6240d7c8265787.jpg"}]}}
--------------------------e9b6fd531fc73c30--
      `,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			r := chi.NewRouter()
			r.Post("/hook", hookHandler(nil))
			ts := httptest.NewServer(r)
			defer ts.Close()

			res, err := http.Post(ts.URL+"/hook", tc.contentType, strings.NewReader(tc.body))
			if err != nil {
				t.Errorf("could not post to /hook: %+v", err)
			}

			body, err := io.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				t.Errorf("could not read response body: %+v", err)
			}

			if res.StatusCode != http.StatusOK {
				t.Errorf("Want status '%d', got '%d'", http.StatusOK, res.StatusCode)
			}

			if string(body) != "" {
				t.Errorf("Wanted empty response, got %q", body)
			}
		})
	}
}
