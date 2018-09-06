package javascript

import (
	log "github.com/helmutkemper/seelog"
	"io/ioutil"
	"net/http"
)

func GetRoute(url, method string, header map[string]string) string {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Criticalf("getRoute.http.NewRequest.error: %v - %v", err.Error(), url)
		return ""
	}
	// GET https://www.ahgora.com.br/login/jwt
	// Authorization: basic base64_encode( user + ':' + senha );
	//req.Header.Add( "Authorization", "Bearer " + components.LoadJwtFromAhpi() )
	for k, v := range header {
		req.Header.Add(k, v)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Criticalf("getRoute.httpClient.Do.error: %v - %v", err.Error(), url)
		return ""
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Criticalf("getRoute.ioutil.ReadAll.error: %v - %v", err.Error(), url)
		return ""
	}
	res.Body.Close()

	return string(body)
}
