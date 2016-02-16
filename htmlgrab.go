package htmlcrawler
import (
	"encoding/xml"
	"fmt"
)

type HtmlTag struct {
	Content   string `xml:",chardata"`
	Inner     string `xml:",innerxml"`
	ChildTags []HtmlTag `xml:",any"`
	Class     string `xml:"class,attr"`
	Id        string `xml:"id,attr"`
	XMLName   xml.Name
}

func getTagsFromTag(tagName string, className string, idName string, htmlTag *HtmlTag, matchingDivs *[]HtmlTag){
	fmt.Println(htmlTag.Class)
	if htmlTag.XMLName.Local == tagName {
		if htmlTag.Class == className && htmlTag.Id == idName {

			*matchingDivs = append(*matchingDivs, *htmlTag)
		}

	}

	for _, tag := range htmlTag.ChildTags {
		getTagsFromTag(tagName, className, idName, &tag, matchingDivs)
	}
}

func GetHtmlTagsFromHtmlString(tagName string, className string, idName string, html string) []HtmlTag {
	root := new(HtmlTag)
	err := xml.Unmarshal([]byte(html), root)
	if err != nil {
		panic(err)
	}
	err, bodyTag := GrabBodyTag(root)
	if err != nil {
		panic(err)
	}

	return GetTagSliceFromTag(tagName, className, idName, bodyTag)

}

func GetTagSliceFromTag(tagName string, className string, idName string, htmlTag *HtmlTag) []HtmlTag {
	var matchingDivs []HtmlTag
	getTagsFromTag(tagName, className, idName, htmlTag, &matchingDivs)
	return matchingDivs
}
