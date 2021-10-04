package json_new

import (
	"fmt"
	"regexp"
	"strings"
)

func Detect(resp *FetchResult) {
	products := make([]string, 0)
	//获取网页返回数据并赋值
	web_Content := strings.ToLower(string(resp.Content))
	//web_Headers := resp.Headers
	certString := string(resp.Certs)
	web_Certs := resp.Certs
	web_HeaderString := resp.HeaderString
	headerServerString := fmt.Sprintf("Server : %v\n", resp.Headers["Server"])

	fofajson, _ := Parse("fofa.json")
	for _, fp := range fofajson {
		//fofa指纹中的最后一项
		rules := fp.Rules
		matchFlag := false
		//其中的match
		//matchFlag := false
		//对每个json的最后一项进行迭代
		for _, onerule := range rules {

			//控制继续器
			ruleMatchContinueFlag := true

			for _, rule := range onerule {
				if !ruleMatchContinueFlag {
					break
				}
				lowerRuleContent := strings.ToLower(rule.Content)

				switch strings.Split(rule.Match, "_")[0] {

				case "banner":
					reBanner := regexp.MustCompile(`(?im)<\s*banner.*>(.*?)<\s*/\s*banner>`)
					matchResults := reBanner.FindAllString(web_Content, -1)
					if len(matchResults) == 0 {
						ruleMatchContinueFlag = false
						break
					}

					for _, matchResult := range matchResults {
						if !strings.Contains(strings.ToLower(matchResult), lowerRuleContent) {
							ruleMatchContinueFlag = false
							break
						}

					}

				case "title":
					reTitle := regexp.MustCompile(`(?im)<\s*title.*>(.*?)<\s*/\s*title>`)
					matchResults := reTitle.FindAllString(web_Content, -1)
					if len(matchResults) == 0 {
						ruleMatchContinueFlag = false
						break
					}

					for _, matchResult := range reTitle.FindAllString(web_Content, -1) {
						if !strings.Contains(strings.ToLower(matchResult), lowerRuleContent) {
							ruleMatchContinueFlag = false
						}
					}
				case "body":
					if !strings.Contains(web_Content, lowerRuleContent) {
						ruleMatchContinueFlag = false
					}
				case "header":
					if !strings.Contains(web_HeaderString, rule.Content) {
						ruleMatchContinueFlag = false
					}
				case "server":
					if !strings.Contains(headerServerString, rule.Content) {
						ruleMatchContinueFlag = false
					}
				case "cert":
					if (web_Certs == nil) || (web_Certs != nil && !strings.Contains(certString, rule.Content)) {
						ruleMatchContinueFlag = false
					}
				default:
					ruleMatchContinueFlag = false

				}
				// 单个rule之间是AND关系，匹配成功
				if ruleMatchContinueFlag {
					matchFlag = true
					break
				}

			}

		}
		// 多个rule之间是OR关系，匹配成功
		if matchFlag {
			products = append(products, fp.Product)
		}

	}
	PrintResult(resp.Url, products)

}

func PrintResult(target string, products []string) {
	fmt.Printf("[+] %s %s\n", target, strings.Join(products, ", "))
}
