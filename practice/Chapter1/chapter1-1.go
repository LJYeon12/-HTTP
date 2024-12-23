package main

import (
	"fmt" // 입출력 기능 제공
	"log"
	"net/http"          // 서버와 클라이언트를 구현하기 위한 기본적인 HTTP 기능을 제공하는 패키지
	"net/http/httputil" // HTTP 메시지와 디버깅 작업을 돕기 위한 유틸리티 패키지. 요청 및 응답 데이터를 자세히 분석하거나 디버깅할 때 유용.
)

func handler(w http.ResponseWriter, r *http.Request) { // 브라우저나 curl 커맨드 등 클라이언트가 접속했을 때 호출
	dump, err := httputil.DumpRequest(r, true) // 클라이언트의 요청정보
	if err != nil {                            // nil : 어떠한 것도 참조하지 않은 상태
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body><html>\n")
}

func main() { // main : Http 서버 초기화
	var httpServer http.Server
	http.HandleFunc("/", handler) // "/"에 있을 때 handler 함수 호출
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888" // 18888포트로 시작 -> 원래는 80포트가 디폴트
	log.Println(httpServer.ListenAndServe())
}
