package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func main() {
	//解析html
	//查找元素，find
	//查找子元素childernfiltered
	//获取内容text//html
	//获取属性attr
	//遍历元素each
	tags := "goquery"
	url := "https://pkg.go.dev/search?q=" + tags
	//发起http请求并返回document
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatalln(err)
	} else {
		//在document中通过选择器去找元素《tagName》
		sele := doc.Find("a") //获取所有a标签
		sele.Each(func(index int, tag *goquery.Selection) {
			//href, _ := tag.Attr("href")
			tag.Html() //如果获取的内容包含其他标签的话也会将其他标签的内容显示出来
			tag.Text() //如果要获取的标签内的内容包含其他标签的话不会获取其他标签
			//fmt.Println(tag.Html())
		})
	}
	//class选择器
	fmt.Println("class=================")
	//selection := doc.Find(".main") //class选择器
	//子孙选择器写法
	doc.Find(".go-Container .go-Content.SearchResults div:nth-child(2) a").Each(func(index int, tag *goquery.Selection) {
		href, _ := tag.Attr("href")
		fmt.Println(href)
	})
	//子选择器
	//doc.Find(select1).Children(select2)
}
