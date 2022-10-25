package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

func main() {

	CookieTest()
}
func CookieTest() {
	//fmt.Println(login())
	html1 := Gethtml1("http://office.chaoxing.com/front/third/apps/seat/list?deptIdEnc=f524a3b5b616c8f7")
	//fmt.Println(html1)
	pattern := "&pageToken.*&fidEnc"
	re := regexp.MustCompile(pattern)
	token1 := re.FindStringSubmatch(html1)[0][16:48]
	//fmt.Println(token1)
	url2 := fmt.Sprintf("http://office.chaoxing.com/front/third/apps/seat/select?id=2590&day=%s&backLevel=2&pageToken=%s&fidEnc=f524a3b5b616c8f7", GetodayTime(), token1)
	html2 := Gethtml1(url2)
	//fmt.Println(url2)
	pattern2 := "token: '.*"
	re2 := regexp.MustCompile(pattern2)
	token2 := re2.FindStringSubmatch(html2)[0][8:40]
	//fmt.Println(len(token2))
	if len(token2) != 32 {
		fmt.Println("cookie挂了，要重新生成，先手动抢吧")
	} else {
		fmt.Println("可以开抢了")
	}
	//return "0"
}
func Gethtml1(html string) string {
	client := http.Client{}
	defer client.CloseIdleConnections()
	req, _ := http.NewRequest("GET", html, nil)
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 SLBrowser/8.0.0.2242 SLBChan/10")
	req.Header.Add("Cookie", "JSESSIONID=314FFFD0193B4AB64B9345706D72B2D5.reserve_web_127; route=e37b564281be6bf42f59a301ebbe05b5; lv=0; fid=2275; _uid=108003249; uf=da0883eb5260151ed549c1af6676580c58468f2ee833d3480244ad59c4a450ff7c5eea09435d2184884f9f0de45ac9d0913b662843f1f4ad6d92e371d7fdf64419c2e6aa68aac736fd68be96b6183b1a5986f6450bbdeedc0d784ae784f986cf7821914bdc659a34; _d=1666102160517; UID=108003249; vc=1FE1684ED127A4EB5464BEDED937A955; vc2=27F3243C893D6135217DBCAF3A5345D1; vc3=YaNKOVA7qMNew9arUBngmHogJfDyg6VyF9jH0%2BSuu%2BJAjvk%2FwBMjVkOSttxQyoWTzwfwdvWCIbBKoJdGFU1%2Foe%2FVPRp%2B0kyOXMMswEm2YDdWvrjjtAEoerIV%2FyaaVwnZXfjDA2VobBbMj2WIvDd1mKk97rbqz7Kl%2BP4Ze6tLz9w%3D118dada68c6c4884b9ce06f6d7ba6c97; xxtenc=0b234e854effcda3f9cacd4a543dc955; DSSTASH_LOG=C_38-UN_863-US_108003249-T_1666102160518")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("遇到错误：", err)
	}
	body, _ := io.Copy(os.Stdout, resp.Body)

	return string(body)
}

func GetodayTime() string {
	time1 := time.Now()
	timeday := time1.Format("2006-01-02")
	return timeday
}
