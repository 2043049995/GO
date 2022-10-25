package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	for i := 30000; i < 32767; i++ {
		for j := 1; j < 254; j++ {
			//html := Gethtml("http://192.168.80.%d:%d")
			html := fmt.Sprintf("http://192.168.80.%d:%d", j, i)
			html1 := Gethtml(html)
			pattern := "全部职位分类"
			re := regexp.MustCompile(pattern)
			token1 := re.FindStringSubmatch(html1)
			a := token1[0]
			if a == "全部职位分类" {
				fmt.Println(a, j, i)
			}

		}

	}
}
func Gethtml(html string) string {
	client := http.Client{}
	defer client.CloseIdleConnections()
	req, _ := http.NewRequest("GET", html, nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 SLBrowser/8.0.0.2242 SLBChan/10")
	req.Header.Add("Cookie", "lv=0; _uid=87521496; UID=87521496; vc=808D0FD1A89747244391D55F6DE06F9A; vc2=DFA30C87226B44B24836494B52BEE86D; xxtenc=0440b56fa6420d5e044d1fdbd451f4d2; uname=1805040233; fid=11254; uf=b2d2c93beefa90dc63092c4d62ff7663d7346345223c81bfdbb5f7e35556341108f4aad531127f0caacdb654d55a14d0913b662843f1f4ad6d92e371d7fdf64467c3000f6e78d08dc04fd8193a78cec6c992d031135b2ab71327bc3c9dfe4dcfe488250a65360ab5aa2ebad65cd196bb; _d=1665800212825; vc3=ZSPZPObcOf2XyYAp7gsAb6UPfVU5isJT9zs%2BTTUrl88neucLjrzjzDBALl8DC1149WwIRXmnnvVEEPA5XvrYmLmMKCOqHO6UwRFiGJmJkCCEIRmwIrca%2BNg3nugBLj3dCSzuKM2xoLCZ9qt%2BUrU78In%2BYaN%2BHVtopBxqWsS6c1k%3D9614338ef414352052a55a2b46cb96cd; DSSTASH_LOG=C_38-UN_10726-US_87521496-T_1665800212827; oa_uid=87521496; oa_name=%E6%9D%8E%E6%9C%9D%E9%98%B3; route=3b84754f32aaf54a995f7d41f1a28848; JSESSIONID=AB57CBCED547E7C45E6A7F3DF6F0AE53.reserve_web_124; oa_deptid=11254; oa_enc=6f85c610b0008803ccbe9dc516fe07b7")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("遇到错误：", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
